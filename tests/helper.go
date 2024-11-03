package tests

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/vincentandreas/dealls/api/handlers"
	"github.com/vincentandreas/dealls/models"
	"gorm.io/gorm"

	"github.com/vincentandreas/dealls/utilities"
)

func ReadRegisterCSV(filePath string) ([]models.UserRegister, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var users []models.UserRegister

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		user := models.UserRegister{
			Name:           record[0],
			Email:          record[1],
			Password:       record[2],
			Gender:         record[3],
			InterestGender: record[4],
			Hobby:          record[5],
		}

		users = append(users, user)
	}

	return users, nil
}

func CleanDb(db *gorm.DB){
	tblNames := []string{
		"histories",
        "user_preferences",
        "users",
	}
	for _, name := range tblNames {
        err := db.Exec("DROP TABLE IF EXISTS " + name).Error
		if err!= nil {
			log.Fatal(err)
		}
    }
}

func PopulateData() {
	maleRegister, _ := ReadRegisterCSV("../testdata/register_male.csv")
	femaleRegister, _ := ReadRegisterCSV("../testdata/register_female.csv")	

	router,_ := handlers.SetupServer()
	
	userJson, err := json.Marshal(maleRegister[0])
	utilities.CheckError(err)
	
	req, err := http.NewRequest("POST","/register",  bytes.NewBuffer(userJson))
	utilities.CheckError(err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	for _, r := range femaleRegister {
		userJson, err := json.Marshal(r)
		utilities.CheckError(err)
		
		req, err := http.NewRequest("POST","/register",  bytes.NewBuffer(userJson))
		utilities.CheckError(err)
	
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}