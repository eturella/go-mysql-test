package analyzer

import (
	"os"

	"github.com/eturella/go-mysql-test/sql"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// ExternalAnalyzerBuilder provides an easy way to generate Analyzer with custom rules and options.
type ExternalAnalyzerBuilder struct {
	catalog     *sql.Catalog
	debug       bool
	parallelism int
}

// NewExternalAnalyzerBuilder creates a new ExternalAnalyzerBuilder from a specific catalog.
// This builder allow us add custom Rules and modify some internal properties.
func NewExternalAnalyzerBuilder(c *sql.Catalog) *ExternalAnalyzerBuilder {
	return &ExternalAnalyzerBuilder{catalog: c}
}

// WithDebug activates debug on the Analyzer.
func (ab *ExternalAnalyzerBuilder) WithDebug() *ExternalAnalyzerBuilder {
	ab.debug = true
	return ab
}

// Build creates a new Analyzer using all previous data setted to the Builder
func (ab *ExternalAnalyzerBuilder) Build() *ExternalAnalyzer {
	_, debug := os.LookupEnv(debugAnalyzerKey)
	// var batches = []*Batch{}

	return &ExternalAnalyzer{
		Debug: debug || ab.debug,
		// Batches:     batches,
		Catalog:     ab.catalog,
		Parallelism: ab.parallelism,
	}
}

// ExternalAnalyzer analyzes nodes of the execution plan and applies rules and validations
// to them.
type ExternalAnalyzer struct {
	Debug       bool
	Parallelism int
	// Batches of Rules to apply.
	Batches []*Batch
	// Catalog of databases and registered functions.
	Catalog *sql.Catalog
}

// NewExternalDefault creates a default Analyzer instance with all default Rules and configuration.
// To add custom rules, the easiest way is use the Builder.
func NewExternalDefault(c *sql.Catalog) *ExternalAnalyzer {
	return NewExternalAnalyzerBuilder(c).Build()
}

// Log prints an INFO message to stdout with the given message and args
// if the analyzer is in debug mode.
func (a *ExternalAnalyzer) Log(msg string, args ...interface{}) {
	if a != nil && a.Debug {
		logrus.Infof(msg, args...)
	}
}

// Analyze the node and all its children.
func (a *ExternalAnalyzer) Analyze(ctx *sql.Context, n sql.Node) (sql.Node, error) {
	span, ctx := ctx.Span("analyze", opentracing.Tags{
		"plan": n.String(),
	})

	prev := n
	var err error
	a.Log("starting analysis of node of type: %T", n)

	prev, err = resolveExternalTables(ctx, a, n)

	defer func() {
		if prev != nil {
			span.SetTag("IsResolved", prev.Resolved())
		}
		span.Finish()
	}()

	return prev, err
}
