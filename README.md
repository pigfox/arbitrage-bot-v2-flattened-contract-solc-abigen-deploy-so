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
`
Setting up...
Wallet...
Config...
Connection...
Contract compiled successfully
Go bindings generated successfully
Contract deploying...
auth.GasLimit:  1070172
calculatedCost:  1070537484071268
pendingBalance:  3038069886154048302
suggestedGasPrice:  1000341519
gasLimit:  1070172
auth.GasLimit:  1070172
auth.GasPrice:  1000341519
nonce:  6858
----After deploy----
Token deployed at: 0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2
Token tx: 0xe38e6a42251705b2355dca08049332b88ceda27a467dd9497398d5d3e22e7417
Token cost: 1070537484071268
Token gas limit: 1070172
Token gas price: 1000341519
Waiting to be mined...
Mined...
Token deployed successfully
Contract verifying...
****************************ERROR****************************************
The deployed contract bytecode does not match the compiled bytecode.
Length deployed contract bytecode 935 Length compiled bytecode 1394
****************************END ERROR****************************************
Running code...
Add result: &{824633950592 {13947092382033123707 14682017799 11704640} {[] {} 0} {{} {} 0} {[] {} 0}}
Concat result: abcqwBBBBBhuguty
SetQ transaction hash: 0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7
Waiting for transaction to be mined...
Transaction 0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7 mined in block 6002728
Transaction Receipt: &{Type:2 Status:1 CumulativeGasUsed:11068867 Bloom:[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 64 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Logs:[0xc0000f8580] TransactionHash:0x797eb898d43f8d37d214ac27e3d4eba15c0a80df7d386ff6f9c5b9bc0b9e27b7 ContractAddress:0x0000000000000000000000000000000000000000 GasUsed:44652 BlockHash:0x3883a2bdb09dfb3e5ae71ad1d0c0202fe4abeddc222979d3955911ae7a42d320 BlockNumber:+6002728 TransactionIndex:80}
GetQ result: 88888999999
`

6. In this case the contract was deployed on the Sepolia chain with the Go code at the address `0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2`, then `sol/Base.sol` was verified with Compiler Version: v0.8.26+commit.8a97fa7a, Optimization Enabled: 1, Runs: 200.
Here's the problem, the Go code fails to verify the contract that can be verified on Etherscan using the Etherscan verification process.
https://sepolia.etherscan.io/address/0x7ede1c76aD55d8c67c8a9E3786C63e00711e04a2#code
7. When you replicate you will get a different address, that is the only differnce form this example.
8. Can you find the error?






