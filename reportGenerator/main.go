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

	for _, r := range allReminders {
		rs := strings.Split(r, "|")
		dates := strings.Split(rs[0], " ")

		if dates[0] != "*" {
			dom, err := strconv.Atoi(dates[0]) // Day of month
			check(err)
			if dom != time.Now().Day() {
				continue
			}
		}

		if dates[1] != "*" {
			m, err := strconv.Atoi(dates[1]) // Month
			check(err)
			if m != int(time.Now().Month()) {
				continue
			}
		}

		if dates[2] != "*" {
			dow, err := strconv.Atoi(dates[2]) // Day of week
			check(err)
			if dow != int(time.Now().Weekday()) {
				continue
			}
		}

		reminders = append(reminders, rs[1])
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
