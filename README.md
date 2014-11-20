#blockchainImport
blockchainImport is a little script that let you import your bitcoin private keys from your [blockchain.info](https://blockchain.info/) wallet into the bitcoin core desktop application

##Installation
````bash
go get -u github.com/zaibon/blockchainImport
````

##Usage
1. First you need to export your private keys from your blockchain.info wallet
2. Then simply run the script
````bash
blockchainImport -key keys.json -conf /path/to/bitcoin.conf -bin /path/to/bitcoin-cli/binary
````