package main
import (
    "fmt"
        "encoding/json"
        "log"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type User struct {
        Name string
}
type UGreetings struct {
        Greeting string
}


func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func handlePost(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
        res.Header().Set("Content-Type", "application/json")
        user := new(User)
        decoder := json.NewDecoder(req.Body)
        error := decoder.Decode(&user)
        if error != nil {
                        log.Println(error.Error())
                        http.Error(res, error.Error(), http.StatusInternalServerError)
                        return
                }
        uname := user.Name
        userGreet := "Hello, "+uname+"!"
        user1 := UGreetings{userGreet}
        outgoingJSON, err := json.Marshal(user1)
        if err != nil {
                        log.Println(error.Error())
                        http.Error(res, err.Error(), http.StatusInternalServerError)
                        return
                }
        res.WriteHeader(http.StatusCreated)
        fmt.Fprint(res, string(outgoingJSON))
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", handlePost)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}


