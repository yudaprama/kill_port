package killport

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Do(port string) {
	var command string
	switch runtime.GOOS {
	case "windows":
		command = fmt.Sprintf("netstat -ano | findstr :%s | findstr LISTENING", port)

		// execute the command to get the PIDs listening to port 8086
		out, err := exec.Command("cmd", "/C", command).Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1)
		}

		// split the output by newlines and trim whitespace
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")

		// kill each PID listening to port 8086
		for _, line := range lines {
			pid := strings.TrimSpace(line)

			if pid != "" {
				cmd := exec.Command("taskkill", "/F", "/PID", pid)
				if err := cmd.Run(); err != nil {
					fmt.Println("Error killing process:", err)
				} else {
					fmt.Println("Successfully killed process with PID", pid)
				}
			}
		}
	case "linux", "darwin":
		command = fmt.Sprintf("lsof -i :%s -t", port)

		// Execute the command to get the PIDs listening to port.
		out, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1)
		}

		// Split the output by newline characters to get each PID.
		pids := strings.Split(strings.TrimSpace(string(out)), "\n")

		// Kill each PID listening to port 8086.
		for _, pid := range pids {
			cmd := exec.Command("kill", pid)
			if err := cmd.Run(); err != nil {
				fmt.Println("Error killing process:", err)
			} else {
				fmt.Println("Successfully killed process with PID", pid)
			}
		}
	default:
		fmt.Println("Unsupported operating system:", runtime.GOOS)
		os.Exit(1)
	}
}
