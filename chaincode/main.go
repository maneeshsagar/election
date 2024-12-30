package main

import (
	"chaincode/contracts"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	voterRegistrationContract := contracts.VoterRegistrationContract{}
	voteContract := contracts.VoteContract{}

	chaincode, err := contractapi.NewChaincode(&voterRegistrationContract, &voteContract)
	if err != nil {
		fmt.Println("unable to initialize the chaincode", err)
		return
	}
	err = chaincode.Start()
	if err != nil {
		fmt.Println("unable to start the chaincode")
		return
	}
}
