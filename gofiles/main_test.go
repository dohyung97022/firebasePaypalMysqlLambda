package main

import (
	"testing"
	// "context"
	"fmt"
	"time"
)

func TestFuncs(t *testing.T){
	//firebase testing
	// ctx := context.Background()
	// firebase, err := newFirebase(&ctx)
	// if err!=nil{
	// 	logger.Printf(err.Error())
	// 	return
	// }
	// start := time.Now()
	// resToken, err := firebase.getToken.fromTokenStr("")
	// if err!=nil {
	// 	//토큰이 valid하지 않을 경우 / 자체적인 err인 경우
	// 	logger.Printf(err.Error())
	// 	return
	// }
	// UID := firebase.getStr.UIDFromToken(resToken)
	// fmt.Printf("UID := %v\n",UID)
	// timeTrack(start, "varifyLoginToken")

	//mysql testing
	// mysql, err := newMysql()
	// if err!=nil{
	// 	return
	// }
	// paymentID, err = mysql.getStr.paymentIDFromUID("uidStr")
	// if (err!=nil || paymentID == ""){
	// 	return
	// }

	//paypal testing
	paypal, err := newPaypal()
	if err!=nil{
		return
	}
	resStr, err := paypal.getStr.accessToken()
	if err!=nil{
		return
	}
	fmt.Printf(resStr)
}
func TestMain(t *testing.T){
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}