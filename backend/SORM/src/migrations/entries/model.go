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

func (m *Model) WriteToSQLTable(table ISQLTable) error {
	table.AddTableName(m.tableName)
	for _, field := range m.fields {
		err := field.WriteToSQLTable(table)
		if err != nil {
			return fmt.Errorf("Unable to create SQL table for model %s, got error: %s", m.tableName, err.Error())
		}
	}
	return nil
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
