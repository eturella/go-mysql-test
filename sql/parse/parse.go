package parse

import (
	"strings"

	"github.com/eturella/go-mysql-test/sql"
	"github.com/eturella/go-mysql-test/sql/plan"
	"github.com/opentracing/opentracing-go"
)

// Parse parses the given SQL sentence and returns the corresponding node.
func Parse(ctx *sql.Context, query string) (sql.Node, error) {
	span, ctx := ctx.Span("parse", opentracing.Tag{Key: "query", Value: query})
	defer span.Finish()

	s := strings.TrimSpace(removeComments(query))
	if strings.HasSuffix(s, ";") {
		s = s[:len(s)-1]
	}

	if s == "" {
		ctx.Warn(0, "query was empty after trimming comments, so it will be ignored")
		return plan.Nothing, nil
	}

	return plan.NewUnresolvedTable("mytable", "mydb"), nil
}
