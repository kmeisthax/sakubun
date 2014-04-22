package core

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
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