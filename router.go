package "code.fantranslation.org/sakubun"

import (
    "github.com/gorilla/mux"
)

var r mux.Router

func setup_router() {
    r = mux.NewRouter()
    
    
}