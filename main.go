package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	t := NewTrie()
	dat, _ := ioutil.ReadFile("words.txt")
	t.PushText(string(dat))
	// t.Repr()
	fmt.Println("completing: tha")
	t1 := time.Now()
	res := t.Complete("tha")
	t2 := time.Now()
	fmt.Printf("found %v results in %v\n",len(res), t2.Sub(t1))
}