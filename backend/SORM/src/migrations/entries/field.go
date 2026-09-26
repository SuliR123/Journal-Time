package entries

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

var validationRegex = regexp.MustCompile(`^[^;]+(?:;[^;]+)*$`)

var constraints []string = []string{
	"not null",
	"unique",
	"primary key",
}

var mappedConstraints []string = []string{
	"foreign key",
	"default",
	"check",
	"exclusion",
	"type",
}

type Field struct {
	name       string
	tags       []string
	mappedTags map[string]string
}

func NewField(tagString string) (*Field, error) {
	name, mappedSormTags, sormTags, err := createTagList(tagString)
	if err != nil {
		return nil, err
	}

	return &Field{
		name:       name,
		tags:       sormTags,
		mappedTags: mappedSormTags,
	}, nil
}

func (f *Field) WriteToSQL() {
	/*
		TAG LIST:
			- NOT NULL
			- UNIQUE
			- PRIMARY KEY
			- FOREIGN KEY (ON UPDATE, ON DELETE, CASCADE, NO ACTION)
			- CHECK
			- EXCLUSION
			- DEFAULT

		TYPES: prolly will make a txt file of all the types and load it in
	*/
}

func (f *Field) String() string {
	return fmt.Sprintf("Name: %s\n Tags: %s, Mapped Tags: %s", f.name, f.tags, f.mappedTags)
}

/*
HELPER FUNCTIONS
*/

// Parse the name of the field based on the json tag and database values for the field based on the sorm tags
func createTagList(tagString string) (string, map[string]string, []string, error) {
	var tagList []string
	var mappedTags map[string]string
	var name string
	tags := splitTagString(tagString)
	for _, tag := range tags {
		tagKey, tagValue, parseErr := parseTag(tag)
		if parseErr != nil {
			return "", nil, nil, parseErr
		}
		switch tagKey {
		case "json":
			name = tagValue
		case "sorm":
			mappedSormTagValues, sormValues, err := parseSormTags(tagValue)
			if err != nil {
				return "", nil, nil, err
			}
			tagList = sormValues
			mappedTags = mappedSormTagValues
		default:
			continue
		}
	}
	if name == "" {
		return "", nil, nil, fmt.Errorf("Cannot create field without 'json' tag, required to identify the field name in database")
	}

	if len(tagList) == 0 {
		return "", nil, nil, fmt.Errorf("Cannot create field without 'sorm' tag, required to configure migrations for the field")
	}
	return name, mappedTags, tagList, nil
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
func parseTag(tag string) (string, string, error) {
	before, _, ok := strings.Cut(tag, ":")
	if !ok {
		return "", "", fmt.Errorf("Tag is formatted incorrectly, `:` is required in the tag : %s", tag)
	}
	key := before
	_, after, found := strings.Cut(tag, `"`)
	if !found {
		return "", "", fmt.Errorf(`Tag is formatted incorrectly, expected "" around the tag value (e.x: json:"id"): %s`, tag)
	}
	value := after[:strings.Index(after, `"`)]
	return key, value, nil
}

// Parses and validates a sorm tag value to get the information to create a field migration entry
func parseSormTags(tagValue string) (map[string]string, []string, error) {
	validFormat := validationRegex.MatchString(tagValue)
	if !validFormat {
		return nil, nil, fmt.Errorf("sorm tag is formatted incorrectly: %s", tagValue)
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
				return nil, nil, err
			}
			mappedSormValues[before] = after
		} else {
			err := validateTagValue(tag, constraints)
			if err != nil {
				return nil, nil, err
			}
			sormValues = append(sormValues, tag)
		}
	}

	_, ok := mappedSormValues["type"]
	if !ok {
		return nil, nil, fmt.Errorf("Cannot create the field with the given sorm value: %s. The field is missing the type value", tagValue)
	}

	return mappedSormValues, sormValues, nil
}

func validateMappedTagValue(key string, value string) error {
	err := validateTagValue(key, mappedConstraints)
	if err != nil {
		return err
	}

	switch key {
	case "type": // check if the type is a possible type for a postgres database
		dir, _ := os.Getwd()
		typesFilepath := filepath.Join(dir, "SORM", "src", "migrations", "entries", "postgres_types.txt")
		contentBytes, err := os.ReadFile(typesFilepath)
		if err != nil {
			return fmt.Errorf("Unable to parse postgres_types.txt file, recieved the following error when opening the file %e", err)
		}
		types := strings.Split(string(contentBytes), " ")
		if !slices.Contains(types, value) {
			return fmt.Errorf("The given type %s is not a possible type in a postgres database", value)
		}
		return nil
	case "foreign key": // idk how to validate this one yet... might just be up to the db
		return nil
	case "default": // ^ same idea with this one
		return nil
	default:
		return fmt.Errorf("Cannot process the given constraint %s yet, I'm too lazy to cover all cases...", key)
	}
}

func validateTagValue(constraint string, constraintSlice []string) error {
	if slices.Contains(constraintSlice, constraint) {
		return nil
	} else {
		return fmt.Errorf("Cannot process the given tag %s, not a possible constraint", constraint)
	}
}
