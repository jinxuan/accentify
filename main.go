package main

import (
	// "io/ioutil"
	"bufio"
	"fmt"
	"os"
)

const (
	ukWordsPath = "/wordlists/uk"
	usWordsPath = "/wordlists/us"
)

func usage() {
	fmt.Println("=========================================================")
	fmt.Println("This is a cli for accenting your text files")
	fmt.Println("Usage: ./accentify filepath uk||us")
	fmt.Println("example to UK-ize your textfile: ./accentify original.md uk")
	fmt.Println("=========================================================")
	os.Exit(1)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Need more arguments")
		usage()
	}

	fmt.Println("Loading word lists...")
	curDir, _ := os.Getwd()
	ukWords := readLine(curDir + ukWordsPath)
	usWords := readLine(curDir + usWordsPath)

	// sanity check
	if len(ukWords) != len(usWords) || len(ukWords) == 0 || len(usWords) == 0 {
		fmt.Println("Failed to load word lists")
		return
	}
	fmt.Println("Loaded...\nUK words:", len(ukWords), "US words:", len(usWords))

	// scan through the file

}

func readLine(path string) (lines []string) {
	fmt.Println(path)
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
