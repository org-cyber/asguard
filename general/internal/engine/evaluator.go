// internal/engine/evaluator.go
package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

// Evaluator handles CEL expression compilation and execution
type Evaluator struct {
	env *cel.Env
}

// NewEvaluator creates a CEL environment with our standard types
func NewEvaluator() (*Evaluator, error) {
	// Create CEL environment
	// We declare that 'input' will be a map with string keys and any values
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("input", decls.NewMapType(decls.String, decls.Dyn)),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL environment: %w", err)
	}

	return &Evaluator{env: env}, nil
}

// CompileRule parses a rule's condition into an executable program
// We cache this - don't re-parse on every request
func (e *Evaluator) CompileRule(condition string) (cel.Program, error) {
	// Parse the expression string
	ast, issues := e.env.Parse(condition)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("parse error: %w", issues.Err())
	}

	// Type-check (catches errors like "string > int")
	checked, issues := e.env.Check(ast)
	if issues != nil && issues.Err() != nil {
		return nil, fmt.Errorf("type check error: %w", issues.Err())
	}

	// Compile to bytecode
	program, err := e.env.Program(checked)
	if err != nil {
		return nil, fmt.Errorf("compile error: %w", err)
	}

	return program, nil
}

// Evaluate checks if input matches the compiled rule
func (e *Evaluator) Evaluate(program cel.Program, input map[string]interface{}) (bool, error) {
	// Add timeout to prevent infinite loops
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Run the program with our input data
	out, _, err := program.ContextEval(ctx, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return false, fmt.Errorf("evaluation error: %w", err)
	}

	// Convert CEL result to Go bool
	if boolVal, ok := out.Value().(bool); ok {
		return boolVal, nil
	}

	return false, fmt.Errorf("rule returned non-boolean: %T", out.Value())
}
