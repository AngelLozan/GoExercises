package main

import (
	"github.com/AngelLozan/learningpackages/helpers"
	"fmt"
)


const numPool = 10
func CalculateValue(intChan chan int){
	randomNum := helpers.RandomNumber(numPool)
	intChan <- randomNum // pass the random number to the channel
}
func main(){
	intChan := make(chan int) // Make a channel (place to send info) that can be received by one or more places in my program
	defer close(intChan) // Whatever comes after this, execute that after the current function is done. So close the channel.

	//Set up concurrent operation in it's own go routine. They can all run at same time
	go CalculateValue(intChan)

	// Listen for response to the channel (get value from channel)
	num := <-intChan
	fmt.Println(num)
}
