// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type ReceiptChainCode struct {
}

type Receipt struct {
	ReceiptId  string  `json:"receiptId"`
	Message    string  `json:"message"`
	Image      string  `json:"image"`
	WaterImage string  `json:"waterImage"`
	Owner      Owner   `json:"owner"`
	ShareList  []Owner `json:"shareList"`
	Time       string  `json:"time"`
}

type Owner struct {
	OwnerId   string `json:"ownerId"`
	OwnerName string `json:"ownerName"`
	OwnerPw   string `json:"ownerPw"`
}

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
		return shim.Error("wrong arguments,")
	}

	fmt.Println("json -- ", args[0])

	fmt.Println("Init() is running")
	fmt.Println("Transaction ID:", txId)

	receiptJson := []byte(args[0])
	err = json.Unmarshal(receiptJson, &receipt)
	if err != nil {
		fmt.Println("json is wrong,json is: " + args[0])
		return shim.Error(err.Error())
	}

	owner := Owner{}
	owner.OwnerId = receipt.Owner.OwnerId
	owner.OwnerName = receipt.Owner.OwnerName
	owner.OwnerPw = receipt.Owner.OwnerPw
	argsJson, err := json.Marshal(owner)
	argsString := string(argsJson[:])
	argsOwner := []string{argsString}
	initOwener(stub, argsOwner)
	receipt.Time = time.Now().String()
	receiptJson, err = json.Marshal(receipt)
	if err != nil {
		fmt.Println("json is wrong")
		return shim.Error(err.Error())
	}

	err = stub.PutState(receipt.ReceiptId, receiptJson)
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
	switch {
	case function == "init":
		return t.Init(stub)

	case function == "readByReciptId":
		return readByReciptId(stub, args)
	case function == "readByReciptIdAndOwner":
		return readByReciptIdAndOwner(stub, args)
	case function == "readByMessage":
		return readByMessage(stub, args)
	case function == "readByMessageAndOwner":
		return readByMessageAndOwner(stub, args)
	case function == "readShareList":
		return readSharelist(stub, args)
	case function == "readByOwner":
		return readByOwner(stub, args)
	case function == "readShareListByOwner":
		return readShareListByOwner(stub, args)
	case function == "readByOidAll":
		return readByOidAll(stub, args)
	case function == "readOwnerAll":
		return readOwnerAll(stub)
	case function == "readReceiptAll":
		return readAll(stub)

	case function == "write":
		return write(stub, args)
	case function == "delete":
		return deleteReceipt(stub, args)
	case function == "update":
		return update(stub, args)
	case function == "initOwner":
		return initOwener(stub, args)
	case function == "setOwner":
		return setOwner(stub, args)
		/*	case function == "history":
			return getHistory(stub, args)*/
	case function == "setShareList":
		return setSharelist(stub, args)
	case function == "verityOwner":
		return verityOwner(stub, args)
	default:
		return shim.Error("Received unknown invoke function name - '" + function + "'")
	}

	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}

// Query - legacy function
//func (t *ReceiptChainCode) Query(stub shim.ChaincodeStubInterface) pb.Response {
//	return shim.Error("Unknown supported call - Query()")
//}
