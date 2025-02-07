package main

import (
	"fmt"
	"time"

	"github.com/kolibrid/kvik/r"
)

func main() {

	c := r.Client{"localhost:8181", "", ""}
	t0 := time.Now()

	key, err := c.Call("stats", "rnorm", "n=100")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := c.Get(key, "json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	res, err = c.Rpc("stats", "rnorm", "n=100", "json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))

}
