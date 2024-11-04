package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vincentandreas/dealls/api/handlers"
	"github.com/vincentandreas/dealls/models"
	"github.com/vincentandreas/dealls/utilities"
)


var maleRegister []models.UserRegister

func setup() {
    maleRegister, _ = ReadRegisterCSV("../testdata/register_male.csv")
}



func TestMain(m *testing.M) {
    setup()
    m.Run()
}

func Test_register_should_success(t *testing.T) {
	router,db := handlers.SetupServer()
	defer CleanDb(db)
	userJson, err := json.Marshal(maleRegister[0])
	utilities.CheckError(err)
	
	req, err := http.NewRequest("POST","/register",  bytes.NewBuffer(userJson))
	utilities.CheckError(err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Result().StatusCode)
}



func Test_login_wrong_password_should_failed(t * testing.T){
	router,db := handlers.SetupServer()
	defer CleanDb(db)
	
	dataReq := &models.UserLogin{
		Email:     maleRegister[0].Email,
        Password: "wrongpass",
	}
	
	userJson, err := json.Marshal(maleRegister[0])
	utilities.CheckError(err)
	req, err := http.NewRequest("POST","/register",  bytes.NewBuffer(userJson))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)	
	utilities.CheckError(err)


	dataJson, err := json.Marshal(dataReq)
	utilities.CheckError(err)
	req, err = http.NewRequest("POST","/login",  bytes.NewBuffer(dataJson))
	utilities.CheckError(err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Result().StatusCode)
}

func Test_login_should_success(t * testing.T){
	router,db := handlers.SetupServer()
	defer CleanDb(db)
	
	dataReq := &models.UserLogin{
		Email:     maleRegister[0].Email,
        Password: maleRegister[0].Password,
	}
	
	userJson, err := json.Marshal(maleRegister[0])
	utilities.CheckError(err)
	req, err := http.NewRequest("POST","/register",  bytes.NewBuffer(userJson))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)	
	utilities.CheckError(err)


	dataJson, err := json.Marshal(dataReq)
	utilities.CheckError(err)
	req, err = http.NewRequest("POST","/login",  bytes.NewBuffer(dataJson))
	utilities.CheckError(err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Result().StatusCode)
}