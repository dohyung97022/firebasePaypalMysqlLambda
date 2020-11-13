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
	IDTokenStr, err := server.getStr.header("IDToken")
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond("Need IDToken in header GET request")
		return
	}
	context := r.Context()
	firebase, err := newFirebase(&context)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond("Firebase could not be initiated")
		return
	}
	firebaseToken, err := firebase.getToken.fromTokenStr(IDTokenStr)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond("firebaseToken is not valid")
		return
	}
	UID := firebase.getStr.UIDFromToken(firebaseToken)
	
	newMysql()

	//아래 줄은 보안상 나중에 꼭 삭제하기!!!!!
	server.setStr.header("Access-Control-Allow-Origin","*")
	server.execStr.respond("We have ran out of youtube api quotas. Please try again tomorrow.")
}