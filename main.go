package main

// POST /tasks – создание задачи.
// GET /tasks – получение списка всех задач.
// PUT /tasks/:id – обновление задачи.
// DELETE /tasks/:id – удаление задачи.

import (
	"log"

	"github.com/8Air/SkillsRockTestTask/db"
	_ "github.com/8Air/SkillsRockTestTask/docs"
	"github.com/8Air/SkillsRockTestTask/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	db.InitDB()
	defer db.CloseConnection()
	app := fiber.New()

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	//----------------

	//CRUD requests
	app.Post("/tasks", handlers.CreateTask)
	app.Get("/tasks", handlers.GetTasksList)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
	//----------------

	log.Fatal(app.Listen(":3000"))
}
