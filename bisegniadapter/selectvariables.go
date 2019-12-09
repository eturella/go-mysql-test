package bisegniadapter

import (
	"github.com/eturella/go-mysql-test/sql"
)

// SelectVariables is a node that shows the global and session variables
type SelectVariables struct {
	config map[string]sql.TypedValue
	query  string
}

// NewSelectVariables returns a new SelectVariables reference.
// config is a variables lookup table
// like is a "like pattern". If like is an empty string it will return all variables.
func NewSelectVariables(config map[string]sql.TypedValue, q string) *SelectVariables {
	return &SelectVariables{
		config: config,
		query:  q,
	}
}

// Resolved implements sql.Node interface. The function always returns true.
func (sv *SelectVariables) Resolved() bool {
	return true
}

// WithChildren implements the Node interface.
func (sv *SelectVariables) WithChildren(children ...sql.Node) (sql.Node, error) {
	if len(children) != 0 {
		return nil, sql.ErrInvalidChildrenNumber.New(sv, len(children), 0)
	}

	return sv, nil
}

// String implements the Stringer interface.
func (sv *SelectVariables) String() string {
	return sv.query
}

// Schema returns a new Schema reference for "SHOW VARIABLES" query.
func (sv *SelectVariables) Schema() sql.Schema {
	s := sql.Schema{}
	for k, v := range sv.config {
		s = append(s, &sql.Column{Name: k, Type: v.Typ, Nullable: true})
	}
	return s
}

// Children implements sql.Node interface. The function always returns nil.
func (*SelectVariables) Children() []sql.Node { return nil }

// RowIter implements the sql.Node interface.
// The function returns an iterator for filtered variables (based on like pattern)
func (sv *SelectVariables) RowIter(ctx *sql.Context) (sql.RowIter, error) {
	var (
		rows []sql.Row
	)

	for k, v := range sv.config {
		// if like != nil {
		// 	b, err := like.Eval(ctx, sql.NewRow(k, sv.pattern))
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	if !b.(bool) {
		// 		continue
		// 	}
		// }

		rows = append(rows, sql.NewRow(k, v.Value))
	}

	return sql.RowsToRowIter(rows...), nil
}
