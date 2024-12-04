package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Serve(portSource, portDest string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Validate the Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, err := validateToken(token, portDest)
		if err != nil {
			log.Printf("Error validating token: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if !isValid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Forward the request to the destination port
		destURL := fmt.Sprintf("http://localhost:%s%s", portDest, r.URL.Path)
		req, err := http.NewRequest(r.Method, destURL, r.Body)
		if err != nil {
			log.Printf("Error creating request: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Copy headers
		for key, values := range r.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// Send the request to the destination
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error forwarding request: %v\n", err)
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copy the response to the client
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	log.Printf("Starting server on port %s...\n", portSource)
	return http.ListenAndServe(fmt.Sprintf(":%s", portSource), nil)
}

func validateToken(token, port string) (bool, error) {
	db, err := OpenDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := `SELECT COUNT(*) FROM tokens WHERE token = ? AND port = ?`
	var count int
	err = db.QueryRow(query, token, port).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return count > 0, nil
}
