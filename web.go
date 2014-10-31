package gofizzbuzz

import (
	"fmt"
	"net/http"
	"strconv"
)

var tmpl = `<!DOCTYPE html>
<html>
	<head>
		<title>FizzBuzz</title>
	</head>
	<body>
		<span>Say: %s</span>
	</body>
</html>
`

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} else {
		if err := fizzBuzzHandlerInternal(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func fizzBuzzHandlerInternal(w http.ResponseWriter, r *http.Request) error {
	var nstr = r.FormValue("number")
	var n, err = strconv.Atoi(nstr)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(tmpl, FizzBuzz(n).Say())))

	return nil
}
