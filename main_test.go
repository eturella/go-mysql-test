package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(t *testing.T) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM dual")
	if err != nil {
		t.Fatal(err.Error())
	}

	columns, _ := rows.Columns()
	t.Logf("%+v \n", columns)
	length := len(columns)
	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		t.Logf("*********************** \n")
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			panic(err)
		}
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			key := columns[i]
			val := *(current[i]).(*interface{})
			if val == nil {
				value[key] = nil
				continue
			}
			vType := reflect.TypeOf(val)
			switch vType.String() {
			case "int64":
				value[key] = val.(int64)
			case "string":
				value[key] = val.(string)
			case "time.Time":
				value[key] = val.(time.Time)
			case "[]uint8":
				value[key] = string(val.([]uint8))
			default:
				fmt.Printf("unsupport data type '%s' now\n", vType)
				// TODO remember add other data type
			}
			t.Logf("%s = %v \n", key, value[key])
		}
		results = append(results, value)
	}

}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
