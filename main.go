package main

import (
	"fmt"
	"time"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(s)
}


func main() {
	wg := sync.WaitGroup{}
	fmt.Println("start")
	wg.Add(1)
	go say("hello", &wg)
	wg.Wait()
	fmt.Println("done")
}