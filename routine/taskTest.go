package main

import "fmt"

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c)=", len(c))

	go func() {
		c <- 1

	}()
	num := <-c
	println(num)
}
