package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/quadrosh/user-list/user"
)

type postgresRepository struct {
	DB *sql.DB
}

// ConnectPostgres connects to postgres database and returns type (postgresRepository) which implements the user.RepoInterface interface
func ConnectPostgres(dbHost, dbPort, dbUser, dbName, dbPass string) (user.RepoInterface, error) {
	repo := postgresRepository{}
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	log.Println("connectionString: ", connectionString)

	var err error
	repo.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = repo.DB.Ping(); err != nil {
		return nil, err
	}
	return &repo, nil
}

// pingDB tries to ping the database
func pingDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// CreateUserTableIfNotExists creates users table for local environnent,
// in Docker db and table creates file /database/create_fixture.sql
func (m *postgresRepository) CreateUserTableIfNotExists() error {

	if err := m.DB.Ping(); err != nil {
		return err
	}

	query := `
	 CREATE TABLE IF NOT EXISTS users
	 (
		id VARCHAR(255) NOT NULL CONSTRAINT unique_id UNIQUE,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		age INTEGER DEFAULT 0,
		recording_date BIGINT NOT NULL
	
	)`

	if _, err := m.DB.Exec(query); err != nil {
		return err
	}
	log.Println("Database table created")

	return nil
}

// CreateUser creates new user
func (m *postgresRepository) CreateUser(user *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID string

	query := `insert into users(
		id,
		first_name, 
		last_name, 
		age, 
		recording_date)
		values($1, $2, $3, $4, $5) returning id`

	err := m.DB.QueryRowContext(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Age,
		user.RecordingDate,
	).Scan(&newID)

	if err != nil {
		return err
	}

	return nil
}

// FindUsersByID finds users by id
func (m *postgresRepository) FindUsersByID(value string) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []user.User

	query := `	
		select
			id,
			first_name, 
			last_name, 
			age, 
			recording_date
		from
			users
		where
			id = $1`

	rows, err := m.DB.QueryContext(ctx, query, value)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RecordingDate,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindUsersByFirstName finds users by first_name
func (m *postgresRepository) FindUsersByFirstName(value string) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []user.User

	query := `	
		select
			id,
			first_name, 
			last_name, 
			age, 
			recording_date
		from
			users
		where
			first_name = $1`

	rows, err := m.DB.QueryContext(ctx, query, value)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RecordingDate,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

// FindUsersByLastName finds users by last_name
func (m *postgresRepository) FindUsersByLastName(value string) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []user.User
	query := `	
		select
			id,
			first_name, 
			last_name, 
			age, 
			recording_date
		from
			users
		where
			last_name = $1`
	rows, err := m.DB.QueryContext(ctx, query, value)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RecordingDate,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindUsersByAge finds users by age
func (m *postgresRepository) FindUsersByAge(value int) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []user.User
	query := `	
		select
			id,
			first_name, 
			last_name, 
			age, 
			recording_date
		from
			users
		where
			age = $1`
	rows, err := m.DB.QueryContext(ctx, query, value)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RecordingDate,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindUsersByAge finds users by age
func (m *postgresRepository) FilterUsersByRange(dateFrom, dateTo, ageFrom, ageTo int) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []user.User

	query := `	
		SELECT
			id,
			first_name, 
			last_name, 
			age, 
			recording_date

		FROM
			users
			
		WHERE
			id IS NOT NULL
		AND (cast($1 as bigint)>0 AND recording_date >= $1 OR $1=0)	
		AND (cast($2 as bigint)>0 AND recording_date <= $2 OR $2=0)	
		AND ($3>0 AND age >= $3 OR $3=0)	
		AND ($4>0 AND age <= $4 OR $4=0)	
			`

	rows, err := m.DB.QueryContext(ctx, query, dateFrom, dateTo, ageFrom, ageTo)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user user.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.RecordingDate,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
