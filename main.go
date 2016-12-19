// chaincode_insurance project main.go
package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type InsuranceChaincode struct {
}
type Policy struct {
	PolicyNo     string //保单号码
	PolicyType   string //险种
	Startdate    string //保险生效时间
	Enddate      string //保险失效时间
	Status       string //保单状态
	PolicyHolder string //投保人
	Assured      string //被保险人
	Beneficiary  string //保险受益人
	Premium      string //保费
	Amount       string //保险金
}

func (insurance *InsuranceChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("---UserMsg---ChainCode has been initilazed!")
	return nil, nil
}

func (insurance *InsuranceChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("---UserMsg---", "Function=", function)
	fmt.Println("---UserMsg---", "args=", args)

	var item Policy
	item.PolicyNo = args[0]

	if function == "Insert" {
		return insurance.Insert(stub, args)
		//	} else if function == "Change" {
		//		return insurance.change(stub, args)
	} else if function == "Delete" {
		return insurance.Delete(stub, args)
	}

	return nil, errors.New("Incorrect function! They may be one of Insert,Delete,Change.")
}

func (insurance *InsuranceChaincode) Insert(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var item Policy
	item.PolicyNo = args[0]

	if len(args) != 10 {
		return nil, errors.New("Incorrect number of arguments. Deposit function expecting 10")
	}

	//解析参数
	item.PolicyType = args[1]
	item.Startdate = args[2]
	item.Enddate = args[3]
	item.Status = args[4]
	item.PolicyHolder = args[5]
	item.Assured = args[6]
	item.Beneficiary = args[7]
	item.Premium = args[8]
	item.Amount = args[9]

	detailbytes, err := stub.GetState(item.PolicyNo)

	//保单号重复CHECK
	if detailbytes != nil {
		errmsg := "The Policy No is duplicate."
		fmt.Println(errmsg)
		return nil, errors.New(errmsg)
	}

	if detailbytes == nil && err == nil {
		//添加check处理等功能

		//将结构体（保单信息）转化成json数据，然后保存到账本中
		p, err := json.Marshal(item)
		err = stub.PutState(item.PolicyNo, p)

		// Write the state to the ledger
		//err = stub.PutState(args[0], []byte(args))
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// Deletes an entity from state
func (insurance *InsuranceChaincode) Delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	PolicyNo := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(PolicyNo)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (insurance *InsuranceChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return insurance.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}
func (insurance *InsuranceChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var PolicyNo string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	PolicyNo = args[0]

	// Get the state from the ledger
	policybytes, err := stub.GetState(PolicyNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + PolicyNo + "\"}"
		return nil, errors.New(jsonResp)
	}

	if policybytes == nil {
		jsonResp := "{\"Result\":\"Nil policy for " + PolicyNo + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return nil, nil
	} else {
		jsonResp := "{\"PolicyNo\":\"" + PolicyNo + "\",\"Detail\":\"" + string(policybytes) + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return policybytes, nil
	}
}

func main() {
	err := shim.Start(new(InsuranceChaincode))
	if err != nil {
		fmt.Printf("Error starting insurance chaincode: %s", err)
	}
}
