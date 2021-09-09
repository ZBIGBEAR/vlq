package main

import (
	"flag"
	"fmt"
	"vlq/decode"
	"vlq/encode"
)


var v = flag.Bool("v",false,"input v")
var input = flag.String("input","","please input nubmers")

func main(){
	flag.Parse()
	var ret string
	var err error
	if *v {
		// 解码
		ret, err = decode.Decode(*input)
	}else{
		// 编码
		ret, err = encode.Encode(*input)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
