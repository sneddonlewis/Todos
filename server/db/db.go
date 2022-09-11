package db

import (
	"database/sql"

	"github.com/sneddonlewis/todos/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
	GetTodos() ([]*model.Todo, error)
}

type PostgresDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

func (d PostgresDB) GetTodos() ([]*model.Todo, error) {
	rows, err := d.db.Query("select title, description from todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*model.Todo
	for rows.Next() {
		t := new(model.Todo)
		err = rows.Scan(&t.Title, &t.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (d PostgresDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := d.db.Query("select name, details from technologies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tech []*model.Technology
	for rows.Next() {
		t := new(model.Technology)
		err = rows.Scan(&t.Name, &t.Details)
		if err != nil {
			return nil, err
		}
		tech = append(tech, t)
	}
	return tech, nil
}
