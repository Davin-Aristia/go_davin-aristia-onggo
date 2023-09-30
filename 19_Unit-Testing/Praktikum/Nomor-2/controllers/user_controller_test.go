package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"test/mvc/config"
	"test/mvc/controllers/responses"
	"test/mvc/controllers/mocking"
	"test/mvc/models"
	"testing"
	"strings"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() (models.User, error) {
	user := models.User{
		Name:     "Davin",
		Password: "123",
		Email:    "davin@gmail.com",
	}

	var err error
	if err = config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all user",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Users   []responses.UserResponse `json:"users"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(responseData.Users))
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "get user by id found",
			id:         1,
			path:       "/users",
			expectCode: http.StatusOK,
		},
		{
			name:       "get user by id not found",
			id:			2,
			path:       "/users",
			expectCode: http.StatusInternalServerError,
		},
	}

	e := InitEchoTestAPI()
	dummy, _ := InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				User    responses.UserResponse `json:"user"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.User.ID, dummy.ID)
				assert.Equal(t, responseData.User.Name, dummy.Name)
				assert.Equal(t, responseData.User.Email, dummy.Email)
			} else {
				assert.Equal(t, responseData.Message, "error get data")
			}
		}
	}
}

func TestPostUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		user	   mocking.User
		expectCode int
	}{
		{
			name:       "create new user",
			path:       "/users",
			user:		mocking.User{
							Name:     "Davin2",
							Password: "123456",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "create new user invalid password",
			path:       "/users",
			user:		mocking.User{
							Name:     "Davin Error",
							Password: 123,
							Email:    "davinError@gmail.com",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				User   responses.UserResponse    `json:"user"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.User.Name, testCase.user.Name)
				assert.Equal(t, responseData.User.Email, testCase.user.Email)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			}
		}
	}
}

func TestPutUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		id         int
		user	   mocking.User
		expectCode int
	}{
		{
			name:       "update user",
			path:       "/users",
			id:			1,
			user:		mocking.User{
							Name:     "Davin2",
							Password: "123456",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "update user id not found",
			path:       "/users",
			id:			2,
			user:		mocking.User{
							Name:     "Davin2",
							Password: "123456",
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusNotFound,
		},
		{
			name:       "update user invalid password",
			path:       "/users",
			id:			2,
			user:		mocking.User{
							Name:     "Davin2",
							Password: 123,
							Email:    "davin2@gmail.com",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		userJSON, _ := json.Marshal(testCase.user)
		
		req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(userJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				User   responses.UserResponse    `json:"user"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.User.Name, testCase.user.Name)
				assert.Equal(t, responseData.User.Email, testCase.user.Email)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			}else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no user found for id "+strconv.Itoa(testCase.id))
				// assert.Equal(t, responseData.Message, "record not found")
			}
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "delete user by id found",
			id:         1,
			path:       "/users",
			expectCode: http.StatusOK,
		},
		{
			name:       "delete user by id not found",
			id:         2,
			path:       "/users",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				User    responses.UserResponse `json:"user"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Message, "success delete user with id " + strconv.Itoa(testCase.id))
			} else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no user found for id "+strconv.Itoa(testCase.id))
				// assert.Equal(t, responseData.Message, "record not found")
			} else {
				assert.Equal(t, responseData.Message, "error delete data")
			}
		}
	}
}