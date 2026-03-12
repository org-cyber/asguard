// internal/engine/engine.go
package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"general/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

// RuleEngine orchestrates validation
type RuleEngine struct {
	queries   *db.Queries
	evaluator *Evaluator
}


// NewRuleEngine creates the engine
func NewRuleEngine(queries *db.Queries) (*RuleEngine, error) {
	eval, err := NewEvaluator()
	if err != nil {
		return nil, err
	}

	return &RuleEngine{
		queries:   queries,
		evaluator: eval,
	}, nil
}

// ValidationRequest is what users send
type ValidationRequest struct {
	Context string                 `json:"context"`
	Input   map[string]interface{} `json:"input"`
}

// ValidationResult is what we return
type ValidationResult struct {
	Decision         string   `json:"decision"`      // allow, block, challenge
	Score            int      `json:"score"`         // cumulative score
	Reason           string   `json:"reason"`        // explanation
	RulesMatched     []string `json:"rules_matched"` // which rules fired
	ProcessingTimeMs int      `json:"processing_time_ms"`
}

// RuleRequest is the body for creating/updating a rule via the API
type RuleRequest struct {
	Name      string `json:"name"`
	Context   string `json:"context"`
	Condition string `json:"condition"`
	Action    string `json:"action"`
	Score     *int   `json:"score,omitempty"`
	Priority  int    `json:"priority"`
	Enabled   bool   `json:"enabled"`
}

// RuleResponse is the API-facing representation of a rule
type RuleResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Context   string `json:"context"`
	Condition string `json:"condition"`
	Action    string `json:"action"`
	Score     *int   `json:"score,omitempty"`
	Priority  int    `json:"priority"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// RuleToResponse converts a db.Rule into an API-facing RuleResponse
func RuleToResponse(rule db.Rule) RuleResponse {
	resp := RuleResponse{
		ID:        pgUUIDToString(rule.ID),
		Name:      rule.Name,
		Context:   rule.Context,
		Condition: rule.Condition,
		Action:    rule.Action,
		Priority:  int(rule.Priority),
		Enabled:   rule.Enabled,
		CreatedAt: rule.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: rule.UpdatedAt.Time.Format(time.RFC3339),
	}
	if rule.Score != nil {
		s := int(*rule.Score)
		resp.Score = &s
	}
	return resp
}

// pgUUIDToString converts a pgtype.UUID to its string representation
func pgUUIDToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// Validate runs all rules for a context and returns decision
func (re *RuleEngine) Validate(ctx context.Context, req ValidationRequest) (*ValidationResult, error) {
	start := time.Now()

	// 1. Fetch active rules for this context (already sorted by priority)
	zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

	rules, err := re.queries.ListRulesByContext(ctx, db.ListRulesByContextParams{
		TenantID: zeroUUID,
		Context:  req.Context,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load rules: %w", err)
	}

	result := &ValidationResult{
		Decision:     "allow", // Default: allow unless rule says otherwise
		Score:        0,
		RulesMatched: []string{},
	}

	// 2. Evaluate each rule in priority order
	for _, rule := range rules {
		// Compile the CEL expression (in production, cache this)
		program, err := re.evaluator.CompileRule(rule.Condition)
		if err != nil {
			// Log error but continue with other rules
			fmt.Printf("Failed to compile rule %s: %v\n", rule.Name, err)
			continue
		}

		// Evaluate against input
		matched, err := re.evaluator.Evaluate(program, req.Input)
		if err != nil {
			fmt.Printf("Failed to evaluate rule %s: %v\n", rule.Name, err)
			continue
		}

		if !matched {
			continue // Rule didn't match, check next
		}

		// Rule matched!
		result.RulesMatched = append(result.RulesMatched, rule.Name)

		// Apply action
		switch rule.Action {
		case "allow":
			result.Decision = "allow"
			result.Reason = fmt.Sprintf("Rule '%s' explicitly allowed", rule.Name)
			return result, nil // Stop evaluation

		case "block":
			result.Decision = "block"
			result.Reason = fmt.Sprintf("Rule '%s' blocked: %s", rule.Name, rule.Condition)
			return result, nil // Stop evaluation

		case "challenge":
			result.Decision = "challenge"
			result.Reason = fmt.Sprintf("Rule '%s' requires challenge", rule.Name)
			// Continue to see if something blocks, but minimum is challenge

		case "score":
			if rule.Score != nil {
				result.Score += int(*rule.Score)
			}

		case "flag":
			// Just log it, continue evaluation
		}
	}

	// 3. Check cumulative score thresholds (simplified)
	if result.Score >= 100 {
		result.Decision = "block"
		result.Reason = fmt.Sprintf("Cumulative score %d exceeded threshold", result.Score)
	} else if result.Score >= 50 && result.Decision == "allow" {
		result.Decision = "challenge"
		result.Reason = fmt.Sprintf("Cumulative score %d requires verification", result.Score)
	}

	// 4. Log the decision
	processingTime := int(time.Since(start).Milliseconds())

	// Convert matched rules to UUIDs for storage
	matchedRuleIDs := make([]pgtype.UUID, len(result.RulesMatched))
	// ... (simplified, would map names to IDs)

	inputBytes, _ := json.Marshal(req.Input)
	finalScore := int32(result.Score)

	_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
		TenantID:         zeroUUID,
		Context:          req.Context,
		Input:            inputBytes, // JSONB
		MatchedRules:     matchedRuleIDs,
		Decision:         result.Decision,
		FinalScore:       &finalScore,
		Reason:           &result.Reason,
		ProcessingTimeMs: int32(processingTime),
	})
	if err != nil {
		// Log but don't fail the request
		fmt.Printf("Failed to log decision: %v\n", err)
	}

	result.ProcessingTimeMs = processingTime
	return result, nil
}
