package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)
func unzip(name string) {
	nameWithType := name + ".zip"
	out, err := exec.Command("unzip", nameWithType, "-d", name).Output()
}
func execute() {

	out, err := exec.Command("ls", "/").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")

	s := strings.Split(string(out),"\n")
	output := string(out[:])

	fmt.Println(output)
	fmt.Println(s)
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}