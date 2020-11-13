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
		server.execStr.respond(err.Error())
		return
	}
	context := r.Context()
	firebase, err := newFirebase(&context)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	firebaseToken, err := firebase.getToken.fromTokenStr(IDTokenStr)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	mysql, err := newMysql()
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	UID := firebase.getStr.UIDFromToken(firebaseToken)
	paymentID, err := mysql.getStr.paymentIDFromUID(UID)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	paypal, err := newPaypal()
	accessToken, err := paypal.getStr.accessToken()
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	varifiedPayment, err := paypal.getBool.varifyPaymentFromPaymentID(paymentID,accessToken)
	if err!=nil{
		logger.Printf(err.Error())
		server.execStr.respond(err.Error())
		return
	}
	if varifiedPayment{
		//아래 줄은 보안상 나중에 꼭 삭제하기!!!!! cors!!!!
		server.setStr.header("Access-Control-Allow-Origin","*")
		//all is varified. Respond and give result!!!!
		server.execStr.respond("results")	
	}
}