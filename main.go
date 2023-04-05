package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("./rust-code/target/release/rust-code", "--name", "John").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output)) // will print "Hello, John!"
}
