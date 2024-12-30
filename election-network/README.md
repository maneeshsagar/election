# Voting Application

## Quick SetUp Of the network
Clone the repositry

To start the network
``` bash
cd election-network

chmod +x startNetwork.sh

./startNetwork.sh
```

to stop the network 

```bash
chmod +x stopNetwork.sh
./stopNetwork.sh
```

To start the client (this will start the client on gin server to intract with the chain code)

```bash

cd client

go mod tidy

go run .
```


## Complete setup, step by step 

  

## Bring Up the Fabric Network

```bash

docker compose -f docker/docker-compose-ca.yaml up -d

  

source registerEnroll.sh

  

docker compose -f docker/docker-compose-org.yaml up -d

```

  
  
  

``` bash

export FABRIC_CFG_PATH=./config

  

export CHANNEL_NAME=electionchannel

  

configtxgen -profile ThreeOrgsChannel -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME

  

```

  
  

``` bash

export ORDERER_CA=./organizations/ordererOrganizations/election.com/orderers/orderer.election.com/msp/tlscacerts/tlsca.election.com-cert.pem

  

export ORDERER_ADMIN_TLS_SIGN_CERT=./organizations/ordererOrganizations/election.com/orderers/orderer.election.com/tls/server.crt

  

export ORDERER_ADMIN_TLS_PRIVATE_KEY=./organizations/ordererOrganizations/election.com/orderers/orderer.election.com/tls/server.key

```

  

## Join the Channel

  

``` bash

osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/$CHANNEL_NAME.block -o localhost:7053 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

  

osnadmin channel list -o localhost:7053 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

```

  

## Peer0_ECI terminal

  

``` bash

export FABRIC_CFG_PATH=./peercfg

export CHANNEL_NAME=electionchannel

export CORE_PEER_LOCALMSPID=EciMSP

export CORE_PEER_TLS_ENABLED=true

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/eci.election.com/peers/peer0.eci.election.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/eci.election.com/users/Admin@eci.election.com/msp

export CORE_PEER_ADDRESS=localhost:7051

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/election.com/orderers/orderer.election.com/msp/tlscacerts/tlsca.election.com-cert.pem

export ECI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/eci.election.com/peers/peer0.eci.election.com/tls/ca.crt

export UIDAI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/uidai.election.com/peers/peer0.uidai.election.com/tls/ca.crt

export PARTY_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/party.election.com/peers/peer0.party.election.com/tls/ca.crt

  
  
  
  
  

peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block

  

peer channel list

```

  
  

## peer0_uidai

  

```bash

export FABRIC_CFG_PATH=./peercfg

export CHANNEL_NAME=electionchannel

export CORE_PEER_LOCALMSPID=UidaiMSP

export CORE_PEER_TLS_ENABLED=true

export CORE_PEER_ADDRESS=localhost:9051

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/uidai.election.com/peers/peer0.uidai.election.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/uidai.election.com/users/Admin@uidai.election.com/msp

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/election.com/orderers/orderer.election.com/msp/tlscacerts/tlsca.election.com-cert.pem

export ECI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/eci.election.com/peers/peer0.eci.election.com/tls/ca.crt

export UIDAI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/uidai.election.com/peers/peer0.uidai.election.com/tls/ca.crt

export PARTY_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/party.election.com/peers/peer0.party.election.com/tls/ca.crt

  
  

peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block

  

peer channel list

```

  
  

## peer0_party

  

```bash

export FABRIC_CFG_PATH=./peercfg

export CHANNEL_NAME=electionchannel

export CORE_PEER_LOCALMSPID=PartyMSP

export CORE_PEER_TLS_ENABLED=true

export CORE_PEER_ADDRESS=localhost:11051

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/party.election.com/peers/peer0.party.election.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/party.election.com/users/Admin@party.election.com/msp

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/election.com/orderers/orderer.election.com/msp/tlscacerts/tlsca.election.com-cert.pem

export ECI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/eci.election.com/peers/peer0.eci.election.com/tls/ca.crt

export UIDAI_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/uidai.election.com/peers/peer0.uidai.election.com/tls/ca.crt

export PARTY_PEER_TLSROOTCERT=${PWD}/organizations/peerOrganizations/party.election.com/peers/peer0.party.election.com/tls/ca.crt

  
  
  

peer channel join -b ./channel-artifacts/$CHANNEL_NAME.block

  

peer channel list

```

  
  

## peer0_eci anchor peer update

``` bash

peer channel fetch config channel-artifacts/config_block.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com -c $CHANNEL_NAME --tls --cafile $ORDERER_CA

  
  
  
  

cd channel-artifacts

  
  
  
  

configtxlator proto_decode --input config_block.pb --type common.Block --output config_block.json

jq '.data.data[0].payload.data.config' config_block.json > config.json

  

cp config.json config_copy.json

  

jq '.channel_group.groups.Application.groups.EciMSP.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "peer0.eci.election.com","port": 7051}]},"version": "0"}}' config_copy.json > modified_config.json

  

configtxlator proto_encode --input config.json --type common.Config --output config.pb

configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb

configtxlator compute_update --channel_id ${CHANNEL_NAME} --original config.pb --updated modified_config.pb --output config_update.pb

  

configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json

echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat config_update.json)'}}}' | jq . > config_update_in_envelope.json

configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output config_update_in_envelope.pb

  
  

cd ..

  

peer channel update -f channel-artifacts/config_update_in_envelope.pb -c $CHANNEL_NAME -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --tls --cafile $ORDERER_CA

```

  
  

## peer0_Uidai for anchor update terminal

```bash

peer channel fetch config channel-artifacts/config_block.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com -c $CHANNEL_NAME --tls --cafile $ORDERER_CA

  

cd channel-artifacts

  

configtxlator proto_decode --input config_block.pb --type common.Block --output config_block.json

jq '.data.data[0].payload.data.config' config_block.json > config.json

cp config.json config_copy.json

  

jq '.channel_group.groups.Application.groups.UidaiMSP.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "peer0.uidai.election.com","port": 9051}]},"version": "0"}}' config_copy.json > modified_config.json

  

configtxlator proto_encode --input config.json --type common.Config --output config.pb

configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb

configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output config_update.pb

  

configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json

echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat config_update.json)'}}}' | jq . > config_update_in_envelope.json

configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output config_update_in_envelope.pb

  
  

cd ..

  

peer channel update -f channel-artifacts/config_update_in_envelope.pb -c $CHANNEL_NAME -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --tls --cafile $ORDERER_CA

  

```

  
  

## peer0_party for anchor update

  

``` bash

peer channel fetch config channel-artifacts/config_block.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com -c $CHANNEL_NAME --tls --cafile $ORDERER_CA

  

cd channel-artifacts

  
  

configtxlator proto_decode --input config_block.pb --type common.Block --output config_block.json

jq '.data.data[0].payload.data.config' config_block.json > config.json

cp config.json config_copy.json

  

jq '.channel_group.groups.Application.groups.PartyMSP.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "peer0.party.election.com","port": 11051}]},"version": "0"}}' config_copy.json > modified_config.json

  

configtxlator proto_encode --input config.json --type common.Config --output config.pb

configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb

configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output config_update.pb

  

configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json

echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat config_update.json)'}}}' | jq . > config_update_in_envelope.json

configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output config_update_in_envelope.pb

  

cd ..

  
  

peer channel update -f channel-artifacts/config_update_in_envelope.pb -c $CHANNEL_NAME -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --tls --cafile $ORDERER_CA

  
  

peer channel getinfo -c $CHANNEL_NAME

```

## Install the chaincode

  

### peer_0 ECI terminal

  

```bash

peer lifecycle chaincode package election.tar.gz --path ../chaincode/ --lang golang --label election_1.0

  

peer lifecycle chaincode install election.tar.gz

  

peer lifecycle chaincode queryinstalled

  
  

export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid election.tar.gz)

  

```

### peer_0 UIDAI terminal

```bash

peer lifecycle chaincode install election.tar.gz

  

export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid election.tar.gz)

  
  

```

  
  

### peer_0 PARTY terminal

  

```bash

peer lifecycle chaincode install election.tar.gz

export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid election.tar.gz)

  

```

  
  

## Approval

### ECI approval

  

```bash

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --channelID $CHANNEL_NAME --name ELECTION --version 1.0 --collections-config ../chaincode/collection.json --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $ORDERER_CA --waitForEvent

  

```

### UIDAI approval

  

```bash

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --channelID $CHANNEL_NAME --name ELECTION --version 1.0 --collections-config ../chaincode/collection.json --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $ORDERER_CA --waitForEvent

  

```

### PARTY approval

  

```bash

peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --channelID $CHANNEL_NAME --name ELECTION --version 1.0 --collections-config ../chaincode/collection.json --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile $ORDERER_CA --waitForEvent

  

```

  

### ECI Terminal

```bash

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name ELECTION --version 1.0 --sequence 1 --collections-config ../chaincode/collection.json --tls --cafile $ORDERER_CA --output json

```

```bash

peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --channelID $CHANNEL_NAME --name ELECTION --version 1.0 --sequence 1 --collections-config ../chaincode/collection.json --tls --cafile $ORDERER_CA --peerAddresses localhost:7051 --tlsRootCertFiles $ECI_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $UIDAI_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $PARTY_PEER_TLSROOTCERT

```

  

```bash

peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ELECTION --cafile $ORDERER_CA

```

  

```bash

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ELECTION --peerAddresses localhost:7051 --tlsRootCertFiles $ECI_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $UIDAI_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $PARTY_PEER_TLSROOTCERT -c '{"function":"VoteContract:CastVote","Args":["vote-01", "voter-01", "elec-01", "BJP"]}'

  
  
  

peer chaincode query -C $CHANNEL_NAME -n ELECTION -c '{"Args":["VoteContract:GetAllVote"]}'

  

```

  

### Peer0 UIDAI terminal

  

```bash

export NAME=$(echo -n "Maneesh" | base64 | tr -d \\n)

  

export AADHARID=$(echo -n "XXXXXXXXXXXX" | base64 | tr -d \\n)

  

export STATE=$(echo -n "UP" | base64 | tr -d \\n)

  

export DISTRICT=$(echo -n "Gorakhpur" | base64 | tr -d \\n)

  
  

peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.election.com --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n ELECTION --peerAddresses localhost:7051 --tlsRootCertFiles $ECI_PEER_TLSROOTCERT --peerAddresses localhost:9051 --tlsRootCertFiles $UIDAI_PEER_TLSROOTCERT --peerAddresses localhost:11051 --tlsRootCertFiles $PARTY_PEER_TLSROOTCERT -c '{"Args":["VoterRegistrationContract:AddVoter","voter-01"]}' --transient "{\"name\":\"$NAME\",\"aadharId\":\"$AADHARID\",\"state\":\"$STATE\",\"district\":\"$DISTRICT\"}"

  
  

peer chaincode query -C $CHANNEL_NAME -n ELECTION -c '{"Args":["VoterRegistrationContract:GetVoter","voter-01"]}'

  

```


