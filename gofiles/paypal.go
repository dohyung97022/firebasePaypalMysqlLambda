package main

import (
	"github.com/joho/godotenv"
	"os"
	"bytes"
)

//Paypal ----------------------------------------------------------
type Paypal struct {
	getBool paypalGetBool
	getStr paypalGetStr
}
//paypal constructor
func newPaypal()(paypal Paypal, err error){
	err = godotenv.Load("../paypal.env")
	return paypal, err
}
//paypal.getBool
type paypalGetBool struct {
}
//paypal.getBool.varifyPaymentFromPaymentID
func (paypal *paypalGetBool) varifyPaymentFromPaymentID (paymentID string) (resBool bool, err error) {
	request, err := curl.getRequest(
		"https://api-m.sandbox.paypal.com/v1/billing/subscriptions/I-BW452GLLEP1G",
		"GET",
		nil,
	)
	if err != nil{
		logger.Printf(err.Error())
		return false, err
	}
	curl.setRequest.header(request, "Content-Type", "application/json")
	curl.setRequest.header(request, "Authorization", "Bearer Access-Token")
	return resBool, nil
}

//paypal.getStr
type paypalGetStr struct {
}
func (paypal *paypalGetStr) accessToken() (resStr string, err error){
    bodyIOReader := bytes.NewBuffer([]byte("grant_type=client_credentials"))
	request, err := curl.getRequest("https://api-m.sandbox.paypal.com/v1/oauth2/token", "POST", bodyIOReader)
	if err!= nil{
		logger.Printf(err.Error())
		return "", err
	}
	curl.setRequest.header(request, "Accept", "application/json")
	curl.setRequest.header(request, "Accept-Language", "en_US")
	curl.setRequest.userNamePassword(request,os.Getenv("CLIENT_ID"),os.Getenv("SECRET_ID"))
	reqResStr, err := curl.getStr.requestResults(request)
	if err!= nil{
		logger.Printf(err.Error())
		return "", err
	}
	return reqResStr, nil
}

