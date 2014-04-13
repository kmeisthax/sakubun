package main

import (
    "github.com/gorilla/mux"
    "code.fantranslation.org/sakubun"
)

func main() {
    sakubun.SetupConfig()
    sakubun.SetupRouter()
    sakubun.SetupCgi()
}