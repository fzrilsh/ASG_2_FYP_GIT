package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()

	texts := []string{
		"Hello Ms. Nikita Ananda Putri Masaling, S.Kom., M.Kom.",
		"",
		"My name is Fazril Syaveral Hillaby",
		"NIM: 2902659922",
		"",
		"My favorite programming language are: PHP and Javascript",
	}

	for _, line := range texts {
		fmt.Println(line)
	}
}
