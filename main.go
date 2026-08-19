package main

import (
	"fmt"
	"os"
)

// capacity_planner - Plan infrastructure capacity
func capacity_planner(path string) {
	fmt.Println("========================================")
	fmt.Println("  Capacity-Planner")
	fmt.Println("  Plan infrastructure capacity")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	capacity_planner(path)
}
