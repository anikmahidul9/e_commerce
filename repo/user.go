package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(usr User) (*User, error)
	Find(email, pass string) (*User, error)
	// List() ([]*User, error)
	Get(uId int) (*User, error)
	Update(usr User) (*User, error)
	Delete(uId int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(usr User) (*User, error) {
	query := `
        INSERT INTO users
		 (first_name, last_name, email, password, is_shop_owner)
        VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
        RETURNING id
    `
	var userID int
	rows, err := r.db.NamedQuery(query, usr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userID)
	}
	usr.ID = userID
	return &usr, nil
}

func (r userRepo) Find(email, pass string) (*User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE email = :email AND password = :password
    `
	var usr User
	err := r.db.Get(&usr, query, sql.Named("email", email), sql.Named("password", pass))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &usr, nil
}

func (r userRepo) Get(uId int) (*User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE id = :id
    `
	var usr User
	err := r.db.Get(&usr, query, sql.Named("id", uId))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &usr, nil
}

func (r userRepo) Update(usr User) (*User, error) {
	query := `
        UPDATE users
        SET first_name = :first_name, last_name = :last_name, email = :email, password = :password, is_shop_owner = :is_shop_owner
        WHERE id = :id
    `
	_, err := r.db.NamedExec(query, usr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &usr, nil
}

func (r userRepo) Delete(uId int) error {
	query := `
        DELETE FROM users
        WHERE id = :id
    `
	_, err := r.db.NamedExec(query, sql.Named("id", uId))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
