package core

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    
    "reflect"
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
    DBTYPE_VARCHAR,
    DBTYPE_TEXT,
    DBTYPE_BLOB,
    DBTYPE_INT,
    DBTYPE_FLOAT,
    DBTYPE_NUMERIC
);

/* Represents information about a field in an in-database table.
 */
type SchemaField struct {
    FieldName string
    FieldType DbType
}

/* Represents information about an in-database table from which we can pull and
 * push data from.
 */
type Schema struct {
    TableName string
    Fields map[string]SchemaField
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