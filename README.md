# arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so
A repo for a Stackoverflow question, too much code to paste.

1. Make sure you have the following dependencies.

$ go version
go version go1.22.3 linux/amd64

$ solc --version
solc, the solidity compiler commandline interface
Version: 0.8.26+commit.8a97fa7a.Linux.g++

$ abigen --version
abigen version 1.14.4-unstable-cc22e0cd-20240528

2. Save .env.example as .env and enter your credentials.
3. Run `go mod tidy`
4. Run `go run *.go`
5. The termial will show the following:

Setting up...<br>
Wallet...<br>
Config...<br>
Connection...<br>
Contract compiled successfully<br>
Go bindings generated successfully<br>
Contract deploying...<br>
auth.GasLimit:  1070172<br>
calculatedCost:  1070537484071268<br>
pendingBalance:  3038069886154048302<br>
suggestedGasPrice:  1000341519<br>
gasLimit:  1070172<br>
auth.GasLimit:  1070172<br>
auth.GasPrice:  1000341519<br>
nonce:  6858<br>
----After deploy----<br>
Token deployed at: 0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2<br>
Token tx: 0xe38e6a42251705b2355dca08049332b88ceda27a467dd9497398d5d3e22e7417<br>
Token cost: 1070537484071268<br>
Token gas limit: 1070172<br>
Token gas price: 1000341519<br>
Waiting to be mined...<br>
Mined...<br>
Token deployed successfully<br>
Contract verifying...<br>
****************************ERROR****************************************<br>
The deployed contract bytecode does not match the compiled bytecode.<br>
Length deployed contract bytecode 935 Length compiled bytecode 1394<br>
****************************END ERROR****************************************<br>
Running code...<br>
Add result: &{824633950592 {13947092382033123707 14682017799 11704640} {[] {} 0} {{} {} 0} {[] {} 0}}<br>
Concat result: abcqwBBBBBhuguty<br>
SetQ transaction hash: 0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7<br>
Waiting for transaction to be mined...<br>
Transaction 0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7 mined in block 6002728<br>
Transaction Receipt: &{Type:2 Status:1 CumulativeGasUsed:11068867 Bloom:[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 64 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 
Logs:[0xc0000f8580]<br>
TransactionHash:0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7<br> 
ContractAddress:0x0000000000000000000000000000000000000000<br> 
GasUsed:44652<br> 
BlockHash:0x3883a2bdb09dfb3e5ae71ad1d0c0202fe4abeddc222979d3955911ae7a42d320<br> 
BlockNumber:+6002728 TransactionIndex:80}<br>
GetQ result: 88888999999<br>

6. In this case the contract was deployed on the Sepolia chain with the Go code at the address `0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2`, then `sol/Base.sol` was verified with Compiler Version: v0.8.26+commit.8a97fa7a, Optimization Enabled: 1, Runs: 200.
Here's the problem, the Go code fails to verify the contract that can be verified on Etherscan using the Etherscan verification process.
https://sepolia.etherscan.io/address/0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2#code
7. When you replicate you will get a different address and transaction hash, that is the only difference from this example.
8. Can you find the error?


