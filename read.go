package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("stable-store-replica0")
	check(err)
	//fmt.Print(string(dat))

	fmt.Println(dat)

	//str := string(dat)

	//fmt.Println(str)

}