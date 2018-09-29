// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type ReceiptChainCode struct {
}

// Asset Definitions - The ledger will store BusModel and Detail
// ----- Business Model ----- //
type BusModel struct {
	ObjectType      string `json:"docType"`
	Name            string `json:"name"`
	TaxpayerNumber  string `json:"taxId"`
	AddressAndPhone string `json:"addressAndPhone"`
	BankAndAccount  string `json:"bankAndAccount"`
}

type Transaction struct {
	ServiceName  string  `json:"serviceName"`
	ProductModel string  `json:"productModel"`
	Unit         string  `json:"unit"`
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
	Taxrate      float64 `json:"taxrate"`
	TaxAmount    float64 `json:"taxAmount"`
}

type Detail struct {
	TradeName    string      `json:"tradeName"`
	Code         string      `json:"code"`
	Number       string      `json:"number"`
	Date         string      `json:"date"`
	SecurityCode string      `json:"securityCode"`
	PasswordAera string      `json:"passwordAera"`
	Transaction  Transaction `json:"transaction"`
	Sum          float64     `json:"sum"`
	Remark       string      `json:"remark"`
	Payee        string      `json:"payee"`
	Recheck      string      `json:"recheck"`
	Drawer       string      `json:"drawer"`
}

type Receipt struct {
	Id            string        `json:"id"`
	Purchaser     string        `json:"purchaser"`
	Seller        string        `json:"seller"`
	Sum           int           `json:"sum"`
	ImageByte     string        `json:"imageByte""`
	OwnerRelation OwnerRelation `json:"ownerRelation"`
	ShareList     []Owner       `json:"shareList"`
}

type OwnerRelation struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type Owner struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Enabled  bool   `json:"enabled"`
}

//type RealReceipt struct {
//	Purchaser BusModel `json:"purchaser"`
//	Seller    BusModel `json:"seller"`
//	Detail    Detail   `json:"detail"`
//	OcrImage  []byte   `json:"ocrImage"`
//}

// Main
func main() {
	err := shim.Start(new(ReceiptChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// Init - initialize the chaincode
// Returns - shim.Success or error
func (t *ReceiptChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Receipt Is Starting Up")
	_, args := stub.GetFunctionAndParameters()
	var receipt Receipt
	var err error
	txId := stub.GetTxID()

	if len(args[0]) == 0 {
		return shim.Error("wrong arguments")
	}

	fmt.Println("Init() is running")
	fmt.Println("Transaction ID:", txId)

	receiptJson := []byte(args[0])
	err = json.Unmarshal(receiptJson, &receipt)
	if err != nil {
		fmt.Println("json is wrong,json is: " + args[0])
		return shim.Error(err.Error())
	}

	owner := Owner{}
	owner.Id = receipt.OwnerRelation.Id
	owner.Username = receipt.OwnerRelation.Username
	argsJson,err := json.Marshal(owner)
	argsString := string(argsJson[:])
	args_owner := []string{argsString}
	initOwener(stub, args_owner)

	err = stub.PutState(receipt.Id, receiptJson)
	if err != nil {
		fmt.Println("Could not store receipt")
		return shim.Error(err.Error())
	}

	fmt.Println("Ready for action") //self-test pass
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
func (t *ReceiptChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("starting invoke, for - " + function)

	//Handle different functions
	if function == "init" {
		return t.Init(stub)
	} else if function == "read" {
		return read(stub, args)
	} else if function == "write" {
		return write(stub, args)
	} else if function == "update" {
		return update(stub, args)
	} else if function == "delete" {
		return deleteReceipt(stub, args)
	} else if function == "initOwner" {
		return initOwener(stub, args)
	} else if function == "setOwner" {
		return setOwner(stub, args)
	} else if function == "readReceiptAll" {
		return readAll(stub)
	} else if function == "readOwnerAll" {
		return readOwnerAll(stub)
	} else if function == "setShareList" {
		return setSharelist(stub, args)
	} else if function == "readShareList" {
		return readSharelist(stub, args)
	} else if function == "history" {
		return getHistory(stub, args)
	} else if function == "initPurchaser" {
		return initPurchaser(stub, args)
	} else if function == "initSeller" {
		return initSeller(stub, args)
	}

	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}

// Query - legacy function
//func (t *ReceiptChainCode) Query(stub shim.ChaincodeStubInterface) pb.Response {
//	return shim.Error("Unknown supported call - Query()")
//}
