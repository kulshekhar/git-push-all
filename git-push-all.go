package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	branch := "master"
	if len(os.Args) == 2 {
		branch = os.Args[1]
	}

	pushAllTo(branch)
}

func pushAllTo(branch string) {
	remotes := getRemotes()

	for _, remote := range remotes {
		fmt.Printf("Pushing %s to %s:\n", branch, remote)
		c := exec.Command("git", "push", remote, branch)
		var out bytes.Buffer
		var stderr bytes.Buffer
		c.Stdout = &out
		c.Stderr = &stderr
		err := c.Run()

		if err != nil {
			fmt.Printf("An error occurred when pushing %s to %s:\n%s: %s\n", branch, remote, err, stderr.String())
		}

		fmt.Println(out.String())

	}
}

func getRemotes() []string {
	out, err := exec.Command("git",
		"remote", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(out), "\n")

	remotes := map[string]struct{}{}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lineParts := strings.Split(line, "\t")
		remotes[lineParts[0]] = struct{}{}
	}

	remoteList := []string{}
	for k := range remotes {
		remoteList = append(remoteList, k)
	}

	return remoteList
}
