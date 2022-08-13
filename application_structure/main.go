package main

import "fmt"

const secondInHour = 3600

func main() {
	fmt.Println("hello go")
	distance := 60.8
	fmt.Printf("distance in miles is %f \n", distance*secondInHour)
}
