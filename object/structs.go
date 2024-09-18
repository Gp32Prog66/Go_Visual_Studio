package object_name

import {
	"bufio"
	"fmt"
	"os"
	"strings"
}

func main()
{
	fmt.Println("Please select an option")
	fmt.Println("1 Print Menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = string.Trimspace(choice)

	//Output Data
	type menuItem struct
	{
		name string
		prices map[string]float64
	}

	menu := []menuItem
	{
		{name: "Coffee", prices: map[string]map[string]float64{"small" : 1.65, "medium : " : 1.90, "large" : 2.05}}
		{name: "Green Tea", prices: map[string]map[string]float64{"small" : 1.50, "medium : " : 1.95, "large" : 2.20}}

	}

	fmt.Println(menu)

	for _, item := range menu
	{
		fmt.println(item.name)
		fmt.println(strings.Repeat("-", 10))

		for size, price := range item.prices
		{
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}