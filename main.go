package main

import (
    "fmt"

    "github.com/leonrbaker/gomilter"
)

// embed third-party struct
type Mymilter struct {
    gomilter.MilterRaw
}

func main() {
    fmt.Println("Demonstration package for recursion bug during dependency fetching with 'govendor'")
}
