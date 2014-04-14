package main

import (
    "github.com/gorilla/mux"
    "code.fantranslation.org/sakubun/core"
)

func main() {
    core.SetupConfig()
    core.SetupRouter()
    core.SetupCgi()
}