package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ChainConfig struct {
	RPCEndpoint string            `json:"rpcEndpoint"`
	Tokens      map[string]string `json:"tokens"`
}

var ChainConfigs = make(map[string]*ChainConfig)

// Constants
var MaxUint64 = new(big.Int).SetUint64(^uint64(0))
var MaxUint256, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
var SecPerYear = big.NewInt(60 * 60 * 24 * 365)
var ETHMantissaFloat = new(big.Float).SetUint64(1000000000000000000) // 10**18
var ETHMantissaInt = new(big.Int).SetUint64(1000000000000000000)     // 10**18
var ETHBlocksPerDay = big.NewFloat(7200)
var Ray = new(big.Int).Exp(big.NewInt(10), big.NewInt(27), nil) // 10**27
var HalfRay = new(big.Int).Div(Ray, big.NewInt(2))
var PercentageFactor = big.NewInt(10000) // Used for AaveV3 math, 10**4
var HalfPercentageFactor = big.NewInt(5000)

func init() {
	// Parse all config json files
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	configPath := filepath.Join(currentDir, "configs")
	configDir, _ := os.ReadDir(configPath)
	for _, config := range configDir {
		log.Printf("Loading %v...", config.Name())
		if err := loadConfig(config.Name()); err != nil {
			log.Printf("Failed to load %v: %v", config.Name(), err)
		}
	}
}

// Loads the chain config including rpc endpoint and token mappings.
func loadConfig(configName string) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "configs", configName)
	rawConfig, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open config: %v", err)
		return err
	}
	defer rawConfig.Close()
	var parsedConfig ChainConfig
	readConfig, err := io.ReadAll(rawConfig)
	if err != nil {
		log.Printf("Failed to read config: %v", err)
	}
	err = json.Unmarshal(readConfig, &parsedConfig)
	if err != nil {
		log.Printf("Failed to parse config: %v", err)
		return err
	}

	ChainConfigs[strings.TrimSuffix(configName, ".json")] = &parsedConfig
	return nil
}

func ConvertAddressToSymbol(chain string, address string) (string, error) {
	config, ok := ChainConfigs[chain]
	if !ok {
		return "", fmt.Errorf("could not find %v config", chain)
	}

	// Reverse mapping O(tokens)
	reversedMapping := make(map[string]string)
	for token, address := range config.Tokens {
		reversedMapping[address] = token
	}

	// Convert address to symbols O(addresses)
	symbol, ok := reversedMapping[address]
	if ok {
		return symbol, nil
	}
	return "", fmt.Errorf("could not find symbol for %v", address)
}

// Convert 18 decimals to 27 decimals
func WadToRay(wad *big.Int) *big.Int {
	return new(big.Int).Mul(wad, big.NewInt(1000000000))
}

// Multiplies two rays, returns a*b/10^27
func RayMul(a, b *big.Int) *big.Int {
	result := new(big.Int).Mul(a, b)
	result.Add(result, HalfRay)
	return result.Quo(result, Ray)
}

// Divides two rays, returns a*10^27/b
func RayDiv(a, b *big.Int) *big.Int {
	result := new(big.Int).Mul(a, Ray)
	quo := new(big.Int).Quo(b, big.NewInt(2))
	result.Add(result, quo)
	return result.Quo(result, b)
}

// AaveV3 math, returns a*b/10^4
func PercentMul(value, percentage *big.Int) *big.Int {
	threshold := new(big.Int).Sub(MaxUint256, HalfPercentageFactor)
	threshold.Div(threshold, percentage)

	if percentage.Cmp(big.NewInt(0)) == 0 || value.Cmp(threshold) == 1 {
		panic("Invalid PercentMul args")
	}

	prod := new(big.Int).Mul(value, percentage)
	sum := prod.Add(prod, HalfPercentageFactor)
	return sum.Quo(prod, PercentageFactor)
}

// Multiplies two mantissas, returns a*b/10^18
func ManMul(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, b)
	return prod.Div(prod, ETHMantissaInt)
}

// Divides two mantissas, returns a*10^18/b
func ManDiv(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, ETHMantissaInt)
	return prod.Div(prod, b)
}

// Multiplies two basis points, returns a*b/10^4
func BasisMul(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, b)
	return prod.Div(prod, PercentageFactor)
}

// Divides two basis points, returns a*10^4/b
func BasisDiv(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, PercentageFactor)
	return prod.Div(prod, b)
}

// Multiplies a price, returns a*b/10^8
func PriceMul(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, b)
	return prod.Div(prod, ETHMantissaInt)
}

// Divides a price, returns a*10^8/b
func PriceDiv(a, b *big.Int) *big.Int {
	prod := new(big.Int).Mul(a, ETHMantissaInt)
	return prod.Div(prod, b)
}
