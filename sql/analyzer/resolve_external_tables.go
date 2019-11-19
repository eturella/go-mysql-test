package analyzer

import (
	"github.com/eturella/go-mysql-test/sql"
	"github.com/eturella/go-mysql-test/sql/plan"

	extdb "github.com/bisegni/go-c-interface-test/pkg"
)

// const dualTableName = "dual"

// var dualTable = func() sql.Table {
// 	t := memory.NewTable(dualTableName, sql.Schema{
// 		{Name: "dummy", Source: dualTableName, Type: sql.Text, Nullable: false},
// 	})
// 	_ = t.Insert(sql.NewEmptyContext(), sql.NewRow("x"))
// 	return t
// }()

func resolveExternalTables(ctx *sql.Context, a *ExternalAnalyzer, n sql.Node) (sql.Node, error) {
	span, _ := ctx.Span("dynamic_tables_generation")
	defer span.Finish()

	a.Log("resolve table, node of type: %T", n)
	return plan.TransformUp(n, func(n sql.Node) (sql.Node, error) {
		a.Log("transforming node of type: %T", n)
		

		qe, excErr := extdb.Query();
		if excErr != nil {
			return nil, excErr
		}
		
		name := qe.uuid()
		db := "external"

		rt, err := a.Catalog.Table(db, name)
		if err != nil {
			if sql.ErrTableNotFound.Is(err) && name == dualTableName {
				rt = dualTable
				name = dualTableName
			} else {
				return nil, err
			}
		}

		a.Log("table resolved: %q", t.Name())

		return plan.NewResolvedTable(rt), nil
	})
}
