package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := ""
	
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(os.Args); i++ {
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}

	return nil
}
