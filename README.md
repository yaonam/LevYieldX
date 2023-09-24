# LevYieldX
An API written in Go that finds cross-chain leveraged yield farming strategies based on your risk tolerance and portfolio.

This project is an API that provides all the information necessary to identify the best cross-chain leveraged yield farming strategies and to execute them. Users can input the liquidity amount, preferred base collateral tokens, leverage usage, chains, and protocols. The API will present all possible strategies, and once the user has selected one, all the transactions required to execute it.

```
Example: Bob has $10,000 USDT. The API found two profitable strategies.

// Total APY = 5%, Gas Cost = $0.05
1. Supply $10,000 USDT to AaveV3 on Arbitrum for 5% APY.

// Total APY = 7.4%, Gas Cost = $15.00
1. Supply $10,000 USDT to AaveV3 on Arbitrum for 5% APY.
2. Borrow $8,000 WETH from AaveV3 on Arbitrum for 3% APY.
3. Bridge WETH to Ethereum.
4. Supply $8,000 WETH to CompoundV3 on Ethereum for 6% APY.
```

## API

### GET /strats

Returns all strategies sorted by APY in descending order.

### POST /transactions

Returns all transactions required to execute the listed strategies.

## Notes

Used GETH Abigen tool to generate bindings.

Modules -> Packages -> Files

`nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go`

For Windows:

`nodemon --watch './**/*.go' -e go,json --signal SIGKILL --exec go run .`

`go mod init {MODULE_NAME}` to create project dependency file
`go get .` to automatically track dependencies (based on imports)
`go run .` to run server
`go mode tidy` to remove unused dependencies

Uses Multicall to reduce RPC calls. (https://github.com/mds1/multicall)

If need to retrieve non-view function output, just set the stateMutability in the json to view.