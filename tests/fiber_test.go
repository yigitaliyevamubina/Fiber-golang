package tests

import (
	"backend/models"
	"backend/tests/handlers"
	"backend/tests/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberAPI(t *testing.T) {
	app := fiber.New()

	h := handlers.MockHandlers{MockUserService: storage.MockStorage{}}

	requestBody := `{
		"id": "12345",
		"first_name": "Test firstname",
		"last_name": "Test lastname"
	}`

	//Create user
	req := httptest.NewRequest(http.MethodPost, "/users/create", strings.NewReader(requestBody))
	req.Header.Set("Content-type", "application/json")

	app.Post("/users/create", h.CreateUser)

	res, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	responseBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var user models.User
	assert.NoError(t, json.Unmarshal(responseBody, &user))
	assert.Equal(t, "Test firstname", user.FirstName)
	assert.Equal(t, "Test lastname", user.LastName)
	user.ID = "12345"

	//Get user
	getReq := httptest.NewRequest(http.MethodGet, "/users/get", nil)

	app.Get("/users/get", h.GetIdByPath)

	q := getReq.URL.Query()
	q.Add("id", "12345")
	getReq.URL.RawQuery = q.Encode()

	getRes, err := app.Test(getReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRes.StatusCode)

	responseBody, err = io.ReadAll(getRes.Body)
	assert.NoError(t, err)

	var getUserResp models.User
	assert.NoError(t, json.Unmarshal(responseBody, &getUserResp))
	assert.Equal(t, "12345", getUserResp.ID)
	assert.Equal(t, "Test firstname", getUserResp.FirstName)
	assert.Equal(t, "Test lastname", getUserResp.LastName)

	//List users
	listReq := httptest.NewRequest(http.MethodGet, "/users/list", nil)
	app.Get("/users/list", h.GetUsers)

	q = listReq.URL.Query()
	q.Add("page", "1")
	q.Add("limit", "10")
	listReq.URL.RawQuery = q.Encode()

	listRes, err := app.Test(listReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, listRes.StatusCode)

	responseBody, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, responseBody)

	//Update user
	user.FirstName = "Updated firstname"
	user.LastName = "Updated lastname"
	bodyBytes, err := json.Marshal(user)
	assert.NoError(t, err)

	updateReq := httptest.NewRequest(http.MethodPut, "/users/update", strings.NewReader(string(bodyBytes)))
	updateReq.Header.Set("Content-type", "application/json")
	app.Put("/users/update", h.UpdateUser)

	q = updateReq.URL.Query()
	q.Add("id", user.ID)
	updateReq.URL.RawQuery = q.Encode()

	updateRes, err := app.Test(updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updateRes.Body)

	responseBody, err = io.ReadAll(updateRes.Body)
	assert.NoError(t, err)

	var updatedUser models.User
	assert.NoError(t, json.Unmarshal(responseBody, &updatedUser))
	assert.Equal(t, user.ID, updatedUser.ID)
	assert.Equal(t, "Updated firstname", updatedUser.FirstName)
	assert.Equal(t, "Updated lastname", updatedUser.LastName)

	//Delete user
	delReq := httptest.NewRequest(http.MethodDelete, "/users/delete", nil)
	app.Delete("/users/delete", h.DeleteUser)

	q = delReq.URL.Query()
	q.Add("id", user.ID)
	delReq.URL.RawQuery = q.Encode()

	delRes, err := app.Test(delReq)
	assert.NoError(t, err)
	assert.NotNil(t, delReq.Body)

	responseBody, err = io.ReadAll(delRes.Body)
	assert.NoError(t, err)

	var deletedUser models.User
	assert.NoError(t, json.Unmarshal(responseBody, &deletedUser))
	assert.Equal(t, user.ID, deletedUser.ID)
	assert.Equal(t, "Test firstname", deletedUser.FirstName)
	assert.Equal(t, "Test lastname", deletedUser.LastName)

}
