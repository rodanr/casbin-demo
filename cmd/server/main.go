package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rodanr/casbin-demo/internal/authz"
	"github.com/rodanr/casbin-demo/internal/httpx"
)

func main() {
	e, err := authz.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("new enforcer: %v", err)
	}

	mux := http.NewServeMux()

	// /docs (GET/POST)
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintln(w, "list docs (stub)")
		case http.MethodPost:
			fmt.Fprintln(w, "create doc (stub)")
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /docs/:id (GET/PUT/DELETE)
	mux.HandleFunc("/docs/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/docs/")
		if id == "" {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "get doc %s (stub)\n", id)
		case http.MethodPut:
			fmt.Fprintf(w, "update doc %s (stub)\n", id)
		case http.MethodDelete:
			fmt.Fprintf(w, "delete doc %s (stub)\n", id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Wrap with Casbin authorization
	handler := httpx.Authorize(e)(mux)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
