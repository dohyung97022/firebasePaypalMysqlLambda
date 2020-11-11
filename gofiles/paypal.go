package main

import (
)

//Paypal ----------------------------------------------------------
type Paypal struct {
	getBool paypalGetBool
}
//paypal constructor
func newPaypal()(paypal Paypal, err error){
	return paypal, nil
}

//paypal.getBool
type paypalGetBool struct {
}
//paypal.getBool.varifyPaymentFromPaymentID
func (paypal *paypalGetBool) varifyPaymentFromPaymentID (paymentID string) (resBool bool, err error) {
	return resBool, nil
}