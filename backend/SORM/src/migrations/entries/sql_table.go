package entries

import (
	"fmt"
	"strings"
)

func NewSQLTable() ISQLTable {
	return &SQLTable{
		fields:      [][]string{},
		primaryKeys: []string{},
		foreignKeys: []string{},
	}
}

type ISQLTable interface {
	// Set the name of this table
	AddTableName(tableName string)

	// Add a column with the given name and postgres type
	AddColumn(columnName string, columnType string) error

	// Make the given column a primary key, the order in which columns are added as primary keys determine which column is the search index
	AddPrimaryKey(columnName string) error

	// Adds a foreign key constraint for the given column name. Takes in the reference table/column and onUpdate/onDelete action
	AddForeignKey(columnName string, table string, referenceColumn string, onUpdate string, onDelete string) error

	// Adds a default value for the given column
	AddDefault(columnName string, defaultValue string) error

	// Specifies that the given column cannot be null
	AddNotNull(columnName string) error

	// Specifies that the given column can be null
	AddNull(columnName string) error

	// Specifies that the given column must be unique in the table
	AddUnique(columnName string) error

	// Writes SQL definition for the table with the stored table name and fields with their associated constraints
	WriteToSQL() string
}

/*
Represents an migration entry for one table in a Postgres database
Contains the functionality to add the name of the table, create columns, & add constraints to the columns
*/
type SQLTable struct {
	tableName   string
	fields      [][]string
	primaryKeys []string
	foreignKeys []string
}

func (s *SQLTable) AddTableName(tableName string) {
	s.tableName = tableName
}

func (s *SQLTable) AddColumn(columnName string, columnType string) error {
	_, _, ok := s.getColumn(columnName)
	if ok {
		return fmt.Errorf("Cannot add column, column with the same name `%s` already exists in the table", columnName)
	}
	s.fields = append(s.fields, []string{columnName, columnType})
	return nil
}

func (s *SQLTable) AddPrimaryKey(columnName string) error {
	_, _, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot add primary key to field %s because the field does not exist in the table", columnName)
	}
	s.primaryKeys = append(s.primaryKeys, fmt.Sprintf(`"%s"`, columnName))
	return nil
}

func (s *SQLTable) AddForeignKey(columnName string, table string, referenceColumn string, onUpdate string, onDelete string) error {
	_, _, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot create foreign key for column %s because the column does not exist within this table", columnName)
	}
	// Format: CONSTRAINT "role_permissions_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
	fkeyName := columnName + "->" + table + "_" + referenceColumn + "_fkey"
	fullConstraint := fmt.Sprintf(`CONSTRAINT "%s" FOREIGN KEY ("%s") REFERENCES "%s" ("%s") ON UPDATE %s ON DELETE %s`, fkeyName, columnName, table, referenceColumn, strings.ToUpper(onUpdate), strings.ToUpper(onDelete))
	s.foreignKeys = append(s.foreignKeys, fullConstraint)
	return nil
}

func (s *SQLTable) AddDefault(columnName string, defaultValue string) error {
	values, index, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot create default value for the given column %s because it doesn't exist in the table", columnName)
	}
	typeStr := values[1]
	formattedDefault := defaultValue
	if typeStr == "text" || strings.Contains(typeStr, "varchar") || typeStr == "char" {
		formattedDefault = fmt.Sprintf(`'%s'`, defaultValue)
	}
	s.fields[index] = append(values, fmt.Sprintf("DEFAULT %s", formattedDefault))
	return nil
}

func (s *SQLTable) AddNotNull(columnName string) error {
	values, index, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot add not null constraint to field %s because it doesn't exist in the table", columnName)
	}
	s.fields[index] = append(values, "NOT NULL")
	return nil
}

func (s *SQLTable) AddNull(columnName string) error {
	values, index, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot add not null constraint to field %s because it doesn't exist in the table", columnName)
	}
	s.fields[index] = append(values, "NULL")
	return nil
}

func (s *SQLTable) AddUnique(columnName string) error {
	values, index, ok := s.getColumn(columnName)
	if !ok {
		return fmt.Errorf("Cannot add not null constraint to field %s because it doesn't exist in the table", columnName)
	}
	s.fields[index] = append(values, "UNIQUE")
	return nil
}

func (s *SQLTable) WriteToSQL() string {
	tableHeader := fmt.Sprintf(`CREATE TABLE "%s"`, s.tableName)
	fields := make([]string, 0)
	for _, values := range s.fields {
		values[0] = fmt.Sprintf(`"%s"`, values[0])
		fieldValues := strings.Join(values, " ")
		fields = append(fields, fieldValues)
	}
	if len(s.primaryKeys) > 0 {
		primaryKeysFormatted := fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(s.primaryKeys, ", "))
		fields = append(fields, primaryKeysFormatted)
	}
	if len(s.foreignKeys) > 0 {
		foreignKeysFormatted := strings.Join(s.foreignKeys, ", ")
		fields = append(fields, foreignKeysFormatted)
	}
	fieldStrings := strings.Join(fields, ", ")
	return fmt.Sprintf("%s (%s);", tableHeader, fieldStrings)
}

func (s *SQLTable) getColumn(columnName string) (fieldSlice []string, index int, contains bool) {
	for i, fieldSlice := range s.fields {
		if fieldSlice[0] == columnName {
			return fieldSlice, i, true
		}
	}
	return nil, -1, false
}
