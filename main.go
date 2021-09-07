package main

import (
	"flag"
	"fmt"
	"vlq/decode"
)


var v = flag.String("v","a","input v")
var input = flag.String("input","","please input nubmers")

func main(){
	flag.Parse()
	fmt.Println("end:",*v)

	ret, err := decode.Decode(*input)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
