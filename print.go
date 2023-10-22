package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func printData(data string) {
	// Datetime
	curtime := createDate()

	// Colorize the data
	data = colorizeData(data)

	// Render a table
	tableRender(curtime, data)
}

func createDate() string {
	// Format the date
	dateFormat := "02/01/2006 15:04:05" // FIXME: Hard coded data
	curtime := time.Now().Format(dateFormat)

	// Colorize the time
	timeColor := color.New(color.FgCyan).Add(color.Bold)
	curtime = timeColor.Sprint(curtime)

	return curtime
}

func colorizeData(data string) string {
	// Check if the data starts with ^XA and ends with ^XZ
	if strings.HasPrefix(data, "^XA") && strings.HasSuffix(data, "^XZ") {
		// Create colors for the template variables and the rule-matching text
		templateVarColor := color.New(color.FgGreen)
		zplVarColor := color.New(color.FgYellow)

		// Define a regular expression pattern to match the rule
		zplVarPattern := regexp.MustCompile(`[~^][A-Z]{1,2}`)

		// Replace all matches of the rule with the colored version
		data = zplVarPattern.ReplaceAllStringFunc(data, func(match string) string {
			return zplVarColor.Sprint(match)
		})

		// Define a regular expression pattern to match template variables
		templateVarPattern := regexp.MustCompile(`{{\..*?}}`)

		// Replace all template variables with the templateVarColor
		data = templateVarPattern.ReplaceAllStringFunc(data, func(match string) string {
			return templateVarColor.Sprint(match)
		})
	}

	return data
}

func tableRender(curtime string, data string) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{curtime})
	t.AppendRow(table.Row{data})

	fmt.Println(t.Render())
}
