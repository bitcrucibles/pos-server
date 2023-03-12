sleep 20s
./generate_macaroon_hex LND /root/breez_node/data/chain/bitcoin/simnet/admin.macaroon /root/breez_node/tls.cert >> .env


/go/bin/migrate   -source file:///go/lspd/postgresql/migrations/ --database postgres://postgres:test@10.5.0.8:5432/postgres?sslmode=disable up 10
#wait for breez rpc
until cat /root/breez_node/logs/bitcoin/simnet/lnd.log | grep 'RPC server listening on' > /dev/null;
do
    sleep 1
    echo "lspd waiting for breez RPC..."
done

#Get node pub key
source .env

/go/bin/lncli --rpcserver=$LND_ADDRESS --macaroonpath=/root/breez_node/data/chain/bitcoin/simnet/admin.macaroon --tlscertpath=/root/breez_node/tls.cert -network=simnet getinfo
export LND_PUB_KEY=$(/go/bin/lncli --rpcserver=$LND_ADDRESS --macaroonpath=/root/breez_node/data/chain/bitcoin/simnet/admin.macaroon --tlscertpath=/root/breez_node/tls.cert -network=simnet getinfo | jq -r '.identity_pubkey')
echo "LND_PUB_KEY: $LND_PUB_KEY"

echo NODES=\'[ {\"name\": \"Breez\", \"nodePubkey\": \"$LND_PUB_KEY\", \"lspdPrivateKey\": \"$LSPD_PRIVATE_KEY\", \"token\": \"$TOKEN\", \"host\": \"$LISTEN_ADDRESS\", \"publicChannelAmount\": \"1000183\", \"channelAmount\": \"100000\", \"channelPrivate\": false, \"targetConf\": \"6\", \"minHtlcMsat\": \"600\", \"baseFeeMsat\": \"1000\", \"feeRate\": \"0.000001\", \"timeLockDelta\": \"144\", \"channelFeePermyriad\": \"40\", \"channelMinimumFeeMsat\": \"2000000\", \"additionalChannelCapacity\": \"100000\", \"maxInactiveDuration\": \"3888000\", \"lnd\": { \"address\": \"$LND_ADDRESS\", \"cert\": \"$LND_CERT\", \"macaroon\": \"$LND_MACAROON_HEX\" } } ]\' \
>> .env

echo "#### cat .env:"

cat .env

exec /go/bin/godotenv ./lspd/lspd