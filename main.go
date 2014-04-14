package main

import (
    "code.fantranslation.org/sakubun/core"
)

func main() {
    core.SetupConfig()
    core.SetupDB()
    core.SetupRouter()
    core.SetupCgi()
}