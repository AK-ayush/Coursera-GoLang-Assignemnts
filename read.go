package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type name struct {
	fname string
	lname string
}

func trimName(str string) string {
	if len(str) > 20 {
		return string(str[0:20])
	}
	return str
}

func main() {
	var fname string
	fmt.Println("Please enter the full path of file: ")
	fmt.Scanln(&fname)
	fd, err := os.Open(fname)
	check(err)

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)

	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	fd.Close()

	nameArr := make([]name, 0)
	var cnt int
	for _, line := range txtLines {
		words := strings.Fields(line)

		if len(words[0]) > 20 || len(words[1]) > 20 {
			cnt++
		}
		newName := name{
			trimName(words[0]),
			trimName(words[1]),
		}

		nameArr = append(nameArr, newName)
	}
	if cnt > 0 {
		fmt.Printf("\nFound %d lines with First or Last name bigger than 20 chars!\nTrimmed to 20 chars.\n\n", cnt)
	}
	// fmt.Println(len(nameArr))

	for i, val := range nameArr {
		fmt.Printf("Name %d :\n", i+1)
		fmt.Printf("\tFirst Name: %s \n\tLast Name: %s\n", val.fname, val.lname)
	}
}
