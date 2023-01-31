package checks

import (
	"io"
	"log"
	"net/http"
)

type Err struct {
	MSG string
}

func (e Err) Error() string {
	return e.MSG
}

func IsPOST(w *http.ResponseWriter, method, msg string) bool {
	if method == "POST" {
		return true
	}
	(*w).WriteHeader(http.StatusBadRequest)
	log.Println(msg)
	return false
}

func ReturnIntServErr(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte(err.Error()))
	log.Println(err)
}

func ReadContent(w *http.ResponseWriter, body io.ReadCloser) ([]byte, bool) {
	status := true
	content, err := io.ReadAll(body)
	if err != nil {
		ReturnIntServErr(w, err)
		status = false
	}
	return content, status
}

func IsPOSTAndGetContent(w *http.ResponseWriter, r *http.Request, msg string) ([]byte, bool) {
	ok := true
	content := []byte("")
	if ok = IsPOST(w, r.Method, msg); !ok {
		return content, ok
	}
	content, ok = ReadContent(w, r.Body)
	return content, ok
}
