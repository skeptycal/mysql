package mysql

import (
	"fmt"
	"reflect"
)

// Notes
/* https://github.com/arnehormann/sqlinternals
   https://godoc.org/github.com/arnehormann/sqlinternals/mysqlinternals
    https://github.com/go-sql-driver/mysql/issues/585
    https://go-review.googlesource.com/c/go/+/29961

*/

var columnPointers []interface{}

func fake() {
	columnPointers := make([]interface{}, columnCount)

	for rows.Next() && previewRowCount < utils.PREVIEW_ROW_COUNT {
		if err = rows.Scan(columnPointers...); err != nil {
			// error handling
		}
		for i := 0; i < columnCount; i++ {
			columnValue := *columnPointers[i]
			// switch value := (*columnPointers[i]).(type)
			switch value := (*columnPointers[i]).(type) {
			case string:
				fmt.Println(value)
				fmt.Println("string")
			case int64:
				fmt.Println(value)
				fmt.Println("int64")
			case int:
				fmt.Println(value)
				fmt.Println("int")
			default:
				fmt.Println(value)
				fmt.Println(reflect.TypeOf(value))
			}
			val := parseValue(columnPointers[i].(*interface{}), "")
			preview.Columns[i].Data = append(preview.Columns[i].Data, val)
		}
	}
}
