package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VoterRegistrationContract struct {
	contractapi.Contract
}

type Voter struct {
	VoterId  string `json:"voterId"`
	Name     string `json:"name"`
	AadharId string `json:"aadharId"`
	State    string `json:"state"`
	District string `json:"district"`
	Status   string `json:"status"`
}

const CollectionString string = "VoterCollection"

// will add the voter data will be shared between eci and uidai / uidai can add the voter
func (v *VoterRegistrationContract) AddVoter(ctx contractapi.TransactionContextInterface, voterId string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not read the client identity. %s", err)
	}

	if clientOrgID != "UidaiMSP" {
		return "", fmt.Errorf("only uidai can add the voter")
	}

	transientData, err := ctx.GetStub().GetTransient()
	if err != nil {
		return "", err
	}

	name, exists := transientData["name"]
	if !exists {
		return "", fmt.Errorf("name not present")
	}

	aadharId, exists := transientData["aadharId"]
	if !exists {
		return "", fmt.Errorf("aadharId not present")

	}

	state, exists := transientData["state"]
	if !exists {
		return "", fmt.Errorf("state not present")

	}

	district, exists := transientData["district"]
	if !exists {
		return "", fmt.Errorf("district not present")
	}

	voter := Voter{
		VoterId:  voterId,
		Name:     string(name),
		AadharId: string(aadharId),
		State:    string(state),
		District: string(district),
		Status:   "enabled",
	}

	voterDataByte, err := json.Marshal(voter)
	if err != nil {
		return "error while marshalling data", err
	}

	err = ctx.GetStub().PutPrivateData(CollectionString, voterId, voterDataByte)
	if err != nil {
		return "can't put on the chain code", err
	}

	return "Success", nil
}

// this is to delete the voter
func (v *VoterRegistrationContract) DeleteVoter(ctx contractapi.TransactionContextInterface, voterId string) error {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("could not read the client identity. %s", err)
	}

	if clientOrgID != "UidaiMSP" {
		return fmt.Errorf("only uidai can add the voter")
	}

	exists, err := v.VoterExists(ctx, voterId)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", voterId)
	}

	return ctx.GetStub().DelPrivateData(CollectionString, voterId)
}

func (v *VoterRegistrationContract) VoterExists(ctx contractapi.TransactionContextInterface, voterId string) (bool, error) {
	voterByte, err := ctx.GetStub().GetPrivateData(CollectionString, voterId)
	if err != nil {
		return false, err
	}
	if voterByte == nil || len(voterByte) <= 0 {
		return false, nil
	}

	return true, nil
}

// to get the single voter
func (v *VoterRegistrationContract) GetVoter(ctx contractapi.TransactionContextInterface, voterId string) (*Voter, error) {

	voterData, err := ctx.GetStub().GetPrivateData(CollectionString, voterId)
	if err != nil {
		return nil, err
	}
	if voterData == nil || len(voterData) <= 0 {
		return nil, fmt.Errorf("voter doesn't exists")
	}

	voter := &Voter{}
	err = json.Unmarshal(voterData, voter)
	if err != nil {
		return nil, err
	}

	return voter, nil
}

// to get all voter
func (v *VoterRegistrationContract) GetAllVoters(ctx contractapi.TransactionContextInterface, voterId string) ([]*Voter, error) {

	queryString := `{"selector":{"status":"enabled"}}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(CollectionString, queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return OrderResultIteratorFunction(resultsIterator)
}

func OrderResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Voter, error) {
	var voters []*Voter
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of result iterator. %s", err)
		}
		var voter Voter
		err = json.Unmarshal(queryResult.Value, &voter)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		voters = append(voters, &voter)
	}

	return voters, nil
}

// to get voter by the range
func (o *VoterRegistrationContract) GetVoterByRange(ctx contractapi.TransactionContextInterface, startKey string, endKey string) ([]*Voter, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(CollectionString, startKey, endKey)

	if err != nil {
		return nil, fmt.Errorf("could not fetch the private data by range. %s", err)
	}
	defer resultsIterator.Close()

	return OrderResultIteratorFunction(resultsIterator)
}
