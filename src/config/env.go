package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load() {
	filePath := ".env"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}

	for _, e := range lines {
		pair := strings.Split(e, "=")
		os.Setenv(pair[0], pair[1])
	}
}
