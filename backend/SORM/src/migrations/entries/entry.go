package entries

import "fmt"

type IEntry interface {
	WriteToSQLTable(table ISQLTable) error

	fmt.Stringer
}
