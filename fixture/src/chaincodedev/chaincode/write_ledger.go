// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"fmt"

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

	owner.Enabled = true

	ownerAsBytes, _ := json.Marshal(owner)
	err = stub.PutState(owner.Id, ownerAsBytes)
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

	res.OwnerRelation.Id = owner.Id
	res.OwnerRelation.Username = owner.Username

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
	owner, err := getOwner(stub, receipt.OwnerRelation.Id)
	if err != nil {
		fmt.Println("Failed to find owner - " + owner.Id)
		return shim.Error(err.Error())
	}

	//check if new receipt exists
	err = checkReceipt(stub, receipt.Id)
	if err != nil {
		fmt.Println("Failed to create receipt - " + receipt.Id)
		return shim.Error(err.Error())
	}

	err = stub.PutState(receipt.Id, receiptJson) //write the variable into the ledger
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
	owner, err := getOwner(stub, receipt.OwnerRelation.Id)
	if err != nil {
		fmt.Println("Failed to find owner - " + owner.Id)
		return shim.Error(err.Error())
	}

	err = stub.PutState(receipt.Id, receiptJson) //write the variable into the ledger
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

//create a new purchaser, store into chaincode state
func initPurchaser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting init_purchaser")

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//input sanitation
	err = sanitizeArguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	name := args[0]
	taxid := args[1]
	phone := args[2]
	account := args[3]

	//build the receipt json string manually
	str := `{
		"docType":"purchaser", 
		"name": "` + name + `", 
		"taxid": "` + taxid + `", 
		"phone": ` + phone + `, 
		"account": "` + account + `"
	}`
	err = stub.PutState(taxid, []byte(str))
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init_purchaser")
	return shim.Success(nil)
}

//create a new Serller, store into chaincode state
func initSeller(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("starting init_seller")

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//input sanitation
	err = sanitizeArguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	name := args[0]
	taxid := args[1]
	phone := args[2]
	account := args[3]

	//build the receipt json string manually
	str := `{
		"docType":"seller", 
		"name": "` + name + `", 
		"taxid": "` + taxid + `", 
		"phone": ` + phone + `, 
		"account": "` + account + `"
	}`
	err = stub.PutState(taxid, []byte(str))
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end init_purchaser")
	return shim.Success(nil)
}

// Set Owner on Receipt
func setSharelist(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var err error
	fmt.Println("starting set_sharelist")

	if len(args) !=3 {
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
	owner.Id = ownerId
	owner.Username = ownerName
	res.ShareList = append(res.ShareList, owner)

	resAsBytes, _ := json.Marshal(res)
	err = stub.PutState(receiptId, resAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("end set_sharelist")
	return shim.Success(nil)
}
