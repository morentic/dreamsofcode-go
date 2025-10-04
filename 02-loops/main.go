package main

import "fmt"

func main() {
	names := []string{"Julchen", "Kathi", "Dennis", "Hendrik"}

	for _, name := range names {
		fmt.Println("Hello", name)
	}
}
