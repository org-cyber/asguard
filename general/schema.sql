-- engines/general/schema.sql

-- Enable UUID extension (PostgreSQL native)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Rules: User-defined validation logic
CREATE TABLE rules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    -- For future multi-tenancy. Single tenant uses default '00000000-0000-0000-0000-000000000000'
    tenant_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    
    -- Human-readable name. Must be unique per tenant.
    name TEXT NOT NULL,
    
    -- What scenario? Examples: 'user_login', 'api_request', 'content_submit'
    context TEXT NOT NULL,
    
    -- The CEL expression. Example: 'input.failed_attempts > 5'
    condition TEXT NOT NULL,
    
    -- What to do when condition matches: 'allow', 'block', 'challenge', 'flag', 'score'
    action TEXT NOT NULL CHECK (action IN ('allow', 'block', 'challenge', 'flag', 'score')),
    
    -- For 'score' action: how many points to add. NULL for other actions.
    score INTEGER,
    
    -- Evaluation order. Higher number = evaluated first. Default 0.
    priority INTEGER NOT NULL DEFAULT 0,
    
    -- Soft delete. false = rule is active.
    enabled BOOLEAN NOT NULL DEFAULT true,
    
    -- When created/modified
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Prevent duplicate names within a tenant
    UNIQUE(tenant_id, name)
);

-- Fast lookup: Find active rules for a context, highest priority first
CREATE INDEX idx_rules_lookup ON rules(tenant_id, context, enabled, priority DESC);

-- Decisions: Audit log of every validation request
CREATE TABLE decisions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    tenant_id UUID NOT NULL,
    
    -- What was being validated? Same as rules.context
    context TEXT NOT NULL,
    
    -- Input data as JSONB. Flexible, indexable, fast.
    input JSONB NOT NULL,
    
    -- Which rules matched? Array of rule IDs.
    matched_rules UUID[],
    
    -- Final result: 'allow', 'block', 'challenge'
    decision TEXT NOT NULL CHECK (decision IN ('allow', 'block', 'challenge')),
    
    -- Total score if using cumulative scoring
    final_score INTEGER,
    
    -- Human-readable explanation
    reason TEXT,
    
    -- Performance tracking
    processed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    processing_time_ms INTEGER NOT NULL,
    
    -- For debugging
    engine_version TEXT NOT NULL DEFAULT '1.0.0'
);

-- Query recent decisions fast (for dashboard/debugging)
CREATE INDEX idx_decisions_recent ON decisions(tenant_id, context, processed_at DESC);

-- Index JSONB for searching specific input patterns (advanced feature)
CREATE INDEX idx_decisions_input ON decisions USING gin(input);