package router

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	. "github.com/evandroad/gomyadm/internal/respond"
)

var Debug = false

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[DKM] %s | %sPANIC%s %s %s: %v\n",
					time.Now().Format("06/01/02 15:04:05.000"),
					red, reset, r.Method, r.URL.Path, err)
				if Debug {
					fmt.Printf("%s", debug.Stack())
				}
				Error(w, http.StatusInternalServerError, "Internal Server Error", nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}