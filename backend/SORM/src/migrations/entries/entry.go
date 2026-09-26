package entries

import "fmt"

type IEntry interface {
	WriteToSQL()

	fmt.Stringer
}
