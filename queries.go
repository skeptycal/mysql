package mysql

import (
	"fmt"
	"strings"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

// these are not exact conversions, but should eliminate most confusion
type (
	UINT      uint32 // unsigned
	UTINYINT  uint8  //unsigned
	USMALLINT uint16 //unsigned
	// UMEDIUMIN = uint24 //unsigned
	UBIGINT uint64 // unsigned

	INT      int32 // signed
	TINYINT  int8  // signed
	SMALLINT int16 // signed
	// MEDIUMIN = int24 // signed
	BIGINT int64 // signed

	FLOAT   float32 // always signed
	DOUBLE  float64 // always signed
	DECIMAL float64 // always signed

	DATE string
)

func Values(value interface{}){
 switch value := columnPointers[i].(type) {
case string:
   // value has type string and contains the string value
   return
case int64:
   // value has type int64 and contains the value as int64
case float64:
  ...
}
}
func NewDATE(date string) DATE {
	d := new(DATE)
	t, _ := time.
	d = new(DATE(t.Format(layoutISO)))

	return d
}

func create(name string, args ...string) error {
	arglist := strings.Join(args, ", ")
	query := fmt.Sprintf(`CREATE TABLE %s(%s);`, name, arglist)

	Db.Exec(query)
	// id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), population INT
}

const (
	USE = `USE %s;`
	DIE = `DROP TABLE IF EXISTS %s;`
)
