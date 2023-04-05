package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("./rust-cli-code/target/release/rust-cli-code", "--name", "John").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output)) // will print "Hello John!"
}
