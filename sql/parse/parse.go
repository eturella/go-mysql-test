package parse

import (
	"fmt"
	"strings"

	"github.com/araddon/qlbridge/rel"

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
		// return plan.Nothing, nil
		panic("")
	}

	// l := lex.NewSqlLexer(query)
	// tok := l.NextToken()
	// for tok.T != lex.TokenEOF {
	// 	fmt.Printf("got:%v  \n", tok)
	// 	tok = l.NextToken()
	// }

	sqlRequest, err := rel.ParseSql(query)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	fmt.Printf("%v \n", sqlRequest)
	// fmt.Printf("%s \n", json.MarshalIndent(sqlRequest, "", "    "))
	switch sqlRequest.(type) {
	case *rel.SqlShow:
		stmt := sqlRequest.(*rel.SqlShow)
		fmt.Printf("%+v \n", *stmt)
		showType := strings.ToLower(stmt.ShowType)
		fmt.Printf("%+v \n", showType)
		// {Raw:SHOW full tables from qualcosa like 'utente@%' Db:qualcosa Full:true Scope:
		// ShowType:tables From: Identity: Create:false CreateWhat: Where:<nil> Like:Table LIKE "utente@%"}

		// {Raw:SHOW databases like 'utente@%' Db: Full:false Scope: S
		// howType:databases From: Identity: Create:false CreateWhat: Where:<nil> Like:<nil>}
	}

	return plan.NewUnresolvedTable("mytable", "mydb"), nil
}
