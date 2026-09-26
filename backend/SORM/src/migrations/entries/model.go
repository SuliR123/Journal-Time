package entries

import (
	"fmt"
	"strings"
)

type Model struct {
	tableName string
	fields    []IEntry
}

func NewModel(tableName string, fields []IEntry) *Model {
	return &Model{
		tableName: strings.ToLower(tableName),
		fields:    fields,
	}
}

func (m *Model) WriteToSQL() {

}

func (m *Model) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')
	for _, field := range m.fields {
		sb.WriteString(field.String())
		sb.WriteByte('\n')
	}
	return fmt.Sprintf("Table Name: %s\n Fields: %s", m.tableName, sb.String())
}
