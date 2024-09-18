package object_name


import "fmt"

func object_name(){
	var arr [3] string
	//[][][][]  Elements
	//0 1 2 3   Index Values

	fmt.Println(arr)

arr = [3]string{"Tea", "Coffee", "Water"}

fmt.Println(arr)

fmt.Println(arr[2]) //Pay Attention to Output Here
arr[2] = "Green Tea" //Similar to pop and shift in arrays

fmt.Println(arr)

} 