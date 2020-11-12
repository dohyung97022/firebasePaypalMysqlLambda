package main

import(
	"net/http"
)

func main(){
	http.HandleFunc("/search", handler)
	http.ListenAndServe("", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	server:= newServer(&w,r)
	server.setStr.header("Access-Control-Allow-Origin","*")
	server.execStr.respond("We have ran out of youtube api quotas. Please try again tomorrow.")
}