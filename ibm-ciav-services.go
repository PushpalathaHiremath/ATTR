/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
//	"github.com/op/go-logging"
)

//var myLogger = logging.MustGetLogger("customer_CIAV_details")

type ServicesChaincode struct {
}

func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
//	myLogger.Debug("Hi Abhishek . . . ")
	err := stub.PutState("counter", []byte("0"))
	return nil, err
}


func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function != "increment" {
		return nil, errors.New("Invalid invoke function name. Expecting \"increment\"")
	}
	val, err := stub.ReadCertAttribute("position")
	fmt.Printf("Position => %v error %v \n", string(val), err)
	isOk, _ := stub.VerifyAttribute("position", []byte("Software Engineer")) // Here the ABAC API is called to verify the attribute, just if the value is verified the counter will be incremented.
	if isOk {
		counter, err := stub.GetState("counter")
		if err != nil {
			return nil, err
		}
		var cInt int
		cInt, err = strconv.Atoi(string(counter))
		if err != nil {
			return nil, err
		}
		cInt = cInt + 1
		counter = []byte(strconv.Itoa(cInt))
		stub.PutState("counter", counter)
	}
	return nil, nil

}

/*
 		Get Customer record by customer id or PAN number
*/
func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "read" {
		return read(stub, args)
	}else {
	 	return readAttr(stub, args)
	}
	return nil, errors.New("Invalid query function name. Expecting \"read\"")
}

func readAttr(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	attrVal1, err := stub.ReadCertAttribute("position")
	isPresent, err := stub.VerifyAttribute("position", []byte("Software Engineer")) // Here the ABAC API is called to verify the attribute, just if the value is verified the counter will be incremented.
	if err != nil {
		return nil, err
	}
	jsonResp := "{ " +
					"Attribute Name  01: "+ string(attrVal1) +
					"Attribute Value  01 : "+ strconv.FormatBool(isPresent) +

				 "}"
	fmt.Printf("Query Response:%s\n", jsonResp)

	bytes, err := json.Marshal(jsonResp)
	if err != nil {
		return nil, errors.New("Error converting kyc record")
	}
	return bytes, nil
}


func read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	// Get the state from the ledger
	Avalbytes, err := stub.GetState("counter")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for counter\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for counter\"}"
		return nil, errors.New(jsonResp)
	}

	isOk, _ := stub.VerifyAttribute("position123", []byte("Software Engineer"))
	var jsonResp string
	if isOk {
		jsonResp = "{\"counter1\" : " + string(Avalbytes) +
								"\"}"
	}else{
		jsonResp = "{\"counter2\" : " + string(Avalbytes) +
								"\"}"
	}

	fmt.Printf("Query Response:%s\n", jsonResp)

	bytes, err := json.Marshal(jsonResp)
	if err != nil {
		return nil, errors.New("Error converting kyc record")
	}
	return bytes, nil
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
