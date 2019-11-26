package analyzer

// import (
// 	"os"

// 	"github.com/eturella/go-mysql-test/sql"
// 	opentracing "github.com/opentracing/opentracing-go"
// 	"github.com/sirupsen/logrus"
// 	"gopkg.in/src-d/go-errors.v1"
// )

// const debugAnalyzerKey = "DEBUG_ANALYZER"

// // const maxAnalysisIterations = 1000

// // ErrMaxAnalysisIters is thrown when the analysis iterations are exceeded
// var ErrMaxAnalysisIters = errors.NewKind("exceeded max analysis iterations (%d)")

// // Builder provides an easy way to generate Analyzer with custom rules and options.
// type Builder struct {
// 	// preAnalyzeRules     []Rule
// 	// postAnalyzeRules    []Rule
// 	// preValidationRules  []Rule
// 	// postValidationRules []Rule
// 	catalog     *sql.Catalog
// 	debug       bool
// 	parallelism int
// }

// // NewBuilder creates a new Builder from a specific catalog.
// // This builder allow us add custom Rules and modify some internal properties.
// func NewBuilder(c *sql.Catalog) *Builder {
// 	return &Builder{catalog: c}
// }

// // WithDebug activates debug on the Analyzer.
// func (ab *Builder) WithDebug() *Builder {
// 	ab.debug = true

// 	return ab
// }

// // Build creates a new Analyzer using all previous data setted to the Builder
// func (ab *Builder) Build() *Analyzer {
// 	_, debug := os.LookupEnv(debugAnalyzerKey)
// 	// var batches = []*Batch{}

// 	return &Analyzer{
// 		Debug: debug || ab.debug,
// 		// Batches:     batches,
// 		Catalog:     ab.catalog,
// 		Parallelism: ab.parallelism,
// 	}
// }

// Analyzer analyzes nodes of the execution plan and applies rules and validations
// to them.
type Analyzer struct {
	Debug       bool
	Parallelism int
	// // Batches of Rules to apply.
	// Batches []*Batch
	// // Catalog of databases and registered functions.
	// Catalog *sql.Catalog
}

// // NewDefault creates a default Analyzer instance with all default Rules and configuration.
// // To add custom rules, the easiest way is use the Builder.
// func NewDefault(c *sql.Catalog) *Analyzer {
// 	return NewBuilder(c).Build()
// }

// // Log prints an INFO message to stdout with the given message and args
// // if the analyzer is in debug mode.
// func (a *Analyzer) Log(msg string, args ...interface{}) {
// 	if a != nil && a.Debug {
// 		logrus.Infof(msg, args...)
// 	}
// }

// // Analyze the node and all its children.
// func (a *Analyzer) Analyze(ctx *sql.Context, n sql.Node) (sql.Node, error) {
// 	span, ctx := ctx.Span("analyze", opentracing.Tags{
// 		"plan": n.String(),
// 	})

// 	prev := n
// 	var err error
// 	a.Log("starting analysis of node of type: %T", n)

// 	prev, err = resolveTables(ctx, a, n)

// 	defer func() {
// 		if prev != nil {
// 			span.SetTag("IsResolved", prev.Resolved())
// 		}
// 		span.Finish()
// 	}()

// 	return prev, err
// }
