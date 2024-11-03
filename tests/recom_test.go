package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/vincentandreas/dealls/api/handlers"
// 	"github.com/vincentandreas/dealls/models"
// 	"github.com/vincentandreas/dealls/utilities"
// )


// func Test_should_get_different_recom(t *testing.T) {
// 	PopulateData()
// 	router,db := handlers.SetupServer()
// 	defer CleanDb(db)

// 	req, err := http.NewRequest("GET","/register",  bytes.NewBuffer(userJson))
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)	
// 	utilities.CheckError(err)


// }
