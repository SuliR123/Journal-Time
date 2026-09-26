# Migrations Setup

## Workflow

Configure .env file in SORM/.env with the directory storing the table object definitions, the directory to output the migrations, and the Database connection string

```
MODEL_DIRECTORY=/internal/models
OUTPUT_DIRECTORY=/internal/migrations
DB_CONNECTION_STRING=postgresql://postgres:postgres@127.0.0.1:54322/postgres
```

## Formatting Fields

The table name will default to the name of the struct. Camel case struct names will be converted to snake case & the name will be converted to all lowercase, ex: BruhBurger would become table bruh_burger

To configure the fields within the table the migrations scripts will utilize the *tags* attatched to each field. Each field is required to have the "json:<name>" tag and "sorm:" tag

The sorm tag stores all of the information related how the struct will be parsed into a database table. Multiple values will go within this singular tag, and for that reason, it must be formatted in the following way sorm:"<value1>;<value2>". Values that require an input (such as type, foreign key, default, etc) will be formatted as following: sorm:"<key>:<value>;<value1>". Here's an example: 

``` go
type Bruh struct {
    Balls string `json:"balls" sorm:"type:text;not null"`
}
```

## Constraints

### Primary Key: 

A specific column or set of columns that uniqley identifies each row

Ex:
``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"primary key"`
}
```

### Foreign Key:

Represents a connection from this field to another table in the database

When creating the struct in the models folder, the type of the field in go should match the struct of the table this field is refering too.

However, the type for the entry in the database should match the field the foreign key is referring to. For example if the there was a field referring to the "balls" table via the id (type:uuid) on my struct "Bruh" the type of the field in go would correspond to the Balls object (which corresponds to the Balls table) while the sorm:"type:_" would be uuid as the "bruh" table will store the uuid corresponding to the balls object

Ex:
``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"type:uuid"`
}

type Bruh struct {
    Stinky Balls `json:"stinky" sorm:"type:uuid;foreign key:balls.id"`
}
```

#### On Update/Delete

The table that contains the foreign key is the child (referencer) table and the table its referring to is the parent table. The "On Update" and "On Delete" commands refer to what happens to the child table if the parent table is updated or deleted. 

Here is a list of the possible values these commands can take in: No Action, Restrict, Cascade, Set Null, Set Default

If these commands are not included in the sorm tag, they will default to "No Action". For "On Update" this means the referenced key cannot be updated in the parent row if there are any child tables referencing the column. For "On Delete" this means the parent row the foreign key is referencing cannot be deleted if there are any child tables referencing the row.

``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"type:uuid"`
}

type Bruh struct {
    Stinky Balls `json:"stinky" sorm:"type:uuid;foreign key:balls.id;on update:cascade"`
}
```

### Default: 

If a value is not provided for this field you can specify what the column should default too on creation

``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"type:uuid;default:gen_random_uuid()"`
}

```

### Not Null:

Ensures the column cannot be empty or missing a value

``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"type:uuid;not null"`
}

```

### Unique:

Ensures there can only be one row with a given value for this column in a table

``` go
type Balls struct {
    Id uuid.UUID `json:"id" sorm:"type:uuid;unique"`
}

```

### Type (*REQUIRED*):

Represents the type the field will be stored as within the Postgress database. This type must be one of the supported types within the Postgres database. To see an exhaustive list of supported migration types check SORM/src/migrations/entries/postgres_types.txt (I might've missed some but here's the working list I have so far)

The "type" is a mapped value, meaning it must take in the type value you expect it to be in the database

Ex:
``` go
type Balls struct {
    Bruh string `json:"bruh" sorm:"type:text"`
}
```

### Check: 
Not implemented yet

### Exclusion:
Not implemented yet