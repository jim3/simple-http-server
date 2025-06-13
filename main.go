package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p>Simple HHTP Server</p>
</body>
</html>
`
	w.Write([]byte(page))
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlePage)

	// Create a server instance and initialize the `Server` struct fields
	const addr = ":8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
