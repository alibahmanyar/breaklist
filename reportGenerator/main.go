package main

import (
	"html/template"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type PageData struct {
	TasksRems  []string
	Interval   intervals
	HNArticles []hnArticle
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

// getLines reads the contents of a file, filters out empty lines and lines starting with "#" (comments),
// and returns a slice containing the non-comment lines.
//
// Parameters:
// filename string - The name of the file to be read.
//
// Returns:
// []string - A slice containing non-comment lines from the file.
func getLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	sdata := string(data)
	allLines := strings.Split(sdata, "\n")
	var lines []string

	for _, line := range allLines {
		if !strings.HasPrefix(line, "#") && len(line) > 0 { // skip empty lines and comments
			lines = append(lines, line)
		}
	}

	return lines, err
}

// matchesCronPart checks if a given value matches a cron part.
// Parameters:
//   - value: The value to match against the cron part.
//   - cronPart: The cron part (DOM, M, or DOW) to match against.
//
// Returns true if the value matches any part of the cronPart, otherwise false.
func matchesCronPart(value int, cronPart string) bool {
	if cronPart == "*" {
		return true // Wildcard always matches
	}

	values := strings.Split(cronPart, ",")
	for _, v := range values {
		if strings.HasPrefix(v, "*/") {
			intervalStr := strings.TrimPrefix(v, "*/")
			interval, err := strconv.Atoi(intervalStr)
			check(err)
			if value%interval == 0 {
				return true // Value matches an interval
			}
		} else {
			expectedValue, err := strconv.Atoi(v)
			check(err)
			if value == expectedValue {
				return true // Value matches an expected specific value
			}
		}
	}

	return false // No match found
}

// matchCronExpression checks if a given date matches a cron expression.
// The cron expression is in the format "DOM M DOW".
// Parameters:
//   - date: The date to check against the cron expression.
//   - cronExpression: The cron expression to match against.
//
// Returns true if the date matches the cron expression, otherwise false.
func matchCronExpression(date time.Time, cronExpression string) bool {
	parts := strings.Split(cronExpression, " ")

	// Check Day of Month (DOM)
	if !matchesCronPart(date.Day(), parts[0]) {
		return false
	}

	// Check Month (M)
	if !matchesCronPart(int(date.Month()), parts[1]) {
		return false
	}

	// Check Day of Week (DOW)
	if !matchesCronPart(int(date.Weekday()), parts[2]) {
		return false
	}

	return true
}

func main() {
	godotenv.Load()
	var err error

	// Get weather data
	interval := getWeatherForecast()

	// Get tasks list
	tasks, err := getLines(os.Getenv("TASKS_LIST_PATH"))
	check(err)

	// Get reminders list and check which ones should be reminded today
	allReminders, err := getLines(os.Getenv("REMINDERS_LIST_PATH"))
	check(err)
	var reminders []string

	now := time.Now()

	for _, r := range allReminders {
		rs := strings.Split(r, "|")

		if matchCronExpression(now, rs[0]) {
			reminders = append(reminders, rs[1])
		}
	}

	// Get HN articles
	articles := getHNArticles()[:8]

	// Rendering the HTML template
	tmpl, _ := template.ParseFiles("template.html")
	f, _ := os.Create("temp.html")
	err = tmpl.Execute(f, PageData{TasksRems: append(tasks, reminders...), Interval: interval, HNArticles: articles})
	f.Close()
	check(err)

	f.Close()

	cmd := exec.Command("sh", "-c", "wkhtmltopdf --encoding utf-8 --margin-top 1mm --margin-bottom 7mm --margin-left 0mm --margin-right 0mm --page-height 500mm --page-width 47mm --grayscale --enable-local-file-access \"temp.html\" \"breaklist.pdf\"")
	_, err = cmd.Output()
	check(err)

	log.Print("Created new report")
}
