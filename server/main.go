package main

func main() {

	app := App{}
	app.Init()
	app.Run(":8080")

}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {

// 	router := mux.NewRouter()
// 	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "selam")
// 	})
// 	addr := ":8080"
// 	fmt.Println("listening on " + addr)
// 	http.ListenAndServe(addr, router)

// }
