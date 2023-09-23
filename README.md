# LevYieldX
An API written in Go that finds cross-chain leveraged yield farming strategies based on your risk tolerance and portfolio.

Used GETH Abigen tool to generate bindings.

## Notes

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