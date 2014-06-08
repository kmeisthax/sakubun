package core

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    
    "reflect"
    "fmt"
    "strings"
)

var databases map[string]*sql.DB

func connect_database(name string) error {
    thedb, ok := sql.Open(config.Databases[name].Driver, config.Databases[name].DSN)
    if (ok != nil) {
        return ok
    }
    
    databases[name] = thedb
    return nil
}

func SetupDB() error {
    for _, dbname := range config.Sites[config_site].Databases {
        ok := connect_database(dbname)
        if (ok != nil) {
            return ok
        }
    }
    
    return nil
}

type DbType int;

const (
    DBTYPE_CHAR DbType = iota
    DBTYPE_VARCHAR
    DBTYPE_TEXT
    DBTYPE_BLOB
    DBTYPE_INT
    DBTYPE_FLOAT
    DBTYPE_NUMERIC
);

/* Represents information about a field in an in-database table.
 */
type SchemaField struct {
    /* The SQL type of this field. */
    FieldType DbType
    
    /* If the field is allowed to be null. */
    IsNullable bool
    
    /* String-likes: Number of characters/bytes allowed in field. */
    Length int
    
    /* Numeric: Number of significant digits to store. */
    Precision int
    
    /* Numeric: Number of fractional digits. */
    Scale int
    
    /* Default value for the field.
     * 
     * Default values must be of a primitive Go type compatible with the DbType
     * specified for this field.
     */
    Default *interface{}
}

/* Represents information about a foreign table which should have fields that
 * match our own. */
type ForeignKey struct {
    /* List of local columns that compromise the key. */
    LocalColumnKey []string
    
    ForeignTableName string
    
    /* List of foreign columns that the key must match. */
    ForeignColumnKey []string
}

/* Represents information about an in-database table from which we can pull and
 * push data from.
 */
type Schema struct {
    /* The name of this database table. */
    TableName string
    
    /* Fields present within the table. */
    Fields map[string]SchemaField
    
    /* List of field names which form the identity of each row. */
    PrimaryKey []string
    
    /* List of additional sets of fields which must be collectively unique.
     * 
     * The primary key is implicitly unique.
     */
    UniqueKeys map[string][]string
    
    /* List of fields that map to fields in another table.
     */
    ForeignKeys map[string]ForeignKey
}

/* Given a schema, create a Stmt suitable for installing the schema in a DB. */
func (sch *Schema) CreateTableStmt(forDB *sql.DB) (*sql.Stmt, error) {
    colDefs := make([]string, 0, len(sch.Fields) + 1 + len(sch.UniqueKeys) + len(sch.ForeignKeys))
    
    for fieldName, fieldSpec := range sch.Fields {
        ourDef := ""
        
        ourDef += fieldName;
        
        switch fieldSpec.FieldType {
            case DBTYPE_CHAR:
                ourDef += " CHAR"
            case DBTYPE_VARCHAR:
                ourDef += " VARCHAR"
            case DBTYPE_TEXT:
                ourDef += " TEXT"
            case DBTYPE_BLOB:
                ourDef += " BLOB"
            case DBTYPE_INT:
                ourDef += " INT"
            case DBTYPE_FLOAT:
                ourDef += " FLOAT"
            case DBTYPE_NUMERIC:
                ourDef += " NUMERIC"
            default:
                ourDef += " INT"
        }
        
        colDefs = append(colDefs, ourDef)
    }
    
    colDefs = append(colDefs, fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(sch.PrimaryKey, ", ")))
    
    for _, uniqueKeyFields := range sch.UniqueKeys {
        colDefs = append(colDefs, fmt.Sprintf("UNIQUE KEY (%s)", strings.Join(uniqueKeyFields, ", ")))
    }
    
    for _, foreignKeySpec := range sch.ForeignKeys {
        colDefs = append(colDefs, fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s (%s)", 
                                             strings.Join(foreignKeySpec.LocalColumnKey, ", "),
                                             foreignKeySpec.ForeignTableName,
                                             strings.Join(foreignKeySpec.ForeignColumnKey, ", ")))
    }
    
    stmtQuery := fmt.Sprintf("CREATE TABLE %s (%s)", sch.TableName, strings.Join(colDefs, ", "))
    
    return forDB.Prepare(stmtQuery)
}

/* Create a prepared statement for inserting data into the table.
 * The FieldNames provided will be the list of arguments to pipe into the query
 * at Exec time. 
 */
func (sch *Schema) InsertStmt(forDB *sql.DB, FieldNames []string) (*sql.Stmt, error) {
    paramsString := make([]string, 0, len(FieldNames))
    
    for i := 0; i < len(FieldNames); i++ {
        paramsString[i] = "?"
    }
    
    stmtQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
                             sch.TableName,
                             strings.Join(FieldNames, ", "),
                             strings.Join(paramsString, ", "))
    
    return forDB.Prepare(stmtQuery)
}

/* Given a struct type, generate a database Schema for it.
 * 
 * You cannot generate schemata from arbitrary types. Types must instead be
 * presented in serializable form. That is, all fields on the struct must be
 * of the following types:
 * 
 *      Intrinsic integer and float types
 *      Boolean types
 *      Byte slices
 *      String slices
 *      database/sql's various nullable versions of the above types
 *      Embedded or anonymous structs in serializable form
 * 
 * 
 */
func SchemaFromType(t reflect.Type) *Schema {
    /* TODO: Implement */
    return nil
}