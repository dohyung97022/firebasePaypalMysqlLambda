package main

import (
	"context"
	firebasePackage "firebase.google.com/go"
	"google.golang.org/api/option"
)

//Firebase ----------------------------------------------------------
type Firebase struct {
	getBool firebaseGetBool
	getStr firebaseGetStr
	app *firebasePackage.App
	context *context.Context
}
//firebase constructor
func newFirebase(serverReaderContext *context.Context)(firebase Firebase, err error){
	credentialClientOption := option.WithCredentialsFile("../firebaseCredential.json")
	app, err := firebasePackage.NewApp(*serverReaderContext, nil, credentialClientOption)
	if err != nil {
		logger.Printf(err.Error())
		return firebase, err
	}
	firebase.app = app
	firebase.context = serverReaderContext
	firebase.getBool.app = app
	firebase.getBool.context = serverReaderContext
	return firebase, nil
}

//firebase.getBool
type firebaseGetBool struct {
	app *firebasePackage.App
	context *context.Context
}
//firebase.getBool.varifyLoginToken
func (firebase *firebaseGetBool) varifyLoginToken (tokenStr string) (resBool bool,err error) {
	client, err := firebase.app.Auth(*firebase.context)
	if err != nil {
		logger.Printf(err.Error())
		return false, err
	}
	token, err := client.VerifyIDToken(*firebase.context, tokenStr)
	if err != nil {
		logger.Printf(err.Error())
		return false, err
	}
	logger.Printf("Verified ID token: %v\n", token)
	return true, nil
}

//firebase.getStrAry
type firebaseGetStr struct {
}
//Mysql.getStrAry.UIDFromLoginToken
func (*firebaseGetStr) UIDFromLoginToken (token string)(UID string, err error){
	return "UID", nil
}