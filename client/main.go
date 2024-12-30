package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Vote struct {
	VoteId     string `json:"voteId"`
	VoterId    string `json:"voterId"`
	ElectionId string `json:"electionId"`
	Party      string `json:"party"`
}

type Voter struct {
	VoterId  string `json:"voterId"`
	Name     string `json:"name"`
	AadharId string `json:"aadharId"`
	State    string `json:"state"`
	District string `json:"district"`
	Status   string `json:"status"`
}

type VoteHistory struct {
	Record    *Vote  `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

type Register struct {
	CarId     string `json:"carId"`
	CarOwner  string `json:"carOwner"`
	RegNumber string `json:"regNumber"`
}

type VoteByRangeRequest struct {
	StartKey string `json:"startKey"`
	EndKey   string `json:"endKey"`
}

type VoterByRangeRequest struct {
	StartKey string `json:"startKey"`
	EndKey   string `json:"endKey"`
}

type VoteByPagination struct {
	Pagesize string `json:"pagesize"`
	Bookmark string `json:"bookmark"`
}

func main() {
	router := gin.Default()

	var wg sync.WaitGroup
	wg.Add(1)
	go ChaincodeEventListener("uidai", "electionchannel", "ELECTION", &wg)

	router.GET("/votes", func(ctx *gin.Context) {
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "query", make(map[string][]byte), "GetAllVote")

		var votes []Vote

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &votes); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(200, votes)
	})

	router.GET("/vote", func(ctx *gin.Context) {
		var req VoteByRangeRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "query", make(map[string][]byte), "GetVotesByRange", req.StartKey, req.EndKey)

		var cars []Vote

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(200, cars)
	})

	router.GET("/votePage", func(ctx *gin.Context) {
		var req VoteByPagination
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "query", make(map[string][]byte), "GetVoteWithPagination", req.Pagesize, req.Bookmark)

		var cars []Vote

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &cars); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(200, cars)
	})

	router.POST("/vote", func(ctx *gin.Context) {
		var req Vote
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "invoke", make(map[string][]byte), "CastVote", req.VoteId, req.VoterId, req.ElectionId, req.Party)

		ctx.JSON(200, result)
	})

	router.DELETE("/vote/:id", func(ctx *gin.Context) {
		voteId := ctx.Param("id")
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "invoke", make(map[string][]byte), "RevokeVote", voteId)
		ctx.JSON(200, result)
	})

	router.POST("/api/voter", func(ctx *gin.Context) {
		var req Voter
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		fmt.Printf("order  %s", req)
		req.Status = "Enabled"

		privateData := map[string][]byte{
			"name":     []byte(req.Name),
			"voterId":  []byte(req.VoterId),
			"aadharId": []byte(req.AadharId),
			"state":    []byte(req.State),
			"district": []byte(req.District),
			"status":   []byte(req.Status),
		}

		submitTxnFn("uidai", "electionchannel", "ELECTION", "VoterRegistrationContract", "private", privateData, "AddVoter", req.VoterId)

		ctx.JSON(http.StatusOK, req)
	})

	router.GET("/api/voter/:id", func(ctx *gin.Context) {
		voterId := ctx.Param("id")

		result := submitTxnFn("uidai", "electionchannel", "ELECTION", "VoterRegistrationContract", "query", make(map[string][]byte), "GetVoter", voterId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.DELETE("/api/voter/:id", func(ctx *gin.Context) {
		voterId := ctx.Param("id")

		result := submitTxnFn("uidai", "electionchannel", "ELECTION", "VoterRegistrationContract", "invoke", make(map[string][]byte), "DeleteVoter", voterId)

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/api/voters", func(ctx *gin.Context) {
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoterRegistrationContract", "query", make(map[string][]byte), "GetAllVoters", "")

		var voters []Voter

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &voters); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(200, voters)
	})

	router.GET("/api/votersRange", func(ctx *gin.Context) {

		var req VoterByRangeRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoterRegistrationContract", "query", make(map[string][]byte), "GetVoterByRange", req.StartKey, req.EndKey)

		var voters []Voter

		if len(result) > 0 {
			// Unmarshal the JSON array string into the cars slice
			if err := json.Unmarshal([]byte(result), &voters); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(200, voters)
	})

	router.GET("/api/vote/history", func(ctx *gin.Context) {
		voteId := ctx.Query("voteId")
		result := submitTxnFn("eci", "electionchannel", "ELECTION", "VoteContract", "query", make(map[string][]byte), "GetVoteHistory", voteId)

		// fmt.Printf("result %s", result)

		var votes []VoteHistory

		if len(result) > 0 {
			// Unmarshal the JSON array string into the orders slice
			if err := json.Unmarshal([]byte(result), &votes); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		ctx.JSON(http.StatusOK, votes)
	})

	router.GET("/api/event", func(ctx *gin.Context) {
		result := getEvents()
		fmt.Println("result:", result)

		ctx.JSON(http.StatusOK, gin.H{"VoteEvent": result})

	})

	router.Run("localhost:8080")
}
