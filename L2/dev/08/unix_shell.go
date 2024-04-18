
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	goPs "github.com/mitchellh/go-ps"
)

func cd(request []string) {
	if len(request) == 1 {
		fmt.Fprintln(os.Stderr, "where dir?")
	} else if len(request) == 2 {
		err := os.Chdir(request[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func pwd(request []string) {
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "pwd command takes no arguments")
		return
	}
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(path)
	}
}

func echo(request []string) {
	fmt.Println(strings.Join(request[1:], " "))
}

func kill(request []string) {
	if len(request) != 2 {
		fmt.Fprintln(os.Stderr, "usage: kill <pid>")
		return
	}
	pid, err := strconv.Atoi(request[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid PID")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func ps(request []string) {
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "ps command takes no arguments")
		return
	}
	sliceProc, err := goPs.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to fetch processes:", err)
		return
	}
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v, Process ID: %v\n", proc.Executable(), proc.Pid())
	}
}

func executePipeline(commands []string) {
	var cmd *exec.Cmd
	processes := make([]*exec.Cmd, len(commands))

	for i, command := range commands {
		parts := strings.Fields(command)
		cmd = exec.Command(parts[0], parts[1:]...)
		processes[i] = cmd
		if i > 0 {
			if stdout, err := processes[i-1].StdoutPipe(); err == nil {
				cmd.Stdin = stdout
			} else {
				fmt.Fprintln(os.Stderr, "Failed to create pipe")
				return
			}
		}
	}

	if len(processes) > 0 {
		processes[len(processes)-1].Stdout = os.Stdout
	}

	for _, cmd := range processes {
		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Error starting command:", err)
			return
		}
	}

	for _, cmd := range processes {
		if err := cmd.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for command:", err)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("my-shell> ")
	for scanner.Scan() {
		input := scanner.Text()
		if strings.Contains(input, "|") {
			commands := strings.Split(input, "|")
			executePipeline(commands)
		} else {
			request := strings.Fields(input)
			if len(request) == 0 {
				continue
			}
			switch request[0] {
			case "cd":
				cd(request)
			case "pwd":
				pwd(request)
			case "echo":
				echo(request)
			case "kill":
				kill(request)
			case "ps":
				ps(request)
			}
		}
		fmt.Print("my-shell> ")
	}
}
