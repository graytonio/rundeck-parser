package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var inputFile string
var outputFile string

func init() {
	flag.StringVar(&inputFile, "i", "", "Path to input file")
	flag.StringVar(&outputFile, "o", "./output.txt", "Path to output file")
}

func main() {
	flag.Parse()
	if inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	body, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error Reading Input File:\n%s", err.Error())
	}
	errors := parseFile(string(body))

	outputString := ""
	for key, lines := range errors {
		// fmt.Printf("\n%s\n", key)
		outputString += fmt.Sprintf("\n%s\n", key)
		for _, line := range lines {
			// fmt.Printf("\t%s\n", line)
			outputString += fmt.Sprintf("\t%s\n", line)
		}
	}
	outputString = strings.Trim(outputString, " \n\t")

	ioutil.WriteFile(outputFile, []byte(outputString), 0644)
}

func stripLineInfo(line string) string {
	reg := regexp.MustCompile(`([0-9]{2}:[0-9]{2}:[0-9]{2}) (\[(.*?)\])(\[(.*?)\]) `)
	res := reg.ReplaceAllString(line, "")
	return res
}

func parseFile(text string) map[string][]string {
	lines := strings.Split(text, "\n")
	errors := make(map[string][]string)

	var current_task string
	for _, line := range lines {
		if strings.Contains(line, "[print a handy formatted list of failed and unreachable groups]") {
			break
		}

		clean_line := stripLineInfo(line)
		if strings.Contains(clean_line, "TASK") {
			current_task = clean_line
			continue
		}

		if strings.Contains(clean_line, "fatal") || strings.Contains(clean_line, "failure") {
			errors[current_task] = append(errors[current_task], clean_line)
		}
	}

	return errors
}
