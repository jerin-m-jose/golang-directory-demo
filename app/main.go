package main

import (
	"bufio"
	"filesystemdemo/internal/filesystem"
	"filesystemdemo/internal/inmemoryfs"
	"fmt"
	"os"
	"strings"
)

const allowedCmds = "CREATE/MOVE/LIST/DELETE/EXIT"

func main() {
	// create in memory filesystem
	fs := inmemoryfs.New()

	//tell user what to do
	fmt.Println("Enter cmd :", allowedCmds)

	// get input from user
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			fields := strings.Fields(line)
			bExit := handleCmd(fs, fields)
			if bExit {
				fmt.Println("Exiting app.. ")
				return
			}
		} else {
			// Exit if scanner is closed or encounters an error
			fmt.Println("Input closed. Exiting app..")
			return
		}
	}
}

// handleCmd handles input command and call filesystem
func handleCmd(fs filesystem.FileSystem, cmd []string) (exit bool) {
	switch strings.ToLower(cmd[0]) {
	case "create":
		if len(cmd) < 2 {
			fmt.Println("Provide a folder name to create ")
			return
		}
		err := fs.Create(cmd[1])
		if err != nil {
			fmt.Println("Cannot create ", cmd[1])
		}

	case "list":
		fs.List(filesystem.DefaultSort)

	case "move":
		if len(cmd) < 3 {
			fmt.Println("Provide a src & dest folder name")
			return
		}
		err := fs.Move(cmd[1], cmd[2])
		if err != nil {
			fmt.Printf("Cannot move from %s to %s\n", cmd[1], cmd[2])
		}
	case "delete":
		if len(cmd) > 1 {
			err := fs.Delete(cmd[1])
			if err != nil {
				fmt.Println("Cannot delete ", cmd[1], "-", err.Error())
			}
		}
	case "exit":
		return true
	default:
		fmt.Println("Unknown cmd. Choose from ", allowedCmds)
	}

	return false
}
