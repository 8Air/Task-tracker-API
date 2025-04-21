package handlers

// POST /tasks – создание задачи.
// GET /tasks – получение списка всех задач.
// PUT /tasks/:id – обновление задачи.
// DELETE /tasks/:id – удаление задачи.

import (
	"log"
	"strconv"
	"time"

	"github.com/8Air/SkillsRockTestTask/db"
	"github.com/8Air/SkillsRockTestTask/models"
	"github.com/gofiber/fiber/v2"
)

// CreateTask godoc
// @Summary Создать новую задачу
// @Description Добавляет новую задачу в базу данных
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Новая задача"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]interface{}
// @Router /tasks [post]
func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	var err error
	if err = c.BodyParser(&task); err != nil {
		log.Printf("failed to parse request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if task.Status == "" {
		task.Status = models.StatusNew
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if task.ID, err = db.CreateTask(task); err != nil {
		log.Printf("failed to create task: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("task has been added")
	return c.Status(fiber.StatusCreated).JSON(task)
}

// GetTasksList godoc
// @Summary Получить список всех задач
// @Description Возвращает все задачи
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 400 {object} map[string]interface{}
// @Router /tasks [get]
func GetTasksList(c *fiber.Ctx) error {
	tasks, err := db.GetAllTasks()
	if err != nil {
		log.Printf("failed to get tasks list: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	log.Printf("tasks list has been sent")
	return c.JSON(tasks)
}

// UpdateTask godoc
// @Summary Обновить задачу
// @Description Обновляет задачу по ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body models.Task true "Обновляемые поля задачи"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tasks/{id} [put]
func UpdateTask(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("invalid task id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}

	var changes models.Task
	if err := c.BodyParser(&changes); err != nil {
		log.Printf("invalid request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	changes.UpdatedAt = time.Now()

	err = db.UpdateTask(id, changes)
	if err != nil {
		log.Printf("failed to update task: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update task",
		})
	}

	log.Printf("task updated successfully")
	return c.JSON(fiber.Map{
		"message": "task updated successfully",
	})
}

// DeleteTask  godoc
// @Summary Удалить задачу
// @Description Удаляет задачу по ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 204 "Задача успешно удалена"
// @Failure 400 {object} map[string]interface{}
// @Router /tasks/{id} [delete]
func DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("failed to delete task task: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := db.DeleteTask(id); err != nil {
		log.Printf("failed to delete task task: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("task has been deleted: %v", err)
	return c.SendStatus(fiber.StatusNoContent)
}
