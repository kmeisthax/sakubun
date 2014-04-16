package core

import (
    "flag"
    "os"
    "encoding/json"
)

type DatabaseConfig struct {
    Driver string
    DSN string
}

type SiteConfig struct {
    Databases []string
    SessionEntropy string
}

type Config struct {
    Databases map[string]DatabaseConfig
    Sites map[string]SiteConfig
}

var config_location string;
var config_site string;
var config_cgi string;
var config Config;

func SetupConfig() error {
    flag.StringVar(&config_location, "conf", "sakubun-conf.json", "Define where the master configuration for Sakubun is.")
    flag.StringVar(&config_location, "site", "default", "Specify what site to serve requests for.")
    flag.StringVar(&config_location, "cgi", "cgi", "Specify what gateway interface to use.")
    
    flag.Parse()
    
    file, ok := os.Open(config_location)
    if ok != nil {
        return ok
    }
    
    decoder := json.NewDecoder(file)
    decoder.Decode(&config)
    
    return nil
}