// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Read - read a generic variable from ledger
func read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error
	fmt.Println("starting read")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}

	// input sanitation
	err = sanitizeArguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key) //get the var from ledger
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("- end read")
	return shim.Success(valAsbytes) //send it onward
}

// Get all of receipt
func readAll(stub shim.ChaincodeStubInterface) pb.Response {
	type All struct {
		Receipt []Receipt `json:"receipt"`
	}
	var all All

	resultsIterator, err := stub.GetStateByRange("r0", "r9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on receipt id - ", queryKeyAsStr)
		var receipt Receipt
		json.Unmarshal(queryValAsBytes, &receipt) //un stringify it aka JSON.parse()
		all.Receipt = append(all.Receipt, receipt)
	}
	fmt.Println("receipt array - ", all.Receipt)

	//change to array of bytes
	everythingAsBytes, _ := json.Marshal(all) //convert to array of bytes
	return shim.Success(everythingAsBytes)
}

// Get all of owner
func readOwnerAll(stub shim.ChaincodeStubInterface) pb.Response {
	type All struct {
		Owner []Owner `json:"receipt"`
	}
	var all All

	resultsIterator, err := stub.GetStateByRange("o0", "o9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on owner id - ", queryKeyAsStr)
		var owner Owner
		json.Unmarshal(queryValAsBytes, &owner) //un stringify it aka JSON.parse()
		all.Owner = append(all.Owner, owner)
	}
	fmt.Println("owner array - ", all.Owner)

	//change to array of bytes
	everythingAsBytes, _ := json.Marshal(all) //convert to array of bytes
	return shim.Success(everythingAsBytes)
}

// Get history of receipt
func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type AuditHistory struct {
		TxId  string  `json:"txId"`
		Value Receipt `json:"value"`
	}
	var history []AuditHistory
	var receipt Receipt

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	receiptId := args[0]
	fmt.Printf("- start getHistoryForReceipt: %s\n", receiptId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(receiptId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx AuditHistory
		tx.TxId = historyData.TxId                  //copy transaction id over
		json.Unmarshal(historyData.Value, &receipt) //un stringify it aka JSON.parse()
		if historyData.Value == nil {               //receipt has been deleted
			var emptyReceipt Receipt
			tx.Value = emptyReceipt //copy nil receipt
		} else {
			json.Unmarshal(historyData.Value, &receipt) //un stringify it aka JSON.parse()
			tx.Value = receipt                          //copy receipt over
		}
		history = append(history, tx) //add this tx to the list
	}
	fmt.Printf("- getHistoryForReceipt returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return shim.Success(historyAsBytes)
}

// Get sharelist of receipt
func readSharelist (stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var key, jsonResp string
	var err error
	fmt.Println("starting read_sharelist")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}

	// input sanitation
	err = sanitizeArguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key) //get the var from ledger
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}
	res := Receipt{}
	json.Unmarshal(valAsbytes,&res)
	ownerAsByte,_ := json.Marshal(res.ShareList)

	fmt.Println("- end read_sharelist")
	return shim.Success(ownerAsByte) //send it onward
}