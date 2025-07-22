package repository

import (
	"apiRestFulSimpes/internal/model"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrRegistroNaoEncontrado = errors.New("nenhum registro com o id encontrado")

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) FindAllTask() ([]model.Task, error) {
	rows, err := s.db.Query(context.Background(), "SELECT * FROM TASKS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []model.Task
	var u model.Task

	for rows.Next() {

		if err := rows.Scan(&u.ID, &u.Title, &u.Description, &u.Done); err != nil {
			return nil, err
		}

		tarefas = append(tarefas, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tarefas, nil
}

func (s *Repository) FindTaskByID(id int) (model.Task, error) {

	var tarefa model.Task

	row := s.db.QueryRow(context.Background(), "SELECT * FROM TASKS WHERE ID = $1", id)

	if err := row.Scan(&tarefa.ID, &tarefa.Title, &tarefa.Description, &tarefa.Done); err != nil {
		return tarefa, err
	}

	return tarefa, nil
}

func (s *Repository) InsertTask(task model.Task) (int, error) {
	err := s.db.QueryRow(context.Background(), "INSERT INTO tasks (title, description, done) VALUES ($1, $2, $3) RETURNING id", task.Title, task.Description, task.Done).Scan(&task.ID)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (s *Repository) UpdateTask(id int, task model.Task) error {
	result, err := s.db.Exec(context.Background(), "UPDATE tasks SET title=$1, description=$2, done=$3 WHERE id=$4", task.Title, task.Description, task.Done, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}

func (s *Repository) DeleteTask(id int) error {
	result, err := s.db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}
