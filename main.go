package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)
func unzip(name string) {
	nameWithType := name + ".zip"
	//unzip
	_,errUnzip := exec.Command("unzip", nameWithType, "-d", name).Output()
	if errUnzip != nil {
		fmt.Printf("%s", errUnzip)
	}

	//ls and recursive
	outLs, errLs := exec.Command("ls", name).Output()
	if errLs != nil {
		fmt.Printf("%s", errLs)
	}
	s := strings.Split(string(outLs),"\n")
	for i:=0 ; i < len(s)-1 ; i++ {
		if len(s[i]) >= 4 {
			if s[i][len(s[i])-4:] == ".zip" {
				unzip(name + "/" + s[i][:len(s[i])-4])
			}
		}
	}

	//remove zips
	_, errRm := exec.Command("rm", nameWithType).Output()
	if errRm != nil {
		fmt.Printf("%s", errRm)
	}


}
func execute() {
	address := ""
	fmt.Println("Enter zip address:")
	fmt.Scanf("%s",&address)
	unzip(address[:len(address)-4])

}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}