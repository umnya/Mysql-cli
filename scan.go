package main

import (
	"fmt"
	"os/exec"
)

func main() {

	getSarInfo()

}

func getSarInfo() error {
	IP_ADDR := "10.80.51.232"
	//command := fmt.Sprintf("ssh %s sar 1 10", IP_ADDR)
	//fmt.Println(command)
	out, _ := exec.Command("ssh", IP_ADDR, "sar", "1", "5").Output()
	output := string(out[:])

	fmt.Println(output)

	return nil
}
