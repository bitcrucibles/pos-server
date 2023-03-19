# breez
In order to build breez you will need to install [gomobile](https://github.com/golang/go/wiki/Mobile) and [go 1.17.x](https://go.dev/dl/). If you install go from homebrew, you will have to ensure the GOPATH environment variable is set yourself.
## Prepare your environment
```
git clone https://github.com/breez/breez.git
go install golang.org/x/mobile/cmd/gomobile
go install golang.org/x/mobile/cmd/gobind
export PATH=$PATH:$GOPATH/bin
gomobile init
```

## Building `breez` for Android
You need to install the ndk as part of your sdk Tools.
If you have a separate ndk installed then make sure to set the ANDROID_NDK_HOME environment variable to your ndk install location.
```
export ANDROID_HOME=<your android sdk directory>
```
Or in case you want to use a direct ndk path
```
export ANDROID_NDK_HOME=<your android ndk directory>
```
If you are using NDK 24+, install [NDK r19](https://github.com/android/ndk/wiki/Unsupported-Downloads#r19c) and point ANDROID_NDK_HOME to it's folder due to [gomobile incompatibility](https://github.com/golang/go/issues/35030). Alternatively, you can add [-androidapi 19](https://github.com/golang/go/issues/52470#issuecomment-1203998993) to gomobile command in build script.

If the library will be run on an emulator target, add `-ldflags=-extldflags=-Wl,-soname,libgojni.so` to gomobile command in build script.

Then you are ready to run the build:
```
./build.sh
```
The file breez.aar will be built in build/android/
## Building `breez` for iOS
```
./build-ios.sh
```
The bindings.framework will be built in build/ios/

# Running Breez dev environment on simnet
This article describes how to run breez infra in simnet: https://doc.breez.technology/Running-Breez-in-simnet.html#running-breez-in-simnet

Plus to the above fill in GOOGLE_APPLICATION_CREDENTIALS with Firebase details

Below picture shows interaction of all breez components:
![BreezComponentsInteraction](https://user-images.githubusercontent.com/124206069/219502847-cd37795d-c4bc-40bc-a8be-e791194c8b19.jpg)


```
sudo docker exec dev-breez "/lnd/lncli" -network=simnet getinfo

{
    "version": "0.15.4-beta commit=",
    "commit_hash": "",
    "identity_pubkey": "02add562a76e8ed044fb3dfb565dc94f12e7e52b19e7ed1cf76073ff9c936725e7",
```


```
sudo docker exec dev-btcd /start-btcctl.sh generate 6
sudo docker exec dev-alice "/root/go/bin/lncli" --network simnet newaddress np2wkh
```
Alice address: 
```
{
    "address": "rcSP2ujRMWjqaSPDG43CZetXQvqkZAP8Mr"
}
```

```
sudo docker exec dev-breez "/lnd/lncli" -network=simnet sendcoins rrEGRsE5mrqgnK1dV3HziAHj4NMA7MpkST 20000000
```

{
    "txid": "d0db9d1acaaab94824d374ce28c92d86b9d588a3fcf1120359d10da43229e5a7"
}
```

```
sudo docker exec dev-btcd /start-btcctl.sh generate 6
```

```
sudo docker exec dev-alice "/root/go/bin/lncli" -network=simnet openchannel --node_key 02add562a76e8ed044fb3dfb565dc94f12e7e52b19e7ed1cf76073ff9c936725e7 --local_amt=1000000 --connect=10.5.0.3:9735 --private

{
	"funding_txid": "1da38f92586fb739d5aa15c2e01bdce375de08b2c6f202007b415ba98cc73f75"
}
```

```
sudo docker exec dev-btcd /start-btcctl.sh generate 6
```

```
sudo docker exec dev-alice "/root/go/bin/lncli" -network=simnet payinvoice -f --pay_req <payment_request>
```

```
sudo docker exec dev-btcd /start-btcctl.sh generate 6
```

