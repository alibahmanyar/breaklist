package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/exp/slices"
)

type response struct {
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type reqData struct {
	Data []string `json:"data"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTasks(c *fiber.Ctx) error {
	data, err0 := os.ReadFile("tasks.list")
	sdata := string(data)

	tasks := strings.Split(sdata, "\n")

	err1 := c.JSON(&response{Message: "success",
		Data: tasks})

	if err0 != nil {
		return err0
	} else {
		return err1
	}

}

func getReminders(c *fiber.Ctx) error {
	data, err0 := os.ReadFile("reminders.list")
	sdata := string(data)

	rems := strings.Split(sdata, "\n")

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

	tasksf, _ := os.ReadFile("tasks.list")
	tasks := strings.Split(string(tasksf), "\n")

	f, err1 := os.OpenFile("tasks.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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

	remsf, err1 := os.ReadFile("reminders.list")
	rems := strings.Split(string(remsf), "\n")

	f, err1 := os.OpenFile("reminders.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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

	tasksf, err1 := os.ReadFile("tasks.list")
	tasks := strings.Split(string(tasksf), "\n")

	f, err1 := os.Create("tasks.list")
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

	remsf, err1 := os.ReadFile("reminders.list")
	rems := strings.Split(string(remsf), "\n")

	f, err1 := os.Create("reminders.list")
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
	// Create tasks and reminders files if they don't exist
	f1, _ := os.OpenFile("tasks.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f2, _ := os.OpenFile("reminders.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f1.Close()
	f2.Close()

	app := fiber.New()
	log.Debug("start")

	app.Get("/task", getTasks)
	app.Get("/reminder", getReminders)

	app.Post("/task", addTasks)
	app.Post("/reminder", addReminders)

	app.Delete("/task", delTasks)
	app.Delete("/reminder", delReminders)

	app.Listen(":3000")
}
