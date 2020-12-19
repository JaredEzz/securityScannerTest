package main

import (
	"cataloger/applist"
	"fmt"
)

func main() {
	files := applist.Files()
	fmt.Println("Executable Files:")
	for _, file := range files {
			fmt.Println(file)
	}

	fmt.Println("Verified Files:")
	for _, file := range files {
		if applist.IsVerified(file){
			fmt.Println(file)
		}
	}
} //end main