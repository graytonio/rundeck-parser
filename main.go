package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
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

	for lines := errors.Oldest(); lines != nil; lines = lines.Next() {
		outputString += fmt.Sprintf("\n%s\n", lines.Key)
		for _, line := range lines.Value {
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

func parseFile(text string) *orderedmap.OrderedMap[string, []string] {
	lines := strings.Split(text, "\n")
	// errors := make(map[string][]string)
	errors := orderedmap.New[string, []string]()

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
			// errors[current_task] = append(errors[current_task], clean_line)
			current_array, present := errors.Get(current_task)
			if !present {
				errors.Set(current_task, []string{clean_line})
			} else {
				errors.Set(current_task, append(current_array, clean_line))
			}

		}
	}

	return errors
}
