package models

import (
	"github.com/youssefsafwat2/event-booking/db"
	"github.com/youssefsafwat2/event-booking/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password, name)
	VALUES (?, ?, ?);
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	result, err := stmt.Exec(u.Email, u.Password, u.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetUsers() ([]User, error) {
	var users = []User{}
	query := `
	SELECT id, email, name
	FROM users;
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}
