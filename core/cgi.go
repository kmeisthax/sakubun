package core

import (
    "net/http"
    "net/http/cgi"
    "net/http/fcgi"
)

func serve(handler http.Handler, cgiName string) error {
    switch cgiName {
        case "cgi":
            return cgi.Serve(handler)
        case "fcgi":
            return fcgi.Serve(nil, handler)
        case "standalone":
            return http.ListenAndServe(":8080", handler)
    }
    
    //TODO: Actually return new error
    return nil
}

func SetupCgi() {
    serve(SiteHandler(r), config_cgi)
}