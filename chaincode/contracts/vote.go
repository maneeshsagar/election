package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VoteContract struct {
	contractapi.Contract
}

type Vote struct {
	VoteId     string `json:"voteId"`
	ElectionId string `json:"electionId"`
	Party      string `json:"party"`
}

type PaginatedQueryResult struct {
	Records             []*Vote `json:"records"`
	FetchedRecordsCount int32   `json:"fetchedRecordsCount"`
	Bookmark            string  `json:"bookmark"`
}

type HistoryQueryResult struct {
	Record    *Vote  `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

// this is to cast the vote
func (v *VoteContract) CastVote(ctx contractapi.TransactionContextInterface, voteId, voterId, electionId, party string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID != "EciMSP" {
		return "", fmt.Errorf("can't cast vote from here")
	}

	// here will check the identity

	bytes, err := ctx.GetStub().GetPrivateData(CollectionString, voterId)
	if err != nil {
		return "", fmt.Errorf("could not get the private data: %s", err)
	}

	if len(bytes) <= 0 {
		return "", fmt.Errorf("no voter exists")
	}

	voter := &Voter{}
	err = json.Unmarshal(bytes, voter)
	if err != nil {
		return "", fmt.Errorf("error while unmarshalling: %s", err)
	}

	vote := Vote{
		VoteId:     voteId,
		ElectionId: electionId,
		Party:      party,
	}

	voteBytes, err := json.Marshal(vote)
	if err != nil {
		return "unable to marshal", err
	}

	err = ctx.GetStub().PutState(voteId, voteBytes)
	if err != nil {
		return "unable to marshal", err
	}
	return "success", nil
}

// this is to revoke the vote
func (v *VoteContract) RevokeVote(ctx contractapi.TransactionContextInterface, voteId string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID != "EciMSP" {
		return "", fmt.Errorf("can't cast vote from here")
	}

	err = ctx.GetStub().DelState(voteId)
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}
	return "Success", nil
}

func (v *VoteContract) VoteExists(ctx contractapi.TransactionContextInterface, voteId string) (bool, error) {

	voteBytes, err := ctx.GetStub().GetState(voteId)
	if err != nil {
		return false, err
	}
	if len(voteBytes) == 0 {
		return false, nil
	}
	return true, nil
}

// to get all votes
func (v *VoteContract) GetAllVote(ctx contractapi.TransactionContextInterface) ([]*Vote, error) {
	queryString := `{"selector":{"party":"BJP"}}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return voteResultIteratorFunction(resultsIterator)
}

func voteResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Vote, error) {
	var votes []*Vote
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of the result iterator. %s", err)
		}
		var vote Vote
		err = json.Unmarshal(queryResult.Value, &vote)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		votes = append(votes, &vote)
	}
	return votes, nil
}

// to get voteByrange
func (c *VoteContract) GetVotesByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Vote, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the  data by range. %s", err)
	}
	defer resultsIterator.Close()

	return voteResultIteratorFunction(resultsIterator)
}

// to get the vote history
func (c *VoteContract) GetVoteHistory(ctx contractapi.TransactionContextInterface, voteId string) ([]*HistoryQueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(voteId)
	if err != nil {
		return nil, fmt.Errorf("could not get the data. %s", err)
	}
	defer resultsIterator.Close()

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not get the value of resultsIterator. %s", err)
		}

		var vote Vote
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &vote)
			if err != nil {
				return nil, err
			}
		} else {
			vote = Vote{
				VoteId: voteId,
			}
		}

		timestamp := response.Timestamp.AsTime()

		formattedTime := timestamp.Format(time.RFC1123)

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &vote,
			IsDelete:  response.IsDelete,
		}
		records = append(records, &record)
	}

	return records, nil
}

// get vote with pagination
func (c *VoteContract) GetVoteWithPagination(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {

	queryString := `{"selector":{"party":"bjp"}}`

	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, fmt.Errorf("could not get the vote records. %s", err)
	}
	defer resultsIterator.Close()

	cars, err := voteResultIteratorFunction(resultsIterator)
	if err != nil {
		return nil, fmt.Errorf("could not return the vote records %s", err)
	}

	return &PaginatedQueryResult{
		Records:             cars,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}
