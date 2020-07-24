package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/services/{service}/imagehubs/{imageHubName:[A-Za-z][A-Za-z0-9|/|\\-|_|\\.]+[A-Za-z0-9]+[^(/tags)]}/tags/{tag}", TestRepo).Methods(http.MethodGet)
	r.HandleFunc("/v1/services/{service}/imagehubs/{imageHubName:[A-Za-z][A-Za-z0-9|/|\\-|_|\\.]+[A-Za-z0-9]+[^(/tags)]}/tags", TestRepo_NoTag).Methods(http.MethodGet)
	r.HandleFunc("/v1/services/{service}/imagehubs/{imageHubName:[A-Za-z][A-Za-z0-9|/|\\-|_|\\.]{0,}[A-Za-z0-9]}", Test).Methods(http.MethodGet)

	err := http.ListenAndServe( ":19000" , r)
	if err!=nil{
		fmt.Println(err)
	}
}

func Test(w http.ResponseWriter,req *http.Request)  {
	imageHubName:=	mux.Vars(req)["imageHubName"]
	writeResult(fmt.Sprintf(`{"imageHubName":"%s","method":"test"}`,imageHubName),w)

}

func TestRepo(w http.ResponseWriter,req *http.Request)  {
	imageHubName:=	mux.Vars(req)["imageHubName"]
	tag:=	mux.Vars(req)["tag"]
	writeResult(fmt.Sprintf(`{"imageHubName":"%s","tag":"%s","method":"testRepo"}`,imageHubName,tag),w)
}
func TestRepo_NoTag(w http.ResponseWriter,req *http.Request)  {
	imageHubName:=	mux.Vars(req)["imageHubName"]
	writeResult(fmt.Sprintf(`{"imageHubName":"%s","method":"testRepo_NoTag"}`,imageHubName),w)
}

func writeError(statusCode int, message string, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	errorString := fmt.Sprintf("{\"error\":\"%s\"}", message)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(errorString))

}

func writeResult(dataString string, w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(dataString))

}