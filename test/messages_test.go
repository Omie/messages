package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/omie/messages/api"
	_ "github.com/omie/messages/api/messages"
	"github.com/omie/messages/lib/db"
)

type CreateMessageResponse struct {
	Id string `json:"id"`
}

var (
	server         *httptest.Server
	reader         io.Reader
	messagesUrl    string
	messageText    string
	createdMessage *CreateMessageResponse
)

func init() {
	// log.SetLevel(log.WarnLevel)

	// get parameters and initialize the database
	dbDriver := os.Getenv("MESSAGES_DB_DRIVER")
	dbName := os.Getenv("MESSAGES_DB_NAME")
	if dbDriver == "" || dbName == "" {
		log.Error("main: db driver or db name not set")
		return
	}

	if err := db.InitDB(dbDriver, dbName); err != nil {
		log.Error("Error initializing database", err.Error())
		return
	}

	server = httptest.NewServer(api.Container)
	messagesUrl = fmt.Sprintf("%s/messages", server.URL)
	messageText = "sample text"
}

// Test that message is saved properly
func TestCreateMessage(t *testing.T) {
	data := url.Values{}
	data.Set("text", messageText)

	u, _ := url.ParseRequestURI(messagesUrl)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)

	// save received createdMessage uuid to test get message operation
	createdMessage = new(CreateMessageResponse)
	json.NewDecoder(resp.Body).Decode(createdMessage)
	fmt.Println(createdMessage)

	if createdMessage.Id == "" {
		t.Errorf("Expected UUID to be not empty")
	}
	if resp.StatusCode != 201 {
		t.Errorf("Expected 201 to match: %d", resp.StatusCode)
	}

}

// test that message is retrieved properly
func TestGetMessage(t *testing.T) {
	u, _ := url.ParseRequestURI(messagesUrl)
	u.Path = "/messages/" + createdMessage.Id
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println(urlStr)
	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)

	resp, _ := client.Do(r)

	if resp.StatusCode != 200 {
		t.Errorf("Expected 200 to match: %d", resp.StatusCode)
	}

	if txt, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fail()
	} else {
		if string(txt) != messageText {
			t.Errorf("invalid message returned:", txt)
		}
	}
}

// test that invalid message id is returned as 404
func TestInvalidMessageId(t *testing.T) {
	u, _ := url.ParseRequestURI(messagesUrl)
	u.Path = "/messages/12345"
	urlStr := fmt.Sprintf("%v", u)
	fmt.Println(urlStr)
	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)

	resp, _ := client.Do(r)

	if resp.StatusCode != 404 {
		t.Errorf("Expected 404 to match: %d", resp.StatusCode)
	}
}

// Test that unexpected parameter in request result in bad request
func TestUnexpectedParamInRequest(t *testing.T) {
	data := url.Values{}
	data.Set("text", messageText)
	data.Set("unexpected", "blah")

	u, _ := url.ParseRequestURI(messagesUrl)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)

	if resp.StatusCode != 400 {
		t.Errorf("Expected 400 to match: %d", resp.StatusCode)
	}

}

// Test that empty value for text parameter in request result in bad request
func TestEmptyParamInRequest(t *testing.T) {
	data := url.Values{}
	data.Set("text", "")

	u, _ := url.ParseRequestURI(messagesUrl)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)

	if resp.StatusCode != 400 {
		t.Errorf("Expected 400 to match: %d", resp.StatusCode)
	}

}
