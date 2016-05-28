package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/fjukstad/kvik/r"
)

func main() {
	t0 := time.Now()
	var server = flag.String("server", "localhost:8181", "ip:port of server")
	flag.Parse()

	s := r.Server{*server, "", ""}
	for i := 0; i < 100; i++ {
		out, err := s.Call("stats", "rnorm", "n="+strconv.Itoa(i)+",mean=1")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(out)
	}

	//err = s.UploadPackage("packages/addman_0.1.tgz")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//out, err = s.Call("addman", "hello", "")
	//fmt.Println(out, err)
	//ag, err := s.Call("mixt", "getGeneList", "tissue='blood',module='blue'")
	//if err != nil {
	//	fmt.Println("Call failed", err)
	//	return
	//}

	//genes, err := s.Get(ag, "csv")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(genes))

	//fmt.Println("GETALLGENES OUTPUT", string(genes))
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

}
