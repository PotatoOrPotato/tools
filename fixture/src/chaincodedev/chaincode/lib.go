// Copyright © 2018 shiyu xu <xushiyu@sinodata.net.cn>.
package main

import (
	"encoding/json"
	"errors"
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

	if len(owner.Username) == 0 { //test if owner is actually here or just nil
		return owner, errors.New("Owner does not exist - " + id + ", '" + owner.Username + "'")
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
