package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserRepo interface {
	Insert(u *User) User
	GetByName(name string) (User, error)
	Update(id int, u User) error
	Delete(id int) error
}

type UserRepoImpl struct {
	connection_setting string
}

func NewUserRepo(conn_setting string) *UserRepoImpl {
	return &UserRepoImpl{
		connection_setting: conn_setting,
	}
}

func (s *UserRepoImpl) getConnection() *sql.DB {
	db, err := sql.Open("postgres", s.connection_setting)

	if err != nil {
		panic(err)
	}

	fmt.Println(db.Stats())

	return db
}

func (s *UserRepoImpl) Insert(u *User) User {
	var db *sql.DB = s.getConnection()
	defer db.Close()

	var user User
	db.QueryRow("INSERT INTO \"go\".users (name) VALUES ($1)", u.Name).Scan(&user.Id, &user.Name)

	return user
}

func (s *UserRepoImpl) GetByName(name string) (User, error) {
	var db *sql.DB = s.getConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM \"go\".users WHERE name = $1", name)

	var user User = User{}

	err := row.Scan(&user.Id, &user.Name)

	if err != nil {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (s *UserRepoImpl) Update(id int, u User) error {
	var db *sql.DB = s.getConnection()
	defer db.Close()

	_, err := db.Exec("UPDATE \"go\".users set name = $1 where id = $2", u.Name, id)

	if err != nil {
		return errors.New("Cannot update user")
	}

	return nil
}

func (s *UserRepoImpl) Delete(id int) error {
	var db *sql.DB = s.getConnection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM \"go\".users WHERE id = $1", id)

	if err != nil {
		return errors.New("Cannot delete user")
	}

	return nil
}
