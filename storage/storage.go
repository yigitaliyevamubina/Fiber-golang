package storage

import (
	"backend/models"
	"database/sql"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	connection := "user=postgres password=mubina2007 dbname=backend sslmode=disable"
	mydb, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return mydb, err
}

func CreateUser(user models.User) (*models.User, error) {
	mydb, err := connect()
	if err != nil {
		return nil, err
	}
	defer mydb.Close()
	query := `INSERT INTO users(id, first_name, last_name) VALUES($1, $2, $3) RETURNING id, first_name, last_name`

	var respUser models.User

	rowUser := mydb.QueryRow(query, user.ID, user.FirstName, user.LastName)
	if err := rowUser.Scan(&respUser.ID, &respUser.FirstName, &respUser.LastName); err != nil {
		return nil, err
	}

	return &respUser, nil
}

func GetUserById(userID string) (*models.User, error) {
	mydb, err := connect()
	if err != nil {
		return nil, err
	}
	defer mydb.Close()

	query := `SELECT id, first_name, last_name FROM users WHERE id = $1`

	var respUser models.User

	rowUser := mydb.QueryRow(query, userID)
	if err := rowUser.Scan(&respUser.ID, &respUser.FirstName, &respUser.LastName); err != nil {
		return nil, err
	}
	return &respUser, nil
}

func GetAllUsers(page, limit int) ([]*models.User, error) {
	mydb, err := connect()
	if err != nil {
		return nil, err
	}
	defer mydb.Close()

	offset := limit * (page - 1)

	respUsers := []*models.User{}

	query := `SELECT id, first_name, last_name FROM users LIMIT $1 OFFSET $2`

	rows, err := mydb.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userId sql.NullString
		var user models.User

		if err := rows.Scan(&userId, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		if userId.Valid {
			user.ID = userId.String
		}

		respUsers = append(respUsers, &user)
	}
	return respUsers, nil
}

func UpdateUserById(userID string, user models.User) (*models.User, error) {
	mydb, err := connect()
	if err != nil {
		return nil, err
	}
	defer mydb.Close()

	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3 RETURNING id, first_name, last_name`
	rowUser := mydb.QueryRow(query, user.FirstName, user.LastName, userID)

	var respUser models.User
	if err := rowUser.Scan(&respUser.ID, &respUser.FirstName, &respUser.LastName); err != nil {
		return nil, err
	}
	return &respUser, nil
}

func DeleteUserById(userID string) (*models.User, error) {
	mydb, err := connect()
	if err != nil {
		return nil, err
	}
	defer mydb.Close()

	query := `DELETE FROM users WHERE id = $1 RETURNING id, first_name, last_name`

	var respUser models.User
	rowUser := mydb.QueryRow(query, userID)
	if err := rowUser.Scan(&respUser.ID, &respUser.FirstName, &respUser.LastName); err != nil {
		return nil, err
	}
	return &respUser, nil
}
