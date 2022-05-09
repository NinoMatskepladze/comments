package main

import "fmt"

// Run function is responsible for instaniation and stasrtup of the application
func Run() error {
	fmt.Println("Starting up the application")
	return nil
}

func main() {
	fmt.Println("Comments app")
	if err := Run(); err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
