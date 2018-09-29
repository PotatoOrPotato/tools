// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Read - read a generic variable from ledger By receiptid
func readByReciptId(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlRead(args)
	fmt.Println("starting readByReciptId")
	return getReceipt(stub, sql)
}

// Read - read a generic variable from ledger By receiptid and ownerid
func readByReciptIdAndOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForReadByIdAndOwner(args)
	fmt.Println("starting readByReciptIdAndOwner")
	return getReceipt(stub, sql)
}

// Read - read a generic variable from ledger By message
func readByMessage(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForRead(args)
	fmt.Println("starting readByMessage")
	return getReceipt(stub, sql)
}

// Read - read a generic variable from ledger By message and ownerid
func readByMessageAndOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForReadByMessAndOwner(args)
	fmt.Println("starting readByMessageAndOwner")
	return getReceipt(stub, sql)
}

//Get receipt list from owner
func readByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForReadByOwner(args)
	fmt.Println("starting readByOwner")
	return getReceipt(stub, sql)
}

//Get receipt list from owner
func readShareListByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForReadSharelistByOwner(args)
	fmt.Println("starting readShareListByOwner")
	return getReceipt(stub, sql)
}

// Get all of receipt
func readAll(stub shim.ChaincodeStubInterface) pb.Response {
	sql := sqlForAllReceipt()
	fmt.Println("starting readALl")
	return getReceipt(stub, sql)
}

// Get all of receipt by ownerid or ownerid in sharelist
func readByOidAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}
	sql := sqlForOidReceipt(args)
	fmt.Println("starting readOidAll")
	return getReceipt(stub, sql)
}

// Abstract function of receipt
func getReceipt(stub shim.ChaincodeStubInterface, sql string) pb.Response {
	type All struct {
		Receipt []Receipt `json:"receipt"`
	}
	var all All
	var err error

	resultsIterator, err := stub.GetQueryResult(sql)
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
		fmt.Println("key - ", queryKeyAsStr)
		fmt.Println("value - ", string(queryValAsBytes[:]))
		var receipt Receipt
		json.Unmarshal(queryValAsBytes, &receipt) //un stringify it aka JSON.parse()
		all.Receipt = append(all.Receipt, receipt)
	}

	everythingAsBytes, _ := json.Marshal(all)
	fmt.Println("- end read")
	return shim.Success(everythingAsBytes)
}

// Get all of owner
func readOwnerAll(stub shim.ChaincodeStubInterface) pb.Response {
	sql := sqlForAllOwner()
	fmt.Println("starting readOwnerAll")
	return getOwnerAll(stub, sql)
}

// Abstract function of owner
func getOwnerAll(stub shim.ChaincodeStubInterface, sql string) pb.Response {
	type All struct {
		Owner []Owner `json:"owner"`
	}
	var all All
	var err error

	resultsIterator, err := stub.GetQueryResult(sql)
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
		fmt.Println("key - ", queryKeyAsStr)
		fmt.Println("value - ", string(queryValAsBytes[:]))
		var owner Owner
		json.Unmarshal(queryValAsBytes, &owner) //un stringify it aka JSON.parse()
		all.Owner = append(all.Owner, owner)
	}

	everythingAsBytes, _ := json.Marshal(all)
	fmt.Println("- end owner")
	return shim.Success(everythingAsBytes)
}

// Get sharelist of receipt
func readSharelist(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type All struct {
		Receipt []Receipt `json:"receipt"`
	}
	var all All
	var err error

	sql := sqlRead(args)
	resultsIterator, err := stub.GetQueryResult(sql)
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
		fmt.Println("key - ", queryKeyAsStr)
		fmt.Println("value - ", string(queryValAsBytes[:]))
		var receipt Receipt
		json.Unmarshal(queryValAsBytes, &receipt) //un stringify it aka JSON.parse()
		all.Receipt = append(all.Receipt, receipt)
	}

	var everythingAsBytes []byte
	if len(all.Receipt) > 0 {
		everythingAsBytes, _ = json.Marshal(all.Receipt[0].ShareList)
	} else {
		everythingAsBytes, _ = json.Marshal(all.Receipt)
	}

	fmt.Println("- end read")

	return shim.Success(everythingAsBytes) //send it onward
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

//verify the owner
func verityOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	flag := "true"

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sql := sqlForVerityOwner(args)
	resultsIterator, err := stub.GetQueryResult(sql)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		flag = "false"
	}

	str := `{"isExist": "` + flag + `"}`
	return shim.Success([]byte(str))
}
