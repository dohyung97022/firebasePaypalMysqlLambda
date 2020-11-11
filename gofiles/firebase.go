package main

import (
)

//Firebase ----------------------------------------------------------
type Firebase struct {
	getBool firebaseGetBool
	getStr firebaseGetStr
}
//firebase constructor
func newFirebase()(firebase Firebase, err error){
	return firebase, nil
}

//firebase.getBool
type firebaseGetBool struct {
}
//firebase.getBool.varifyLoginToken
func (firebase *firebaseGetBool) varifyLoginToken (token string) (resBool bool,err error) {
	return resBool, nil
}

//firebase.getStrAry
type firebaseGetStr struct {
}
//Mysql.getStrAry.UIDFromLoginToken
func (firebase *firebaseGetStr) UIDFromLoginToken (token string)(UID string, err error){
	return "UID", nil
}