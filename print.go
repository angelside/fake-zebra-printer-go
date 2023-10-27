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
	currentTime := createDate()

	// Colorize the data
	data = colorizeData(data)

	// Render a table
	tableRender(currentTime, data)
}

func createDate() string {
	// Format the date
	dateFormat := "02/01/2006 15:04:05" // FIXME: Hard coded data
	currentTime := time.Now().Format(dateFormat)

	// Colorize the time
	currentTime = color.New(color.FgCyan, color.Bold).Sprint(currentTime)

	return currentTime
}

func colorizeData(data string) string {
	if !isValidZplData(data) {
		return data
	}

	templateVarColor := color.New(color.FgGreen)
	zplVarColor := color.New(color.FgYellow)

	data = colorizePattern(data, `[~^][A-Z]{1,2}`, zplVarColor)
	data = colorizePattern(data, `{{\..*?}}`, templateVarColor)

	return data
}

func isValidZplData(data string) bool {
	return strings.HasPrefix(data, "^XA") && strings.HasSuffix(data, "^XZ")
}

func colorizePattern(data string, pattern string, c *color.Color) string {
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllStringFunc(data, func(match string) string {
		return c.Sprint(match)
	})
}

func tableRender(currentTime, data string) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{currentTime})
	t.AppendRow(table.Row{data})

	fmt.Println(t.Render())
}
