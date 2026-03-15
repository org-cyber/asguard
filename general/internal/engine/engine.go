// internal/engine/engine.go
package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"general/internal/db"
	"sync"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/jackc/pgx/v5/pgtype"
)

// RuleEngine orchestrates validation
type RuleEngine struct {
	queries       *db.Queries
	Evaluator     *Evaluator
	programeCache sync.Map //// map[string]cel.Program (key = rule ID as string)
}

// NewRuleEngine creates the engine
func NewRuleEngine(queries *db.Queries) (*RuleEngine, error) {
	eval, err := NewEvaluator()
	if err != nil {
		return nil, err
	}

	return &RuleEngine{
		queries:   queries,
		Evaluator: eval,
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
// Validate runs all rules for a context and returns decision
func (re *RuleEngine) Validate(ctx context.Context, tenantID pgtype.UUID, req ValidationRequest) (*ValidationResult, error) {
	start := time.Now()

	// 1. Fetch active rules for this context (already sorted by priority)
	rules, err := re.queries.ListRulesByContext(ctx, db.ListRulesByContextParams{
		TenantID: tenantID,
		Context:  req.Context,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load rules: %w", err)
	}

	result := &ValidationResult{
		Decision:     "allow",
		Score:        0,
		RulesMatched: []string{},
	}
	var matchedIDs []pgtype.UUID

	// 2. Evaluate each rule in priority order
	for _, rule := range rules {
		// Build cache key
		tenantKey := pgUUIDToString(tenantID)
		ruleKey := pgUUIDToString(rule.ID)
		cacheKey := tenantKey + ":" + ruleKey

		// Get from cache or compile
		progIface, ok := re.programeCache.Load(cacheKey)
		var program cel.Program
		if !ok {
			prog, err := re.Evaluator.CompileRule(rule.Condition)
			if err != nil {
				fmt.Printf("Failed to compile rule %s: %v\n", rule.Name, err)
				continue
			}
			re.programeCache.Store(cacheKey, prog)
			program = prog
		} else {
			program = progIface.(cel.Program)
		}

		// Evaluate
		matched, err := re.Evaluator.Evaluate(program, req.Input)
		if err != nil {
			fmt.Printf("Failed to evaluate rule %s: %v\n", rule.Name, err)
			continue
		}

		if !matched {
			continue
		}

		// Rule matched! - add its ID and name
		matchedIDs = append(matchedIDs, rule.ID)
		result.RulesMatched = append(result.RulesMatched, rule.Name)

		// Apply action (unchanged)
		switch rule.Action {
		case "allow":
			result.Decision = "allow"
			result.Reason = fmt.Sprintf("Rule '%s' explicitly allowed", rule.Name)
			processingTime := int(time.Since(start).Milliseconds())
			inputBytes, _ := json.Marshal(req.Input)
			finalScore := int32(result.Score)
			_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
				TenantID:         tenantID,
				Context:          req.Context,
				Input:            inputBytes,
				MatchedRules:     matchedIDs,
				Decision:         result.Decision,
				FinalScore:       &finalScore,
				Reason:           &result.Reason,
				ProcessingTimeMs: int32(processingTime),
			})
			if err != nil {
				fmt.Printf("Failed to log decision: %v\n", err)
			}
			result.ProcessingTimeMs = processingTime
			return result, nil
		case "block":
			result.Decision = "block"
			result.Reason = fmt.Sprintf("Rule '%s' blocked: %s", rule.Name, rule.Condition)
			processingTime := int(time.Since(start).Milliseconds())
			inputBytes, _ := json.Marshal(req.Input)
			finalScore := int32(result.Score)
			_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
				TenantID:         tenantID,
				Context:          req.Context,
				Input:            inputBytes,
				MatchedRules:     matchedIDs,
				Decision:         result.Decision,
				FinalScore:       &finalScore,
				Reason:           &result.Reason,
				ProcessingTimeMs: int32(processingTime),
			})
			if err != nil {
				fmt.Printf("Failed to log decision: %v\n", err)
			}
			result.ProcessingTimeMs = processingTime
			return result, nil
		case "challenge":
			result.Decision = "challenge"
			result.Reason = fmt.Sprintf("Rule '%s' requires challenge", rule.Name)
			processingTime := int(time.Since(start).Milliseconds())
			inputBytes, _ := json.Marshal(req.Input)
			finalScore := int32(result.Score)
			_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
				TenantID:         tenantID,
				Context:          req.Context,
				Input:            inputBytes,
				MatchedRules:     matchedIDs,
				Decision:         result.Decision,
				FinalScore:       &finalScore,
				Reason:           &result.Reason,
				ProcessingTimeMs: int32(processingTime),
			})
			if err != nil {
				fmt.Printf("Failed to log decision: %v\n", err)
			}
			result.ProcessingTimeMs = processingTime
			return result, nil
		case "score":
			if rule.Score != nil {
				result.Score += int(*rule.Score)
			}
		case "flag":
			processingTime := int(time.Since(start).Milliseconds())
			inputBytes, _ := json.Marshal(req.Input)
			finalScore := int32(result.Score)
			_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
				TenantID:         tenantID,
				Context:          req.Context,
				Input:            inputBytes,
				MatchedRules:     matchedIDs,
				Decision:         result.Decision,
				FinalScore:       &finalScore,
				Reason:           &result.Reason,
				ProcessingTimeMs: int32(processingTime),
			})
			if err != nil {
				fmt.Printf("Failed to log decision: %v\n", err)
			}
			result.ProcessingTimeMs = processingTime
			return result, nil
		}
	}

	// 3. Score thresholds (unchanged)
	if result.Score >= 100 {
		result.Decision = "block"
		result.Reason = fmt.Sprintf("Cumulative score %d exceeded threshold", result.Score)
	} else if result.Score >= 50 && result.Decision == "allow" {
		result.Decision = "challenge"
		result.Reason = fmt.Sprintf("Cumulative score %d requires verification", result.Score)
	}

	// 4. Log decision (unchanged)
	processingTime := int(time.Since(start).Milliseconds())

	// ... (you still need to populate matchedRuleIDs properly)
	inputBytes, _ := json.Marshal(req.Input)
	finalScore := int32(result.Score)

	_, err = re.queries.CreateDecision(ctx, db.CreateDecisionParams{
		TenantID:         tenantID,
		Context:          req.Context,
		Input:            inputBytes,
		MatchedRules:     matchedIDs,
		Decision:         result.Decision,
		FinalScore:       &finalScore,
		Reason:           &result.Reason,
		ProcessingTimeMs: int32(processingTime),
	})
	if err != nil {
		fmt.Printf("Failed to log decision: %v\n", err)
	}

	result.ProcessingTimeMs = processingTime
	return result, nil
}

// InvalidateRule removes a specific rule from the cache.
func (re *RuleEngine) InvalidateRule(tenantID pgtype.UUID, ruleID pgtype.UUID) {
	key := pgUUIDToString(tenantID) + ":" + pgUUIDToString(ruleID)
	re.programeCache.Delete(key)
}

// StoreRule compiles and caches a rule. Returns error if compilation fails.
func (re *RuleEngine) StoreRule(tenantID pgtype.UUID, ruleID pgtype.UUID, condition string) error {
	prog, err := re.Evaluator.CompileRule(condition)
	if err != nil {
		return err
	}
	key := pgUUIDToString(tenantID) + ":" + pgUUIDToString(ruleID)
	re.programeCache.Store(key, prog)
	return nil
}
