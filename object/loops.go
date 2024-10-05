package main

import "fmt"


func main() {
	i := 99
	fmt.Println("Printing ",i)


	//No Breakpoint. Will loop forever. Hit Ctrl + C to break
	for {
		fmt.Println(i)
		i += 1
		break
	} 

	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	
}