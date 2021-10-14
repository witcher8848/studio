package main

import (
	"fmt"
)

func main()  {
	count:=0
	fmt.Scanln(&count)
	arr:=make([]int,count)
	temp:=0
	for i := 0; i < count; i++ {
		arr[i]=temp
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr[count/2])
}
