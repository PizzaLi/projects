package main

import "fmt"

func main() {
	s := "Pending"
	var s1 string
	switch s {
	case "Pending", "Running", "Succeeded", "Failed":
		s1 = s
	default:
		s1 = "Unknown"
	}
	fmt.Println(s1)
}
