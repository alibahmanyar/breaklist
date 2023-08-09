package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type PageData struct {
	Tasks    []string
	Interval intervals
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	tasks, err := getLines("tasks.list")
	check(err)

	// Get reminders list and check which ones should be reminded today
	allReminders, err := getLines("reminders.list")
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

	// Rendering the HTML template
	tmpl, _ := template.ParseFiles("template.html")
	f, _ := os.Create("temp.html")
	err = tmpl.Execute(f, PageData{Tasks: append(tasks, reminders...), Interval: interval})
	f.Close()
	check(err)

	f.Close()

	cmd := exec.Command("sh", "-c", "./wkhtmltopdf --encoding utf-8 --margin-top 1mm --margin-bottom 7mm --margin-left 0mm --margin-right 0mm --page-height 210mm --page-width 47mm --grayscale --enable-local-file-access \"temp.html\" \"to_print.pdf\"")
	_, err = cmd.Output()
	check(err)

	fmt.Println("Done")
}
