package main

import (
    "net/http"
    "log"
    "math/rand"
    "fmt"
    "time"
    "os"
)


type rootHandler struct {
    id int
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    log.Println("got a request")

    simpleEnv := os.Getenv("SIMPLE_ENV")
    config := os.Getenv("CONFIG_VALUE")
    password := os.Getenv("SECRET_PASSWORD")
    w.Write([]byte(fmt.Sprintf(`<html>
<head></head>
<body>
<h1>Response from server %d</h1>
<h2>Simple Env: %s</h2>
<h2>Config: %s</h2>
<h2>Secret: %s</h2>
<iframe width="560" height="315" src="https://www.youtube.com/embed/fn3KWM1kuAw" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
</body>
</html>`, h.id, simpleEnv, config, password)))
}

func main() {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    id := r1.Intn(100)
    handler := &rootHandler{id}
    http.HandleFunc("/", handler.ServeHTTP)

    log.Println("starting server on port 8080")
    http.ListenAndServe(":8080", nil)
}
