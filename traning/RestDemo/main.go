/*package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloHandler)

	log.Println("Listening...")
	http.ListenAndServe("localhost:9000", r)
}

func HelloHandler(resp http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(resp, "hi i am here!")
}
*/
package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
type person struct {
    Name     string
    LastName string
    Age      uint8
}
 
func sendResponse(response http.ResponseWriter, request *http.Request) {
    person := person{Name: "Monika", LastName: "Tiwari", Age: 22}
 
    jsonResponse, jsonError := json.Marshal(person)
 
    if jsonError != nil {
        fmt.Println("Unable to encode JSON")
    }
 
    fmt.Println(string(jsonResponse))
 
    response.Header().Set("Content-Type", "application/json")
    response.WriteHeader(http.StatusOK)
    response.Write(jsonResponse)
}
 
func main() {
    route := mux.NewRouter()
 
    route.HandleFunc("/", sendResponse)
 
    http.ListenAndServe(":8000", route)
}