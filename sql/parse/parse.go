package parse

import (
	"fmt"
	"strings"

	pkg "github.com/bisegni/go-c-interface-test/query"

	"github.com/araddon/qlbridge/rel"

	"github.com/eturella/go-mysql-test/bisegniadapter"
	"github.com/eturella/go-mysql-test/sql"
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

	fmt.Println(" -------> 1 ")
	sqlRequest, err := rel.ParseSql(query)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	fmt.Printf("%+v \n", sqlRequest)
	// fmt.Printf("%s \n", json.MarshalIndent(sqlRequest, "", "    "))
	// fmt.Println(" -------> 2 ")
	switch sqlRequest.(type) {
	case *rel.SqlShow:
		// fmt.Println(" -------> 3 ")
		stmt := sqlRequest.(*rel.SqlShow)
		fmt.Printf("%+v \n", *stmt)
		showType := strings.ToLower(stmt.ShowType)
		fmt.Printf("%+v \n", showType)
		switch showType {
		case "databases":
			// {Raw:SHOW databases like 'utente@%' Db: Full:false Scope: S
			// howType:databases From: Identity: Create:false CreateWhat: Where:<nil> Like:<nil>}
			return bisegniadapter.NewShowDatabases(), nil
		case "tables":
			// {Raw:SHOW full tables from qualcosa like 'utente@%' Db:qualcosa Full:true Scope:
			// ShowType:tables From: Identity: Create:false CreateWhat: Where:<nil> Like:Table LIKE "utente@%"}
			ftm := pkg.NewFileTable("information_schema", "tables")
			return bisegniadapter.NewExternalTable("tables", ftm)
		case "variables":
			// {Raw:SHOW variables Db: Full:false Scope: ShowType:variables From: Identity: Create:false CreateWhat: Where:<nil> Like:<nil>}
			config := ctx.Session.GetAll()
			return bisegniadapter.NewShowVariables(config, ""), nil
		}
	case *rel.SqlSelect:
		// fmt.Println(" -------> 4 ")
		if strings.HasSuffix(strings.ToLower(query), "select @@global.max_allowed_packet") {
			// m := ctx.Session.GetAll()
			m := map[string]sql.TypedValue{
				"@@global.max_allowed_packet": {Typ: sql.Int64, Value: "4194304"}, // 16777216
			}
			// r := bisegniadapter.NewShowVariables(m, query)
			r := bisegniadapter.NewSelectVariables(m, query)
			return r, nil
		}
	}

	ftm := pkg.NewFileTable("data", "test")
	return bisegniadapter.NewExternalTable("test", ftm)
}
