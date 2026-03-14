package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"general/internal/db"
	"general/internal/engine"
	"general/internal/middleware"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	connString := "postgres://asguard:devpassword@localhost:5433/general_engine?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)
	ruleEngine, err := engine.NewRuleEngine(queries)
	if err != nil {
		log.Fatalf("Failed to create rule engine: %v", err)
	}

	// Create router
	mux := http.NewServeMux()

	// Routes (same as before)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/v1/validate", validateHandler(ruleEngine))

	mux.HandleFunc("/v1/rules", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listRulesHandler(queries)(w, r)
		case http.MethodPost:
			createRuleHandler(queries, ruleEngine)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/rules/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/v1/rules/"):]
		if id == "" {
			http.Error(w, "Rule ID required", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getRuleHandler(queries, id)(w, r)
		case http.MethodPut:
			updateRuleHandler(queries, ruleEngine, id)(w, r)
		case http.MethodDelete:
			deleteRuleHandler(queries, ruleEngine, id)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// WRAP WITH MIDDLEWARE - This is the key line
	handler := middleware.AuthMiddleware(mux)

	port := "8083"
	log.Printf("General Validation Engine starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler)) // Use handler, not nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"engine": "general-validation",
	})
}

// tenantNamespace is a fixed UUID v5 namespace for deriving tenant UUIDs from strings.
// This lets tokens with non-UUID tenant IDs (e.g. "tenant-acme-corp") work seamlessly.
var tenantNamespace = uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8") // uuid.NameSpaceDNS

// tenantPgUUID converts any tenant ID string (UUID or slug) to a deterministic pgtype.UUID.
// If the string is already a valid UUID it is used directly; otherwise a UUID v5 is derived.
func tenantPgUUID(tenantIDStr string) pgtype.UUID {
	if id, err := uuid.Parse(tenantIDStr); err == nil {
		return pgtype.UUID{Bytes: id, Valid: true}
	}
	// Derive a deterministic UUID v5 from the slug so the same string always maps to the same UUID.
	return pgtype.UUID{Bytes: uuid.NewSHA1(tenantNamespace, []byte(tenantIDStr)), Valid: true}
}

func validateHandler(ruleEngine *engine.RuleEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req engine.ValidationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if req.Context == "" {
			http.Error(w, "Missing required field: context", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "missing tenat ID in token", http.StatusUnauthorized)
			return
		}
		tenantID := tenantPgUUID(tenantIDStr)

		result, err := ruleEngine.Validate(ctx, tenantID, req)
		if err != nil {
			log.Printf("Validation error: %v", err)
			http.Error(w, "Validation failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

// GET /v1/rules?context=optional_filter
func listRulesHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Extract tenant ID from JWT context (works with UUID strings and slugs alike)
		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "Missing tenant ID in token", http.StatusUnauthorized)
			return
		}
		tenantID := tenantPgUUID(tenantIDStr)

		// Optional context filter from query param
		contextFilter := r.URL.Query().Get("context")

		// ListAllRules fetches all rules for the tenant; filter by context in-memory if specified
		rules, err := queries.ListAllRules(ctx, tenantID)
		if err != nil {
			log.Printf("Failed to list rules: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Optional: filter by context in-memory if query param was provided
		var response []engine.RuleResponse
		for _, rule := range rules {
			if contextFilter == "" || rule.Context == contextFilter {
				response = append(response, engine.RuleToResponse(rule))
			}
		}
		if response == nil {
			response = []engine.RuleResponse{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// POST /v1/rules
func createRuleHandler(queries *db.Queries, ruleEngine *engine.RuleEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req engine.RuleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate required fields
		if req.Name == "" || req.Context == "" || req.Action == "" {
			http.Error(w, "Mising required fields", http.StatusBadRequest)
			return
		}

		validActions := map[string]bool{"allow": true, "block": true, "challenge": true, "flag": true, "score": true}
		if !validActions[req.Action] {
			http.Error(w, "Invalid action. Must be: allow, block, challenge, flag, score", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		// Extract tenant ID from JWT context (works with UUID strings and slugs alike)
		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "Missing tenant ID in token", http.StatusUnauthorized)
			return
		}

		tenantID := tenantPgUUID(tenantIDStr)

		// Convert score to database type (*int32)
		var score *int32
		if req.Score != nil {
			s := int32(*req.Score)
			score = &s
		}

		rule, err := queries.CreateRule(ctx, db.CreateRuleParams{
			TenantID:  tenantPgUUID(tenantIDStr),
			Name:      req.Name,
			Context:   req.Context,
			Condition: req.Condition,
			Action:    req.Action,
			Score:     score,
			Priority:  int32(req.Priority),
			Enabled:   req.Enabled,
		})
		if err != nil {
			log.Printf("Failed to create rule: %v", err)
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := ruleEngine.StoreRule(tenantID, rule.ID, req.Condition); err != nil {
			log.Printf("Warning: failed to cache new rule: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(engine.RuleToResponse(rule))
	}
}

// GET /v1/rules/{id}
func getRuleHandler(queries *db.Queries, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse rule UUID
		ruleID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid rule ID format", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		// Extract tenant ID from JWT context (works with UUID strings and slugs alike)
		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "Missing tenant ID in token", http.StatusUnauthorized)
			return
		}

		rule, err := queries.GetRule(ctx, db.GetRuleParams{
			ID:       pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID: tenantPgUUID(tenantIDStr),
		})
		if err != nil {
			http.Error(w, "Rule not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(engine.RuleToResponse(rule))
	}
}

// PUT /v1/rules/{id}
func updateRuleHandler(queries *db.Queries, ruleEngine *engine.RuleEngine, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ruleID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid rule ID format", http.StatusBadRequest)
			return
		}

		var req engine.RuleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if req.Name == "" || req.Context == "" || req.Condition == "" || req.Action == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "Missing tenant ID in token", http.StatusUnauthorized)
			return
		}
		tenantID := tenantPgUUID(tenantIDStr)

		var score *int32
		if req.Score != nil {
			s := int32(*req.Score)
			score = &s
		}

		// Perform update
		rule, err := queries.UpdateRule(ctx, db.UpdateRuleParams{
			ID:        pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID:  tenantID,
			Name:      req.Name,
			Context:   req.Context,
			Condition: req.Condition,
			Action:    req.Action,
			Score:     score,
			Priority:  int32(req.Priority),
			Enabled:   req.Enabled,
		})
		if err != nil {
			log.Printf("Failed to update rule: %v", err)
			http.Error(w, "Rule not found or database error", http.StatusInternalServerError)
			return
		}

		// Invalidate old cache entry (if any)
		ruleEngine.InvalidateRule(tenantID, pgtype.UUID{Bytes: ruleID, Valid: true})

		// Cache the new compiled program
		if err := ruleEngine.StoreRule(tenantID, rule.ID, req.Condition); err != nil {
			log.Printf("Warning: failed to cache updated rule %s: %v", rule.Name, err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(engine.RuleToResponse(rule))
	}
}

// DELETE /v1/rules/{id}
func deleteRuleHandler(queries *db.Queries, ruleEngine *engine.RuleEngine, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ruleID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid rule ID format", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		// Extract tenant ID from JWT context (works with UUID strings and slugs alike)
		tenantIDStr := middleware.GetTenantID(ctx)
		if tenantIDStr == "" {
			http.Error(w, "Missing tenant ID in token", http.StatusUnauthorized)
			return
		}

		err = queries.DeleteRule(ctx, db.DeleteRuleParams{
			ID:       pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID: tenantPgUUID(tenantIDStr),
		})
		if err != nil {
			http.Error(w, "Rule not found", http.StatusNotFound)
			return
		}

		// Remove from cache
		tenantID := tenantPgUUID(tenantIDStr)
		ruleEngine.InvalidateRule(tenantID, pgtype.UUID{Bytes: ruleID, Valid: true})

		w.WriteHeader(http.StatusNoContent) // 204 - success, no body
	}
}
