package storage

import "backend/models"

type MockStorage struct{}

var user = &models.User{
	ID: "12345",
	FirstName: "Test firstname",
	LastName: "Test lastname",
}

func (m *MockStorage) CreateUser(user models.User) (*models.User, error) {
	user.ID = "12345"
	return &user, nil
}

func (m *MockStorage) GetAllUsers(page, limit int) ([]*models.User, error) {
	return []*models.User{user, user, user}, nil
}

func (m *MockStorage) UpdateUserById(userID string, body models.User) (*models.User, error) {
	return &body, nil
}

func (m *MockStorage) DeleteUserById(userID string) (*models.User, error) {
	return user, nil
}

func (m *MockStorage) GetUserById(userID string) (*models.User, error) {
	return user, nil
}