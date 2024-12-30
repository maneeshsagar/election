package main

// Config represents the configuration for a role.
type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}

// Create a Profile map
var profile = map[string]Config{

	"eci": {
		CertPath:     "../election-network/organizations/peerOrganizations/eci.election.com/users/User1@eci.election.com/msp/signcerts/cert.pem",
		KeyDirectory: "../election-network/organizations/peerOrganizations/eci.election.com/users/User1@eci.election.com/msp/keystore/",
		TLSCertPath:  "../election-network/organizations/peerOrganizations/eci.election.com/peers/peer0.eci.election.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.eci.election.com",
		MSPID:        "EciMSP",
	},

	"uidai": {
		CertPath:     "../election-network/organizations/peerOrganizations/uidai.election.com/users/User1@uidai.election.com/msp/signcerts/cert.pem",
		KeyDirectory: "../election-network/organizations/peerOrganizations/uidai.election.com/users/User1@uidai.election.com/msp/keystore/",
		TLSCertPath:  "../election-network/organizations/peerOrganizations/uidai.election.com/peers/peer0.uidai.election.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.uidai.election.com",
		MSPID:        "UidaiMSP",
	},

	"party": {
		CertPath:     "../election-network/organizations/peerOrganizations/party.election.com/users/User1@party.election.com/msp/signcerts/cert.pem",
		KeyDirectory: "../election-network/organizations/peerOrganizations/party.election.com/users/User1@party.election.com/msp/keystore/",
		TLSCertPath:  "../election-network/organizations/peerOrganizations/party.election.com/peers/peer0.party.election.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.party.election.com",
		MSPID:        "PartyMSP",
	},
	"org1": {
		CertPath:     "../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem",
		KeyDirectory: "../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/",
		TLSCertPath:  "../../fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.org1.example.com",
		MSPID:        "Org1MSP",
	},

	"org2": {
		CertPath:     "../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/signcerts/cert.pem",
		KeyDirectory: "../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/users/User1@org2.example.com/msp/keystore/",
		TLSCertPath:  "../../fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.org2.example.com",
		MSPID:        "Org2MSP",
	},

	"org3": {
		CertPath:     "../../fabric-samples/test-network/organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/signcerts/cert.pem",
		KeyDirectory: "../../fabric-samples/test-network/organizations/peerOrganizations/org3.example.com/users/User1@org3.example.com/msp/keystore/",
		TLSCertPath:  "../../fabric-samples/test-network/organizations/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.org3.example.com",
		MSPID:        "Org3MSP",
	},
}
