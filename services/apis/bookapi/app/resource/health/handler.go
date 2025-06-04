package health

import "net/http"

func Read(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("."))
}
