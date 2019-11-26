package main

import (
	"github.com/sirupsen/logrus"

	"github.com/eturella/go-mysql-test/auth"
	"github.com/eturella/go-mysql-test/engine"
	"github.com/eturella/go-mysql-test/server"
)

var serverInstance server.Server

// Example of how to implement a MySQL server based on a Engine:
//
// ```
// > mysql --host=127.0.0.1 --port=5123 -u user -ppass db -e "SELECT * FROM mytable"
// +----------+-------------------+-------------------------------+---------------------+
// | name     | email             | phone_numbers                 | created_at          |
// +----------+-------------------+-------------------------------+---------------------+
// | John Doe | john@doe.com      | ["555-555-555"]               | 2018-04-18 09:41:13 |
// | John Doe | johnalt@doe.com   | []                            | 2018-04-18 09:41:13 |
// | Jane Doe | jane@doe.com      | []                            | 2018-04-18 09:41:13 |
// | Evil Bob | evilbob@gmail.com | ["555-666-555","666-666-666"] | 2018-04-18 09:41:13 |
// +----------+-------------------+-------------------------------+---------------------+
// ```
func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Avvio...")
	engine := engine.NewDefault()
	// logrus.Debug("Creazione DB...")
	// engine.AddDatabase(createTestDatabase())
	// logrus.Debug("Creazione CATALOG...")
	// engine.AddDatabase(sql.NewInformationSchemaDatabase(engine.Catalog))

	logrus.Debug("Definizione del server...")
	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("root", "", auth.AllPermissions),
	}

	logrus.Debug("Avvio del server...")
	serverInstance, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}
	serverInstance.Start()
}

func stop() {
	serverInstance.Close()
}

// func createTestDatabase() *memory.Database {
// 	const (
// 		dbName    = "mydb"
// 		tableName = "mytable"
// 	)

// 	db := memory.NewDatabase(dbName)
// 	table := memory.NewTable(tableName, sql.Schema{
// 		{Name: "name", Type: sql.Text, Nullable: false, Source: tableName},
// 		{Name: "email", Type: sql.Text, Nullable: false, Source: tableName},
// 		{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
// 		{Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
// 	})

// 	db.AddTable(tableName, table)
// 	ctx := sql.NewEmptyContext()
// 	table.Insert(ctx, sql.NewRow("Test1 Doe", "john@doe.com", []string{"555-555-555"}, time.Now()))
// 	table.Insert(ctx, sql.NewRow("Test2 Doe", "johnalt@doe.com", []string{}, time.Now()))
// 	table.Insert(ctx, sql.NewRow("Test3 Doe", "jane@doe.com", []string{}, time.Now()))
// 	table.Insert(ctx, sql.NewRow("Test4 Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()))
// 	return db
// }
