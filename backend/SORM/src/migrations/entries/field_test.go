package entries

import (
	"fmt"
	"testing"
)

var expectedField1 *Field = &Field{
	name: "id",
	tags: []string{"primary key"},
	mappedTags: map[string]string{
		"type":    "uuid",
		"default": "gen_random_uuid()",
	},
}

var expectedField2 *Field = &Field{
	name: "fk",
	tags: []string{"not null"},
	mappedTags: map[string]string{
		"type": "uuid",
	},
}

func TestNewField(t *testing.T) {
	tag1 := `json:"id" sorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
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
		t.Fatalf("Recieved the incorrect error when given tag with no color %s", errColon)
	}

	noQuotes := "json:value"
	_, errNoQuotes := NewField(noQuotes)
	if errNoQuotes == nil {
		t.Fatalf("Expected field creation to error out due to missing quotes but it didn't :(")
	}
	if errNoQuotes.Error() != fmt.Sprintf(`Tag is formatted incorrectly, expected "" around the tag value (e.x json:"id"): %s`, noQuotes) {
		t.Fatalf("Recieved the incorrect error when given tag with no quotes %s", errNoQuotes)
	}
}

// Takes in the actual field and the expected field and checks if they are equal
func testEquality(f1 *Field, f2 *Field, t *testing.T) {
	if f1.name != f2.name {
		t.Fatalf("Exepcted field names to match but got: Field1.name: %s != Field2.name: %s", f1.name, f2.name)
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
