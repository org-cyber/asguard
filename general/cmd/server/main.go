package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"general/internal/db"
	"general/internal/engine"
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

	// Routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/v1/validate", validateHandler(ruleEngine))

	// Rule management routes
	http.HandleFunc("/v1/rules", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listRulesHandler(queries)(w, r)
		case http.MethodPost:
			createRuleHandler(queries)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/v1/rules/", func(w http.ResponseWriter, r *http.Request) {
		// Extract ID from path: /v1/rules/{id}
		id := r.URL.Path[len("/v1/rules/"):]
		if id == "" {
			http.Error(w, "Rule ID required", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getRuleHandler(queries, id)(w, r)
		case http.MethodPut:
			updateRuleHandler(queries, id)(w, r)
		case http.MethodDelete:
			deleteRuleHandler(queries, id)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := "8083"
	log.Printf("General Validation Engine starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"engine": "general-validation",
	})
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

		result, err := ruleEngine.Validate(ctx, req)
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
		zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

		// Optional context filter from query param
		contextFilter := r.URL.Query().Get("context")
		if contextFilter == "" {
			contextFilter = "" // Empty string means no filter in SQL
		}

		// ListAllRules fetches all rules for the tenant; filter by context in-memory if specified
		rules, err := queries.ListAllRules(ctx, zeroUUID)
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
func createRuleHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req engine.RuleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate required fields
		if req.Name == "" || req.Context == "" || req.Condition == "" || req.Action == "" {
			http.Error(w, "Missing required fields: name, context, condition, action", http.StatusBadRequest)
			return
		}

		// Validate action type
		validActions := map[string]bool{"allow": true, "block": true, "challenge": true, "flag": true, "score": true}
		if !validActions[req.Action] {
			http.Error(w, "Invalid action. Must be: allow, block, challenge, flag, score", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

		// Convert score to database type (*int32)
		var score *int32
		if req.Score != nil {
			s := int32(*req.Score)
			score = &s
		}

		rule, err := queries.CreateRule(ctx, db.CreateRuleParams{
			TenantID:  zeroUUID,
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(engine.RuleToResponse(rule))
	}
}

// GET /v1/rules/{id}
func getRuleHandler(queries *db.Queries, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse UUID
		ruleID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid rule ID format", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

		rule, err := queries.GetRule(ctx, db.GetRuleParams{
			ID:       pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID: zeroUUID,
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
func updateRuleHandler(queries *db.Queries, id string) http.HandlerFunc {
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

		// Validate
		if req.Name == "" || req.Context == "" || req.Condition == "" || req.Action == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

		var score *int32
		if req.Score != nil {
			s := int32(*req.Score)
			score = &s
		}

		// Note: Context is not updated (SQL UPDATE does not modify the context column)
		rule, err := queries.UpdateRule(ctx, db.UpdateRuleParams{
			ID:        pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID:  zeroUUID,
			Name:      req.Name,
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(engine.RuleToResponse(rule))
	}
}

// DELETE /v1/rules/{id}
func deleteRuleHandler(queries *db.Queries, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ruleID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, "Invalid rule ID format", http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		zeroUUID := pgtype.UUID{Bytes: [16]byte{}, Valid: true}

		err = queries.DeleteRule(ctx, db.DeleteRuleParams{
			ID:       pgtype.UUID{Bytes: ruleID, Valid: true},
			TenantID: zeroUUID,
		})
		if err != nil {
			http.Error(w, "Rule not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204 - success, no body
	}
}
