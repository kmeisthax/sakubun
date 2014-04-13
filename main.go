package "code.fantranslation.org/sakubun"

import (
    "github.com/gorilla/mux"
)

func Bootstrap() {
    setup_config()
    setup_router()
    setup_cgi()
}