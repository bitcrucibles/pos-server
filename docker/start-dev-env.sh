export DEV_HOST_IP=192.168.0.8
export TEST_DIR=/home/deniv/tmp/test/docker-volumes

export ALICE_BREEZ_ADDRESS="192.168.0.8:50053"
export ALICE_DIR=$TEST_DIR/alice_node
export ALICE_LND_ADDRESS="192.168.0.8:10009"
export BREEZ_DIR=$TEST_DIR/breez_node
export BREEZ_LND_ADDRESS="192.168.0.8:10010"
export SUBSWAP_DIR=$TEST_DIR/subswap_node
export SUBSWAP_LND_ADDRESS="192.168.0.8:10012"
export BTCD_HOST="192.168.0.8:18556"
export BTCD_CERT_FILE=$TEST_DIR/btcd-rpc.cert
export DEBUG="debug"

rm -rf $TEST_DIR
mkdir $TEST_DIR

# create alice folder
mkdir $ALICE_DIR
cp ./alice/lnd.conf ./alice/breez.conf $ALICE_DIR

# create breez node folder
mkdir $BREEZ_DIR
cp ./breez/lnd.conf $BREEZ_DIR

# create subswap node folder
mkdir $SUBSWAP_DIR
cp ./breez/lnd.conf $SUBSWAP_DIR

# run breez node and get mining address
docker-compose -f dev-simnet.yml run -d --name dev-breez dev-breez

#wait for breez rpc
until docker exec dev-breez "cat" /root/.lnd/logs/bitcoin/simnet/lnd.log | grep 'RPC server listening on' > /dev/null;
do
    sleep 1
    echo "waiting for breez RPC..."
done
docker exec dev-breez "/lnd/lncli" -network=simnet newaddress np2wkh | jq -r '.address'
export MINING_ADDRESS=$(docker exec dev-breez "/lnd/lncli" -network=simnet newaddress np2wkh | jq -r '.address')
echo "Mining Address: $MINING_ADDRESS"
docker exec dev-btcd cat /rpc/rpc.cert > $TEST_DIR/btcd-rpc.cert

docker stop dev-breez
docker rm dev-breez

# restart containers
echo "Restart containers"
docker-compose -f dev-simnet.yml down
echo "dev-simnet is down"
docker-compose -f dev-simnet.yml up -d
echo "dev-simnet is up"
docker exec dev-btcd /start-btcctl.sh generate 10
echo "BTCD: generated 400"

#go test ../itest/tests