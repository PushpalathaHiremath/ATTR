/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the IBM India Pvt Ltd, Version 1.0 (the "License");
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
//	"github.com/op/go-logging"
	"bytes"
	"strings"
	"strconv"
)

// var myLogger = logging.MustGetLogger("customer_address_details")
var dummyValue = "99999"
type ServicesChaincode struct {
}

type Identification struct {
	CustomerId  	string
	IdentityNumber  string
	PoiType 				 string
	PoiDoc          string
	PoiExpiryDate  string
	Source		   		 string
}

type PersonalDetails struct {
	CustomerId  	   string
	FirstName    	 string
	LastName 			 string
	Sex   					 string
	EmailId         string
	Dob              string
	PhoneNumber     string
	Occupation       string
	AnnualIncome    string
	IncomeSource    string
	Source		   		 string
}

type Kyc struct {
	CustomerId  	   string
	KycStatus       string
	LastUpdated     string
	Source		   		 string
}

type Address struct {
	CustomerId  	   string
	AddressId       string
	AddressType     string
	DoorNumber      string
	Street           string
	Locality         string
	City 						 string
	State 					 string
	Pincode          string
	PoaType         string
	PoaDoc          string
	PoaExpiryDate  string
	Source		   		 string
}

type Customer struct {
		Identification    []Identification
		PersonalDetails   PersonalDetails
		Kyc								Kyc
		Address					  []Address
}

/*
   Deploy KYC data model
*/
func (t *ServicesChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	CreateIdentificationTable(stub, args)
	CreateCustomerTable(stub, args)
	CreateKycTable(stub, args)
	CreateAddressTable(stub, args)
	return nil, nil
}

/*
  Add Customer record
*/
func (t *ServicesChaincode) addCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	var Cust Customer
	// rec, err := ioutil.ReadFile("./config.json")
	// rec := "{\"Identification\":[{\"customerId\":\"CUST001\",\"identityNumber\":\"SGDIH8754L\",\"poiType\":\"PAN\",\"poiDoc\":\"some url\",\"expiryDate\":\"01/01/2030\",\"source\":\"HDFC Bank\"},{\"customerId\":\"CUST001\",\"identityNumber\":\"Aadhar\",\"poiType\":\"1646-2648-277\",\"poiDoc\":\"some url\",\"expiryDate\":\"01/01/2000\",\"source\":\"HDFC Bank\"}],\"PersonalDetails\":{\"customerId\":\"CUST001\",\"firstName\":\"Sachin\",\"lastName\":\"Prasad\",\"sex\":\"Male\",\"emailId\":\"saprasad@gmail.com\",\"dob\":\"08/08/88\",\"phoneNumber\":\"9648654973\",\"occupation\":\"1\",\"annualIncome\":\"3\",\"incomeSource\":\"2\",\"source\":\"HDFC Bank\"},\"KYC\":{\"CustomerId\":\"CUST001\",\"KycStatus\":\"non-compliant\",\"LastUpdated\":\"10/08/2016\",\"source\":\"HDFC Bank\"},\"address\":[{\"customerId\":\"CUST001\",\"addressId\":\"12345\",\"addressType\":\"permanent\",\"doorNumber\":\"#01\",\"street\":\"Residency road\",\"locality\":\"Near super Market\",\"city\":\"Bangalore\",\"state\":\"Karnataka\",\"pincode\":\"560107\",\"poaType\":\"Driving Licence\",\"poaDoc\":\"some url\",\"expiryDate\":\"01/01/2020\",\"source\":\"HDFC Bank\"},{\"customerId\":\"CUST001\",\"addressId\":\"addressID456\",\"addressType\":\"communication\",\"doorNumber\":\"#53\",\"street\":\"Queens road\",\"locality\":\"Near BEL circle\",\"city\":\"Bangalore\",\"state\":\"Karnataka\",\"pincode\":\"560119\",\"poaType\":\"Driving Licence\",\"poaDoc\":\"some url\",\"expiryDate\":\"01/01/2021\",\"source\":\"HDFC Bank\"}]}"
	// fmt.Print(string(rec))
	// var rec string
	// rec = "{\"Identification\":[{\"CustomerId\":\"CUST00001\",\"IdentityNumber\":\"SGDIH8754L\",\"PoiType\":\"PAN\",\"PoiDoc\":\"someurl\",\"ExpiryDate\":\"01/01/2030\",\"Source\":\"HDFCBank\"}],\"PersonalDetails\":{\"CustomerId\":\"CUST00001\",\"FirstName\":\"Sachin\",\"LastName\":\"Prasad\",\"Sex\":\"Male\",\"EmailId\":\"saprasad@gmail.com\",\"Dob\":\"08/08/88\",\"PhoneNumber\":\"9648654973\",\"Occupation\":\"1\",\"AnnualIncome\":\"3\",\"IncomeSource\":\"2\",\"Source\":\"HDFCBank\"},\"KYC\":{\"CustomerId\":\"CUST00001\",\"KycStatus\":\"non-compliant\",\"LastUpdated\":\"10/08/2016\",\"Source\":\"HDFCBank\"},\"Address\":[{\"CustomerId\":\"CUST00001\",\"AddressId\":\"12345\",\"AddressType\":\"permanent\",\"DoorNumber\":\"#01\",\"Street\":\"Residencyroad\",\"Locality\":\"NearsuperMarket\",\"City\":\"Bangalore\",\"State\":\"Karnataka\",\"Pincode\":\"560107\",\"PoaType\":\"DrivingLicence\",\"PoaDoc\":\"someurl\",\"ExpiryDate\":\"01/01/2020\",\"Source\":\"HDFCBank\"},{\"CustomerId\":\"CUST00001\",\"AddressId\":\"addressID456\",\"AddressType\":\"communication\",\"DoorNumber\":\"#53\",\"Street\":\"Queensroad\",\"Locality\":\"NearBELcircle\",\"City\":\"Bangalore\",\"State\":\"Karnataka\",\"Pincode\":\"560119\",\"PoaType\":\"DrivingLicence\",\"PoaDoc\":\"someurl\",\"ExpiryDate\":\"01/01/2021\",\"Source\":\"HDFCBank\"}]}"
	err := json.Unmarshal([]byte(string(args[0])), &Cust)
	if err != nil {
		fmt.Println("Error is :",err)
	}
	// res := &Cust.Identification
	// fmt.Println(len(Cust.Identification))
	// res1B, _ := json.Marshal(res)
	// fmt.Println(string(res1B))
	// fmt.Println(string(res1B[0]))

for i := range Cust.Identification {
	AddIdentification(stub, []string{Cust.Identification[i].CustomerId, Cust.Identification[i].IdentityNumber, Cust.Identification[i].PoiType, Cust.Identification[i].PoiDoc,
		Cust.Identification[i].PoiExpiryDate, Cust.Identification[i].Source})
}
	// AddIdentification(stub, []string{Cust.Identification[i].CustomerId, Cust.Identification[i].IdentityNumber, Cust.Identification.PoiType, Cust.Identification.PoiDoc, Cust.Identification.PoiExpiryDate, Cust.Identification.Source})
	AddCustomer(stub, []string{Cust.PersonalDetails.CustomerId, Cust.PersonalDetails.FirstName, Cust.PersonalDetails.LastName,
		Cust.PersonalDetails.Sex, Cust.PersonalDetails.EmailId, Cust.PersonalDetails.Dob, Cust.PersonalDetails.PhoneNumber, Cust.PersonalDetails.Occupation,
		Cust.PersonalDetails.AnnualIncome, Cust.PersonalDetails.IncomeSource, Cust.PersonalDetails.Source})
	AddKYC(stub, []string{Cust.Kyc.CustomerId, Cust.Kyc.KycStatus, Cust.Kyc.LastUpdated, Cust.Kyc.Source})
for i := range Cust.Address {
	AddAddress(stub, []string{Cust.Address[i].CustomerId, Cust.Address[i].AddressId, Cust.Address[i].AddressType,
		Cust.Address[i].DoorNumber, Cust.Address[i].Street, Cust.Address[i].Locality, Cust.Address[i].City, Cust.Address[i].State,
		Cust.Address[i].Pincode, Cust.Address[i].PoaType, Cust.Address[i].PoaDoc, Cust.Address[i].PoaExpiryDate, Cust.Address[i].Source})
}


	// Common
	// customer_id := args[0]
	// source := args[1]
	//
	// // Identification
	// identity_number := args[2]
	// poi_type := args[3]
	// poi_doc := args[4]
	// poi_expiry_date := args[5]
	//
	// identity_number2 := args[28]
	// poi_type2 := args[29]
	// poi_doc2 := args[30]
	// poi_expiry_date2 := args[31]
	//
	// // Customer personal details
	// first_name := args[6]
	// last_name := args[7]
	// sex := args[8]
	// email_id := args[9]
	// dob := args[10]
	// phone_number := args[11]
	// occupation := args[12]
	// annual_income := args[13]
	// income_source := args[14]
	//
	// //kyc
	// kyc_status := args[15]
	// last_updated := args[16]
	//
	// // Address
	// address_id := args[17]
	// address_type := args[18]
	// door_number := args[19]
	// street := args[20]
  // locality := args[21]
	// city :=args[22]
	// state := args[23]
	// pincode := args[24]
	// poa_type := args[25]
	// poa_doc := args[26]
	// poa_expiry_date := args[27]
	//
	// address_id2 := args[32]
	// address_type2 := args[33]
	// door_number2 := args[34]
	// street2 := args[35]
	// locality2 := args[36]
	// city2 :=args[37]
	// state2 := args[38]
	// pincode2 := args[39]
	// poa_type2 := args[40]
	// poa_doc2 := args[41]
	// poa_expiry_date2 := args[42]

	// AddIdentification(stub, []string{customer_id, identity_number, poi_type, poi_doc, poi_expiry_date, source})
	// AddCustomer(stub, []string{customer_id, first_name, last_name, sex, email_id, dob, phone_number, occupation, annual_income, income_source, source})
	// AddKYC(stub, []string{customer_id, kyc_status, last_updated, source})
	// AddAddress(stub, []string{customer_id, address_id, address_type, door_number, street, locality, city, state, pincode, poa_type, poa_doc, poa_expiry_date, source})
	//
	// if args[28] != "" {
	// 	AddIdentification(stub, []string{customer_id, identity_number2, poi_type2, poi_doc2, poi_expiry_date2, source})
	// }
	// if args[32] != "" {
	// 	AddAddress(stub, []string{customer_id, address_id2, address_type2, door_number2, street2, locality2, city2, state2, pincode2, poa_type2, poa_doc2, poa_expiry_date2, source})
	// }
	return nil, nil
}

/*
 Update customer record
*/
func (t *ServicesChaincode) updateCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 43 {
		return nil, errors.New("Incorrect number of arguments. Expecting 43")
	}
	// Common
	customer_id := args[0]
	source := args[1]

	// Identification
	identity_number := args[2]
	poi_type := args[3]
	poi_doc := args[4]
	poi_expiry_date := args[5]

	identity_number2 := args[28]
	poi_type2 := args[29]
	poi_doc2 := args[30]
	poi_expiry_date2 := args[31]

	// Customer personal details
	first_name := args[6]
	last_name := args[7]
	sex := args[8]
	email_id := args[9]
	dob := args[10]
	phone_number := args[11]
	occupation := args[12]
	annual_income := args[13]
	income_source := args[14]

	//kyc
	kyc_status := args[15]
	last_updated := args[16]

	// Address
	address_id := args[17]
	address_type := args[18]
	door_number := args[19]
	street := args[20]
	locality := args[21]
	city := args[22]
	state := args[23]
	pincode := args[24]
	poa_type := args[25]
	poa_doc := args[26]
	poa_expiry_date := args[27]

	address_id2 := args[32]
	address_type2 := args[33]
	door_number2 := args[34]
	street2 := args[35]
	locality2 := args[36]
	city2 :=args[37]
	state2 := args[38]
	pincode2 := args[39]
	poa_type2 := args[40]
	poa_doc2 := args[41]
	poa_expiry_date2 := args[42]

	UpdateIdentification(stub, []string{customer_id, identity_number, poi_type, poi_doc, poi_expiry_date, source})
	UpdateCustomer(stub, []string{customer_id, first_name, last_name, sex, email_id, dob, phone_number, occupation, annual_income, income_source, source})
	UpdateKYC(stub, []string{customer_id, kyc_status, last_updated, source})
	UpdateAddress(stub, []string{customer_id, address_id, address_type, door_number, street, locality, city, state, pincode, poa_type, poa_doc, poa_expiry_date, source})


	if args[28] != "" {
		UpdateIdentification(stub, []string{customer_id, identity_number2, poi_type2, poi_doc2, poi_expiry_date2, source})
	}
	if args[32] != "" {
		UpdateAddress(stub, []string{customer_id, address_id2, address_type2, door_number2, street2, locality2, city2, state2, pincode2, poa_type2, poa_doc2, poa_expiry_date2, source})
	}
	return nil, nil
}

/*
   Invoke : addCIAV and updateCIAV
*/
func (t *ServicesChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	if function == "addCIAV" {
		// add customer
		return t.addCIAV(stub, args)
	} else {
		// update customer
		return t.updateCIAV(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

/*
 		Get Customer record by customer id or PAN number
*/
func (t *ServicesChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "getCIAV" {
		return t.getCIAV(stub, args)
	}else if function == "getKYCStats" {
		return t.GetKYCStats(stub)
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t *ServicesChaincode) getCIAV(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var jsonResp string
	var customerIds []string
	var err error

	var identificationStr string
	var customerStr string
	var kycStr string
	var addressStr string
	if args[0] == "PAN" {
		customerIds, err = GetCustomerID(stub, args[1])
		// jsonResp = "["
		for i := range customerIds {
			customerId := customerIds[i]
			identificationStr, err =GetIdentification(stub, customerId)
			customerStr, err = GetCustomer(stub, customerId)
			kycStr, err = GetKYC(stub, customerId)
			addressStr, err = GetAddress(stub, customerId)

			if i != 0 {
				jsonResp = jsonResp + ","
			}
			jsonResp = jsonResp + "{\"Identification\":" + identificationStr +
				",\"PersonalDetails\":" + customerStr +
				",\"KYC\":" + kycStr +
				",\"address\":" + addressStr + "}"
		}
			// jsonResp = jsonResp + "]"
	} else if args[0] == "CUST_ID" {
		customerId := args[1]
		identificationStr, err = GetIdentification(stub, customerId)
		customerStr, err = GetCustomer(stub, customerId)
		kycStr, err = GetKYC(stub, customerId)
		addressStr, err = GetAddress(stub, customerId)

		jsonResp = "{\"Identification\":" + identificationStr +
			",\"PersonalDetails\":" + customerStr +
			",\"KYC\":" + kycStr +
			",\"address\":" + addressStr + "}"
	} else {
		return nil, errors.New("Invalid arguments. Please query by CUST_ID or PAN")
	}

	bytes, err := json.Marshal(jsonResp)
	if err != nil {
		return nil, errors.New("Error converting kyc record")
	}
	return bytes, nil
}


/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				Address
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*
	Create address table
*/
func CreateAddressTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
//	myLogger.Debug("Creating Address Table ...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	err := stub.CreateTable("Address", []*shim.ColumnDefinition{
    &shim.ColumnDefinition{Name: "customer_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "address_id", Type: shim.ColumnDefinition_STRING, Key: true},
    &shim.ColumnDefinition{Name: "address_type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "door_tumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "street", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "locality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "state", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "pincode", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "poa_type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "poa_doc", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "expiry_date", Type: shim.ColumnDefinition_STRING, Key: false},
    &shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Address table.")
	}

	// myLogger.Debug("Address table created Successfully... !!! ")
	return nil, nil
}

/*
	Add Address record
*/
func AddAddress(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Adding Address record ...")

	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 13")
	}
    customerId := args[0]
    addressId := args[1]
  	addressType := args[2]
  	doorNumber := args[3]
  	street := args[4]
  	locality := args[5]
		city := args[6]
  	state := args[7]
  	pincode := args[8]
  	poaType := args[9]
  	poaDoc := args[10]
  	expiryDate := args[11]
  	source := args[12]

	ok, err := stub.InsertRow("Address", shim.Row{
		Columns: []*shim.Column{
      &shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: addressId}},
			&shim.Column{Value: &shim.Column_String_{String_: addressType}},
			&shim.Column{Value: &shim.Column_String_{String_: doorNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: street}},
			&shim.Column{Value: &shim.Column_String_{String_: locality}},
			&shim.Column{Value: &shim.Column_String_{String_: city}},
			&shim.Column{Value: &shim.Column_String_{String_: state}},
			&shim.Column{Value: &shim.Column_String_{String_: pincode}},
			&shim.Column{Value: &shim.Column_String_{String_: poaType}},
			&shim.Column{Value: &shim.Column_String_{String_: poaDoc}},
      &shim.Column{Value: &shim.Column_String_{String_: expiryDate}},
      &shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding address record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully added",)
	return nil, err
}

func UpdateAddress(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Updating address record ...")

	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}

  customerId := args[0]
  addressId := args[1]
  addressType := args[2]
  doorNumber := args[3]
  street := args[4]
  locality := args[5]
  city := args[6]
  state := args[7]
  pincode := args[8]
  poaType := args[9]
  poaDoc := args[10]
  expiryDate := args[11]
  source := args[12]

	ok, err := stub.ReplaceRow("Address", shim.Row{
		Columns: []*shim.Column{
      &shim.Column{Value: &shim.Column_String_{String_: customerId}},
      &shim.Column{Value: &shim.Column_String_{String_: addressId}},
      &shim.Column{Value: &shim.Column_String_{String_: addressType}},
      &shim.Column{Value: &shim.Column_String_{String_: doorNumber}},
      &shim.Column{Value: &shim.Column_String_{String_: street}},
      &shim.Column{Value: &shim.Column_String_{String_: locality}},
			&shim.Column{Value: &shim.Column_String_{String_: city}},
      &shim.Column{Value: &shim.Column_String_{String_: state}},
      &shim.Column{Value: &shim.Column_String_{String_: pincode}},
      &shim.Column{Value: &shim.Column_String_{String_: poaType}},
      &shim.Column{Value: &shim.Column_String_{String_: poaDoc}},
      &shim.Column{Value: &shim.Column_String_{String_: expiryDate}},
      &shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in updated customer address record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully updated ",)
	return nil, err
}

/*
 Get address record
*/
func GetAddress(stub *shim.ChaincodeStub, customerId string) (string, error) {
	var err error
	// myLogger.Debugf("Getting address record for customer : [%s]", string(customerId))

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: customerId}}
	columns = append(columns, col1)
	rows, err := GetAllRows(stub, "Address", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Address details [%s]: [%s]", string(customerId), err)
		return "", fmt.Errorf("Failed retriving Address details [%s]: [%s]", string(customerId), err)
	}

	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString("[")
	for i := range rows {
		row := rows[i]
		// myLogger.Debugf("Identification rows [%s], is : [%s]", i, row)
		fmt.Println(row)
		if i != 0 {
			jsonRespBuffer.WriteString(",")
		}
		jsonRespBuffer.WriteString("{\"customerId\":\"" + row.Columns[0].GetString_() + "\"" +
			",\"addressId\":\"" + row.Columns[1].GetString_() + "\"" +
			",\"addressType\":\"" + row.Columns[2].GetString_() + "\"" +
			",\"doorNumber\":\"" + row.Columns[3].GetString_() + "\"" +
			",\"street\":\"" + row.Columns[4].GetString_() + "\"" +
			",\"locality\":\"" + row.Columns[5].GetString_() + "\"" +
			",\"city\":\"" + row.Columns[6].GetString_() + "\"" +
			",\"state\":\"" + row.Columns[7].GetString_() + "\"" +
			",\"pincode\":\"" + row.Columns[8].GetString_() + "\"" +
			",\"poaType\":\"" + row.Columns[9].GetString_() + "\"" +
			",\"poaDoc\":\"" + row.Columns[10].GetString_() + "\"" +
	    ",\"expiryDate\":\"" + row.Columns[11].GetString_() + "\"" +
	    ",\"source\":\"" + row.Columns[12].GetString_() + "\"}")
	}
	jsonRespBuffer.WriteString("]")

	return jsonRespBuffer.String(), nil
}

/*
	Get all rows corresponding to the partial keys given
*/
func GetAllRows(stub *shim.ChaincodeStub, tableName string, columns []shim.Column) ([]shim.Row, error) {
	rowChannel, err := stub.GetRows(tableName, columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving address details for : [%s]", err)
		return nil, fmt.Errorf("Failed retriving address details : [%s]", err)
	}
	var rows []shim.Row
	for {
		select {
		case temprow, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				// myLogger.Debugf("Fetching row : [%s]", temprow.Columns[0].GetString_())
				rows = append(rows, temprow)
			}
		}
		if rowChannel == nil {
			break
		}
	}
	return rows, nil
}

/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				Customer
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*
	Create customer table
*/
func CreateCustomerTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Creating Customer table...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create Customer table
	err := stub.CreateTable("Customer", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "customer_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "first_name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "last_name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sex", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email_id", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phone_number", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "occupation", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "annual_income", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "income_source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed to create Customer table.")
	}
	// myLogger.Debug("Customer table created Successfully... !!! ")
	return nil, nil
}

/*
	Add customer record
*/
func AddCustomer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Deb ug("Adding Customer record ...")

	if len(args) != 11 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}
	customerId := args[0]
	firstName := args[1]
	lastName := args[2]
	sex := args[3]
	emailId := args[4]
	dob := args[5]
	phoneNumber := args[6]
	occupation := args[7]
	annualIncome := args[8]
	incomeSource := args[9]
	source := args[10]

	// Add customer entry to Customer table
	// myLogger.Debugf("Adding Customer [%s] ", firstName)

	ok, err := stub.InsertRow("Customer", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: firstName}},
			&shim.Column{Value: &shim.Column_String_{String_: lastName}},
			&shim.Column{Value: &shim.Column_String_{String_: sex}},
			&shim.Column{Value: &shim.Column_String_{String_: emailId}},
			&shim.Column{Value: &shim.Column_String_{String_: dob}},
			&shim.Column{Value: &shim.Column_String_{String_: phoneNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: occupation}},
			&shim.Column{Value: &shim.Column_String_{String_: annualIncome}},
			&shim.Column{Value: &shim.Column_String_{String_: incomeSource}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding customer record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully added [%s] ", firstName)
	return nil, err
}

/*
	Update customer record
*/
func UpdateCustomer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Updating Customer record ...")

	if len(args) != 11 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}

	customerId := args[0]
	firstName := args[1]
	lastName := args[2]
	sex := args[3]
	emailId := args[4]
	dob := args[5]
	phoneNumber := args[6]
	occupation := args[7]
	annualIncome := args[8]
	incomeSource := args[9]
	source := args[10]

	// Update customer entry to Customer table
	// myLogger.Debugf("Updating Customer [%s] ", firstName)

	ok, err := stub.ReplaceRow("Customer", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: firstName}},
			&shim.Column{Value: &shim.Column_String_{String_: lastName}},
			&shim.Column{Value: &shim.Column_String_{String_: sex}},
			&shim.Column{Value: &shim.Column_String_{String_: emailId}},
			&shim.Column{Value: &shim.Column_String_{String_: dob}},
			&shim.Column{Value: &shim.Column_String_{String_: phoneNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: occupation}},
			&shim.Column{Value: &shim.Column_String_{String_: annualIncome}},
			&shim.Column{Value: &shim.Column_String_{String_: incomeSource}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in updated customer record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully updated [%s] ", firstName)
	return nil, err
}

/*
 Get customer record
*/
func GetCustomer(stub *shim.ChaincodeStub, customerId string) (string, error) {
	var err error
	// myLogger.Debugf("Get personal details record for customer : [%s]", string(customerId))

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: customerId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Customer", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Customer details [%s]: [%s]", string(customerId), err)
		return "", fmt.Errorf("Failed retriving Customer details [%s]: [%s]", string(customerId), err)
	}

	jsonResp := "{\"customerId\":\"" + row.Columns[0].GetString_() + "\"" +
		",\"firstName\":\"" + row.Columns[1].GetString_() + "\"" +
		",\"lastName\":\"" + row.Columns[2].GetString_() + "\"" +
		",\"sex\":\"" + row.Columns[3].GetString_() + "\"" +
		",\"emailId\":\"" + row.Columns[4].GetString_() + "\"" +
		",\"dob\":\"" + row.Columns[5].GetString_() + "\"" +
		",\"phoneNumber\":\"" + row.Columns[6].GetString_() + "\"" +
		",\"occupation\":\"" + row.Columns[7].GetString_() + "\"" +
		",\"annualIncome\":\"" + row.Columns[8].GetString_() + "\"" +
		",\"incomeSource\":\"" + row.Columns[9].GetString_() + "\"" +
		",\"source\":\"" + row.Columns[10].GetString_() + "\"}"
	return jsonResp, nil
}


/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				identification
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*
 Create Identification table
*/
func CreateIdentificationTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Init Identification Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create Identification table
	err := stub.CreateTable("Identification", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "customer_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "identity_number", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "poi_type", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "poi_doc", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "expiry_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	// Create identification relation table
	err = stub.CreateTable("IDRelation", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "identity_number", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "customer_id", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	if err != nil {
		return nil, errors.New("Failed creating Identification table.")
	}
	// myLogger.Debug("Identification table initialization done Successfully... !!! ")
	return nil, nil
}

/*
	add Identification record
*/
func AddIdentification(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Add Identification record ...")
	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}

	customerId := args[0]
	identityNumber := args[1]
	poiType := args[2]
	poiDoc := args[3]
	expiryDate := args[4]
	source := args[5]

	// myLogger.Debugf("Adding identity doc : [%s] ", poiType)

	ok, err := stub.InsertRow("Identification", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: identityNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: poiType}},
			&shim.Column{Value: &shim.Column_String_{String_: poiDoc}},
			&shim.Column{Value: &shim.Column_String_{String_: expiryDate}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	// update identification relation
	_, err = updateIDRelation(stub, identityNumber, customerId, "add")
	if !ok && err == nil {
		return nil, errors.New("Error in adding Identification record.")
	}

	// myLogger.Debug("Congratulations !!! Successfully added, [%s]", res)
	return nil, err
}

/*
 Update Identification record
*/
func UpdateIdentification(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Update Identification record ...")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}

	customerId := args[0]
	identityNumber := args[1]
	poiType := args[2]
	poiDoc := args[3]
	expiryDate := args[4]
	source := args[5]

	// myLogger.Debugf("Updating identity : [%s] ", poiType)

	ok, err := stub.ReplaceRow("Identification", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: identityNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: poiType}},
			&shim.Column{Value: &shim.Column_String_{String_: poiDoc}},
			&shim.Column{Value: &shim.Column_String_{String_: expiryDate}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	_, err = updateIDRelation(stub, identityNumber, customerId, "update")
	ok, err = stub.ReplaceRow("IDRelation", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: identityNumber}},
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in updating Identification record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully updated [%s]", res)
	return nil, err
}

/*
	Get Identification record
*/
func GetIdentification(stub *shim.ChaincodeStub, customerId string) (string, error) {
	var err error
	// myLogger.Debugf("Get identification record for customer : [%s]", string(customerId))

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: customerId}}
	columns = append(columns, col1)
	rows, err := GetAllRows(stub, "Identification", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Identification details [%s]: [%s]", string(customerId), err)
		return "", fmt.Errorf("Failed retriving Identification details [%s]: [%s]", string(customerId), err)
	}

	var jsonRespBuffer bytes.Buffer
	jsonRespBuffer.WriteString("[")
	for i := range rows {
		row := rows[i]
		// myLogger.Debugf("Identification rows [%s], is : [%s]", i, row)
		fmt.Println(row)
		if i != 0 {
			jsonRespBuffer.WriteString(",")
		}
		jsonRespBuffer.WriteString("{\"customerId\":\"" + row.Columns[0].GetString_() + "\"" +
			",\"identityNumber\":\"" + row.Columns[1].GetString_() + "\"" +
			",\"poiType\":\"" + row.Columns[2].GetString_() + "\"" +
			",\"poiDoc\":\"" + row.Columns[3].GetString_() + "\"" +
			",\"expiryDate\":\"" + row.Columns[4].GetString_() + "\"" +
			",\"source\":\"" + row.Columns[5].GetString_() + "\"}")
	}
	jsonRespBuffer.WriteString("]")

	return jsonRespBuffer.String(), nil
}

/*
 Get the customer id by PAN number
*/
func GetCustomerID(stub *shim.ChaincodeStub, panId string) ([]string, error) {
	var err error

	// myLogger.Debugf("Get customer id for PAN : [%s]", panId)

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: panId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("IDRelation", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Identification details for PAN [%s]: [%s]", string(panId), err)
		return nil, fmt.Errorf("Failed retriving Identification details  for PAN [%s]: [%s]", string(panId), err)
	}

	custIds := row.Columns[1].GetString_()
	custIdArray := strings.Split(custIds, "|")
	return custIdArray, nil
}

/*
	Update ID relation table
*/
func updateIDRelation(stub *shim.ChaincodeStub, identityNumber string, customerId string, functionType string) (string, error) {
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: identityNumber}}
	columns = append(columns, col1)

	idrow, err := stub.GetRow("IDRelation", columns)
	if err != nil {
		// myLogger.Debugf("Failed retriving Identification relation details for ID [%s]: [%s]", string(identityNumber), err)
		return "", fmt.Errorf("Failed retriving Identification relation details  for ID [%s]: [%s]", string(identityNumber), err)
	}

	var isRowExists bool
	isRowExists = (idrow.Columns != nil)

	var ok bool
	if isRowExists {
		if functionType == "update" {
			if strings.Contains(idrow.Columns[1].GetString_(), customerId) {
				// myLogger.Debugf("Identification relation exists. Do nothing.")
				return "", nil
			}
		}
		customerId = idrow.Columns[1].GetString_() + "|" + customerId
		ok, err = stub.ReplaceRow("IDRelation", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: identityNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			},
		})
	} else {
		ok, err = stub.InsertRow("IDRelation", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: identityNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			},
		})
	}
	if !ok && err == nil {
		return "", errors.New("Error in updating Identification relation record.")
	}
	return "", nil
}

/*
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
																				kyc
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
*/

/*
	Create KYC table
*/
func CreateKycTable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Creating KYC Table...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	err := stub.CreateTable("KYC", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "cummy", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "customerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "kycStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdated", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating KYC table.")
	}
	// myLogger.Debug("KYC table initialization done Successfully... !!! ")
	return nil, nil
}

/*
	Add KYC record
*/
func AddKYC(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Adding KYC record ...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	customerId := args[0]
	kycStatus := args[1]
	lastUpdated := args[2]
	source := args[3]

	ok, err := stub.InsertRow("KYC", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding KYC record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully added")
	return nil, err
}

/*
	Update KYC record
*/
func UpdateKYC(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// myLogger.Debug("Updating KYC record ...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	customerId := args[0]
	kycStatus := args[1]
	lastUpdated := args[2]
	source := args[3]

	ok, err := stub.ReplaceRow("KYC", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: dummyValue}},
			&shim.Column{Value: &shim.Column_String_{String_: customerId}},
			&shim.Column{Value: &shim.Column_String_{String_: kycStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: lastUpdated}},
			&shim.Column{Value: &shim.Column_String_{String_: source}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in updating KYC record.")
	}
	// myLogger.Debug("Congratulations !!! Successfully updated")
	return nil, err
}

/*
 Get KYC record
*/
func GetKYC(stub *shim.ChaincodeStub, customerId string) (string, error) {
	var err error
	// myLogger.Debugf("Get identification record for customer : [%s]", string(customerId))
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: dummyValue}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: customerId}}
	columns = append(columns, col1)
	columns = append(columns, col2)
	row, err := stub.GetRow("KYC", columns)
	if err != nil {
		return "", fmt.Errorf("Failed retriving KYC details [%s]: [%s]", string(customerId), err)
	}
	jsonResp := "{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
		",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
		",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\"" +
		",\"source\":\"" + row.Columns[4].GetString_() + "\"}"

	return jsonResp, nil
}

func (t *ServicesChaincode) GetKYCStats(stub *shim.ChaincodeStub) ([]byte, error) {
	var err error

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: dummyValue}}
	columns = append(columns, col1)
	rows, err := GetAllRows(stub, "KYC", columns)
	if err != nil {
		return nil,fmt.Errorf("Failed retriving KYC details [%s]", err)
	}

	var kycBuffer bytes.Buffer
	// var compliantBuffer bytes.Buffer
	// var noncompliantBuffer bytes.Buffer
	var compliantCustomersCount int
	var noncompliantCustomersCount int
	var totalCustomers int

	for i := range rows {
		row := rows[i]
		totalCustomers++;
		if row.Columns[2].GetString_() == "compliant"{
			compliantCustomersCount++
			// if compliantBuffer.String() != "" {
			// 	compliantBuffer.WriteString(",")
			// }
			// compliantBuffer.WriteString("{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
			// 	",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
			// 	",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\"" +
			// 	",\"source\":\"" + row.Columns[4].GetString_() + "\"}")
		}else if row.Columns[2].GetString_() == "non-compliant" {
			noncompliantCustomersCount++
		// 	if noncompliantBuffer.String() != "" {
		// 		noncompliantBuffer.WriteString(",")
		// 	}
		// 	noncompliantBuffer.WriteString("{\"customerId\":\"" + row.Columns[1].GetString_() + "\"" +
		// 		",\"kycStatus\":\"" + row.Columns[2].GetString_() + "\"" +
		// 		",\"lastUpdated\":\"" + row.Columns[3].GetString_() + "\"" +
		// 		",\"source\":\"" + row.Columns[4].GetString_() + "\"}")
		}
	}
	kycBuffer.WriteString("{" +
			"\"compliant\" : \"" + strconv.Itoa(compliantCustomersCount)+"\"," +
			"\"noncompliant\" : \"" + strconv.Itoa(noncompliantCustomersCount) +"\"," +
			"\"total\" : \"" + strconv.Itoa(totalCustomers) + "\"" +
			"}")

	bytes, err := json.Marshal(kycBuffer.String())
	if err != nil {
		return nil, errors.New("Error converting kyc stats")
	}
	return bytes, nil
}

func main() {
	err := shim.Start(new(ServicesChaincode))
	if err != nil {
		fmt.Printf("Error starting ServicesChaincode: %s", err)
	}
}
