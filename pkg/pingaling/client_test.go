package pingaling

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	var client Client
	var s interface{}
	s = &Session{}
	session, err := client.CreateSession()
	assert.Nil(t, err)
	assert.IsType(t, s, session)
}

func TestURLBase(t *testing.T) {
	var client Client
	URL := client.urlBase("foo")
	assert.Equal(t, "http://localhost:4000/api/foo", URL)
}

func TestDoReqURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c := Client{
		BaseURL: "http://localhost/api",
	}

	headers := map[string]string{
		"foo": "bar",
		"zoo": "cat",
	}
	statusCode, b, _ := c.doReqURL(ctx, http.MethodGet, ts.URL, headers, nil)

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, "{\"fake json string\"}\n", b.String())

}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}

	b, err := c.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, "{\"fake json string\"}\n", b.String())

}

func TestGetError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}
	b, err := c.Get(ts.URL)
	assert.NotNil(t, err)
	assert.Empty(t, b)

}

func TestDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}

	b, err := c.Delete(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, "{\"fake json string\"}\n", b.String())

}

func TestDeleteError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}
	b, err := c.Delete(ts.URL)
	assert.NotNil(t, err)
	assert.Empty(t, b)

}

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}
	body := []byte(`{"foo": "bar"}`)
	b, err := c.Post(ts.URL, bytes.NewBuffer(body))
	assert.Nil(t, err)
	assert.Equal(t, b.String(), "{\"fake json string\"}\n")

}

func TestPostError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	c := Client{
		BaseURL: "http://localhost/api",
	}
	body := []byte(`{"foo": "bar"}`)
	b, err := c.Post(ts.URL, bytes.NewBuffer(body))
	assert.NotNil(t, err)
	assert.Empty(t, b)

}
