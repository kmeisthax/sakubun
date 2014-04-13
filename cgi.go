package "code.fantranslation.org/sakubun"

import (
    "net/http"
    "net/http/cgi"
    "net/http/fcgi"
)

func serve(handler http.Handler, cgi String) error {
    switch cgi {
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

func setup_cgi() {
    serve(r, config_cgi)
}