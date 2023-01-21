package main

//Note: this style of importing custom local packages ONLY works if your project/workspace is in your GOPATH inside the src folder
// windows my folder is in C:\Promgram Files\Go\src\pal-api
//look into temp changing your gopath when you go to do your work? think that would break all of the other built in packages
import (
	"net/http"
	"pal-api/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe("localhost:3000", mux)

}
