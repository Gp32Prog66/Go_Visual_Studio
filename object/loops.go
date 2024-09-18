package object_name

import "fmt"


func object_name()
{
	fmt.Println(i)
	i := 99


	//No Breakpoint. Will loop forever. Hit Ctrl + C to break
	for {
		fmt.Println(i)
		i += 1
		break
	}

	i = 5
	for i < 10
	{
		fmt.Println(i)
		i++
	}

	
}