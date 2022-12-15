package main

import (
	"fmt"
	"strconv"
	"time"
)

func greeting(myChannel chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		myChannel <- "hi " + strconv.Itoa(i)
	}
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	for i := 0; i < 10; i++ {
		fmt.Println(<-myChannel)
	}
}
