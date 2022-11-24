package main

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		out, _ := exec.Command("tasklist").Output()

		if strings.Contains(string(out), "DiskAlert.exe") {
			log.Println("DiskAlert.exe is running, killing it")
			exec.Command("taskkill", "/F", "/IM", "DiskAlert.exe").Run()
		}

		time.Sleep(1 * time.Second)
	}
}
