package sakubun

import (
    "github.com/gorilla/mux"
)

var r mux.Router

func SetupRouter() {
    r = mux.NewRouter()
    
    
}