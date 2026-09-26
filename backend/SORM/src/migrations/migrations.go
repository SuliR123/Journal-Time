package migrations

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	entries "github.com/SuliR123/Journal-Time/SORM/src/migrations/entries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MigrationCommands struct {
	dbConnection   *pgxpool.Pool
	models         []entries.IEntry
	prevMigrations []string
}

func New(dbConnection *pgxpool.Pool, modelsDirPath string) MigrationCommands {
	fset := token.NewFileSet()
	modelEntries, err := os.ReadDir(modelsDirPath)
	if err != nil {
		log.Fatalf("Could not read models from the given directory path %s got error %s", modelsDirPath, err)
	}
	var models []entries.IEntry = make([]entries.IEntry, len(modelEntries))
	for i, modelEntry := range modelEntries {
		pathToModel := filepath.Join(modelsDirPath, modelEntry.Name())
		file, err := parser.ParseFile(fset, pathToModel, nil, parser.SkipObjectResolution)
		if err != nil {
			log.Fatalf("Unable to find model file: %s, got error: %s", pathToModel, err)
		}

		var fields []entries.IEntry
		ast.Inspect(file, func(n ast.Node) bool {
			typeSpec, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}

			structSpec, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return true
			}

			for _, field := range structSpec.Fields.List {
				var fieldTags string = field.Tag.Value
				fieldEntry, err := entries.NewField(fieldTags)
				if err != nil {
					fieldNameInStruct := field.Names[0].Name
					log.Fatalf("Could not create field %s, recieved error: %s", fieldNameInStruct, err)
				}
				fields = append(fields, fieldEntry)
			}
			models[i] = entries.NewModel(typeSpec.Name.Name, fields)
			return true
		})
	}

	// if migrations folder doesn't exist at specified location create it

	// load previous migration files if they exist

	return MigrationCommands{
		dbConnection: dbConnection,
		models:       models,
	}
}

func (m *MigrationCommands) CreateMigrations() {}

func (m *MigrationCommands) HashMigrations() {}

func (m *MigrationCommands) ApplyMigrations() {}

func (m *MigrationCommands) RevertMigrations() {}

// load model files, any error handling associated with loading in the model files
// deal with any out of order migrations, hashing, generating migrations based on an existing migrations folder (in the case models are changed, deleted, etc)

func (m *MigrationCommands) String() string {
	var sb strings.Builder

	sb.WriteByte('\n')
	for _, model := range m.models {
		sb.WriteString(model.String())
		sb.WriteByte('\n')
	}

	return sb.String()
}
