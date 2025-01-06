package token_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	router "github.com/gody-server/router"
	token "github.com/gody-server/adapters/jwt"

	/*
		- Uncomment this when you are testing real data coming from database.
		db "github.com/app/gody-server/infraestructure"
	*/
)

// Setup and Teardown
func setup(t *testing.T) func(t *testing.T) {
	// Setup
	t.Log("setup sub test")

	// For test db
	t.Setenv("ENV", "TEST")

	/*
		- Uncomment this when you are testing real data coming from test database.
		db.Connect()
	*/

	// Teardown
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}

func TestGetToken(t *testing.T) {
	// Call Setup/Teardown
	teardown := setup(t)
	defer teardown(t)

	tokenData := token.GenerateToken("test") //Your token data

	router := router.Router()

	w := httptest.NewRecorder()

	values := map[string]interface{}{"token": tokenData} // this is the body in case you make a post, put
	jsonValue, _ := json.Marshal(values)

	req, _ := http.NewRequest("GET", "/token", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenData)

	// In case you use cookies like for example token
	req.Header.Set("Cookie", "token="+tokenData+";")

	router.ServeHTTP(w, req) 

	expected := `{"data":[{"Id":1}]}` // Your expected data inside backquote 
	expectedStatus := "200 OK"

	assert.Contains(t, w.Body.String(), expected, "🔴 Expected %v 🔴 got %v", expected, w.Body.String())
	assert.Contains(t, w.Result().Status, expectedStatus, "🔴 Expected %v 🔴 got %v", expectedStatus, w.Result().Status)
	fmt.Println("🟢")
}