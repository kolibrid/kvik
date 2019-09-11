package main

import (
	"fmt"

	"github.com/kolibrid/kvik/genecards"
)

func main() {
	sum := genecards.Summary("BRCA1")
	fmt.Println(sum)
}
