package entries

import "testing"

var expectedTable1 string = `CREATE TABLE "roles" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "name" varchar(50) NOT NULL, PRIMARY KEY ("id"));`

func TestSQLTable(t *testing.T) {
	table := NewSQLTable()

	table.AddTableName("roles")
	table.AddColumn("id", "uuid")
	table.AddNotNull("id")
	table.AddDefault("id", "gen_random_uuid()")

	table.AddColumn("created_at", "timestamptz")
	table.AddNull("created_at")

	table.AddColumn("updated_at", "timestamptz")
	table.AddNull("updated_at")

	table.AddColumn("deleted_at", "timestamptz")
	table.AddNull("deleted_at")

	table.AddColumn("name", "varchar(50)")
	table.AddNotNull("name")

	table.AddPrimaryKey("id")

	actualOutput := table.WriteToSQL()
	if expectedTable1 != actualOutput {
		t.Fatalf("Generated table does not match expected table: \nExpected: %s \nActual:   %s", expectedTable1, actualOutput)
	}
}

func TestSQLTable_WithDefaultStringType(t *testing.T) {
	table := NewSQLTable()

	table.AddTableName("bruh")

	table.AddColumn("bruh", "varchar(50)")
	table.AddDefault("bruh", "bruh")

	expected := `CREATE TABLE "bruh" ("bruh" varchar(50) DEFAULT 'bruh');`
	actual := table.WriteToSQL()
	if expected != actual {
		t.Fatalf("Generated table does not match expected table: \nExpected: %s \nActual:   %s", expected, actual)
	}
}
