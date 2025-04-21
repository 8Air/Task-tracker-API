package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/8Air/SkillsRockTestTask/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() error {
	err := connectToDb()
	if err != nil {
		return fmt.Errorf("error connecting to data base: %w", err)
	}

	err = migration()
	if err != nil {

		return fmt.Errorf("error running migration: %w", err)
	}
	return nil
}

func CloseConnection() {
	Pool.Close()
}

func GetAllTasks() ([]models.Task, error) {
	rows, err := Pool.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func CreateTask(task models.Task) (int, error) {
	fmt.Println(task)

	var id int
	query := `
        INSERT INTO tasks (title, description, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := Pool.QueryRow(
		context.Background(),
		query,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&id)

	return id, err
}

func UpdateTask(id int, changes models.Task) error {
	var (
		setClauses []string
		args       []interface{}
		argPos     = 1
	)

	if changes.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf("title = $%d", argPos))
		args = append(args, changes.Title)
		argPos++
	}

	if changes.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argPos))
		args = append(args, changes.Description)
		argPos++
	}

	if changes.Status != "" {
		switch changes.Status {
		case models.StatusNew, models.StatusInProgress, models.StatusDone:
			setClauses = append(setClauses, fmt.Sprintf("status = $%d", argPos))
			args = append(args, changes.Status)
			argPos++
		default:
			return fmt.Errorf("wrong task status")
		}
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argPos))
	args = append(args, time.Now())
	argPos++

	query := fmt.Sprintf(`
        UPDATE tasks
        SET %s
        WHERE id = $%d
    `, strings.Join(setClauses, ", "), argPos)

	args = append(args, id)

	cmd, err := Pool.Exec(context.Background(), query, args...)

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("task not found")
	}

	return err
}

func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	cmdTag, err := Pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
