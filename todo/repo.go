package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(connStr string) (*Repo, error) {
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &Repo{db: db}, nil
}

func (r *Repo) Add(title string) (Task, error) {
	var t Task
	err := r.db.QueryRow(context.Background(),
		"INSERT INTO tasks (title) VALUES ($1) RETURNING id, title, done, created_at",
		title).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	return t, err
}

func (r *Repo) List() ([]Task, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, title, done, created_at FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *Repo) ToggleComplete(id int) error {
	_, err := r.db.Exec(context.Background(),
		"UPDATE tasks SET done = NOT done WHERE id=$1", id)
	return err
}

func (r *Repo) Init() error {
	_, err := r.db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`)
	return err
}

func (r *Repo) GetByID(id int) (Task, error) {
	var t Task
	err := r.db.QueryRow(context.Background(),
		"SELECT id, title, done, created_at FROM tasks WHERE id=$1", id).
		Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	return t, err
}

func (r *Repo) Update(id int, title string, done bool) (Task, error) {
	var t Task
	err := r.db.QueryRow(context.Background(),
		"UPDATE tasks SET title=$1, done=$2 WHERE id=$3 RETURNING id, title, done, created_at",
		title, done, id).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt)
	return t, err
}

func (r *Repo) Delete(id int) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	return err
}
