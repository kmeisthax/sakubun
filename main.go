package main

import (
    "github.com/gorilla/mux"
    "code.fantranslation.org/sakubun/core"
)

func main() {
    sakubun.SetupConfig()
    sakubun.SetupRouter()
    sakubun.SetupCgi()
}