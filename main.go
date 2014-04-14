package main

import (
    "code.fantranslation.org/sakubun/core"
)

func main() {
    core.SetupConfig()
    core.SetupRouter()
    core.SetupCgi()
}