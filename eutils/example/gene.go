package main

import "fmt"
import "github.com/kolibrid/kvik/eutils"

func main() {
	ds, err := eutils.GeneSummary("627")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ds)
}
