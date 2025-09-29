package repository

import (
	"fmt"

	"github.com/alamgir-ahosain/e-commerce-project/internal/models"
	"github.com/jmoiron/sqlx"
)

// Repository Design Pattern
type UserRepo interface {
	CreateUserFunc(u models.User) (*models.User, error)
	LoginUserFunc(login models.RequestLogin) (*models.User, error)
}

type userRepo struct {
	//userList []models.User
	db *sqlx.DB
}

// Public constructor
func NewUserRepo(db *sqlx.DB) UserRepo {
	repo := &userRepo{
		db: db,
	}
	//initialUser(repo)
	return repo
}

// CreateUserFunc implements UserRepo.
func (ur *userRepo) CreateUserFunc(u models.User) (*models.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
	          values(:first_name,:last_name,:email,:password,:is_shop_owner) RETURNING id`

	//Executed Query
	var userId int
	rows, err := ur.db.NamedQuery(query, u)
	if err != nil {
		fmt.Println("executed query error: ", err)
		return nil, err
	}
	defer rows.Close()

	//Read the returned ID
	if rows.Next() {
		if err := rows.Scan(&userId); err != nil { //If no rows are returned, userId remains 0
			fmt.Println("rows error:", err)
			return nil, err
		}
	}
	//Assign the ID and return
	u.ID = userId
	return &u, nil
}

// LoginUserFunc implements UserRepo: using the database.
func (ur *userRepo) LoginUserFunc(login models.RequestLogin) (*models.User, error) {
	if login.Email == "" || login.Password == "" {
		return nil, fmt.Errorf("email and password required")
	}

	var user models.User
	query := `SELECT * FROM users WHERE email = $1 and password=$2`
	err := ur.db.Get(&user, query, login.Email,login.Password)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("user not found")
	}

	// Compare password
	if user.Password != login.Password {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}
