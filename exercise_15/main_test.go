package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestSourceCodeHandler test
func TestSourceCodeHandler(t *testing.T) {
	t.Run("It returns status 200", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/debug/?line=string&path=%2Fhome%2Fsagar%2F.gvm%2Fgos%2Fgo1.13.4%2Fsrc%2Fruntime%2Fdebug%2Fstack.go", nil)
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
	})

	t.Run("It returns status 500 when file does not exists", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/debug/?line=24&path=path_does_not_exists", nil)
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)

		assert.Equal(t, 500, response.Code)
	})

	t.Run("It returns status 500 when only folder path given instead file path", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/debug/?line=24&path=/home/sagar/.gvm/pkgsets/go1.13.4/global/src/gophercises/exercise_15", nil)
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)

		assert.Equal(t, 500, response.Code)
	})

	t.Run("it return 200 when line number is not int", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/debug/?line=24&path=%2Fhome%2Fsagar%2F.gvm%2Fgos%2Fgo1.13.4%2Fsrc%2Fruntime%2Fdebug%2Fstack.go", nil)
		response := httptest.NewRecorder()
		Router().ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
	})
}

func TestPanicDemo(t *testing.T) {
	request, _ := http.NewRequest("GET", "/panic/", nil)
	response := httptest.NewRecorder()

	assert.Panics(t, func() { PanicDemo(response, request) }, "The code did not panic")
}

func TestMakeLinks(t *testing.T) {
	path := "/home/sagar/.gvm/pkgsets/go1.13.4/global/src/gophercises/exercise_15/test_stack_data.txt"
	file, _ := os.Open(path)

	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	makeLinks(b.String())
}

func TestStackStraceMiddleware(t *testing.T) {
	req, _ := http.NewRequest("GET", "/panic/", nil)
	res := httptest.NewRecorder()
	var mux = http.NewServeMux()
	mux.HandleFunc("/panic/", PanicDemo)
	handler := stackStraceMiddleware(mux)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("Something went worng! it should throw internal server error.")
	}
}
