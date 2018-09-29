// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// init_owener() - genric write variable into ledger
func initOwener(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var owner Owner
	var err error
	fmt.Println("starting init_owner")

	if len(args[0]) == 0 {
		return shim.Error("wrong arguments")
	}

	ownerJson := []byte(args[0])
	err = json.Unmarshal(ownerJson, &owner)
	if err != nil {
		fmt.Println("json is wrong,json is: " + args[0])
		return shim.Error(err.Error())
	}

	ownerAsBytes, _ := json.Marshal(owner)
	err = stub.PutState(owner.OwnerId, ownerAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init_owner")
	return shim.Success(nil)
}

// Set Owner on Receipt
func setOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting set_owner")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	receiptId := args[0]
	ownerId := args[1]
	fmt.Println(receiptId + "->" + ownerId)

	// check if user already exists
	owner, err := getOwner(stub, ownerId)
	if err != nil {
		return shim.Error("This owner does not exist - " + ownerId)
	}

	// get receipt's current state
	receiptAsBytes, err := stub.GetState(receiptId)
	if err != nil {
		return shim.Error("Failed to get Receipt")
	}
	res := Receipt{}
	json.Unmarshal(receiptAsBytes, &res)

	res.Owner.OwnerId = owner.OwnerId
	res.Owner.OwnerName = owner.OwnerName

	jsonAsBytes, _ := json.Marshal(res)
	err = stub.PutState(args[0], jsonAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end set owner")
	return shim.Success(nil)
}

//genric write variable into ledger
func write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var receipt Receipt
	var err error
	fmt.Println("starting write")

	if len(args[0]) == 0 {
		return shim.Error("wrong arguments")
	}

	receiptJson := []byte(args[0])
	err = json.Unmarshal(receiptJson, &receipt)
	if err != nil {
		fmt.Println("json is wrong,json is: " + args[0])
		return shim.Error(err.Error())
	}

	//check if new owner exists
	owner, err := getOwner(stub, receipt.Owner.OwnerId)
	if err != nil {
		fmt.Println("Failed to find owner - " + owner.OwnerId)
		return shim.Error(err.Error())
	}

	//check if new receipt exists
	err = checkReceipt(stub, receipt.ReceiptId)
	if err != nil {
		fmt.Println("Failed to create receipt - " + receipt.ReceiptId)
		return shim.Error(err.Error())
	}

	receipt.Time = time.Now().String()
	receiptJson, err = json.Marshal(receipt)
	if err != nil {
		fmt.Println("json is wrong")
		return shim.Error(err.Error())
	}

	err = stub.PutState(receipt.ReceiptId, receiptJson) //write the variable into the ledger
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end write")
	return shim.Success(nil)
}

//genric update variable into ledger
func update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var receipt Receipt
	var err error
	fmt.Println("starting write")

	if len(args[0]) == 0 {
		return shim.Error("wrong arguments")
	}

	receiptJson := []byte(args[0])
	err = json.Unmarshal(receiptJson, &receipt)
	if err != nil {
		fmt.Println("json is wrong,json is: " + args[0])
		return shim.Error(err.Error())
	}

	//check if new owner exists
	owner, err := getOwner(stub, receipt.Owner.OwnerId)
	if err != nil {
		fmt.Println("Failed to find owner - " + owner.OwnerId)
		return shim.Error(err.Error())
	}

	receipt.Time = time.Now().String()
	receiptJson, err = json.Marshal(receipt)
	if err != nil {
		fmt.Println("json is wrong")
		return shim.Error(err.Error())
	}

	err = stub.PutState(receipt.ReceiptId, receiptJson) //write the variable into the ledger
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end write")
	return shim.Success(nil)
}

//remove a receipt from state and from receipt index
func deleteReceipt(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("starting delete_receipt")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// input sanitation
	err := sanitizeArguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	id := args[0]

	// remove the receipt
	err = stub.DelState(id)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	fmt.Println("- end delete_receipt")
	return shim.Success(nil)
}

// Set Owner on Receipt
func setSharelist(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting set_sharelist")

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	receiptId := args[0]
	ownerId := args[1]
	ownerName := args[2]

	// get receipt's current state
	receiptAsBytes, err := stub.GetState(receiptId)
	if err != nil {
		return shim.Error("Failed to get Receipt")
	}
	res := Receipt{}
	json.Unmarshal(receiptAsBytes, &res)

	owner := Owner{}
	owner.OwnerId = ownerId
	owner.OwnerName = ownerName
	res.ShareList = append(res.ShareList, owner)

	resAsBytes, _ := json.Marshal(res)
	err = stub.PutState(receiptId, resAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("end set_sharelist")
	return shim.Success(nil)
}
