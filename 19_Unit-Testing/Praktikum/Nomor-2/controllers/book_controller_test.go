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

func InsertDataBookForGetBooks() (models.Book, error) {
	book := models.Book{
		Title:     "Harry Potter",
		Writer:    "J.K. Rowling",
		Publisher: "Warner Bros",
	}

	var err error
	if err = config.DB.Create(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all book",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Books   []responses.BookResponse `json:"books"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(responseData.Books))
		}
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "get book by id found",
			id:         1,
			path:       "/books",
			expectCode: http.StatusOK,
		},
		{
			name:       "get book by id not found",
			id:			2,
			path:       "/books",
			expectCode: http.StatusInternalServerError,
		},
	}

	e := InitEchoTestAPI()
	dummy, _ := InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				Book    responses.BookResponse `json:"book"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Book.ID, dummy.ID)
				assert.Equal(t, responseData.Book.Title, dummy.Title)
				assert.Equal(t, responseData.Book.Writer, dummy.Writer)
				assert.Equal(t, responseData.Book.Publisher, dummy.Publisher)
			} else {
				assert.Equal(t, responseData.Message, "error get data")
			}
		}
	}
}

func TestPostBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		book	   mocking.Book
		expectCode int
	}{
		{
			name:       "create new book",
			path:       "/books",
			book:		mocking.Book{
							Title:     "Harry Potter 2",
							Writer:    "J.K. Rowling 2",
							Publisher: "Warner Bros 2",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "create new book invalid writer",
			path:       "/books",
			book:		mocking.Book{
							Title:     "Harry Potter Error",
							Writer:    false,
							Publisher: "Warner Bros Error",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()

	for _, testCase := range testCases {
		bookJSON, _ := json.Marshal(testCase.book)
		
		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(bookJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Book   responses.BookResponse    `json:"book"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Book.Title, testCase.book.Title)
				assert.Equal(t, responseData.Book.Writer, testCase.book.Writer)
				assert.Equal(t, responseData.Book.Publisher, testCase.book.Publisher)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			}
		}
	}
}

func TestPutBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		id         int
		book	   mocking.Book
		expectCode int
	}{
		{
			name:       "update book",
			path:       "/books",
			id:			1,
			book:		mocking.Book{
							Title:     "Harry Potter 3",
							Writer:    "J.K. Rowling 3",
							Publisher: "Warner Bros 3",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "update book id not found",
			path:       "/books",
			id:			2,
			book:		mocking.Book{
							Title:     "Harry Potter 3",
							Writer:    "J.K. Rowling 3",
							Publisher: "Warner Bros 3",
						},
			expectCode: http.StatusNotFound,
		},
		{
			name:       "update book invalid writer",
			path:       "/books",
			id:			2,
			book:		mocking.Book{
							Title:     "Harry Potter 3",
							Writer:    true,
							Publisher: "Warner Bros 3",
						},
			expectCode: http.StatusBadRequest,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		bookJSON, _ := json.Marshal(testCase.book)
		
		req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(bookJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Book   responses.BookResponse    `json:"book"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Book.Title, testCase.book.Title)
				assert.Equal(t, responseData.Book.Writer, testCase.book.Writer)
				assert.Equal(t, responseData.Book.Publisher, testCase.book.Publisher)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			} else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no book found for id "+strconv.Itoa(testCase.id))
				// assert.Equal(t, responseData.Message, "record not found")
			}
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "delete book by id found",
			id:         1,
			path:       "/books",
			expectCode: http.StatusOK,
		},
		{
			name:       "delete book by id not found",
			id:         2,
			path:       "/books",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				Book    responses.BookResponse `json:"book"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Message, "success delete book with id " + strconv.Itoa(testCase.id))
			} else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no book found for id "+strconv.Itoa(testCase.id))
				// assert.Equal(t, responseData.Message, "record not found")
			} else {
				assert.Equal(t, responseData.Message, "error delete data")
			}
		}
	}
}