package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, nil // User not found
		}
		return model.User{}, result.Error // Other error occurred
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	userTaskCategories := []model.UserTaskCategory{}
	err := r.db.Raw("SELECT u.id AS id, u.fullname AS fullname, u.email AS email, t.title AS task, t.deadline AS deadline, t.priority AS priority, t.status AS status, c.name AS category FROM users u, tasks t, categories c WHERE  u.id = t.user_id AND t.category_id = c.id").Scan(&userTaskCategories).Error
	if err != nil {
		return userTaskCategories, err
	}
	return userTaskCategories, nil // TODO: replace this
}
