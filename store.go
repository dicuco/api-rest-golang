package main

import "database/sql"

type Store interface {
	//users
	CreateUser() error
	//tasks
	CreateTask(t *Task) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage { //esto es una función común que crea uno nuevo
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error { //es una función asociada a un receiver (al struct)
	return nil

}
func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}
