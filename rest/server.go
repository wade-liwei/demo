package rest

import (
	"fmt"
	"log"
	"net/http"
	//"strconv"

	"github.com/gorilla/mux"
	"github.com/wade-liwei/demo/rest/swagger"
)

//WebServer init and return the web server instance.
func WebServer() {
	r := mux.NewRouter()
	swagger.RegisterSwaggerUI(r)
	//Routes consist of a path and a handler function.
	// r.HandleFunc("/txs/{page}/{size}", TxsHandler)
	// r.HandleFunc("/tx/{hash}", TxByHashHandler)
	// Bind to a port and pass our router in
	fmt.Println("listen: 0.0.0.0:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
