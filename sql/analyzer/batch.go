package analyzer

import (
	"github.com/eturella/go-mysql-test/sql"
)

// RuleFunc is the function to be applied in a rule.
type RuleFunc func(*sql.Context, *Analyzer, sql.Node) (sql.Node, error)

// Rule to transform nodes.
type Rule struct {
	// Name of the rule.
	Name string
	// Apply transforms a node.
	Apply RuleFunc
}

// Batch executes a set of rules a specific number of times.
// When this number of times is reached, the actual node
// and ErrMaxAnalysisIters is returned.
type Batch struct {
	Desc       string
	Iterations int
	Rules      []Rule
}
