// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Get Owner - get the owner asset from ledger
func getOwner(stub shim.ChaincodeStubInterface, id string) (Owner, error) {
	var owner Owner
	ownerAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                        //this seems to always succeed, even if key didn't exist
		return owner, errors.New("Failed to get owner - " + id)
	}
	json.Unmarshal(ownerAsBytes, &owner) //un stringify it aka JSON.parse()

	if len(owner.OwnerName) == 0 { //test if owner is actually here or just nil
		return owner, errors.New("Owner does not exist - " + id + ", '" + owner.OwnerName + "'")
	}

	return owner, nil
}

// Check receipt is existed
func checkReceipt(stub shim.ChaincodeStubInterface, receiptId string) error {
	var err error

	receiptAsByte, err := stub.GetState(receiptId)

	if len(receiptAsByte) != 0 {
		err = errors.New("The receipt is existed - " + strconv.Itoa(len(receiptAsByte)) + receiptId)
		return err
	}

	return err
}

// Input Sanitation - dumb input checking, look for empty strings
func sanitizeArguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 32 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 32 characters")
		}
	}
	return nil
}

//arg to sql
func sqlRead(args []string) string {
	arg0 := args[0]
	checkSql(arg0)
	sql := "{\"selector\":{\"receiptId\":{\"$eq\":\"" + arg0 + "\"}},\"use_index\":[\"_design/indexReceiptIdDoc\", \"indexReceiptId\"]}"
	fmt.Println(sql)
	return sql
}

func sqlForReadByIdAndOwner(args []string) string {
	arg0 := args[0]
	arg1 := args[1]
	checkSql(arg0)
	sql := "{\"selector\":{\"receiptId\":{\"$eq\":\"" + arg0 + "\"},\"owner.ownerId\":{\"$eq\":\"" + arg1 + "\"}}}"
	fmt.Println(sql)
	return sql
}

func sqlForRead(args []string) string {
	arg0 := args[0]
	checkSql(arg0)
	sql := "{\"selector\":{\"message\":{\"$regex\":\"" + arg0 + "\"}}}"
	fmt.Println(sql)
	return sql
}

func sqlForReadByMessAndOwner(args []string) string {
	arg0 := args[0]
	arg1 := args[1]
	checkSql(arg0)
	sql := "{\"selector\":{\"message\":{\"$regex\":\"" + arg0 + "\"},\"owner.ownerId\":{\"$eq\":\"" + arg1 + "\"}}}"
	fmt.Println(sql)
	return sql
}

func sqlForReadByOwner(args []string) string {
	arg0 := args[0]
	checkSql(arg0)
	sql := "{\"selector\":{\"owner.ownerId\":{\"$regex\":\"" + arg0 + "\"}},\"use_index\":[\"_design/indexReceiptOwnerIdDoc\", \"indexReceiptOwnerId\"]}"
	fmt.Println(sql)
	return sql
}

func sqlForReadSharelistByOwner(args []string) string {
	arg0 := args[0]
	checkSql(arg0)
	sql := "{\"selector\":{\"shareList\":{\"$elemMatch\":{\"ownerId\":{\"$regex\":\"" + arg0 + "\"}}}},\"use_index\":[\"_design/indexShareListIdDoc\", \"indexShareListId\"]}"
	fmt.Println(sql)
	return sql
}

func sqlForAllReceipt() string {
	sql := "{\"selector\":{\"receiptId\":{\"$gt\":\"null\"}},\"use_index\":[\"_design/indexReceiptIdDoc\", \"indexReceiptId\"]}"
	fmt.Println(sql)
	return sql
}

func sqlForAllOwner() string {
	sql := "{\"selector\":{\"ownerId\":{\"$gt\":\"null\"}},\"use_index\":[\"_design/indexOwnerIdDoc\", \"indexOwnerId\"]}"
	fmt.Println(sql)
	return sql
}

func sqlForOidReceipt(args []string) string {
	arg0 := args[0]
	checkSql(arg0)
	sql := "{\"selector\":{\"$or\":[{\"owner.ownerId\":{\"$regex\":\"" + arg0 + "\"}},{\"shareList\":{\"$elemMatch\":{\"ownerId\":{\"$regex\":\"" + arg0 + "\"}}}}]}}"
	fmt.Println(sql)
	return sql
}

func sqlForVerityOwner(args []string) string {
	arg0 := args[0]
	arg1 := args[1]
	checkSql(arg0)
	sql := "{\"selector\":{\"ownerId\":{\"$eq\":\"" + arg0 + "\"},\"ownerPw\":{\"$eq\":\"" + arg1 + "\"}}}"
	fmt.Println(sql)
	return sql
}

//check sql
func checkSql(arg string) {

}
