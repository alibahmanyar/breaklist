package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	log.Info(strings.Join(tasks, ","))

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
	log.Info(strings.Join(rems, ","))

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

	log.Info(data)

	f, err1 := os.OpenFile("tasks.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		log.Error(err1)
	}

	defer f.Close()

	for _, task := range data.Data {
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

	log.Info(data)

	f, err1 := os.OpenFile("reminders.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err1 != nil {
		log.Error(err1)
	}

	defer f.Close()

	for _, rem := range data.Data {
		if _, err1 = f.WriteString("\n" + rem); err1 != nil {
			log.Error(err1)
		}
	}

	c.JSON(&response{Message: "success"})

	return err0
}

func main() {
	app := fiber.New()
	log.Debug("start")

	app.Get("/task", getTasks)
	app.Get("/reminder", getReminders)

	app.Post("/task", addTasks)
	app.Post("/reminder", addReminders)

	app.Listen(":3000")
}
