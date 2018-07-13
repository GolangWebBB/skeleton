package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/husobee/vestigo"
	log "github.com/sirupsen/logrus"
)

func ucFirstNameMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// before
		name := strings.Title(vestigo.Param(r, "name"))
		params := strings.Split(r.URL.RawQuery, "&")
		paramName := url.QueryEscape(":name")

		for i, param := range params {
			split := strings.Split(param, "=")
			if split[0] == paramName {
				params[i] = fmt.Sprintf("%s=%s", split[0], name)
				break
			}
		}

		r.URL.RawQuery = strings.Join(params, "&")

		fmt.Println(r.URL.RawQuery)

		handler(w, r)

		// after
		// (do nothing)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name := vestigo.Param(r, "name")
	w.WriteHeader(200)
	w.Write([]byte("Hello " + name + "!"))
}

func main() {
	router := vestigo.NewRouter()

	// Useful CORS settings
	// router.SetGlobalCors(&vestigo.CorsAccessControl{
	// 	AllowOrigin:      []string{"*", "test.com"},
	// 	AllowCredentials: true,
	// 	ExposeHeaders:    []string{"X-Header", "X-Y-Header"},
	// 	MaxAge:           3600 * time.Second,
	// 	AllowHeaders:     []string{"X-Header", "X-Y-Header"},
	// })

	router.Get("/hello/:name", helloWorld, ucFirstNameMiddleware)

	log.Fatal(http.ListenAndServe(":1234", router))
}
