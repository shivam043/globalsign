package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var regExp = regexp.MustCompile(`\pL+('\pL+)*`)

type WordCount struct {
	Key   string
	Value int
}

// counts the frequency of words in a file
func countWords(file string) ([]WordCount) {
	fileHandler, err := os.Open(file)
	if err != nil {
		fmt.Println("Error reading file " + err.Error())
		os.Exit(1)
	}
	defer fileHandler.Close()
	wcs := make([]WordCount, 1000)
	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file " + err.Error())
			os.Exit(1)
		}
		line := strings.ToLower(scanner.Text())
		words := regExp.FindAllString(line, -1)
		for _, word := range words {

			var found bool
			for i := range wcs {
				if wcs[i].Key == word {
					wcs[i].Value = wcs[i].Value + 1
					found = true
					break
				}

			}

			if !found {
				wc := WordCount{word, 1}

				wcs = append(wcs, wc)

			}

		}

	}

	return wcs
}

func main() {

	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a valid file path")
		os.Exit(1)

	}
	file := flag.Arg(0)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("Please check file path does not exist ")
		os.Exit(1)
	}
	var ss []WordCount

	ss = countWords(file)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	count := 0
	for _, value := range ss {
		fmt.Printf("%d %s\n", value.Value, value.Key)
		count++
		if count == 19 {
			break
		}
	}
}
