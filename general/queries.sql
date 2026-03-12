-- engines/general/queries.sql

-- name: CreateRule :one
-- Create a new validation rule. Returns the created rule with generated ID.
INSERT INTO rules (
    tenant_id, name, context, condition, action, score, priority, enabled
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetRule :one
-- Fetch a single rule by ID.
SELECT * FROM rules 
WHERE id = $1 AND tenant_id = $2;

-- name: ListRulesByContext :many
-- Get all active rules for a specific context, ordered by priority (highest first).
-- This is the main query used during validation.
SELECT * FROM rules 
WHERE tenant_id = $1 
  AND context = $2 
  AND enabled = true 
ORDER BY priority DESC;


-- name: DeleteRule :exec
-- Hard delete a rule.
DELETE FROM rules 
WHERE id = $1 AND tenant_id = $2;

-- name: ListAllRules :many
-- Admin query: Get all rules (including disabled) for a tenant.
SELECT * FROM rules 
WHERE tenant_id = $1 
ORDER BY context, priority DESC;

-- name: CreateDecision :one
-- Log a validation decision for audit trail.
INSERT INTO decisions (
    tenant_id, context, input, matched_rules, decision, final_score, reason, processing_time_ms
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: ListDecisions :many
-- Query recent decisions for debugging/analytics.
SELECT * FROM decisions 
WHERE tenant_id = $1 AND context = $2
ORDER BY processed_at DESC
LIMIT $3 OFFSET $4;

-- name: ListRules :many
SELECT * FROM rules 
WHERE tenant_id = $1 
  AND ($2::text IS NULL OR context = $2)
ORDER BY context, priority DESC;

-- name: UpdateRule :one
UPDATE rules 
SET 
    name = $3,
    context = $4,
    condition = $5,
    action = $6,
    score = $7,
    priority = $8,
    enabled = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND tenant_id = $2
RETURNING *;