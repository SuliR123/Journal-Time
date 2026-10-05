package entries

import (
	"fmt"
	"strings"
	"testing"
)

var expectedField1 *Field = &Field{
	name:    "id",
	typeStr: "uuid",
	tags:    []string{"primary key"},
	mappedTags: map[string]string{
		"default": "gen_random_uuid()",
	},
}

var expectedField2 *Field = &Field{
	name:       "fk",
	typeStr:    "uuid",
	tags:       []string{"not null"},
	mappedTags: map[string]string{},
}

var fieldWithMultipleMappedTags *Field = &Field{
	name:    "bruh",
	typeStr: "text",
	tags:    []string{"primary key", "not null", "unique"},
	mappedTags: map[string]string{
		"default":     "balls",
		"foreign key": "bruh.id",
	},
}

// field example with onUpdate/onDelete but no foreign key

// on update/delete with illegal action

func TestNewField(t *testing.T) {
	tag1 := `json:"id" sorm:"primary key;type:uuid;default:gen_random_uuid()"`
	actualField1, err1 := NewField(tag1)
	if err1 != nil {
		t.Fatalf("Could not create field with the following string %s, recieved this error: %s", tag1, err1)
	}
	testEquality(actualField1, expectedField1, t)
	tag2 := `json:"fk" sorm:"type:uuid;not null"`
	actualField2, err2 := NewField(tag2)
	if err2 != nil {
		t.Fatalf("Could not create field with the following string %s, recieved this error: %s", tag2, err2)
	}
	testEquality(actualField2, expectedField2, t)
}

func TestParseTagErrorCases(t *testing.T) {
	noColon := `json"value"`
	_, errColon := NewField(noColon)
	if errColon == nil {
		t.Fatalf("Expected field creation to error out due to missing colon but it didn't :(")
	}
	if errColon.Error() != fmt.Sprintf("Tag is formatted incorrectly, `:` is required in the tag: %s", noColon) {
		t.Fatalf("Recieved the incorrect error when given tag with no colon: %s", errColon)
	}

	noQuotes := "json:value"
	_, errNoQuotes := NewField(noQuotes)
	if errNoQuotes == nil {
		t.Fatalf("Expected field creation to error out due to missing quotes but it didn't :(")
	}
	if errNoQuotes.Error() != fmt.Sprintf(`Tag is formatted incorrectly, expected "" around the tag value (ex: json:"id"): %s`, noQuotes) {
		t.Fatalf("Recieved the incorrect error when given tag with no quotes: %s", errNoQuotes)
	}
}

func TestSORMFormattingRegex(t *testing.T) {
	pass := validationRegex.MatchString("foreign key:balls.ball;")
	if pass {
		t.Fatalf("Trailing ; should fail...")
	}

	pass = validationRegex.MatchString("")
	if pass {
		t.Fatalf("Blank should fail")
	}

	pass = validationRegex.MatchString("primary key,foreign key:balls.ball")
	if pass {
		t.Fatalf("Should fail, sorm tags can only be seperated by ;")
	}

	pass = validationRegex.MatchString("primary key1")
	if pass {
		t.Fatalf("Should fail, no constraint contains a number")
	}

	pass = validationRegex.MatchString("primary key")
	if !pass {
		t.Fatalf("Should not fail, allowed to have one single constraint")
	}

	pass = validationRegex.MatchString("not null;type:uuid")
	if !pass {
		t.Fatalf("Should not fail, allowed to have multiple constraints joined by ;")
	}

	pass = validationRegex.MatchString("primary key foreign key:balls.balls")
	if pass {
		t.Fatalf("Constraints seperated by space should fail")
	}

	pass = validationRegex.MatchString("default ;foreign key:balls.balls")
	if pass {
		t.Fatalf("Cannot have space before ;")
	}

	pass = validationRegex.MatchString("default;unique;primary key")
	if !pass {
		t.Fatalf("Okay this one would just be so wrong")
	}

	pass = validationRegex.MatchString(";default")
	if pass {
		t.Fatalf("Cannot start sorm tag value with ;")
	}

	pass = validationRegex.MatchString("type:text;on update:no action")
	if !pass {
		t.Fatalf("Regex failed for possible value type:text;on update:no action")
	}
}

func TestValidMappedTags(t *testing.T) {
	err := validateMappedTagValue("type", "balls")
	if err == nil {
		t.Fatalf("Balls isn't a Postgres type...")
	}

	err = validateMappedTagValue("type", "text")
	if err != nil {
		t.Fatalf("Was unable to validate the type 'text'... somethings broken: %s", err)
	}

	err = validateMappedTagValue("type", "varchar(50)")
	if err != nil {
		t.Fatalf("Aw man the varchar thing didn't work... %s", err)
	}

	err = validateMappedTagValue("foreign key", "balls.ball")
	if err != nil {
		t.Fatalf("Something is broken with the foreign key validation... %s", err)
	}

	err = validateMappedTagValue("foreign key", "balls")
	if err == nil {
		t.Fatalf("Okay this foreign key thing should've broke...")
	}

	err = validateMappedTagValue("foreign key", ".balls")
	if err == nil {
		t.Fatalf("Must have a table before the period for the foreign key...")
	}

	err = validateMappedTagValue("foreign key", "ball.")
	if err == nil {
		t.Fatalf("Must have a column reference after the period for the foreign key...")
	}

	err = validateMappedTagValue("foreign key", "")
	if err == nil {
		t.Fatalf("Was passed no string, should break")
	}

	err = validateMappedTagValue("on update", "stinky balls?")
	if err == nil {
		t.Fatalf("This is not a possible action and should explode")
	}

	err = validateMappedTagValue("on delete", "stinky balls?")
	if err == nil {
		t.Fatalf("This is not a possible action and should explode")
	}

	err = validateMappedTagValue("on update", "cascade")
	if err != nil {
		t.Fatalf("This is a possible action and shouldn't explode... %s", err)
	}

	err = validateMappedTagValue("on delete", "no action")
	if err != nil {
		t.Fatalf("This is a possible action and shouldn't explode...: %s", err)
	}

	err = validateMappedTagValue("default", "balls")
	if err != nil {
		t.Fatalf("Default should return nil")
	}

	err = validateMappedTagValue("balls", "balls")
	if err.Error() != "The given constraint balls is not a processable mapped constraint, check the formatting of the constraint in the sorm tag" {
		t.Fatalf("Recieved the wrong error message for impossible constraint: %s", err)
	}
}

func Test_ValidateTagValue(t *testing.T) {
	err := validateTagValue("primary key", constraints)
	if err != nil {
		t.Fatalf("primary key is a valid option")
	}
	err = validateTagValue("not null", constraints)
	if err != nil {
		t.Fatalf("not null is also an option")
	}
	err = validateTagValue("balls", constraints)
	if err == nil {
		t.Fatalf("Should throw an error cuz balls ain't an option")
	}
}

func Test_ParseTagErrors(t *testing.T) {
	_, _, _, err := parseSormTags("primary key;foreign key:balls.id")
	if err.Error() != "Cannot create the field with the given sorm value: primary key;foreign key:balls.id. The field is missing the type value" {
		t.Fatalf("Missing sorm type did not throw error")
	}

	_, _, _, err = parseSormTags("type:text;on update:no action")
	if err == nil || err.Error() != `Cannot contain actions "on update" or "on delete" if a foreign key constraint is not present. Given SORM tag value: type:text;on update:no action` {
		t.Fatalf("on update error handling fails when on update is present but foreign key is not, %s", err.Error())
	}
}

func TestWriteToSQL_IncludeDefaultNull(t *testing.T) {
	table := NewSQLTable()
	table.AddTableName("balls")

	err := expectedField1.WriteToSQLTable(table)
	if err != nil {
		t.Fatalf("Could not create entry for field, got error: %s", err.Error())
	}

	expectedTable := `CREATE TABLE "balls" ("id" uuid NULL DEFAULT gen_random_uuid(), PRIMARY KEY ("id"));`
	actualTable := table.WriteToSQL()

	if expectedTable != actualTable {
		t.Fatalf("Actual table doesn't match expected table: \nExpected: %s, \nActual:   %s", expectedTable, actualTable)
	}
}

func TestWriteToSQL_MultipleTags_OnUpdateDefaultsToNoAction(t *testing.T) {
	table := NewSQLTable()
	table.AddTableName("balls")

	err := fieldWithMultipleMappedTags.WriteToSQLTable(table)
	if err != nil {
		t.Fatalf("Could not create entry for field, got error: %s", err.Error())
	}
	actual := table.WriteToSQL()
	foreignKey := `CONSTRAINT "bruh->bruh_id_fkey" FOREIGN KEY ("bruh") REFERENCES "bruh" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION`
	testTableStringContains(actual, t, `PRIMARY KEY ("bruh")`, `NOT NULL`, `UNIQUE`, `DEFAULT 'balls'`, foreignKey)
}

func testTableStringContains(sql string, t *testing.T, constraints ...string) {
	missing := []string{}
	for _, constraint := range constraints {
		if !strings.Contains(sql, constraint) {
			missing = append(missing, constraint)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("Expected all of the given constraints to be present, missing: %s", missing)
	}
}

// Takes in the actual field and the expected field and checks if they are equal
func testEquality(f1 *Field, f2 *Field, t *testing.T) {
	if f1.name != f2.name {
		t.Fatalf("Exepcted field names to match but got: Field1.name: %s != Field2.name: %s", f1.name, f2.name)
	}

	if f1.typeStr != f2.typeStr {
		t.Fatalf("Exepcted field names to match but got: Field1.typeStr: %s != Field2.typeStr: %s", f1.typeStr, f2.typeStr)
	}

	if len(f1.tags) != len(f2.tags) {
		t.Fatalf("Expected tag list to be the same length but got: Field1.tags.size: %d != Field2.tags.size: %d", len(f1.tags), len(f2.tags))
	}

	for i, tag1 := range f1.tags {
		tag2 := f2.tags[i]
		if tag1 != tag2 {
			t.Fatalf("Expected tag list to contain the same elements but got: Field1.tags[%d]: %s != Field2.tags[%d]: %s", i, tag1, i, tag2)
		}
	}

	if len(f1.mappedTags) != len(f2.mappedTags) {
		t.Fatalf("Expected tag list to be the same length but got: Field1.tags.size: %d != Field2.tags.size: %d", len(f1.mappedTags), len(f2.mappedTags))
	}

	for key1, value1 := range f1.mappedTags {
		value2, ok := f2.mappedTags[key1]
		if !ok {
			t.Fatalf("Expected Field2 mapping to contain the following key present in Field1: %s", key1)
		}
		if value1 != value2 {
			t.Fatalf("Expected tag values for key: %s to match but got: Field1.mappedTags[%s]: %s != Field2.mappedTags[%s]: %s", key1, key1, value1, key1, value2)
		}
	}
}
