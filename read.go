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

	fmt.Println("Stable Store 0")
	dat, err := ioutil.ReadFile("stable-store-replica0")
	check(err)
	//fmt.Print(string(dat))

	fmt.Println(dat)

	fmt.Println("Stable Store 1")
	dat1, err := ioutil.ReadFile("stable-store-replica1")
	check(err)

	fmt.Println(dat1)

	fmt.Println("Stable Store 2")
	dat2, err := ioutil.ReadFile("stable-store-replica2")
	check(err)

	fmt.Println(dat2)

}
