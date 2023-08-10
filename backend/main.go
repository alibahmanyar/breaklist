package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
)

type response struct {
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type reqData struct {
	Data []string `json:"data"`
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

func getTasks(c *fiber.Ctx) error {
	tasks, err0 := getLines(os.Getenv("TASKS_LIST_PATH"))

	err1 := c.JSON(&response{Message: "success",
		Data: tasks})

	if err0 != nil {
		return err0
	} else {
		return err1
	}

}

func getReminders(c *fiber.Ctx) error {
	rems, err0 := getLines(os.Getenv("REMINDERS_LIST_PATH"))

	err1 := c.JSON(&response{Message: "success",
		Data: rems})

	if err0 != nil {
		return err0
	} else {
		return err1
	}

}

func addTasks(c *fiber.Ctx) error {
	data := reqData{}
	dec := json.NewDecoder(bytes.NewReader(c.Body()))
	dec.DisallowUnknownFields()

	err0 := dec.Decode(&data)

	tasks, _ := getLines(os.Getenv("TASKS_LIST_PATH"))

	f, err1 := os.OpenFile(os.Getenv("TASKS_LIST_PATH"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		log.Error(err1)
	}

	defer f.Close()

	for _, task := range data.Data {
		if slices.Contains(tasks, task) {
			continue
		}
		if _, err1 = f.WriteString("\n" + task); err1 != nil {
			log.Error(err1)
		}
	}

	c.JSON(&response{Message: "success"})

	return err0
}

func addReminders(c *fiber.Ctx) error {
	data := reqData{}
	dec := json.NewDecoder(bytes.NewReader(c.Body()))
	dec.DisallowUnknownFields()

	err0 := dec.Decode(&data)

	rems, _ := getLines(os.Getenv("REMINDERS_LIST_PATH"))

	f, err1 := os.OpenFile(os.Getenv("REMINDERS_LIST_PATH"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		log.Error(err1)
	}

	defer f.Close()

	for _, rem := range data.Data {
		if slices.Contains(rems, rem) {
			continue
		}
		if _, err1 = f.WriteString("\n" + rem); err1 != nil {
			log.Error(err1)
		}
	}

	c.JSON(&response{Message: "success"})

	return err0
}

func delTasks(c *fiber.Ctx) error {
	data := reqData{}
	dec := json.NewDecoder(bytes.NewReader(c.Body()))
	dec.DisallowUnknownFields()

	err0 := dec.Decode(&data)

	tasksf, _ := os.ReadFile(os.Getenv("TASKS_LIST_PATH"))
	tasks := strings.Split(string(tasksf), "\n")

	f, err1 := os.Create(os.Getenv("TASKS_LIST_PATH"))
	if err1 != nil {
		log.Error(err1)
	}
	defer f.Close()

	for i, t := range tasks {
		if slices.Contains(data.Data, t) {
			continue
		}
		if i != 0 {
			f.WriteString("\n")
		}
		f.WriteString(t)
	}

	c.JSON(&response{Message: "success"})

	if err0 != nil {
		return err0
	} else {
		return err1
	}

}

func delReminders(c *fiber.Ctx) error {
	data := reqData{}
	dec := json.NewDecoder(bytes.NewReader(c.Body()))
	dec.DisallowUnknownFields()

	err0 := dec.Decode(&data)

	remsf, _ := os.ReadFile(os.Getenv("REMINDERS_LIST_PATH"))
	rems := strings.Split(string(remsf), "\n")

	f, err1 := os.Create(os.Getenv("REMINDERS_LIST_PATH"))
	if err1 != nil {
		log.Error(err1)
	}
	defer f.Close()

	for i, t := range rems {
		if slices.Contains(data.Data, t) {
			continue
		}
		if i != 0 {
			f.WriteString("\n")
		}
		f.WriteString(t)
	}

	c.JSON(&response{Message: "success"})

	if err0 != nil {
		return err0
	} else {
		return err1
	}
}

func main() {
	godotenv.Load()

	// Create tasks and reminders files if they don't exist
	os.MkdirAll(filepath.Dir(os.Getenv("TASKS_LIST_PATH")), os.ModePerm)
	os.MkdirAll(filepath.Dir(os.Getenv("REMINDERS_LIST_PATH")), os.ModePerm)

	f1, _ := os.OpenFile(os.Getenv("TASKS_LIST_PATH"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f2, _ := os.OpenFile(os.Getenv("REMINDERS_LIST_PATH"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f1.Close()
	f2.Close()

	app := fiber.New()

	app.Use(helmet.New())
	app.Use(cors.New())

	app.Get("/api/task", getTasks)
	app.Get("/api/reminder", getReminders)

	app.Post("/api/task", addTasks)
	app.Post("/api/reminder", addReminders)

	app.Delete("/api/task", delTasks)
	app.Delete("/api/reminder", delReminders)

	app.Static("/", "./static/")

	app.Listen(":3000")
}
