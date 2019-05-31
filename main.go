package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zserge/webview"
)

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("Error http.listen: %v", err)
	}
	fs := http.FileServer(http.Dir("app/"))
	http.Handle("/app/", http.StripPrefix("/app/", fs))

	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to my website!")
		})
		log.Fatal(http.Serve(ln, nil))
	}()

	return "http://" + ln.Addr().String()
}

func main() {
	url := startServer() + "/app/game.html"
	fmt.Println("Serving game at: " + url)
	w := webview.New(webview.Settings{
		Width:  640,
		Height: 480,
		Title:  "Dark Reign",
		URL:    url,
	})
	defer w.Exit()
	w.Run()
}
