package main

import "fmt"

// run will instantiate and startup the go application
func Run() error {
	fmt.Println("Starting...")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}
