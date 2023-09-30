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

func InsertDataBlogForGetBlogs() (models.Blog, error) {
	blog := models.Blog{
		UserId:  1,
		Title:   "Jalan Pagi",
		Content: "Bangun Pagi untuk Jalan Pagi merupakan hal yang menyenangkan",
	}

	var err error
	if err = config.DB.Create(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func TestGetBlogsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get all blog",
			path:       "/blogs",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	InsertDataBlogForGetBlogs()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBlogsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Blogs   []responses.BlogResponse `json:"blogs"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(responseData.Blogs))
		}
	}
}

func TestGetBlogController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "get blog by id found",
			id:         1,
			path:       "/blogs",
			expectCode: http.StatusOK,
		},
		{
			name:       "get blog by id not found",
			id:			2,
			path:       "/blogs",
			expectCode: http.StatusInternalServerError,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	dummy, _ := InsertDataBlogForGetBlogs()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, GetBlogController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				Blog    responses.BlogResponse `json:"blog"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Blog.ID, dummy.ID)
				assert.Equal(t, responseData.Blog.UserId, dummy.UserId)
				assert.Equal(t, responseData.Blog.Title, dummy.Title)
				assert.Equal(t, responseData.Blog.Content, dummy.Content)
			} else {
				assert.Equal(t, responseData.Message, "error get data")
			}
		}
	}
}

func TestPostBlogController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		blog	   mocking.Blog
		expectCode int
	}{
		{
			name:       "create new blog",
			path:       "/blogs",
			blog:		mocking.Blog{
							UserId:  1,
							Title:   "Jalan Sore",
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "create new blog invalid title",
			path:       "/blogs",
			blog:		mocking.Blog{
							UserId:  1,
							Title:   1,
							Content: "Title seharusnya bertipe data string",
						},
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "create new blog user id not found",
			path:       "/blogs",
			blog:		mocking.Blog{
							UserId:  2,
							Title:   "Jalan Sore",
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		blogJSON, _ := json.Marshal(testCase.blog)
		
		req := httptest.NewRequest(http.MethodPost, "/blogs", strings.NewReader(string(blogJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBlogController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Blog   responses.BlogResponse    `json:"blog"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Blog.UserId, testCase.blog.UserId)
				assert.Equal(t, responseData.Blog.Title, testCase.blog.Title)
				assert.Equal(t, responseData.Blog.Content, testCase.blog.Content)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			} else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no user found for id "+strconv.Itoa(int(testCase.blog.UserId)))
				// assert.Equal(t, responseData.Message, "record not found")
			}
		}
	}
}

func TestPutBlogController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		id         int
		blog	   mocking.Blog
		expectCode int
	}{
		{
			name:       "update blog",
			path:       "/blogs",
			id:			1,
			blog:		mocking.Blog{
							UserId:  1,
							Title:   "Jalan Sore",
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusOK,
		},
		{
			name:       "update blog id not found",
			path:       "/blogs",
			id:			2,
			blog:		mocking.Blog{
							UserId:  1,
							Title:   "Jalan Sore",
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusNotFound,
		},
		{
			name:       "update blog invalid title",
			path:       "/blogs",
			id:			2,
			blog:		mocking.Blog{
							UserId:  1,
							Title:   false,
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "create new blog user id not found",
			path:       "/blogs",
			blog:		mocking.Blog{
							UserId:  2,
							Title:   "Jalan Sore",
							Content: "Bangun Sore untuk Jalan Sore merupakan hal yang menyenangkan",
						},
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	InsertDataBlogForGetBlogs()

	for _, testCase := range testCases {
		blogJSON, _ := json.Marshal(testCase.blog)
		
		req := httptest.NewRequest(http.MethodPut, "/blogs", strings.NewReader(string(blogJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, UpdateBlogController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()
			
			// open file
			// convert struct
			type Response struct {
				Message string                   `json:"message"`
				Blog   responses.BlogResponse    `json:"blog"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Blog.UserId, testCase.blog.UserId)
				assert.Equal(t, responseData.Blog.Title, testCase.blog.Title)
				assert.Equal(t, responseData.Blog.Content, testCase.blog.Content)
			} else if rec.Code == 400 {
				assert.Equal(t, responseData.Message, "error bind data")
			} else if rec.Code == 404 {
				if strings.Contains(responseData.Message, "blog"){
					assert.Equal(t, responseData.Message, "no blog found for id "+strconv.Itoa(testCase.id))
				} else if strings.Contains(responseData.Message, "user"){
					assert.Equal(t, responseData.Message, "no user found for id "+strconv.Itoa(int(testCase.blog.UserId)))
				}
				// assert.Equal(t, responseData.Message, "record not found")
			}
		}
	}
}

func TestDeleteBlogController(t *testing.T) {
	var testCases = []struct {
		name       string
		id         int
		path       string
		expectCode int
	}{
		{
			name:       "delete blog by id found",
			id:         1,
			path:       "/blogs",
			expectCode: http.StatusOK,
		},
		{
			name:       "delete blog by id not found",
			id:         2,
			path:       "/blogs",
			expectCode: http.StatusNotFound,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	InsertDataBlogForGetBlogs()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(testCase.id))

		if assert.NoError(t, DeleteBlogController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			type Response struct {
				Message string                 `json:"message"`
				Blog    responses.BlogResponse `json:"blog"`
			}
			var responseData Response
			err := json.Unmarshal([]byte(body), &responseData)

			if err != nil {
				assert.Error(t, err, "error")
			}
			if rec.Code == 200 {
				assert.Equal(t, responseData.Message, "success delete blog with id " + strconv.Itoa(testCase.id))
			} else if rec.Code == 404 {
				assert.Equal(t, responseData.Message, "no blog found for id "+strconv.Itoa(testCase.id))
				// assert.Equal(t, responseData.Message, "record not found")
			} else {
				assert.Equal(t, responseData.Message, "error delete data")
			}
		}
	}
}