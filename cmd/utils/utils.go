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
var MaxUint256, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)

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
