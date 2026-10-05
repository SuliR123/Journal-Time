package entries

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var validationRegex = regexp.MustCompile(`^[a-z]+( [a-z]+)?(:[a-z_.()]+( [a-z]+)?)?(?:;[a-z]+( [a-z]+)?(:[a-z_.()]+( [a-z]+)?)?)*$`)

var constraints []string = []string{
	"not null",
	"unique",
	"primary key",
	"null",
}

var mappedConstraints []string = []string{
	"foreign key",
	"default",
	"check",
	"exclusion",
	"type",
	"on update",
	"on delete",
}

var actions []string = []string{"no action", "restrict", "cascade", "set null", "set default"}

type Field struct {
	name       string
	typeStr    string
	tags       []string
	mappedTags map[string]string
}

func NewField(tagString string) (*Field, error) {
	name, typeStr, mappedSormTags, sormTags, err := createTagList(tagString)
	if err != nil {
		return nil, err
	}

	return &Field{
		name:       name,
		typeStr:    typeStr,
		tags:       sormTags,
		mappedTags: mappedSormTags,
	}, nil
}

func (f *Field) WriteToSQLTable(table ISQLTable) error {
	table.AddColumn(f.name, f.typeStr)

	if !slices.Contains(f.tags, "null") && !slices.Contains(f.tags, "not null") {
		f.tags = append(f.tags, "null")
	}

	for _, constraint := range f.tags {
		err := mapConstraintToTable(constraint, f.name, table)
		if err != nil {
			return fmt.Errorf("Could not create field %s in table, got error: %s", f.name, err.Error())
		}
	}

	tableReference, ok := f.mappedTags["foreign key"]
	if ok {
		onUpdate := "no action"
		onDelete := "no action"
		updateAction, onUpdatePresent := f.mappedTags["on update"]
		deleteAction, onDeletePresent := f.mappedTags["on delete"]
		if onUpdatePresent {
			onUpdate = updateAction
			delete(f.mappedTags, "on update")
		}

		if onDeletePresent {
			onDelete = deleteAction
			delete(f.mappedTags, "on delete")
		}
		tableName, referenceColumn, _ := strings.Cut(tableReference, ".")
		err := table.AddForeignKey(f.name, tableName, referenceColumn, onUpdate, onDelete)
		if err != nil {
			return fmt.Errorf("Was not able to create foreign key constraint for the field %s, got error: %s", f.name, err.Error())
		}
		delete(f.mappedTags, "foreign key")
	}

	for constraint, value := range f.mappedTags {
		switch constraint {
		case "default":
			err := table.AddDefault(f.name, value)
			if err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("Cannot add field %s to SQL table, cannot process the given constraint: %s", f.name, constraint)
		}
	}

	return nil
}

func (f *Field) String() string {
	return fmt.Sprintf("Name: %s\n Tags: %s, Mapped Tags: %s", f.name, f.tags, f.mappedTags)
}

/*
HELPER FUNCTIONS
*/

// Parse the name of the field based on the json tag and database values for the field based on the sorm tags
func createTagList(tagString string) (parsedName string, parsedType string, parsedMappedTags map[string]string, parsedTags []string, err error) {
	var tagList []string
	var mappedTags map[string]string
	var name string
	var typeStr string
	tags := splitTagString(tagString)
	for _, tag := range tags {
		tagKey, tagValue, parseErr := parseTag(tag)
		if parseErr != nil {
			return "", "", nil, nil, parseErr
		}
		switch tagKey {
		case "json":
			name = tagValue
		case "sorm":
			parsedType, mappedSormTagValues, sormValues, err := parseSormTags(tagValue)
			if err != nil {
				return "", "", nil, nil, err
			}
			typeStr = parsedType
			tagList = sormValues
			mappedTags = mappedSormTagValues
		default:
			continue
		}
	}
	if name == "" {
		return "", "", nil, nil, fmt.Errorf("Cannot create field without 'json' tag, required to identify the field name in database")
	}

	if len(tagList) == 0 && len(mappedTags) == 0 {
		return "", "", nil, nil, fmt.Errorf("Cannot create field without 'sorm' tag, required to configure migrations for the field")
	}
	return name, typeStr, mappedTags, tagList, nil
}

// Retrieve all of the tags within a given tag string, assumes there is only 1 space between each tag
// e.x: Would extract {json:"balls", sorm:"primary key"} from `json:"balls" sorm:"primary key"`
func splitTagString(tagString string) []string {
	var tags []string
	start := -1
	inQuotes := false
	for i, char := range tagString {
		if char == '`' { // ignore the start of the quotes
			continue
		}

		// do not account for any spaces within quotations, does not signify the end of the tag
		if !inQuotes && char == '"' {
			inQuotes = true
		} else if inQuotes && char == '"' {
			inQuotes = false
		}

		if char != ' ' && start == -1 { // Found the first character that isn't a white space
			start = i
		} else if char == ' ' && start != -1 && !inQuotes { // find a white space that signifies the end of a tag string
			tags = append(tags, tagString[start:i])
			start = -1
		}
	}
	if start != -1 {
		tags = append(tags, tagString[start:])
	}
	return tags
}

// Parse a given field tag string and return the key of the tag and the value
func parseTag(tag string) (tagKey string, tagValue string, err error) {
	before, _, ok := strings.Cut(tag, ":")
	if !ok {
		return "", "", fmt.Errorf("Tag is formatted incorrectly, `:` is required in the tag: %s", tag)
	}
	key := before
	_, after, found := strings.Cut(tag, `"`)
	if !found {
		return "", "", fmt.Errorf(`Tag is formatted incorrectly, expected "" around the tag value (ex: json:"id"): %s`, tag)
	}
	value := after[:strings.Index(after, `"`)]
	return key, value, nil
}

// Parses and validates a sorm tag value to get the information to create a field migration entry
func parseSormTags(tagValue string) (typeStr string, mappedTags map[string]string, tags []string, err error) {
	validFormat := validationRegex.MatchString(tagValue)
	if !validFormat {
		return "", nil, nil, fmt.Errorf("sorm tag is formatted incorrectly: %s", tagValue)
	}
	var tagValues []string = strings.Split(tagValue, ";")
	// create a mapping of fields for tags with associated values (example: type, default, on update, etc)
	// Create list of tag values that don't have associated values (primary key, not null, etc)
	var mappedSormValues map[string]string = make(map[string]string)
	var sormValues []string
	for _, tag := range tagValues {
		before, after, found := strings.Cut(tag, ":")
		if found {
			err := validateMappedTagValue(before, after)
			if err != nil {
				return "", nil, nil, err
			}
			mappedSormValues[before] = after
		} else {
			err := validateTagValue(tag, constraints)
			if err != nil {
				return "", nil, nil, err
			}
			sormValues = append(sormValues, tag)
		}
	}

	typeStr, ok := mappedSormValues["type"]
	if !ok {
		return "", nil, nil, fmt.Errorf("Cannot create the field with the given sorm value: %s. The field is missing the type value", tagValue)
	}
	delete(mappedSormValues, "type")

	_, onDelete := mappedSormValues["on delete"]
	_, onUpdate := mappedSormValues["on update"]
	_, foreignKey := mappedSormValues["foreign key"]
	if (onDelete || onUpdate) && !foreignKey {
		return "", nil, nil, fmt.Errorf(`Cannot contain actions "on update" or "on delete" if a foreign key constraint is not present. Given SORM tag value: %s`, tagValue)
	}

	return typeStr, mappedSormValues, sormValues, nil
}

func validateMappedTagValue(key string, value string) error {
	err := validateTagValue(key, mappedConstraints)
	if err != nil {
		return fmt.Errorf("The given constraint %s is not a processable mapped constraint, check the formatting of the constraint in the sorm tag", key)
	}

	switch key {
	case "type": // check if the type is a possible type for a postgres database
		if !slices.Contains(postgresTypes, value) && !strings.Contains(value, "varchar") {
			return fmt.Errorf("The given type %s is not a possible type in a postgres database", value)
		}
		return nil
	case "foreign key": // idk how to validate this one yet... might just be up to the db
		before, after, found := strings.Cut(value, ".")
		if !found {
			return fmt.Errorf(`Expected the value of the "foreign key" constraint to be formatted as <table_name>.<column_name> but got "%s"`, value)
		}
		if len(before) == 0 {
			return fmt.Errorf("Must specify table name for the foreign key, got length 0 string before .")
		}
		if len(after) == 0 {
			return fmt.Errorf("Must specify column reference for the foreign key, got length 0 string after .")
		}
		return nil
	case "default": // idk how to validate it yet
		return nil
	case "on update":
		if !slices.Contains(actions, value) {
			return fmt.Errorf("Given on update action %s is not a possible action, select one of the following: %s", value, actions)
		}
		return nil
	case "on delete":
		if !slices.Contains(actions, value) {
			return fmt.Errorf("Given on delete action %s is not a possible action, select one of the following %s", value, actions)
		}
		return nil
	default:
		return fmt.Errorf("Cannot process the given constraint %s yet, I'm too lazy to cover all cases...", key)
	}
}

func validateTagValue(constraint string, constraintSlice []string) error {
	if slices.Contains(constraintSlice, constraint) {
		return nil
	} else {
		return fmt.Errorf("Cannot process the given constraint %s, not a possible constraint with no value %s", constraint, constraintSlice)
	}
}

func mapConstraintToTable(constraint string, columnName string, table ISQLTable) error {
	switch constraint {
	case "not null":
		err := table.AddNotNull(columnName)
		if err != nil {
			return err
		}
		return nil
	case "null":
		err := table.AddNull(columnName)
		if err != nil {
			return err
		}
		return nil
	case "primary key":
		err := table.AddPrimaryKey(columnName)
		if err != nil {
			return err
		}
		return nil
	case "unique":
		err := table.AddUnique(columnName)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("Cannot process the given single value constraint into a table, check that the constraint exists: %s", constraint)
	}
}
