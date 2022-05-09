package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
	"github.com/webtoor/go-rest-api/config"
	"github.com/webtoor/go-rest-api/controller"
	"github.com/webtoor/go-rest-api/helper"
	"github.com/webtoor/go-rest-api/model/web"
	"github.com/webtoor/go-rest-api/repository"
	"github.com/webtoor/go-rest-api/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME_TEST = "root"
const DB_PASSWORD = "Rahasia123"
const DB_NAME = "go_rest_api_test"
const DB_HOST = "192.168.1.6"
const DB_PORT = "3306"

func setupTestDB() *gorm.DB {
	dsn := DB_USERNAME_TEST + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}

func createTestApp() *fiber.App {
	db := setupTestDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userRoleRepository := repository.NewUserRoleRepository()
	userService := service.NewUserService(&userRepository, &userRoleRepository, db, validate)
	userController := controller.NewUserController(userService)
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	userController.Route(app)
	return app
}

var app = createTestApp()

func TestCreateUserSuccess(t *testing.T) {
	userCreateRequest := web.UserCreateRequest{
		Email:    "webtoor@email.com",
		Password: "RAHASIA",
	}
	requestBody, _ := json.Marshal(userCreateRequest)

	request := httptest.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	jsonResponse := web.JsonResponse{}
	json.Unmarshal(responseBody, &jsonResponse)
	assert.Equal(t, 201, jsonResponse.Code)
	assert.Equal(t, "CREATED", jsonResponse.Status)

	jsonData, _ := json.Marshal(jsonResponse.Data)
	userCreateResponse := web.UserCreateResponse{}
	json.Unmarshal(jsonData, &userCreateResponse)
	assert.NotNil(t, userCreateResponse.Id)
	assert.Equal(t, userCreateRequest.Email, userCreateResponse.Email)
}

func TestCreateUserFailed(t *testing.T) {
	userCreateRequest := web.UserCreateRequest{
		Email:    "",
		Password: "",
	}
	requestBody, _ := json.Marshal(userCreateRequest)

	request := httptest.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	jsonResponse := web.JsonResponse{}
	json.Unmarshal(responseBody, &jsonResponse)
	assert.Equal(t, 400, jsonResponse.Code)
	assert.Equal(t, "BAD REQUEST", jsonResponse.Status)
}
