// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundv3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CometConfigurationAssetConfig is an auto generated low-level Go binding around an user-defined struct.
type CometConfigurationAssetConfig struct {
	Asset                     common.Address
	PriceFeed                 common.Address
	Decimals                  uint8
	BorrowCollateralFactor    uint64
	LiquidateCollateralFactor uint64
	LiquidationFactor         uint64
	SupplyCap                 *big.Int
}

// CometConfigurationConfiguration is an auto generated low-level Go binding around an user-defined struct.
type CometConfigurationConfiguration struct {
	Governor                           common.Address
	PauseGuardian                      common.Address
	BaseToken                          common.Address
	BaseTokenPriceFeed                 common.Address
	ExtensionDelegate                  common.Address
	SupplyKink                         uint64
	SupplyPerYearInterestRateSlopeLow  uint64
	SupplyPerYearInterestRateSlopeHigh uint64
	SupplyPerYearInterestRateBase      uint64
	BorrowKink                         uint64
	BorrowPerYearInterestRateSlopeLow  uint64
	BorrowPerYearInterestRateSlopeHigh uint64
	BorrowPerYearInterestRateBase      uint64
	StoreFrontPriceFactor              uint64
	TrackingIndexScale                 uint64
	BaseTrackingSupplySpeed            uint64
	BaseTrackingBorrowSpeed            uint64
	BaseMinForRewards                  *big.Int
	BaseBorrowMin                      *big.Int
	TargetReserves                     *big.Int
	AssetConfigs                       []CometConfigurationAssetConfig
}

// ConfiguratorMetaData contains all meta data concerning the Configurator contract.
var ConfiguratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigurationAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structCometConfiguration.AssetConfig\",\"name\":\"assetConfig\",\"type\":\"tuple\"}],\"name\":\"AddAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newComet\",\"type\":\"address\"}],\"name\":\"CometDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"GovernorTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"oldBaseBorrowMin\",\"type\":\"uint104\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"newBaseBorrowMin\",\"type\":\"uint104\"}],\"name\":\"SetBaseBorrowMin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"oldBaseMinForRewards\",\"type\":\"uint104\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"newBaseMinForRewards\",\"type\":\"uint104\"}],\"name\":\"SetBaseMinForRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBaseTokenPriceFeed\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBaseTokenPriceFeed\",\"type\":\"address\"}],\"name\":\"SetBaseTokenPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldBaseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBaseTrackingBorrowSpeed\",\"type\":\"uint64\"}],\"name\":\"SetBaseTrackingBorrowSpeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldBaseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBaseTrackingSupplySpeed\",\"type\":\"uint64\"}],\"name\":\"SetBaseTrackingSupplySpeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldKink\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newKink\",\"type\":\"uint64\"}],\"name\":\"SetBorrowKink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRBase\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRBase\",\"type\":\"uint64\"}],\"name\":\"SetBorrowPerYearInterestRateBase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRSlopeHigh\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRSlopeHigh\",\"type\":\"uint64\"}],\"name\":\"SetBorrowPerYearInterestRateSlopeHigh\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRSlopeLow\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRSlopeLow\",\"type\":\"uint64\"}],\"name\":\"SetBorrowPerYearInterestRateSlopeLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structCometConfiguration.Configuration\",\"name\":\"oldConfiguration\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structCometConfiguration.Configuration\",\"name\":\"newConfiguration\",\"type\":\"tuple\"}],\"name\":\"SetConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldExt\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExt\",\"type\":\"address\"}],\"name\":\"SetExtensionDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"SetFactory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"SetGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"SetPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldStoreFrontPriceFactor\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newStoreFrontPriceFactor\",\"type\":\"uint64\"}],\"name\":\"SetStoreFrontPriceFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldKink\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newKink\",\"type\":\"uint64\"}],\"name\":\"SetSupplyKink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRBase\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRBase\",\"type\":\"uint64\"}],\"name\":\"SetSupplyPerYearInterestRateBase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRSlopeHigh\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRSlopeHigh\",\"type\":\"uint64\"}],\"name\":\"SetSupplyPerYearInterestRateSlopeHigh\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldIRSlopeLow\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newIRSlopeLow\",\"type\":\"uint64\"}],\"name\":\"SetSupplyPerYearInterestRateSlopeLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"oldTargetReserves\",\"type\":\"uint104\"},{\"indexed\":false,\"internalType\":\"uint104\",\"name\":\"newTargetReserves\",\"type\":\"uint104\"}],\"name\":\"SetTargetReserves\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structCometConfiguration.AssetConfig\",\"name\":\"oldAssetConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structCometConfiguration.AssetConfig\",\"name\":\"newAssetConfig\",\"type\":\"tuple\"}],\"name\":\"UpdateAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldBorrowCF\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBorrowCF\",\"type\":\"uint64\"}],\"name\":\"UpdateAssetBorrowCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldLiquidateCF\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLiquidateCF\",\"type\":\"uint64\"}],\"name\":\"UpdateAssetLiquidateCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldLiquidationFactor\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLiquidationFactor\",\"type\":\"uint64\"}],\"name\":\"UpdateAssetLiquidationFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPriceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceFeed\",\"type\":\"address\"}],\"name\":\"UpdateAssetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"oldSupplyCap\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newSupplyCap\",\"type\":\"uint128\"}],\"name\":\"UpdateAssetSupplyCap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig\",\"name\":\"assetConfig\",\"type\":\"tuple\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structCometConfiguration.Configuration\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"newBaseBorrowMin\",\"type\":\"uint104\"}],\"name\":\"setBaseBorrowMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"newBaseMinForRewards\",\"type\":\"uint104\"}],\"name\":\"setBaseMinForRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newBaseTokenPriceFeed\",\"type\":\"address\"}],\"name\":\"setBaseTokenPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBaseTrackingBorrowSpeed\",\"type\":\"uint64\"}],\"name\":\"setBaseTrackingBorrowSpeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBaseTrackingSupplySpeed\",\"type\":\"uint64\"}],\"name\":\"setBaseTrackingSupplySpeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBorrowKink\",\"type\":\"uint64\"}],\"name\":\"setBorrowKink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBase\",\"type\":\"uint64\"}],\"name\":\"setBorrowPerYearInterestRateBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSlope\",\"type\":\"uint64\"}],\"name\":\"setBorrowPerYearInterestRateSlopeHigh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSlope\",\"type\":\"uint64\"}],\"name\":\"setBorrowPerYearInterestRateSlopeLow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structCometConfiguration.Configuration\",\"name\":\"newConfiguration\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newExtensionDelegate\",\"type\":\"address\"}],\"name\":\"setExtensionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFactory\",\"type\":\"address\"}],\"name\":\"setFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"setPauseGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newStoreFrontPriceFactor\",\"type\":\"uint64\"}],\"name\":\"setStoreFrontPriceFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSupplyKink\",\"type\":\"uint64\"}],\"name\":\"setSupplyKink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBase\",\"type\":\"uint64\"}],\"name\":\"setSupplyPerYearInterestRateBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSlope\",\"type\":\"uint64\"}],\"name\":\"setSupplyPerYearInterestRateSlopeHigh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newSlope\",\"type\":\"uint64\"}],\"name\":\"setSupplyPerYearInterestRateSlopeLow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"newTargetReserves\",\"type\":\"uint104\"}],\"name\":\"setTargetReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"transferGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig\",\"name\":\"newAssetConfig\",\"type\":\"tuple\"}],\"name\":\"updateAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newBorrowCF\",\"type\":\"uint64\"}],\"name\":\"updateAssetBorrowCollateralFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newLiquidateCF\",\"type\":\"uint64\"}],\"name\":\"updateAssetLiquidateCollateralFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newLiquidationFactor\",\"type\":\"uint64\"}],\"name\":\"updateAssetLiquidationFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newPriceFeed\",\"type\":\"address\"}],\"name\":\"updateAssetPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cometProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"newSupplyCap\",\"type\":\"uint128\"}],\"name\":\"updateAssetSupplyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConfiguratorABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfiguratorMetaData.ABI instead.
var ConfiguratorABI = ConfiguratorMetaData.ABI

// Configurator is an auto generated Go binding around an Ethereum contract.
type Configurator struct {
	ConfiguratorCaller     // Read-only binding to the contract
	ConfiguratorTransactor // Write-only binding to the contract
	ConfiguratorFilterer   // Log filterer for contract events
}

// ConfiguratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfiguratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfiguratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfiguratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfiguratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfiguratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfiguratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfiguratorSession struct {
	Contract     *Configurator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfiguratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfiguratorCallerSession struct {
	Contract *ConfiguratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ConfiguratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfiguratorTransactorSession struct {
	Contract     *ConfiguratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ConfiguratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfiguratorRaw struct {
	Contract *Configurator // Generic contract binding to access the raw methods on
}

// ConfiguratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfiguratorCallerRaw struct {
	Contract *ConfiguratorCaller // Generic read-only contract binding to access the raw methods on
}

// ConfiguratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfiguratorTransactorRaw struct {
	Contract *ConfiguratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfigurator creates a new instance of Configurator, bound to a specific deployed contract.
func NewConfigurator(address common.Address, backend bind.ContractBackend) (*Configurator, error) {
	contract, err := bindConfigurator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Configurator{ConfiguratorCaller: ConfiguratorCaller{contract: contract}, ConfiguratorTransactor: ConfiguratorTransactor{contract: contract}, ConfiguratorFilterer: ConfiguratorFilterer{contract: contract}}, nil
}

// NewConfiguratorCaller creates a new read-only instance of Configurator, bound to a specific deployed contract.
func NewConfiguratorCaller(address common.Address, caller bind.ContractCaller) (*ConfiguratorCaller, error) {
	contract, err := bindConfigurator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorCaller{contract: contract}, nil
}

// NewConfiguratorTransactor creates a new write-only instance of Configurator, bound to a specific deployed contract.
func NewConfiguratorTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfiguratorTransactor, error) {
	contract, err := bindConfigurator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorTransactor{contract: contract}, nil
}

// NewConfiguratorFilterer creates a new log filterer instance of Configurator, bound to a specific deployed contract.
func NewConfiguratorFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfiguratorFilterer, error) {
	contract, err := bindConfigurator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorFilterer{contract: contract}, nil
}

// bindConfigurator binds a generic wrapper to an already deployed contract.
func bindConfigurator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConfiguratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Configurator *ConfiguratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurator.Contract.ConfiguratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Configurator *ConfiguratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurator.Contract.ConfiguratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Configurator *ConfiguratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurator.Contract.ConfiguratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Configurator *ConfiguratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Configurator *ConfiguratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Configurator *ConfiguratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurator.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0x395c0fda.
//
// Solidity: function factory(address ) view returns(address)
func (_Configurator *ConfiguratorCaller) Factory(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "factory", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0x395c0fda.
//
// Solidity: function factory(address ) view returns(address)
func (_Configurator *ConfiguratorSession) Factory(arg0 common.Address) (common.Address, error) {
	return _Configurator.Contract.Factory(&_Configurator.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0x395c0fda.
//
// Solidity: function factory(address ) view returns(address)
func (_Configurator *ConfiguratorCallerSession) Factory(arg0 common.Address) (common.Address, error) {
	return _Configurator.Contract.Factory(&_Configurator.CallOpts, arg0)
}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address cometProxy, address asset) view returns(uint256)
func (_Configurator *ConfiguratorCaller) GetAssetIndex(opts *bind.CallOpts, cometProxy common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "getAssetIndex", cometProxy, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address cometProxy, address asset) view returns(uint256)
func (_Configurator *ConfiguratorSession) GetAssetIndex(cometProxy common.Address, asset common.Address) (*big.Int, error) {
	return _Configurator.Contract.GetAssetIndex(&_Configurator.CallOpts, cometProxy, asset)
}

// GetAssetIndex is a free data retrieval call binding the contract method 0x886fe70b.
//
// Solidity: function getAssetIndex(address cometProxy, address asset) view returns(uint256)
func (_Configurator *ConfiguratorCallerSession) GetAssetIndex(cometProxy common.Address, asset common.Address) (*big.Int, error) {
	return _Configurator.Contract.GetAssetIndex(&_Configurator.CallOpts, cometProxy, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address cometProxy) view returns((address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]))
func (_Configurator *ConfiguratorCaller) GetConfiguration(opts *bind.CallOpts, cometProxy common.Address) (CometConfigurationConfiguration, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "getConfiguration", cometProxy)

	if err != nil {
		return *new(CometConfigurationConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new(CometConfigurationConfiguration)).(*CometConfigurationConfiguration)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address cometProxy) view returns((address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]))
func (_Configurator *ConfiguratorSession) GetConfiguration(cometProxy common.Address) (CometConfigurationConfiguration, error) {
	return _Configurator.Contract.GetConfiguration(&_Configurator.CallOpts, cometProxy)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address cometProxy) view returns((address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]))
func (_Configurator *ConfiguratorCallerSession) GetConfiguration(cometProxy common.Address) (CometConfigurationConfiguration, error) {
	return _Configurator.Contract.GetConfiguration(&_Configurator.CallOpts, cometProxy)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Configurator *ConfiguratorCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Configurator *ConfiguratorSession) Governor() (common.Address, error) {
	return _Configurator.Contract.Governor(&_Configurator.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Configurator *ConfiguratorCallerSession) Governor() (common.Address, error) {
	return _Configurator.Contract.Governor(&_Configurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Configurator *ConfiguratorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Configurator *ConfiguratorSession) Version() (*big.Int, error) {
	return _Configurator.Contract.Version(&_Configurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Configurator *ConfiguratorCallerSession) Version() (*big.Int, error) {
	return _Configurator.Contract.Version(&_Configurator.CallOpts)
}

// AddAsset is a paid mutator transaction binding the contract method 0xea31a447.
//
// Solidity: function addAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig) returns()
func (_Configurator *ConfiguratorTransactor) AddAsset(opts *bind.TransactOpts, cometProxy common.Address, assetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "addAsset", cometProxy, assetConfig)
}

// AddAsset is a paid mutator transaction binding the contract method 0xea31a447.
//
// Solidity: function addAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig) returns()
func (_Configurator *ConfiguratorSession) AddAsset(cometProxy common.Address, assetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.Contract.AddAsset(&_Configurator.TransactOpts, cometProxy, assetConfig)
}

// AddAsset is a paid mutator transaction binding the contract method 0xea31a447.
//
// Solidity: function addAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig) returns()
func (_Configurator *ConfiguratorTransactorSession) AddAsset(cometProxy common.Address, assetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.Contract.AddAsset(&_Configurator.TransactOpts, cometProxy, assetConfig)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address cometProxy) returns(address)
func (_Configurator *ConfiguratorTransactor) Deploy(opts *bind.TransactOpts, cometProxy common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "deploy", cometProxy)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address cometProxy) returns(address)
func (_Configurator *ConfiguratorSession) Deploy(cometProxy common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.Deploy(&_Configurator.TransactOpts, cometProxy)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address cometProxy) returns(address)
func (_Configurator *ConfiguratorTransactorSession) Deploy(cometProxy common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.Deploy(&_Configurator.TransactOpts, cometProxy)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address governor_) returns()
func (_Configurator *ConfiguratorTransactor) Initialize(opts *bind.TransactOpts, governor_ common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "initialize", governor_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address governor_) returns()
func (_Configurator *ConfiguratorSession) Initialize(governor_ common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.Initialize(&_Configurator.TransactOpts, governor_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address governor_) returns()
func (_Configurator *ConfiguratorTransactorSession) Initialize(governor_ common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.Initialize(&_Configurator.TransactOpts, governor_)
}

// SetBaseBorrowMin is a paid mutator transaction binding the contract method 0x9a01fe82.
//
// Solidity: function setBaseBorrowMin(address cometProxy, uint104 newBaseBorrowMin) returns()
func (_Configurator *ConfiguratorTransactor) SetBaseBorrowMin(opts *bind.TransactOpts, cometProxy common.Address, newBaseBorrowMin *big.Int) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBaseBorrowMin", cometProxy, newBaseBorrowMin)
}

// SetBaseBorrowMin is a paid mutator transaction binding the contract method 0x9a01fe82.
//
// Solidity: function setBaseBorrowMin(address cometProxy, uint104 newBaseBorrowMin) returns()
func (_Configurator *ConfiguratorSession) SetBaseBorrowMin(cometProxy common.Address, newBaseBorrowMin *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseBorrowMin(&_Configurator.TransactOpts, cometProxy, newBaseBorrowMin)
}

// SetBaseBorrowMin is a paid mutator transaction binding the contract method 0x9a01fe82.
//
// Solidity: function setBaseBorrowMin(address cometProxy, uint104 newBaseBorrowMin) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBaseBorrowMin(cometProxy common.Address, newBaseBorrowMin *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseBorrowMin(&_Configurator.TransactOpts, cometProxy, newBaseBorrowMin)
}

// SetBaseMinForRewards is a paid mutator transaction binding the contract method 0xd89e834f.
//
// Solidity: function setBaseMinForRewards(address cometProxy, uint104 newBaseMinForRewards) returns()
func (_Configurator *ConfiguratorTransactor) SetBaseMinForRewards(opts *bind.TransactOpts, cometProxy common.Address, newBaseMinForRewards *big.Int) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBaseMinForRewards", cometProxy, newBaseMinForRewards)
}

// SetBaseMinForRewards is a paid mutator transaction binding the contract method 0xd89e834f.
//
// Solidity: function setBaseMinForRewards(address cometProxy, uint104 newBaseMinForRewards) returns()
func (_Configurator *ConfiguratorSession) SetBaseMinForRewards(cometProxy common.Address, newBaseMinForRewards *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseMinForRewards(&_Configurator.TransactOpts, cometProxy, newBaseMinForRewards)
}

// SetBaseMinForRewards is a paid mutator transaction binding the contract method 0xd89e834f.
//
// Solidity: function setBaseMinForRewards(address cometProxy, uint104 newBaseMinForRewards) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBaseMinForRewards(cometProxy common.Address, newBaseMinForRewards *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseMinForRewards(&_Configurator.TransactOpts, cometProxy, newBaseMinForRewards)
}

// SetBaseTokenPriceFeed is a paid mutator transaction binding the contract method 0xe1a533ed.
//
// Solidity: function setBaseTokenPriceFeed(address cometProxy, address newBaseTokenPriceFeed) returns()
func (_Configurator *ConfiguratorTransactor) SetBaseTokenPriceFeed(opts *bind.TransactOpts, cometProxy common.Address, newBaseTokenPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBaseTokenPriceFeed", cometProxy, newBaseTokenPriceFeed)
}

// SetBaseTokenPriceFeed is a paid mutator transaction binding the contract method 0xe1a533ed.
//
// Solidity: function setBaseTokenPriceFeed(address cometProxy, address newBaseTokenPriceFeed) returns()
func (_Configurator *ConfiguratorSession) SetBaseTokenPriceFeed(cometProxy common.Address, newBaseTokenPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTokenPriceFeed(&_Configurator.TransactOpts, cometProxy, newBaseTokenPriceFeed)
}

// SetBaseTokenPriceFeed is a paid mutator transaction binding the contract method 0xe1a533ed.
//
// Solidity: function setBaseTokenPriceFeed(address cometProxy, address newBaseTokenPriceFeed) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBaseTokenPriceFeed(cometProxy common.Address, newBaseTokenPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTokenPriceFeed(&_Configurator.TransactOpts, cometProxy, newBaseTokenPriceFeed)
}

// SetBaseTrackingBorrowSpeed is a paid mutator transaction binding the contract method 0x8d2ed08d.
//
// Solidity: function setBaseTrackingBorrowSpeed(address cometProxy, uint64 newBaseTrackingBorrowSpeed) returns()
func (_Configurator *ConfiguratorTransactor) SetBaseTrackingBorrowSpeed(opts *bind.TransactOpts, cometProxy common.Address, newBaseTrackingBorrowSpeed uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBaseTrackingBorrowSpeed", cometProxy, newBaseTrackingBorrowSpeed)
}

// SetBaseTrackingBorrowSpeed is a paid mutator transaction binding the contract method 0x8d2ed08d.
//
// Solidity: function setBaseTrackingBorrowSpeed(address cometProxy, uint64 newBaseTrackingBorrowSpeed) returns()
func (_Configurator *ConfiguratorSession) SetBaseTrackingBorrowSpeed(cometProxy common.Address, newBaseTrackingBorrowSpeed uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTrackingBorrowSpeed(&_Configurator.TransactOpts, cometProxy, newBaseTrackingBorrowSpeed)
}

// SetBaseTrackingBorrowSpeed is a paid mutator transaction binding the contract method 0x8d2ed08d.
//
// Solidity: function setBaseTrackingBorrowSpeed(address cometProxy, uint64 newBaseTrackingBorrowSpeed) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBaseTrackingBorrowSpeed(cometProxy common.Address, newBaseTrackingBorrowSpeed uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTrackingBorrowSpeed(&_Configurator.TransactOpts, cometProxy, newBaseTrackingBorrowSpeed)
}

// SetBaseTrackingSupplySpeed is a paid mutator transaction binding the contract method 0x240dd96d.
//
// Solidity: function setBaseTrackingSupplySpeed(address cometProxy, uint64 newBaseTrackingSupplySpeed) returns()
func (_Configurator *ConfiguratorTransactor) SetBaseTrackingSupplySpeed(opts *bind.TransactOpts, cometProxy common.Address, newBaseTrackingSupplySpeed uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBaseTrackingSupplySpeed", cometProxy, newBaseTrackingSupplySpeed)
}

// SetBaseTrackingSupplySpeed is a paid mutator transaction binding the contract method 0x240dd96d.
//
// Solidity: function setBaseTrackingSupplySpeed(address cometProxy, uint64 newBaseTrackingSupplySpeed) returns()
func (_Configurator *ConfiguratorSession) SetBaseTrackingSupplySpeed(cometProxy common.Address, newBaseTrackingSupplySpeed uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTrackingSupplySpeed(&_Configurator.TransactOpts, cometProxy, newBaseTrackingSupplySpeed)
}

// SetBaseTrackingSupplySpeed is a paid mutator transaction binding the contract method 0x240dd96d.
//
// Solidity: function setBaseTrackingSupplySpeed(address cometProxy, uint64 newBaseTrackingSupplySpeed) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBaseTrackingSupplySpeed(cometProxy common.Address, newBaseTrackingSupplySpeed uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBaseTrackingSupplySpeed(&_Configurator.TransactOpts, cometProxy, newBaseTrackingSupplySpeed)
}

// SetBorrowKink is a paid mutator transaction binding the contract method 0x5bfb8373.
//
// Solidity: function setBorrowKink(address cometProxy, uint64 newBorrowKink) returns()
func (_Configurator *ConfiguratorTransactor) SetBorrowKink(opts *bind.TransactOpts, cometProxy common.Address, newBorrowKink uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBorrowKink", cometProxy, newBorrowKink)
}

// SetBorrowKink is a paid mutator transaction binding the contract method 0x5bfb8373.
//
// Solidity: function setBorrowKink(address cometProxy, uint64 newBorrowKink) returns()
func (_Configurator *ConfiguratorSession) SetBorrowKink(cometProxy common.Address, newBorrowKink uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowKink(&_Configurator.TransactOpts, cometProxy, newBorrowKink)
}

// SetBorrowKink is a paid mutator transaction binding the contract method 0x5bfb8373.
//
// Solidity: function setBorrowKink(address cometProxy, uint64 newBorrowKink) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBorrowKink(cometProxy common.Address, newBorrowKink uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowKink(&_Configurator.TransactOpts, cometProxy, newBorrowKink)
}

// SetBorrowPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x74b20ab8.
//
// Solidity: function setBorrowPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorTransactor) SetBorrowPerYearInterestRateBase(opts *bind.TransactOpts, cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBorrowPerYearInterestRateBase", cometProxy, newBase)
}

// SetBorrowPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x74b20ab8.
//
// Solidity: function setBorrowPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorSession) SetBorrowPerYearInterestRateBase(cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateBase(&_Configurator.TransactOpts, cometProxy, newBase)
}

// SetBorrowPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x74b20ab8.
//
// Solidity: function setBorrowPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBorrowPerYearInterestRateBase(cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateBase(&_Configurator.TransactOpts, cometProxy, newBase)
}

// SetBorrowPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0xe8623ec5.
//
// Solidity: function setBorrowPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactor) SetBorrowPerYearInterestRateSlopeHigh(opts *bind.TransactOpts, cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBorrowPerYearInterestRateSlopeHigh", cometProxy, newSlope)
}

// SetBorrowPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0xe8623ec5.
//
// Solidity: function setBorrowPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorSession) SetBorrowPerYearInterestRateSlopeHigh(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateSlopeHigh(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetBorrowPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0xe8623ec5.
//
// Solidity: function setBorrowPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBorrowPerYearInterestRateSlopeHigh(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateSlopeHigh(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetBorrowPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0xb2f5ee19.
//
// Solidity: function setBorrowPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactor) SetBorrowPerYearInterestRateSlopeLow(opts *bind.TransactOpts, cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setBorrowPerYearInterestRateSlopeLow", cometProxy, newSlope)
}

// SetBorrowPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0xb2f5ee19.
//
// Solidity: function setBorrowPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorSession) SetBorrowPerYearInterestRateSlopeLow(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateSlopeLow(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetBorrowPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0xb2f5ee19.
//
// Solidity: function setBorrowPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactorSession) SetBorrowPerYearInterestRateSlopeLow(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetBorrowPerYearInterestRateSlopeLow(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc1252f90.
//
// Solidity: function setConfiguration(address cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration) returns()
func (_Configurator *ConfiguratorTransactor) SetConfiguration(opts *bind.TransactOpts, cometProxy common.Address, newConfiguration CometConfigurationConfiguration) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setConfiguration", cometProxy, newConfiguration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc1252f90.
//
// Solidity: function setConfiguration(address cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration) returns()
func (_Configurator *ConfiguratorSession) SetConfiguration(cometProxy common.Address, newConfiguration CometConfigurationConfiguration) (*types.Transaction, error) {
	return _Configurator.Contract.SetConfiguration(&_Configurator.TransactOpts, cometProxy, newConfiguration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xc1252f90.
//
// Solidity: function setConfiguration(address cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration) returns()
func (_Configurator *ConfiguratorTransactorSession) SetConfiguration(cometProxy common.Address, newConfiguration CometConfigurationConfiguration) (*types.Transaction, error) {
	return _Configurator.Contract.SetConfiguration(&_Configurator.TransactOpts, cometProxy, newConfiguration)
}

// SetExtensionDelegate is a paid mutator transaction binding the contract method 0x35e67c10.
//
// Solidity: function setExtensionDelegate(address cometProxy, address newExtensionDelegate) returns()
func (_Configurator *ConfiguratorTransactor) SetExtensionDelegate(opts *bind.TransactOpts, cometProxy common.Address, newExtensionDelegate common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setExtensionDelegate", cometProxy, newExtensionDelegate)
}

// SetExtensionDelegate is a paid mutator transaction binding the contract method 0x35e67c10.
//
// Solidity: function setExtensionDelegate(address cometProxy, address newExtensionDelegate) returns()
func (_Configurator *ConfiguratorSession) SetExtensionDelegate(cometProxy common.Address, newExtensionDelegate common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetExtensionDelegate(&_Configurator.TransactOpts, cometProxy, newExtensionDelegate)
}

// SetExtensionDelegate is a paid mutator transaction binding the contract method 0x35e67c10.
//
// Solidity: function setExtensionDelegate(address cometProxy, address newExtensionDelegate) returns()
func (_Configurator *ConfiguratorTransactorSession) SetExtensionDelegate(cometProxy common.Address, newExtensionDelegate common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetExtensionDelegate(&_Configurator.TransactOpts, cometProxy, newExtensionDelegate)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5e825564.
//
// Solidity: function setFactory(address cometProxy, address newFactory) returns()
func (_Configurator *ConfiguratorTransactor) SetFactory(opts *bind.TransactOpts, cometProxy common.Address, newFactory common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setFactory", cometProxy, newFactory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5e825564.
//
// Solidity: function setFactory(address cometProxy, address newFactory) returns()
func (_Configurator *ConfiguratorSession) SetFactory(cometProxy common.Address, newFactory common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetFactory(&_Configurator.TransactOpts, cometProxy, newFactory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5e825564.
//
// Solidity: function setFactory(address cometProxy, address newFactory) returns()
func (_Configurator *ConfiguratorTransactorSession) SetFactory(cometProxy common.Address, newFactory common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetFactory(&_Configurator.TransactOpts, cometProxy, newFactory)
}

// SetGovernor is a paid mutator transaction binding the contract method 0x786f0ac4.
//
// Solidity: function setGovernor(address cometProxy, address newGovernor) returns()
func (_Configurator *ConfiguratorTransactor) SetGovernor(opts *bind.TransactOpts, cometProxy common.Address, newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setGovernor", cometProxy, newGovernor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0x786f0ac4.
//
// Solidity: function setGovernor(address cometProxy, address newGovernor) returns()
func (_Configurator *ConfiguratorSession) SetGovernor(cometProxy common.Address, newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetGovernor(&_Configurator.TransactOpts, cometProxy, newGovernor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0x786f0ac4.
//
// Solidity: function setGovernor(address cometProxy, address newGovernor) returns()
func (_Configurator *ConfiguratorTransactorSession) SetGovernor(cometProxy common.Address, newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetGovernor(&_Configurator.TransactOpts, cometProxy, newGovernor)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x8145ae7d.
//
// Solidity: function setPauseGuardian(address cometProxy, address newPauseGuardian) returns()
func (_Configurator *ConfiguratorTransactor) SetPauseGuardian(opts *bind.TransactOpts, cometProxy common.Address, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setPauseGuardian", cometProxy, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x8145ae7d.
//
// Solidity: function setPauseGuardian(address cometProxy, address newPauseGuardian) returns()
func (_Configurator *ConfiguratorSession) SetPauseGuardian(cometProxy common.Address, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetPauseGuardian(&_Configurator.TransactOpts, cometProxy, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x8145ae7d.
//
// Solidity: function setPauseGuardian(address cometProxy, address newPauseGuardian) returns()
func (_Configurator *ConfiguratorTransactorSession) SetPauseGuardian(cometProxy common.Address, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.SetPauseGuardian(&_Configurator.TransactOpts, cometProxy, newPauseGuardian)
}

// SetStoreFrontPriceFactor is a paid mutator transaction binding the contract method 0x9340d433.
//
// Solidity: function setStoreFrontPriceFactor(address cometProxy, uint64 newStoreFrontPriceFactor) returns()
func (_Configurator *ConfiguratorTransactor) SetStoreFrontPriceFactor(opts *bind.TransactOpts, cometProxy common.Address, newStoreFrontPriceFactor uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setStoreFrontPriceFactor", cometProxy, newStoreFrontPriceFactor)
}

// SetStoreFrontPriceFactor is a paid mutator transaction binding the contract method 0x9340d433.
//
// Solidity: function setStoreFrontPriceFactor(address cometProxy, uint64 newStoreFrontPriceFactor) returns()
func (_Configurator *ConfiguratorSession) SetStoreFrontPriceFactor(cometProxy common.Address, newStoreFrontPriceFactor uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetStoreFrontPriceFactor(&_Configurator.TransactOpts, cometProxy, newStoreFrontPriceFactor)
}

// SetStoreFrontPriceFactor is a paid mutator transaction binding the contract method 0x9340d433.
//
// Solidity: function setStoreFrontPriceFactor(address cometProxy, uint64 newStoreFrontPriceFactor) returns()
func (_Configurator *ConfiguratorTransactorSession) SetStoreFrontPriceFactor(cometProxy common.Address, newStoreFrontPriceFactor uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetStoreFrontPriceFactor(&_Configurator.TransactOpts, cometProxy, newStoreFrontPriceFactor)
}

// SetSupplyKink is a paid mutator transaction binding the contract method 0x058e4155.
//
// Solidity: function setSupplyKink(address cometProxy, uint64 newSupplyKink) returns()
func (_Configurator *ConfiguratorTransactor) SetSupplyKink(opts *bind.TransactOpts, cometProxy common.Address, newSupplyKink uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setSupplyKink", cometProxy, newSupplyKink)
}

// SetSupplyKink is a paid mutator transaction binding the contract method 0x058e4155.
//
// Solidity: function setSupplyKink(address cometProxy, uint64 newSupplyKink) returns()
func (_Configurator *ConfiguratorSession) SetSupplyKink(cometProxy common.Address, newSupplyKink uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyKink(&_Configurator.TransactOpts, cometProxy, newSupplyKink)
}

// SetSupplyKink is a paid mutator transaction binding the contract method 0x058e4155.
//
// Solidity: function setSupplyKink(address cometProxy, uint64 newSupplyKink) returns()
func (_Configurator *ConfiguratorTransactorSession) SetSupplyKink(cometProxy common.Address, newSupplyKink uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyKink(&_Configurator.TransactOpts, cometProxy, newSupplyKink)
}

// SetSupplyPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x3dd40365.
//
// Solidity: function setSupplyPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorTransactor) SetSupplyPerYearInterestRateBase(opts *bind.TransactOpts, cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setSupplyPerYearInterestRateBase", cometProxy, newBase)
}

// SetSupplyPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x3dd40365.
//
// Solidity: function setSupplyPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorSession) SetSupplyPerYearInterestRateBase(cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateBase(&_Configurator.TransactOpts, cometProxy, newBase)
}

// SetSupplyPerYearInterestRateBase is a paid mutator transaction binding the contract method 0x3dd40365.
//
// Solidity: function setSupplyPerYearInterestRateBase(address cometProxy, uint64 newBase) returns()
func (_Configurator *ConfiguratorTransactorSession) SetSupplyPerYearInterestRateBase(cometProxy common.Address, newBase uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateBase(&_Configurator.TransactOpts, cometProxy, newBase)
}

// SetSupplyPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0x5c9ae499.
//
// Solidity: function setSupplyPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactor) SetSupplyPerYearInterestRateSlopeHigh(opts *bind.TransactOpts, cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setSupplyPerYearInterestRateSlopeHigh", cometProxy, newSlope)
}

// SetSupplyPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0x5c9ae499.
//
// Solidity: function setSupplyPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorSession) SetSupplyPerYearInterestRateSlopeHigh(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateSlopeHigh(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetSupplyPerYearInterestRateSlopeHigh is a paid mutator transaction binding the contract method 0x5c9ae499.
//
// Solidity: function setSupplyPerYearInterestRateSlopeHigh(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactorSession) SetSupplyPerYearInterestRateSlopeHigh(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateSlopeHigh(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetSupplyPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0x4a900c5a.
//
// Solidity: function setSupplyPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactor) SetSupplyPerYearInterestRateSlopeLow(opts *bind.TransactOpts, cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setSupplyPerYearInterestRateSlopeLow", cometProxy, newSlope)
}

// SetSupplyPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0x4a900c5a.
//
// Solidity: function setSupplyPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorSession) SetSupplyPerYearInterestRateSlopeLow(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateSlopeLow(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetSupplyPerYearInterestRateSlopeLow is a paid mutator transaction binding the contract method 0x4a900c5a.
//
// Solidity: function setSupplyPerYearInterestRateSlopeLow(address cometProxy, uint64 newSlope) returns()
func (_Configurator *ConfiguratorTransactorSession) SetSupplyPerYearInterestRateSlopeLow(cometProxy common.Address, newSlope uint64) (*types.Transaction, error) {
	return _Configurator.Contract.SetSupplyPerYearInterestRateSlopeLow(&_Configurator.TransactOpts, cometProxy, newSlope)
}

// SetTargetReserves is a paid mutator transaction binding the contract method 0x7d5e1840.
//
// Solidity: function setTargetReserves(address cometProxy, uint104 newTargetReserves) returns()
func (_Configurator *ConfiguratorTransactor) SetTargetReserves(opts *bind.TransactOpts, cometProxy common.Address, newTargetReserves *big.Int) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setTargetReserves", cometProxy, newTargetReserves)
}

// SetTargetReserves is a paid mutator transaction binding the contract method 0x7d5e1840.
//
// Solidity: function setTargetReserves(address cometProxy, uint104 newTargetReserves) returns()
func (_Configurator *ConfiguratorSession) SetTargetReserves(cometProxy common.Address, newTargetReserves *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetTargetReserves(&_Configurator.TransactOpts, cometProxy, newTargetReserves)
}

// SetTargetReserves is a paid mutator transaction binding the contract method 0x7d5e1840.
//
// Solidity: function setTargetReserves(address cometProxy, uint104 newTargetReserves) returns()
func (_Configurator *ConfiguratorTransactorSession) SetTargetReserves(cometProxy common.Address, newTargetReserves *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.SetTargetReserves(&_Configurator.TransactOpts, cometProxy, newTargetReserves)
}

// TransferGovernor is a paid mutator transaction binding the contract method 0xb8cc9ce6.
//
// Solidity: function transferGovernor(address newGovernor) returns()
func (_Configurator *ConfiguratorTransactor) TransferGovernor(opts *bind.TransactOpts, newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "transferGovernor", newGovernor)
}

// TransferGovernor is a paid mutator transaction binding the contract method 0xb8cc9ce6.
//
// Solidity: function transferGovernor(address newGovernor) returns()
func (_Configurator *ConfiguratorSession) TransferGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.TransferGovernor(&_Configurator.TransactOpts, newGovernor)
}

// TransferGovernor is a paid mutator transaction binding the contract method 0xb8cc9ce6.
//
// Solidity: function transferGovernor(address newGovernor) returns()
func (_Configurator *ConfiguratorTransactorSession) TransferGovernor(newGovernor common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.TransferGovernor(&_Configurator.TransactOpts, newGovernor)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x9a0fd808.
//
// Solidity: function updateAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAsset(opts *bind.TransactOpts, cometProxy common.Address, newAssetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAsset", cometProxy, newAssetConfig)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x9a0fd808.
//
// Solidity: function updateAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig) returns()
func (_Configurator *ConfiguratorSession) UpdateAsset(cometProxy common.Address, newAssetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAsset(&_Configurator.TransactOpts, cometProxy, newAssetConfig)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0x9a0fd808.
//
// Solidity: function updateAsset(address cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAsset(cometProxy common.Address, newAssetConfig CometConfigurationAssetConfig) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAsset(&_Configurator.TransactOpts, cometProxy, newAssetConfig)
}

// UpdateAssetBorrowCollateralFactor is a paid mutator transaction binding the contract method 0xb73585f1.
//
// Solidity: function updateAssetBorrowCollateralFactor(address cometProxy, address asset, uint64 newBorrowCF) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAssetBorrowCollateralFactor(opts *bind.TransactOpts, cometProxy common.Address, asset common.Address, newBorrowCF uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAssetBorrowCollateralFactor", cometProxy, asset, newBorrowCF)
}

// UpdateAssetBorrowCollateralFactor is a paid mutator transaction binding the contract method 0xb73585f1.
//
// Solidity: function updateAssetBorrowCollateralFactor(address cometProxy, address asset, uint64 newBorrowCF) returns()
func (_Configurator *ConfiguratorSession) UpdateAssetBorrowCollateralFactor(cometProxy common.Address, asset common.Address, newBorrowCF uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetBorrowCollateralFactor(&_Configurator.TransactOpts, cometProxy, asset, newBorrowCF)
}

// UpdateAssetBorrowCollateralFactor is a paid mutator transaction binding the contract method 0xb73585f1.
//
// Solidity: function updateAssetBorrowCollateralFactor(address cometProxy, address asset, uint64 newBorrowCF) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAssetBorrowCollateralFactor(cometProxy common.Address, asset common.Address, newBorrowCF uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetBorrowCollateralFactor(&_Configurator.TransactOpts, cometProxy, asset, newBorrowCF)
}

// UpdateAssetLiquidateCollateralFactor is a paid mutator transaction binding the contract method 0x77a7dafd.
//
// Solidity: function updateAssetLiquidateCollateralFactor(address cometProxy, address asset, uint64 newLiquidateCF) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAssetLiquidateCollateralFactor(opts *bind.TransactOpts, cometProxy common.Address, asset common.Address, newLiquidateCF uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAssetLiquidateCollateralFactor", cometProxy, asset, newLiquidateCF)
}

// UpdateAssetLiquidateCollateralFactor is a paid mutator transaction binding the contract method 0x77a7dafd.
//
// Solidity: function updateAssetLiquidateCollateralFactor(address cometProxy, address asset, uint64 newLiquidateCF) returns()
func (_Configurator *ConfiguratorSession) UpdateAssetLiquidateCollateralFactor(cometProxy common.Address, asset common.Address, newLiquidateCF uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetLiquidateCollateralFactor(&_Configurator.TransactOpts, cometProxy, asset, newLiquidateCF)
}

// UpdateAssetLiquidateCollateralFactor is a paid mutator transaction binding the contract method 0x77a7dafd.
//
// Solidity: function updateAssetLiquidateCollateralFactor(address cometProxy, address asset, uint64 newLiquidateCF) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAssetLiquidateCollateralFactor(cometProxy common.Address, asset common.Address, newLiquidateCF uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetLiquidateCollateralFactor(&_Configurator.TransactOpts, cometProxy, asset, newLiquidateCF)
}

// UpdateAssetLiquidationFactor is a paid mutator transaction binding the contract method 0xecb9a875.
//
// Solidity: function updateAssetLiquidationFactor(address cometProxy, address asset, uint64 newLiquidationFactor) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAssetLiquidationFactor(opts *bind.TransactOpts, cometProxy common.Address, asset common.Address, newLiquidationFactor uint64) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAssetLiquidationFactor", cometProxy, asset, newLiquidationFactor)
}

// UpdateAssetLiquidationFactor is a paid mutator transaction binding the contract method 0xecb9a875.
//
// Solidity: function updateAssetLiquidationFactor(address cometProxy, address asset, uint64 newLiquidationFactor) returns()
func (_Configurator *ConfiguratorSession) UpdateAssetLiquidationFactor(cometProxy common.Address, asset common.Address, newLiquidationFactor uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetLiquidationFactor(&_Configurator.TransactOpts, cometProxy, asset, newLiquidationFactor)
}

// UpdateAssetLiquidationFactor is a paid mutator transaction binding the contract method 0xecb9a875.
//
// Solidity: function updateAssetLiquidationFactor(address cometProxy, address asset, uint64 newLiquidationFactor) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAssetLiquidationFactor(cometProxy common.Address, asset common.Address, newLiquidationFactor uint64) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetLiquidationFactor(&_Configurator.TransactOpts, cometProxy, asset, newLiquidationFactor)
}

// UpdateAssetPriceFeed is a paid mutator transaction binding the contract method 0x294be433.
//
// Solidity: function updateAssetPriceFeed(address cometProxy, address asset, address newPriceFeed) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAssetPriceFeed(opts *bind.TransactOpts, cometProxy common.Address, asset common.Address, newPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAssetPriceFeed", cometProxy, asset, newPriceFeed)
}

// UpdateAssetPriceFeed is a paid mutator transaction binding the contract method 0x294be433.
//
// Solidity: function updateAssetPriceFeed(address cometProxy, address asset, address newPriceFeed) returns()
func (_Configurator *ConfiguratorSession) UpdateAssetPriceFeed(cometProxy common.Address, asset common.Address, newPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetPriceFeed(&_Configurator.TransactOpts, cometProxy, asset, newPriceFeed)
}

// UpdateAssetPriceFeed is a paid mutator transaction binding the contract method 0x294be433.
//
// Solidity: function updateAssetPriceFeed(address cometProxy, address asset, address newPriceFeed) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAssetPriceFeed(cometProxy common.Address, asset common.Address, newPriceFeed common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetPriceFeed(&_Configurator.TransactOpts, cometProxy, asset, newPriceFeed)
}

// UpdateAssetSupplyCap is a paid mutator transaction binding the contract method 0xa2ced7fd.
//
// Solidity: function updateAssetSupplyCap(address cometProxy, address asset, uint128 newSupplyCap) returns()
func (_Configurator *ConfiguratorTransactor) UpdateAssetSupplyCap(opts *bind.TransactOpts, cometProxy common.Address, asset common.Address, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "updateAssetSupplyCap", cometProxy, asset, newSupplyCap)
}

// UpdateAssetSupplyCap is a paid mutator transaction binding the contract method 0xa2ced7fd.
//
// Solidity: function updateAssetSupplyCap(address cometProxy, address asset, uint128 newSupplyCap) returns()
func (_Configurator *ConfiguratorSession) UpdateAssetSupplyCap(cometProxy common.Address, asset common.Address, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetSupplyCap(&_Configurator.TransactOpts, cometProxy, asset, newSupplyCap)
}

// UpdateAssetSupplyCap is a paid mutator transaction binding the contract method 0xa2ced7fd.
//
// Solidity: function updateAssetSupplyCap(address cometProxy, address asset, uint128 newSupplyCap) returns()
func (_Configurator *ConfiguratorTransactorSession) UpdateAssetSupplyCap(cometProxy common.Address, asset common.Address, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _Configurator.Contract.UpdateAssetSupplyCap(&_Configurator.TransactOpts, cometProxy, asset, newSupplyCap)
}

// ConfiguratorAddAssetIterator is returned from FilterAddAsset and is used to iterate over the raw logs and unpacked data for AddAsset events raised by the Configurator contract.
type ConfiguratorAddAssetIterator struct {
	Event *ConfiguratorAddAsset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorAddAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorAddAsset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorAddAsset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorAddAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorAddAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorAddAsset represents a AddAsset event raised by the Configurator contract.
type ConfiguratorAddAsset struct {
	CometProxy  common.Address
	AssetConfig CometConfigurationAssetConfig
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddAsset is a free log retrieval operation binding the contract event 0x1f7dcc7122c2fe2d685db789d8cde941d28c9d5bf456dcd260705c8d4aef4ef8.
//
// Solidity: event AddAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig)
func (_Configurator *ConfiguratorFilterer) FilterAddAsset(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorAddAssetIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "AddAsset", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorAddAssetIterator{contract: _Configurator.contract, event: "AddAsset", logs: logs, sub: sub}, nil
}

// WatchAddAsset is a free log subscription operation binding the contract event 0x1f7dcc7122c2fe2d685db789d8cde941d28c9d5bf456dcd260705c8d4aef4ef8.
//
// Solidity: event AddAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig)
func (_Configurator *ConfiguratorFilterer) WatchAddAsset(opts *bind.WatchOpts, sink chan<- *ConfiguratorAddAsset, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "AddAsset", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorAddAsset)
				if err := _Configurator.contract.UnpackLog(event, "AddAsset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddAsset is a log parse operation binding the contract event 0x1f7dcc7122c2fe2d685db789d8cde941d28c9d5bf456dcd260705c8d4aef4ef8.
//
// Solidity: event AddAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) assetConfig)
func (_Configurator *ConfiguratorFilterer) ParseAddAsset(log types.Log) (*ConfiguratorAddAsset, error) {
	event := new(ConfiguratorAddAsset)
	if err := _Configurator.contract.UnpackLog(event, "AddAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorCometDeployedIterator is returned from FilterCometDeployed and is used to iterate over the raw logs and unpacked data for CometDeployed events raised by the Configurator contract.
type ConfiguratorCometDeployedIterator struct {
	Event *ConfiguratorCometDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorCometDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorCometDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorCometDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorCometDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorCometDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorCometDeployed represents a CometDeployed event raised by the Configurator contract.
type ConfiguratorCometDeployed struct {
	CometProxy common.Address
	NewComet   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCometDeployed is a free log retrieval operation binding the contract event 0x3da528dfe78562a1f409134989443b5f21ee92023a64b90dedeb2002415189b6.
//
// Solidity: event CometDeployed(address indexed cometProxy, address indexed newComet)
func (_Configurator *ConfiguratorFilterer) FilterCometDeployed(opts *bind.FilterOpts, cometProxy []common.Address, newComet []common.Address) (*ConfiguratorCometDeployedIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var newCometRule []interface{}
	for _, newCometItem := range newComet {
		newCometRule = append(newCometRule, newCometItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "CometDeployed", cometProxyRule, newCometRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorCometDeployedIterator{contract: _Configurator.contract, event: "CometDeployed", logs: logs, sub: sub}, nil
}

// WatchCometDeployed is a free log subscription operation binding the contract event 0x3da528dfe78562a1f409134989443b5f21ee92023a64b90dedeb2002415189b6.
//
// Solidity: event CometDeployed(address indexed cometProxy, address indexed newComet)
func (_Configurator *ConfiguratorFilterer) WatchCometDeployed(opts *bind.WatchOpts, sink chan<- *ConfiguratorCometDeployed, cometProxy []common.Address, newComet []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var newCometRule []interface{}
	for _, newCometItem := range newComet {
		newCometRule = append(newCometRule, newCometItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "CometDeployed", cometProxyRule, newCometRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorCometDeployed)
				if err := _Configurator.contract.UnpackLog(event, "CometDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCometDeployed is a log parse operation binding the contract event 0x3da528dfe78562a1f409134989443b5f21ee92023a64b90dedeb2002415189b6.
//
// Solidity: event CometDeployed(address indexed cometProxy, address indexed newComet)
func (_Configurator *ConfiguratorFilterer) ParseCometDeployed(log types.Log) (*ConfiguratorCometDeployed, error) {
	event := new(ConfiguratorCometDeployed)
	if err := _Configurator.contract.UnpackLog(event, "CometDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorGovernorTransferredIterator is returned from FilterGovernorTransferred and is used to iterate over the raw logs and unpacked data for GovernorTransferred events raised by the Configurator contract.
type ConfiguratorGovernorTransferredIterator struct {
	Event *ConfiguratorGovernorTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorGovernorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorGovernorTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorGovernorTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorGovernorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorGovernorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorGovernorTransferred represents a GovernorTransferred event raised by the Configurator contract.
type ConfiguratorGovernorTransferred struct {
	OldGovernor common.Address
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovernorTransferred is a free log retrieval operation binding the contract event 0x6fadb1c244276388aee22be93b919985a18748c021e5d48553957a48101a2560.
//
// Solidity: event GovernorTransferred(address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) FilterGovernorTransferred(opts *bind.FilterOpts, oldGovernor []common.Address, newGovernor []common.Address) (*ConfiguratorGovernorTransferredIterator, error) {

	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "GovernorTransferred", oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorGovernorTransferredIterator{contract: _Configurator.contract, event: "GovernorTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernorTransferred is a free log subscription operation binding the contract event 0x6fadb1c244276388aee22be93b919985a18748c021e5d48553957a48101a2560.
//
// Solidity: event GovernorTransferred(address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) WatchGovernorTransferred(opts *bind.WatchOpts, sink chan<- *ConfiguratorGovernorTransferred, oldGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "GovernorTransferred", oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorGovernorTransferred)
				if err := _Configurator.contract.UnpackLog(event, "GovernorTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernorTransferred is a log parse operation binding the contract event 0x6fadb1c244276388aee22be93b919985a18748c021e5d48553957a48101a2560.
//
// Solidity: event GovernorTransferred(address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) ParseGovernorTransferred(log types.Log) (*ConfiguratorGovernorTransferred, error) {
	event := new(ConfiguratorGovernorTransferred)
	if err := _Configurator.contract.UnpackLog(event, "GovernorTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBaseBorrowMinIterator is returned from FilterSetBaseBorrowMin and is used to iterate over the raw logs and unpacked data for SetBaseBorrowMin events raised by the Configurator contract.
type ConfiguratorSetBaseBorrowMinIterator struct {
	Event *ConfiguratorSetBaseBorrowMin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBaseBorrowMinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBaseBorrowMin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBaseBorrowMin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBaseBorrowMinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBaseBorrowMinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBaseBorrowMin represents a SetBaseBorrowMin event raised by the Configurator contract.
type ConfiguratorSetBaseBorrowMin struct {
	CometProxy       common.Address
	OldBaseBorrowMin *big.Int
	NewBaseBorrowMin *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetBaseBorrowMin is a free log retrieval operation binding the contract event 0x2cc4a6aedb45911a1034b2320f16567fa4e2eca3f1f5d3b46804f6bc42b7c3f2.
//
// Solidity: event SetBaseBorrowMin(address indexed cometProxy, uint104 oldBaseBorrowMin, uint104 newBaseBorrowMin)
func (_Configurator *ConfiguratorFilterer) FilterSetBaseBorrowMin(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBaseBorrowMinIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBaseBorrowMin", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBaseBorrowMinIterator{contract: _Configurator.contract, event: "SetBaseBorrowMin", logs: logs, sub: sub}, nil
}

// WatchSetBaseBorrowMin is a free log subscription operation binding the contract event 0x2cc4a6aedb45911a1034b2320f16567fa4e2eca3f1f5d3b46804f6bc42b7c3f2.
//
// Solidity: event SetBaseBorrowMin(address indexed cometProxy, uint104 oldBaseBorrowMin, uint104 newBaseBorrowMin)
func (_Configurator *ConfiguratorFilterer) WatchSetBaseBorrowMin(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBaseBorrowMin, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBaseBorrowMin", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBaseBorrowMin)
				if err := _Configurator.contract.UnpackLog(event, "SetBaseBorrowMin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseBorrowMin is a log parse operation binding the contract event 0x2cc4a6aedb45911a1034b2320f16567fa4e2eca3f1f5d3b46804f6bc42b7c3f2.
//
// Solidity: event SetBaseBorrowMin(address indexed cometProxy, uint104 oldBaseBorrowMin, uint104 newBaseBorrowMin)
func (_Configurator *ConfiguratorFilterer) ParseSetBaseBorrowMin(log types.Log) (*ConfiguratorSetBaseBorrowMin, error) {
	event := new(ConfiguratorSetBaseBorrowMin)
	if err := _Configurator.contract.UnpackLog(event, "SetBaseBorrowMin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBaseMinForRewardsIterator is returned from FilterSetBaseMinForRewards and is used to iterate over the raw logs and unpacked data for SetBaseMinForRewards events raised by the Configurator contract.
type ConfiguratorSetBaseMinForRewardsIterator struct {
	Event *ConfiguratorSetBaseMinForRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBaseMinForRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBaseMinForRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBaseMinForRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBaseMinForRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBaseMinForRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBaseMinForRewards represents a SetBaseMinForRewards event raised by the Configurator contract.
type ConfiguratorSetBaseMinForRewards struct {
	CometProxy           common.Address
	OldBaseMinForRewards *big.Int
	NewBaseMinForRewards *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetBaseMinForRewards is a free log retrieval operation binding the contract event 0x635524bd339ec1b914f9b13607c72592cfa9538a3c56170a59094dde834bb41d.
//
// Solidity: event SetBaseMinForRewards(address indexed cometProxy, uint104 oldBaseMinForRewards, uint104 newBaseMinForRewards)
func (_Configurator *ConfiguratorFilterer) FilterSetBaseMinForRewards(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBaseMinForRewardsIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBaseMinForRewards", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBaseMinForRewardsIterator{contract: _Configurator.contract, event: "SetBaseMinForRewards", logs: logs, sub: sub}, nil
}

// WatchSetBaseMinForRewards is a free log subscription operation binding the contract event 0x635524bd339ec1b914f9b13607c72592cfa9538a3c56170a59094dde834bb41d.
//
// Solidity: event SetBaseMinForRewards(address indexed cometProxy, uint104 oldBaseMinForRewards, uint104 newBaseMinForRewards)
func (_Configurator *ConfiguratorFilterer) WatchSetBaseMinForRewards(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBaseMinForRewards, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBaseMinForRewards", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBaseMinForRewards)
				if err := _Configurator.contract.UnpackLog(event, "SetBaseMinForRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseMinForRewards is a log parse operation binding the contract event 0x635524bd339ec1b914f9b13607c72592cfa9538a3c56170a59094dde834bb41d.
//
// Solidity: event SetBaseMinForRewards(address indexed cometProxy, uint104 oldBaseMinForRewards, uint104 newBaseMinForRewards)
func (_Configurator *ConfiguratorFilterer) ParseSetBaseMinForRewards(log types.Log) (*ConfiguratorSetBaseMinForRewards, error) {
	event := new(ConfiguratorSetBaseMinForRewards)
	if err := _Configurator.contract.UnpackLog(event, "SetBaseMinForRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBaseTokenPriceFeedIterator is returned from FilterSetBaseTokenPriceFeed and is used to iterate over the raw logs and unpacked data for SetBaseTokenPriceFeed events raised by the Configurator contract.
type ConfiguratorSetBaseTokenPriceFeedIterator struct {
	Event *ConfiguratorSetBaseTokenPriceFeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBaseTokenPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBaseTokenPriceFeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBaseTokenPriceFeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBaseTokenPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBaseTokenPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBaseTokenPriceFeed represents a SetBaseTokenPriceFeed event raised by the Configurator contract.
type ConfiguratorSetBaseTokenPriceFeed struct {
	CometProxy            common.Address
	OldBaseTokenPriceFeed common.Address
	NewBaseTokenPriceFeed common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetBaseTokenPriceFeed is a free log retrieval operation binding the contract event 0x1e3daca7fe0d6c7ab58271d196454f1d9e867a09de95666bcad4ab688e750f1a.
//
// Solidity: event SetBaseTokenPriceFeed(address indexed cometProxy, address indexed oldBaseTokenPriceFeed, address indexed newBaseTokenPriceFeed)
func (_Configurator *ConfiguratorFilterer) FilterSetBaseTokenPriceFeed(opts *bind.FilterOpts, cometProxy []common.Address, oldBaseTokenPriceFeed []common.Address, newBaseTokenPriceFeed []common.Address) (*ConfiguratorSetBaseTokenPriceFeedIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldBaseTokenPriceFeedRule []interface{}
	for _, oldBaseTokenPriceFeedItem := range oldBaseTokenPriceFeed {
		oldBaseTokenPriceFeedRule = append(oldBaseTokenPriceFeedRule, oldBaseTokenPriceFeedItem)
	}
	var newBaseTokenPriceFeedRule []interface{}
	for _, newBaseTokenPriceFeedItem := range newBaseTokenPriceFeed {
		newBaseTokenPriceFeedRule = append(newBaseTokenPriceFeedRule, newBaseTokenPriceFeedItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBaseTokenPriceFeed", cometProxyRule, oldBaseTokenPriceFeedRule, newBaseTokenPriceFeedRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBaseTokenPriceFeedIterator{contract: _Configurator.contract, event: "SetBaseTokenPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetBaseTokenPriceFeed is a free log subscription operation binding the contract event 0x1e3daca7fe0d6c7ab58271d196454f1d9e867a09de95666bcad4ab688e750f1a.
//
// Solidity: event SetBaseTokenPriceFeed(address indexed cometProxy, address indexed oldBaseTokenPriceFeed, address indexed newBaseTokenPriceFeed)
func (_Configurator *ConfiguratorFilterer) WatchSetBaseTokenPriceFeed(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBaseTokenPriceFeed, cometProxy []common.Address, oldBaseTokenPriceFeed []common.Address, newBaseTokenPriceFeed []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldBaseTokenPriceFeedRule []interface{}
	for _, oldBaseTokenPriceFeedItem := range oldBaseTokenPriceFeed {
		oldBaseTokenPriceFeedRule = append(oldBaseTokenPriceFeedRule, oldBaseTokenPriceFeedItem)
	}
	var newBaseTokenPriceFeedRule []interface{}
	for _, newBaseTokenPriceFeedItem := range newBaseTokenPriceFeed {
		newBaseTokenPriceFeedRule = append(newBaseTokenPriceFeedRule, newBaseTokenPriceFeedItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBaseTokenPriceFeed", cometProxyRule, oldBaseTokenPriceFeedRule, newBaseTokenPriceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBaseTokenPriceFeed)
				if err := _Configurator.contract.UnpackLog(event, "SetBaseTokenPriceFeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseTokenPriceFeed is a log parse operation binding the contract event 0x1e3daca7fe0d6c7ab58271d196454f1d9e867a09de95666bcad4ab688e750f1a.
//
// Solidity: event SetBaseTokenPriceFeed(address indexed cometProxy, address indexed oldBaseTokenPriceFeed, address indexed newBaseTokenPriceFeed)
func (_Configurator *ConfiguratorFilterer) ParseSetBaseTokenPriceFeed(log types.Log) (*ConfiguratorSetBaseTokenPriceFeed, error) {
	event := new(ConfiguratorSetBaseTokenPriceFeed)
	if err := _Configurator.contract.UnpackLog(event, "SetBaseTokenPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBaseTrackingBorrowSpeedIterator is returned from FilterSetBaseTrackingBorrowSpeed and is used to iterate over the raw logs and unpacked data for SetBaseTrackingBorrowSpeed events raised by the Configurator contract.
type ConfiguratorSetBaseTrackingBorrowSpeedIterator struct {
	Event *ConfiguratorSetBaseTrackingBorrowSpeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBaseTrackingBorrowSpeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBaseTrackingBorrowSpeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBaseTrackingBorrowSpeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBaseTrackingBorrowSpeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBaseTrackingBorrowSpeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBaseTrackingBorrowSpeed represents a SetBaseTrackingBorrowSpeed event raised by the Configurator contract.
type ConfiguratorSetBaseTrackingBorrowSpeed struct {
	CometProxy                 common.Address
	OldBaseTrackingBorrowSpeed uint64
	NewBaseTrackingBorrowSpeed uint64
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetBaseTrackingBorrowSpeed is a free log retrieval operation binding the contract event 0xcd38d79accad9a0b2cec7a7f94a4ceb7d9b0df5f35bdcb6ee4212b695fad0aa1.
//
// Solidity: event SetBaseTrackingBorrowSpeed(address indexed cometProxy, uint64 oldBaseTrackingBorrowSpeed, uint64 newBaseTrackingBorrowSpeed)
func (_Configurator *ConfiguratorFilterer) FilterSetBaseTrackingBorrowSpeed(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBaseTrackingBorrowSpeedIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBaseTrackingBorrowSpeed", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBaseTrackingBorrowSpeedIterator{contract: _Configurator.contract, event: "SetBaseTrackingBorrowSpeed", logs: logs, sub: sub}, nil
}

// WatchSetBaseTrackingBorrowSpeed is a free log subscription operation binding the contract event 0xcd38d79accad9a0b2cec7a7f94a4ceb7d9b0df5f35bdcb6ee4212b695fad0aa1.
//
// Solidity: event SetBaseTrackingBorrowSpeed(address indexed cometProxy, uint64 oldBaseTrackingBorrowSpeed, uint64 newBaseTrackingBorrowSpeed)
func (_Configurator *ConfiguratorFilterer) WatchSetBaseTrackingBorrowSpeed(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBaseTrackingBorrowSpeed, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBaseTrackingBorrowSpeed", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBaseTrackingBorrowSpeed)
				if err := _Configurator.contract.UnpackLog(event, "SetBaseTrackingBorrowSpeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseTrackingBorrowSpeed is a log parse operation binding the contract event 0xcd38d79accad9a0b2cec7a7f94a4ceb7d9b0df5f35bdcb6ee4212b695fad0aa1.
//
// Solidity: event SetBaseTrackingBorrowSpeed(address indexed cometProxy, uint64 oldBaseTrackingBorrowSpeed, uint64 newBaseTrackingBorrowSpeed)
func (_Configurator *ConfiguratorFilterer) ParseSetBaseTrackingBorrowSpeed(log types.Log) (*ConfiguratorSetBaseTrackingBorrowSpeed, error) {
	event := new(ConfiguratorSetBaseTrackingBorrowSpeed)
	if err := _Configurator.contract.UnpackLog(event, "SetBaseTrackingBorrowSpeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBaseTrackingSupplySpeedIterator is returned from FilterSetBaseTrackingSupplySpeed and is used to iterate over the raw logs and unpacked data for SetBaseTrackingSupplySpeed events raised by the Configurator contract.
type ConfiguratorSetBaseTrackingSupplySpeedIterator struct {
	Event *ConfiguratorSetBaseTrackingSupplySpeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBaseTrackingSupplySpeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBaseTrackingSupplySpeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBaseTrackingSupplySpeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBaseTrackingSupplySpeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBaseTrackingSupplySpeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBaseTrackingSupplySpeed represents a SetBaseTrackingSupplySpeed event raised by the Configurator contract.
type ConfiguratorSetBaseTrackingSupplySpeed struct {
	CometProxy                 common.Address
	OldBaseTrackingSupplySpeed uint64
	NewBaseTrackingSupplySpeed uint64
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSetBaseTrackingSupplySpeed is a free log retrieval operation binding the contract event 0xaf80475c4b413d347b2ac9b71b6c734ba09a6f4771aba12edc65346cffe273f5.
//
// Solidity: event SetBaseTrackingSupplySpeed(address indexed cometProxy, uint64 oldBaseTrackingSupplySpeed, uint64 newBaseTrackingSupplySpeed)
func (_Configurator *ConfiguratorFilterer) FilterSetBaseTrackingSupplySpeed(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBaseTrackingSupplySpeedIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBaseTrackingSupplySpeed", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBaseTrackingSupplySpeedIterator{contract: _Configurator.contract, event: "SetBaseTrackingSupplySpeed", logs: logs, sub: sub}, nil
}

// WatchSetBaseTrackingSupplySpeed is a free log subscription operation binding the contract event 0xaf80475c4b413d347b2ac9b71b6c734ba09a6f4771aba12edc65346cffe273f5.
//
// Solidity: event SetBaseTrackingSupplySpeed(address indexed cometProxy, uint64 oldBaseTrackingSupplySpeed, uint64 newBaseTrackingSupplySpeed)
func (_Configurator *ConfiguratorFilterer) WatchSetBaseTrackingSupplySpeed(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBaseTrackingSupplySpeed, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBaseTrackingSupplySpeed", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBaseTrackingSupplySpeed)
				if err := _Configurator.contract.UnpackLog(event, "SetBaseTrackingSupplySpeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBaseTrackingSupplySpeed is a log parse operation binding the contract event 0xaf80475c4b413d347b2ac9b71b6c734ba09a6f4771aba12edc65346cffe273f5.
//
// Solidity: event SetBaseTrackingSupplySpeed(address indexed cometProxy, uint64 oldBaseTrackingSupplySpeed, uint64 newBaseTrackingSupplySpeed)
func (_Configurator *ConfiguratorFilterer) ParseSetBaseTrackingSupplySpeed(log types.Log) (*ConfiguratorSetBaseTrackingSupplySpeed, error) {
	event := new(ConfiguratorSetBaseTrackingSupplySpeed)
	if err := _Configurator.contract.UnpackLog(event, "SetBaseTrackingSupplySpeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBorrowKinkIterator is returned from FilterSetBorrowKink and is used to iterate over the raw logs and unpacked data for SetBorrowKink events raised by the Configurator contract.
type ConfiguratorSetBorrowKinkIterator struct {
	Event *ConfiguratorSetBorrowKink // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBorrowKinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBorrowKink)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBorrowKink)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBorrowKinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBorrowKinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBorrowKink represents a SetBorrowKink event raised by the Configurator contract.
type ConfiguratorSetBorrowKink struct {
	CometProxy common.Address
	OldKink    uint64
	NewKink    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowKink is a free log retrieval operation binding the contract event 0xe4e8221220a251ab6772c13e04f6c54532602ec260a26f48ff23fa8b11f41be7.
//
// Solidity: event SetBorrowKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) FilterSetBorrowKink(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBorrowKinkIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBorrowKink", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBorrowKinkIterator{contract: _Configurator.contract, event: "SetBorrowKink", logs: logs, sub: sub}, nil
}

// WatchSetBorrowKink is a free log subscription operation binding the contract event 0xe4e8221220a251ab6772c13e04f6c54532602ec260a26f48ff23fa8b11f41be7.
//
// Solidity: event SetBorrowKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) WatchSetBorrowKink(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBorrowKink, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBorrowKink", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBorrowKink)
				if err := _Configurator.contract.UnpackLog(event, "SetBorrowKink", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBorrowKink is a log parse operation binding the contract event 0xe4e8221220a251ab6772c13e04f6c54532602ec260a26f48ff23fa8b11f41be7.
//
// Solidity: event SetBorrowKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) ParseSetBorrowKink(log types.Log) (*ConfiguratorSetBorrowKink, error) {
	event := new(ConfiguratorSetBorrowKink)
	if err := _Configurator.contract.UnpackLog(event, "SetBorrowKink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBorrowPerYearInterestRateBaseIterator is returned from FilterSetBorrowPerYearInterestRateBase and is used to iterate over the raw logs and unpacked data for SetBorrowPerYearInterestRateBase events raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateBaseIterator struct {
	Event *ConfiguratorSetBorrowPerYearInterestRateBase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBorrowPerYearInterestRateBaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBorrowPerYearInterestRateBase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBorrowPerYearInterestRateBase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBorrowPerYearInterestRateBaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBorrowPerYearInterestRateBaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBorrowPerYearInterestRateBase represents a SetBorrowPerYearInterestRateBase event raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateBase struct {
	CometProxy common.Address
	OldIRBase  uint64
	NewIRBase  uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowPerYearInterestRateBase is a free log retrieval operation binding the contract event 0x7e237d05d262a573a6b2d126c2065264a3bcabb68e67a1d5a1e88b1dbef6666d.
//
// Solidity: event SetBorrowPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) FilterSetBorrowPerYearInterestRateBase(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBorrowPerYearInterestRateBaseIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBorrowPerYearInterestRateBase", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBorrowPerYearInterestRateBaseIterator{contract: _Configurator.contract, event: "SetBorrowPerYearInterestRateBase", logs: logs, sub: sub}, nil
}

// WatchSetBorrowPerYearInterestRateBase is a free log subscription operation binding the contract event 0x7e237d05d262a573a6b2d126c2065264a3bcabb68e67a1d5a1e88b1dbef6666d.
//
// Solidity: event SetBorrowPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) WatchSetBorrowPerYearInterestRateBase(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBorrowPerYearInterestRateBase, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBorrowPerYearInterestRateBase", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBorrowPerYearInterestRateBase)
				if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateBase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBorrowPerYearInterestRateBase is a log parse operation binding the contract event 0x7e237d05d262a573a6b2d126c2065264a3bcabb68e67a1d5a1e88b1dbef6666d.
//
// Solidity: event SetBorrowPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) ParseSetBorrowPerYearInterestRateBase(log types.Log) (*ConfiguratorSetBorrowPerYearInterestRateBase, error) {
	event := new(ConfiguratorSetBorrowPerYearInterestRateBase)
	if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateBase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator is returned from FilterSetBorrowPerYearInterestRateSlopeHigh and is used to iterate over the raw logs and unpacked data for SetBorrowPerYearInterestRateSlopeHigh events raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator struct {
	Event *ConfiguratorSetBorrowPerYearInterestRateSlopeHigh // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBorrowPerYearInterestRateSlopeHigh)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBorrowPerYearInterestRateSlopeHigh)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBorrowPerYearInterestRateSlopeHigh represents a SetBorrowPerYearInterestRateSlopeHigh event raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateSlopeHigh struct {
	CometProxy     common.Address
	OldIRSlopeHigh uint64
	NewIRSlopeHigh uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowPerYearInterestRateSlopeHigh is a free log retrieval operation binding the contract event 0xec47838861ac4faa72ef1cf4d982f923880203912cd0c53d6fea2e611562f82c.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) FilterSetBorrowPerYearInterestRateSlopeHigh(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBorrowPerYearInterestRateSlopeHigh", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBorrowPerYearInterestRateSlopeHighIterator{contract: _Configurator.contract, event: "SetBorrowPerYearInterestRateSlopeHigh", logs: logs, sub: sub}, nil
}

// WatchSetBorrowPerYearInterestRateSlopeHigh is a free log subscription operation binding the contract event 0xec47838861ac4faa72ef1cf4d982f923880203912cd0c53d6fea2e611562f82c.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) WatchSetBorrowPerYearInterestRateSlopeHigh(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBorrowPerYearInterestRateSlopeHigh, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBorrowPerYearInterestRateSlopeHigh", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBorrowPerYearInterestRateSlopeHigh)
				if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateSlopeHigh", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBorrowPerYearInterestRateSlopeHigh is a log parse operation binding the contract event 0xec47838861ac4faa72ef1cf4d982f923880203912cd0c53d6fea2e611562f82c.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) ParseSetBorrowPerYearInterestRateSlopeHigh(log types.Log) (*ConfiguratorSetBorrowPerYearInterestRateSlopeHigh, error) {
	event := new(ConfiguratorSetBorrowPerYearInterestRateSlopeHigh)
	if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateSlopeHigh", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator is returned from FilterSetBorrowPerYearInterestRateSlopeLow and is used to iterate over the raw logs and unpacked data for SetBorrowPerYearInterestRateSlopeLow events raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator struct {
	Event *ConfiguratorSetBorrowPerYearInterestRateSlopeLow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetBorrowPerYearInterestRateSlopeLow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetBorrowPerYearInterestRateSlopeLow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetBorrowPerYearInterestRateSlopeLow represents a SetBorrowPerYearInterestRateSlopeLow event raised by the Configurator contract.
type ConfiguratorSetBorrowPerYearInterestRateSlopeLow struct {
	CometProxy    common.Address
	OldIRSlopeLow uint64
	NewIRSlopeLow uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowPerYearInterestRateSlopeLow is a free log retrieval operation binding the contract event 0xe5d8ea5f9f9bc1e3132e593561d5b440062de5e0d5585a7046c87ca172a60d46.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) FilterSetBorrowPerYearInterestRateSlopeLow(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetBorrowPerYearInterestRateSlopeLow", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetBorrowPerYearInterestRateSlopeLowIterator{contract: _Configurator.contract, event: "SetBorrowPerYearInterestRateSlopeLow", logs: logs, sub: sub}, nil
}

// WatchSetBorrowPerYearInterestRateSlopeLow is a free log subscription operation binding the contract event 0xe5d8ea5f9f9bc1e3132e593561d5b440062de5e0d5585a7046c87ca172a60d46.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) WatchSetBorrowPerYearInterestRateSlopeLow(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetBorrowPerYearInterestRateSlopeLow, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetBorrowPerYearInterestRateSlopeLow", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetBorrowPerYearInterestRateSlopeLow)
				if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateSlopeLow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetBorrowPerYearInterestRateSlopeLow is a log parse operation binding the contract event 0xe5d8ea5f9f9bc1e3132e593561d5b440062de5e0d5585a7046c87ca172a60d46.
//
// Solidity: event SetBorrowPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) ParseSetBorrowPerYearInterestRateSlopeLow(log types.Log) (*ConfiguratorSetBorrowPerYearInterestRateSlopeLow, error) {
	event := new(ConfiguratorSetBorrowPerYearInterestRateSlopeLow)
	if err := _Configurator.contract.UnpackLog(event, "SetBorrowPerYearInterestRateSlopeLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetConfigurationIterator is returned from FilterSetConfiguration and is used to iterate over the raw logs and unpacked data for SetConfiguration events raised by the Configurator contract.
type ConfiguratorSetConfigurationIterator struct {
	Event *ConfiguratorSetConfiguration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetConfiguration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetConfiguration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetConfiguration represents a SetConfiguration event raised by the Configurator contract.
type ConfiguratorSetConfiguration struct {
	CometProxy       common.Address
	OldConfiguration CometConfigurationConfiguration
	NewConfiguration CometConfigurationConfiguration
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetConfiguration is a free log retrieval operation binding the contract event 0x11750c0169cdcebd42703782ec1a557a8ddebc9c3239aaf46662cd6c81cdb90b.
//
// Solidity: event SetConfiguration(address indexed cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) oldConfiguration, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration)
func (_Configurator *ConfiguratorFilterer) FilterSetConfiguration(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetConfigurationIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetConfiguration", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetConfigurationIterator{contract: _Configurator.contract, event: "SetConfiguration", logs: logs, sub: sub}, nil
}

// WatchSetConfiguration is a free log subscription operation binding the contract event 0x11750c0169cdcebd42703782ec1a557a8ddebc9c3239aaf46662cd6c81cdb90b.
//
// Solidity: event SetConfiguration(address indexed cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) oldConfiguration, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration)
func (_Configurator *ConfiguratorFilterer) WatchSetConfiguration(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetConfiguration, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetConfiguration", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetConfiguration)
				if err := _Configurator.contract.UnpackLog(event, "SetConfiguration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetConfiguration is a log parse operation binding the contract event 0x11750c0169cdcebd42703782ec1a557a8ddebc9c3239aaf46662cd6c81cdb90b.
//
// Solidity: event SetConfiguration(address indexed cometProxy, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) oldConfiguration, (address,address,address,address,address,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint104,uint104,uint104,(address,address,uint8,uint64,uint64,uint64,uint128)[]) newConfiguration)
func (_Configurator *ConfiguratorFilterer) ParseSetConfiguration(log types.Log) (*ConfiguratorSetConfiguration, error) {
	event := new(ConfiguratorSetConfiguration)
	if err := _Configurator.contract.UnpackLog(event, "SetConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetExtensionDelegateIterator is returned from FilterSetExtensionDelegate and is used to iterate over the raw logs and unpacked data for SetExtensionDelegate events raised by the Configurator contract.
type ConfiguratorSetExtensionDelegateIterator struct {
	Event *ConfiguratorSetExtensionDelegate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetExtensionDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetExtensionDelegate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetExtensionDelegate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetExtensionDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetExtensionDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetExtensionDelegate represents a SetExtensionDelegate event raised by the Configurator contract.
type ConfiguratorSetExtensionDelegate struct {
	CometProxy common.Address
	OldExt     common.Address
	NewExt     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetExtensionDelegate is a free log retrieval operation binding the contract event 0xc86eba353c79b422cdb4d71fcc7db3615603457a42c0559fcf1b31dfd66920b3.
//
// Solidity: event SetExtensionDelegate(address indexed cometProxy, address indexed oldExt, address indexed newExt)
func (_Configurator *ConfiguratorFilterer) FilterSetExtensionDelegate(opts *bind.FilterOpts, cometProxy []common.Address, oldExt []common.Address, newExt []common.Address) (*ConfiguratorSetExtensionDelegateIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldExtRule []interface{}
	for _, oldExtItem := range oldExt {
		oldExtRule = append(oldExtRule, oldExtItem)
	}
	var newExtRule []interface{}
	for _, newExtItem := range newExt {
		newExtRule = append(newExtRule, newExtItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetExtensionDelegate", cometProxyRule, oldExtRule, newExtRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetExtensionDelegateIterator{contract: _Configurator.contract, event: "SetExtensionDelegate", logs: logs, sub: sub}, nil
}

// WatchSetExtensionDelegate is a free log subscription operation binding the contract event 0xc86eba353c79b422cdb4d71fcc7db3615603457a42c0559fcf1b31dfd66920b3.
//
// Solidity: event SetExtensionDelegate(address indexed cometProxy, address indexed oldExt, address indexed newExt)
func (_Configurator *ConfiguratorFilterer) WatchSetExtensionDelegate(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetExtensionDelegate, cometProxy []common.Address, oldExt []common.Address, newExt []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldExtRule []interface{}
	for _, oldExtItem := range oldExt {
		oldExtRule = append(oldExtRule, oldExtItem)
	}
	var newExtRule []interface{}
	for _, newExtItem := range newExt {
		newExtRule = append(newExtRule, newExtItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetExtensionDelegate", cometProxyRule, oldExtRule, newExtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetExtensionDelegate)
				if err := _Configurator.contract.UnpackLog(event, "SetExtensionDelegate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetExtensionDelegate is a log parse operation binding the contract event 0xc86eba353c79b422cdb4d71fcc7db3615603457a42c0559fcf1b31dfd66920b3.
//
// Solidity: event SetExtensionDelegate(address indexed cometProxy, address indexed oldExt, address indexed newExt)
func (_Configurator *ConfiguratorFilterer) ParseSetExtensionDelegate(log types.Log) (*ConfiguratorSetExtensionDelegate, error) {
	event := new(ConfiguratorSetExtensionDelegate)
	if err := _Configurator.contract.UnpackLog(event, "SetExtensionDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetFactoryIterator is returned from FilterSetFactory and is used to iterate over the raw logs and unpacked data for SetFactory events raised by the Configurator contract.
type ConfiguratorSetFactoryIterator struct {
	Event *ConfiguratorSetFactory // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetFactoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetFactory)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetFactory)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetFactoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetFactoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetFactory represents a SetFactory event raised by the Configurator contract.
type ConfiguratorSetFactory struct {
	CometProxy common.Address
	OldFactory common.Address
	NewFactory common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetFactory is a free log retrieval operation binding the contract event 0xcc826d20934cb90e9329d09ff55b4e43831c5bb3a3305fb536842ad49041e7d5.
//
// Solidity: event SetFactory(address indexed cometProxy, address indexed oldFactory, address indexed newFactory)
func (_Configurator *ConfiguratorFilterer) FilterSetFactory(opts *bind.FilterOpts, cometProxy []common.Address, oldFactory []common.Address, newFactory []common.Address) (*ConfiguratorSetFactoryIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldFactoryRule []interface{}
	for _, oldFactoryItem := range oldFactory {
		oldFactoryRule = append(oldFactoryRule, oldFactoryItem)
	}
	var newFactoryRule []interface{}
	for _, newFactoryItem := range newFactory {
		newFactoryRule = append(newFactoryRule, newFactoryItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetFactory", cometProxyRule, oldFactoryRule, newFactoryRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetFactoryIterator{contract: _Configurator.contract, event: "SetFactory", logs: logs, sub: sub}, nil
}

// WatchSetFactory is a free log subscription operation binding the contract event 0xcc826d20934cb90e9329d09ff55b4e43831c5bb3a3305fb536842ad49041e7d5.
//
// Solidity: event SetFactory(address indexed cometProxy, address indexed oldFactory, address indexed newFactory)
func (_Configurator *ConfiguratorFilterer) WatchSetFactory(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetFactory, cometProxy []common.Address, oldFactory []common.Address, newFactory []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldFactoryRule []interface{}
	for _, oldFactoryItem := range oldFactory {
		oldFactoryRule = append(oldFactoryRule, oldFactoryItem)
	}
	var newFactoryRule []interface{}
	for _, newFactoryItem := range newFactory {
		newFactoryRule = append(newFactoryRule, newFactoryItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetFactory", cometProxyRule, oldFactoryRule, newFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetFactory)
				if err := _Configurator.contract.UnpackLog(event, "SetFactory", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFactory is a log parse operation binding the contract event 0xcc826d20934cb90e9329d09ff55b4e43831c5bb3a3305fb536842ad49041e7d5.
//
// Solidity: event SetFactory(address indexed cometProxy, address indexed oldFactory, address indexed newFactory)
func (_Configurator *ConfiguratorFilterer) ParseSetFactory(log types.Log) (*ConfiguratorSetFactory, error) {
	event := new(ConfiguratorSetFactory)
	if err := _Configurator.contract.UnpackLog(event, "SetFactory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetGovernorIterator is returned from FilterSetGovernor and is used to iterate over the raw logs and unpacked data for SetGovernor events raised by the Configurator contract.
type ConfiguratorSetGovernorIterator struct {
	Event *ConfiguratorSetGovernor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetGovernor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetGovernor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetGovernor represents a SetGovernor event raised by the Configurator contract.
type ConfiguratorSetGovernor struct {
	CometProxy  common.Address
	OldGovernor common.Address
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetGovernor is a free log retrieval operation binding the contract event 0x54b27c9d187943783592d79e0380d54f9a545d098bd921e5919d4b3305bd603f.
//
// Solidity: event SetGovernor(address indexed cometProxy, address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) FilterSetGovernor(opts *bind.FilterOpts, cometProxy []common.Address, oldGovernor []common.Address, newGovernor []common.Address) (*ConfiguratorSetGovernorIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetGovernor", cometProxyRule, oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetGovernorIterator{contract: _Configurator.contract, event: "SetGovernor", logs: logs, sub: sub}, nil
}

// WatchSetGovernor is a free log subscription operation binding the contract event 0x54b27c9d187943783592d79e0380d54f9a545d098bd921e5919d4b3305bd603f.
//
// Solidity: event SetGovernor(address indexed cometProxy, address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) WatchSetGovernor(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetGovernor, cometProxy []common.Address, oldGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetGovernor", cometProxyRule, oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetGovernor)
				if err := _Configurator.contract.UnpackLog(event, "SetGovernor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetGovernor is a log parse operation binding the contract event 0x54b27c9d187943783592d79e0380d54f9a545d098bd921e5919d4b3305bd603f.
//
// Solidity: event SetGovernor(address indexed cometProxy, address indexed oldGovernor, address indexed newGovernor)
func (_Configurator *ConfiguratorFilterer) ParseSetGovernor(log types.Log) (*ConfiguratorSetGovernor, error) {
	event := new(ConfiguratorSetGovernor)
	if err := _Configurator.contract.UnpackLog(event, "SetGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetPauseGuardianIterator is returned from FilterSetPauseGuardian and is used to iterate over the raw logs and unpacked data for SetPauseGuardian events raised by the Configurator contract.
type ConfiguratorSetPauseGuardianIterator struct {
	Event *ConfiguratorSetPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetPauseGuardian represents a SetPauseGuardian event raised by the Configurator contract.
type ConfiguratorSetPauseGuardian struct {
	CometProxy       common.Address
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetPauseGuardian is a free log retrieval operation binding the contract event 0xb60cc4e8fd32e3f131099a821d4716afdb1e906f90421b5edb21070d209d0ee2.
//
// Solidity: event SetPauseGuardian(address indexed cometProxy, address indexed oldPauseGuardian, address indexed newPauseGuardian)
func (_Configurator *ConfiguratorFilterer) FilterSetPauseGuardian(opts *bind.FilterOpts, cometProxy []common.Address, oldPauseGuardian []common.Address, newPauseGuardian []common.Address) (*ConfiguratorSetPauseGuardianIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldPauseGuardianRule []interface{}
	for _, oldPauseGuardianItem := range oldPauseGuardian {
		oldPauseGuardianRule = append(oldPauseGuardianRule, oldPauseGuardianItem)
	}
	var newPauseGuardianRule []interface{}
	for _, newPauseGuardianItem := range newPauseGuardian {
		newPauseGuardianRule = append(newPauseGuardianRule, newPauseGuardianItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetPauseGuardian", cometProxyRule, oldPauseGuardianRule, newPauseGuardianRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetPauseGuardianIterator{contract: _Configurator.contract, event: "SetPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchSetPauseGuardian is a free log subscription operation binding the contract event 0xb60cc4e8fd32e3f131099a821d4716afdb1e906f90421b5edb21070d209d0ee2.
//
// Solidity: event SetPauseGuardian(address indexed cometProxy, address indexed oldPauseGuardian, address indexed newPauseGuardian)
func (_Configurator *ConfiguratorFilterer) WatchSetPauseGuardian(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetPauseGuardian, cometProxy []common.Address, oldPauseGuardian []common.Address, newPauseGuardian []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var oldPauseGuardianRule []interface{}
	for _, oldPauseGuardianItem := range oldPauseGuardian {
		oldPauseGuardianRule = append(oldPauseGuardianRule, oldPauseGuardianItem)
	}
	var newPauseGuardianRule []interface{}
	for _, newPauseGuardianItem := range newPauseGuardian {
		newPauseGuardianRule = append(newPauseGuardianRule, newPauseGuardianItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetPauseGuardian", cometProxyRule, oldPauseGuardianRule, newPauseGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetPauseGuardian)
				if err := _Configurator.contract.UnpackLog(event, "SetPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetPauseGuardian is a log parse operation binding the contract event 0xb60cc4e8fd32e3f131099a821d4716afdb1e906f90421b5edb21070d209d0ee2.
//
// Solidity: event SetPauseGuardian(address indexed cometProxy, address indexed oldPauseGuardian, address indexed newPauseGuardian)
func (_Configurator *ConfiguratorFilterer) ParseSetPauseGuardian(log types.Log) (*ConfiguratorSetPauseGuardian, error) {
	event := new(ConfiguratorSetPauseGuardian)
	if err := _Configurator.contract.UnpackLog(event, "SetPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetStoreFrontPriceFactorIterator is returned from FilterSetStoreFrontPriceFactor and is used to iterate over the raw logs and unpacked data for SetStoreFrontPriceFactor events raised by the Configurator contract.
type ConfiguratorSetStoreFrontPriceFactorIterator struct {
	Event *ConfiguratorSetStoreFrontPriceFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetStoreFrontPriceFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetStoreFrontPriceFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetStoreFrontPriceFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetStoreFrontPriceFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetStoreFrontPriceFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetStoreFrontPriceFactor represents a SetStoreFrontPriceFactor event raised by the Configurator contract.
type ConfiguratorSetStoreFrontPriceFactor struct {
	CometProxy               common.Address
	OldStoreFrontPriceFactor uint64
	NewStoreFrontPriceFactor uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetStoreFrontPriceFactor is a free log retrieval operation binding the contract event 0x864cdb8aef050e3718f73f8f3512418f04745083156b40523b128e79b411fc2d.
//
// Solidity: event SetStoreFrontPriceFactor(address indexed cometProxy, uint64 oldStoreFrontPriceFactor, uint64 newStoreFrontPriceFactor)
func (_Configurator *ConfiguratorFilterer) FilterSetStoreFrontPriceFactor(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetStoreFrontPriceFactorIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetStoreFrontPriceFactor", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetStoreFrontPriceFactorIterator{contract: _Configurator.contract, event: "SetStoreFrontPriceFactor", logs: logs, sub: sub}, nil
}

// WatchSetStoreFrontPriceFactor is a free log subscription operation binding the contract event 0x864cdb8aef050e3718f73f8f3512418f04745083156b40523b128e79b411fc2d.
//
// Solidity: event SetStoreFrontPriceFactor(address indexed cometProxy, uint64 oldStoreFrontPriceFactor, uint64 newStoreFrontPriceFactor)
func (_Configurator *ConfiguratorFilterer) WatchSetStoreFrontPriceFactor(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetStoreFrontPriceFactor, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetStoreFrontPriceFactor", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetStoreFrontPriceFactor)
				if err := _Configurator.contract.UnpackLog(event, "SetStoreFrontPriceFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetStoreFrontPriceFactor is a log parse operation binding the contract event 0x864cdb8aef050e3718f73f8f3512418f04745083156b40523b128e79b411fc2d.
//
// Solidity: event SetStoreFrontPriceFactor(address indexed cometProxy, uint64 oldStoreFrontPriceFactor, uint64 newStoreFrontPriceFactor)
func (_Configurator *ConfiguratorFilterer) ParseSetStoreFrontPriceFactor(log types.Log) (*ConfiguratorSetStoreFrontPriceFactor, error) {
	event := new(ConfiguratorSetStoreFrontPriceFactor)
	if err := _Configurator.contract.UnpackLog(event, "SetStoreFrontPriceFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetSupplyKinkIterator is returned from FilterSetSupplyKink and is used to iterate over the raw logs and unpacked data for SetSupplyKink events raised by the Configurator contract.
type ConfiguratorSetSupplyKinkIterator struct {
	Event *ConfiguratorSetSupplyKink // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetSupplyKinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetSupplyKink)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetSupplyKink)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetSupplyKinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetSupplyKinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetSupplyKink represents a SetSupplyKink event raised by the Configurator contract.
type ConfiguratorSetSupplyKink struct {
	CometProxy common.Address
	OldKink    uint64
	NewKink    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyKink is a free log retrieval operation binding the contract event 0x35ebb489f572b08259fa0a007e62d8390043159fcf4a6c9a517dd1eb4dc77dfc.
//
// Solidity: event SetSupplyKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) FilterSetSupplyKink(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetSupplyKinkIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetSupplyKink", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetSupplyKinkIterator{contract: _Configurator.contract, event: "SetSupplyKink", logs: logs, sub: sub}, nil
}

// WatchSetSupplyKink is a free log subscription operation binding the contract event 0x35ebb489f572b08259fa0a007e62d8390043159fcf4a6c9a517dd1eb4dc77dfc.
//
// Solidity: event SetSupplyKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) WatchSetSupplyKink(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetSupplyKink, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetSupplyKink", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetSupplyKink)
				if err := _Configurator.contract.UnpackLog(event, "SetSupplyKink", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyKink is a log parse operation binding the contract event 0x35ebb489f572b08259fa0a007e62d8390043159fcf4a6c9a517dd1eb4dc77dfc.
//
// Solidity: event SetSupplyKink(address indexed cometProxy, uint64 oldKink, uint64 newKink)
func (_Configurator *ConfiguratorFilterer) ParseSetSupplyKink(log types.Log) (*ConfiguratorSetSupplyKink, error) {
	event := new(ConfiguratorSetSupplyKink)
	if err := _Configurator.contract.UnpackLog(event, "SetSupplyKink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetSupplyPerYearInterestRateBaseIterator is returned from FilterSetSupplyPerYearInterestRateBase and is used to iterate over the raw logs and unpacked data for SetSupplyPerYearInterestRateBase events raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateBaseIterator struct {
	Event *ConfiguratorSetSupplyPerYearInterestRateBase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetSupplyPerYearInterestRateBaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetSupplyPerYearInterestRateBase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetSupplyPerYearInterestRateBase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetSupplyPerYearInterestRateBaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetSupplyPerYearInterestRateBaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetSupplyPerYearInterestRateBase represents a SetSupplyPerYearInterestRateBase event raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateBase struct {
	CometProxy common.Address
	OldIRBase  uint64
	NewIRBase  uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyPerYearInterestRateBase is a free log retrieval operation binding the contract event 0x56027dd4756fabe0c40c7de2c732b95d170796feaf975883a17a087b3718e6b9.
//
// Solidity: event SetSupplyPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) FilterSetSupplyPerYearInterestRateBase(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetSupplyPerYearInterestRateBaseIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetSupplyPerYearInterestRateBase", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetSupplyPerYearInterestRateBaseIterator{contract: _Configurator.contract, event: "SetSupplyPerYearInterestRateBase", logs: logs, sub: sub}, nil
}

// WatchSetSupplyPerYearInterestRateBase is a free log subscription operation binding the contract event 0x56027dd4756fabe0c40c7de2c732b95d170796feaf975883a17a087b3718e6b9.
//
// Solidity: event SetSupplyPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) WatchSetSupplyPerYearInterestRateBase(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetSupplyPerYearInterestRateBase, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetSupplyPerYearInterestRateBase", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetSupplyPerYearInterestRateBase)
				if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateBase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyPerYearInterestRateBase is a log parse operation binding the contract event 0x56027dd4756fabe0c40c7de2c732b95d170796feaf975883a17a087b3718e6b9.
//
// Solidity: event SetSupplyPerYearInterestRateBase(address indexed cometProxy, uint64 oldIRBase, uint64 newIRBase)
func (_Configurator *ConfiguratorFilterer) ParseSetSupplyPerYearInterestRateBase(log types.Log) (*ConfiguratorSetSupplyPerYearInterestRateBase, error) {
	event := new(ConfiguratorSetSupplyPerYearInterestRateBase)
	if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateBase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator is returned from FilterSetSupplyPerYearInterestRateSlopeHigh and is used to iterate over the raw logs and unpacked data for SetSupplyPerYearInterestRateSlopeHigh events raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator struct {
	Event *ConfiguratorSetSupplyPerYearInterestRateSlopeHigh // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetSupplyPerYearInterestRateSlopeHigh)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetSupplyPerYearInterestRateSlopeHigh)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetSupplyPerYearInterestRateSlopeHigh represents a SetSupplyPerYearInterestRateSlopeHigh event raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateSlopeHigh struct {
	CometProxy     common.Address
	OldIRSlopeHigh uint64
	NewIRSlopeHigh uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyPerYearInterestRateSlopeHigh is a free log retrieval operation binding the contract event 0x9fd77f9236bc26e7c30eaa679ccff4200f9a813e5950b2904c29f11d0b5f2a5e.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) FilterSetSupplyPerYearInterestRateSlopeHigh(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetSupplyPerYearInterestRateSlopeHigh", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetSupplyPerYearInterestRateSlopeHighIterator{contract: _Configurator.contract, event: "SetSupplyPerYearInterestRateSlopeHigh", logs: logs, sub: sub}, nil
}

// WatchSetSupplyPerYearInterestRateSlopeHigh is a free log subscription operation binding the contract event 0x9fd77f9236bc26e7c30eaa679ccff4200f9a813e5950b2904c29f11d0b5f2a5e.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) WatchSetSupplyPerYearInterestRateSlopeHigh(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetSupplyPerYearInterestRateSlopeHigh, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetSupplyPerYearInterestRateSlopeHigh", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetSupplyPerYearInterestRateSlopeHigh)
				if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateSlopeHigh", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyPerYearInterestRateSlopeHigh is a log parse operation binding the contract event 0x9fd77f9236bc26e7c30eaa679ccff4200f9a813e5950b2904c29f11d0b5f2a5e.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeHigh(address indexed cometProxy, uint64 oldIRSlopeHigh, uint64 newIRSlopeHigh)
func (_Configurator *ConfiguratorFilterer) ParseSetSupplyPerYearInterestRateSlopeHigh(log types.Log) (*ConfiguratorSetSupplyPerYearInterestRateSlopeHigh, error) {
	event := new(ConfiguratorSetSupplyPerYearInterestRateSlopeHigh)
	if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateSlopeHigh", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator is returned from FilterSetSupplyPerYearInterestRateSlopeLow and is used to iterate over the raw logs and unpacked data for SetSupplyPerYearInterestRateSlopeLow events raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator struct {
	Event *ConfiguratorSetSupplyPerYearInterestRateSlopeLow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetSupplyPerYearInterestRateSlopeLow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetSupplyPerYearInterestRateSlopeLow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetSupplyPerYearInterestRateSlopeLow represents a SetSupplyPerYearInterestRateSlopeLow event raised by the Configurator contract.
type ConfiguratorSetSupplyPerYearInterestRateSlopeLow struct {
	CometProxy    common.Address
	OldIRSlopeLow uint64
	NewIRSlopeLow uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyPerYearInterestRateSlopeLow is a free log retrieval operation binding the contract event 0xc936b3eb07b584b686d11042214266fed11036dbc226159d67b3cfaa4ce26f78.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) FilterSetSupplyPerYearInterestRateSlopeLow(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetSupplyPerYearInterestRateSlopeLow", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetSupplyPerYearInterestRateSlopeLowIterator{contract: _Configurator.contract, event: "SetSupplyPerYearInterestRateSlopeLow", logs: logs, sub: sub}, nil
}

// WatchSetSupplyPerYearInterestRateSlopeLow is a free log subscription operation binding the contract event 0xc936b3eb07b584b686d11042214266fed11036dbc226159d67b3cfaa4ce26f78.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) WatchSetSupplyPerYearInterestRateSlopeLow(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetSupplyPerYearInterestRateSlopeLow, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetSupplyPerYearInterestRateSlopeLow", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetSupplyPerYearInterestRateSlopeLow)
				if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateSlopeLow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetSupplyPerYearInterestRateSlopeLow is a log parse operation binding the contract event 0xc936b3eb07b584b686d11042214266fed11036dbc226159d67b3cfaa4ce26f78.
//
// Solidity: event SetSupplyPerYearInterestRateSlopeLow(address indexed cometProxy, uint64 oldIRSlopeLow, uint64 newIRSlopeLow)
func (_Configurator *ConfiguratorFilterer) ParseSetSupplyPerYearInterestRateSlopeLow(log types.Log) (*ConfiguratorSetSupplyPerYearInterestRateSlopeLow, error) {
	event := new(ConfiguratorSetSupplyPerYearInterestRateSlopeLow)
	if err := _Configurator.contract.UnpackLog(event, "SetSupplyPerYearInterestRateSlopeLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorSetTargetReservesIterator is returned from FilterSetTargetReserves and is used to iterate over the raw logs and unpacked data for SetTargetReserves events raised by the Configurator contract.
type ConfiguratorSetTargetReservesIterator struct {
	Event *ConfiguratorSetTargetReserves // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorSetTargetReservesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorSetTargetReserves)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorSetTargetReserves)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorSetTargetReservesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorSetTargetReservesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorSetTargetReserves represents a SetTargetReserves event raised by the Configurator contract.
type ConfiguratorSetTargetReserves struct {
	CometProxy        common.Address
	OldTargetReserves *big.Int
	NewTargetReserves *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetTargetReserves is a free log retrieval operation binding the contract event 0x0fcb429b7790258a8e50ab0f30505a67d74d6571119febb9690e6264594e2e29.
//
// Solidity: event SetTargetReserves(address indexed cometProxy, uint104 oldTargetReserves, uint104 newTargetReserves)
func (_Configurator *ConfiguratorFilterer) FilterSetTargetReserves(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorSetTargetReservesIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "SetTargetReserves", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorSetTargetReservesIterator{contract: _Configurator.contract, event: "SetTargetReserves", logs: logs, sub: sub}, nil
}

// WatchSetTargetReserves is a free log subscription operation binding the contract event 0x0fcb429b7790258a8e50ab0f30505a67d74d6571119febb9690e6264594e2e29.
//
// Solidity: event SetTargetReserves(address indexed cometProxy, uint104 oldTargetReserves, uint104 newTargetReserves)
func (_Configurator *ConfiguratorFilterer) WatchSetTargetReserves(opts *bind.WatchOpts, sink chan<- *ConfiguratorSetTargetReserves, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "SetTargetReserves", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorSetTargetReserves)
				if err := _Configurator.contract.UnpackLog(event, "SetTargetReserves", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTargetReserves is a log parse operation binding the contract event 0x0fcb429b7790258a8e50ab0f30505a67d74d6571119febb9690e6264594e2e29.
//
// Solidity: event SetTargetReserves(address indexed cometProxy, uint104 oldTargetReserves, uint104 newTargetReserves)
func (_Configurator *ConfiguratorFilterer) ParseSetTargetReserves(log types.Log) (*ConfiguratorSetTargetReserves, error) {
	event := new(ConfiguratorSetTargetReserves)
	if err := _Configurator.contract.UnpackLog(event, "SetTargetReserves", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetIterator is returned from FilterUpdateAsset and is used to iterate over the raw logs and unpacked data for UpdateAsset events raised by the Configurator contract.
type ConfiguratorUpdateAssetIterator struct {
	Event *ConfiguratorUpdateAsset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAsset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAsset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAsset represents a UpdateAsset event raised by the Configurator contract.
type ConfiguratorUpdateAsset struct {
	CometProxy     common.Address
	OldAssetConfig CometConfigurationAssetConfig
	NewAssetConfig CometConfigurationAssetConfig
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateAsset is a free log retrieval operation binding the contract event 0xf0d2e933bc5a83ab653c27f5ae312ee5f4a394a45c34bb90e8c790bf0ed38341.
//
// Solidity: event UpdateAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) oldAssetConfig, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAsset(opts *bind.FilterOpts, cometProxy []common.Address) (*ConfiguratorUpdateAssetIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAsset", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetIterator{contract: _Configurator.contract, event: "UpdateAsset", logs: logs, sub: sub}, nil
}

// WatchUpdateAsset is a free log subscription operation binding the contract event 0xf0d2e933bc5a83ab653c27f5ae312ee5f4a394a45c34bb90e8c790bf0ed38341.
//
// Solidity: event UpdateAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) oldAssetConfig, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAsset(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAsset, cometProxy []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAsset", cometProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAsset)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAsset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAsset is a log parse operation binding the contract event 0xf0d2e933bc5a83ab653c27f5ae312ee5f4a394a45c34bb90e8c790bf0ed38341.
//
// Solidity: event UpdateAsset(address indexed cometProxy, (address,address,uint8,uint64,uint64,uint64,uint128) oldAssetConfig, (address,address,uint8,uint64,uint64,uint64,uint128) newAssetConfig)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAsset(log types.Log) (*ConfiguratorUpdateAsset, error) {
	event := new(ConfiguratorUpdateAsset)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetBorrowCollateralFactorIterator is returned from FilterUpdateAssetBorrowCollateralFactor and is used to iterate over the raw logs and unpacked data for UpdateAssetBorrowCollateralFactor events raised by the Configurator contract.
type ConfiguratorUpdateAssetBorrowCollateralFactorIterator struct {
	Event *ConfiguratorUpdateAssetBorrowCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetBorrowCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAssetBorrowCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAssetBorrowCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetBorrowCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetBorrowCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAssetBorrowCollateralFactor represents a UpdateAssetBorrowCollateralFactor event raised by the Configurator contract.
type ConfiguratorUpdateAssetBorrowCollateralFactor struct {
	CometProxy  common.Address
	Asset       common.Address
	OldBorrowCF uint64
	NewBorrowCF uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateAssetBorrowCollateralFactor is a free log retrieval operation binding the contract event 0x14bba7d3df0797b2f5143e4f2087f29fcd3ee3ec3197d38a7f77f85dc0d0ec63.
//
// Solidity: event UpdateAssetBorrowCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldBorrowCF, uint64 newBorrowCF)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAssetBorrowCollateralFactor(opts *bind.FilterOpts, cometProxy []common.Address, asset []common.Address) (*ConfiguratorUpdateAssetBorrowCollateralFactorIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAssetBorrowCollateralFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetBorrowCollateralFactorIterator{contract: _Configurator.contract, event: "UpdateAssetBorrowCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateAssetBorrowCollateralFactor is a free log subscription operation binding the contract event 0x14bba7d3df0797b2f5143e4f2087f29fcd3ee3ec3197d38a7f77f85dc0d0ec63.
//
// Solidity: event UpdateAssetBorrowCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldBorrowCF, uint64 newBorrowCF)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAssetBorrowCollateralFactor(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAssetBorrowCollateralFactor, cometProxy []common.Address, asset []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAssetBorrowCollateralFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAssetBorrowCollateralFactor)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAssetBorrowCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAssetBorrowCollateralFactor is a log parse operation binding the contract event 0x14bba7d3df0797b2f5143e4f2087f29fcd3ee3ec3197d38a7f77f85dc0d0ec63.
//
// Solidity: event UpdateAssetBorrowCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldBorrowCF, uint64 newBorrowCF)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAssetBorrowCollateralFactor(log types.Log) (*ConfiguratorUpdateAssetBorrowCollateralFactor, error) {
	event := new(ConfiguratorUpdateAssetBorrowCollateralFactor)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAssetBorrowCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetLiquidateCollateralFactorIterator is returned from FilterUpdateAssetLiquidateCollateralFactor and is used to iterate over the raw logs and unpacked data for UpdateAssetLiquidateCollateralFactor events raised by the Configurator contract.
type ConfiguratorUpdateAssetLiquidateCollateralFactorIterator struct {
	Event *ConfiguratorUpdateAssetLiquidateCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetLiquidateCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAssetLiquidateCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAssetLiquidateCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetLiquidateCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetLiquidateCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAssetLiquidateCollateralFactor represents a UpdateAssetLiquidateCollateralFactor event raised by the Configurator contract.
type ConfiguratorUpdateAssetLiquidateCollateralFactor struct {
	CometProxy     common.Address
	Asset          common.Address
	OldLiquidateCF uint64
	NewLiquidateCF uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateAssetLiquidateCollateralFactor is a free log retrieval operation binding the contract event 0xbdfea73f1d581ded04f608d47cd124146963b03795e30c9f9e16d57bb1fed652.
//
// Solidity: event UpdateAssetLiquidateCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidateCF, uint64 newLiquidateCF)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAssetLiquidateCollateralFactor(opts *bind.FilterOpts, cometProxy []common.Address, asset []common.Address) (*ConfiguratorUpdateAssetLiquidateCollateralFactorIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAssetLiquidateCollateralFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetLiquidateCollateralFactorIterator{contract: _Configurator.contract, event: "UpdateAssetLiquidateCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateAssetLiquidateCollateralFactor is a free log subscription operation binding the contract event 0xbdfea73f1d581ded04f608d47cd124146963b03795e30c9f9e16d57bb1fed652.
//
// Solidity: event UpdateAssetLiquidateCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidateCF, uint64 newLiquidateCF)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAssetLiquidateCollateralFactor(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAssetLiquidateCollateralFactor, cometProxy []common.Address, asset []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAssetLiquidateCollateralFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAssetLiquidateCollateralFactor)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAssetLiquidateCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAssetLiquidateCollateralFactor is a log parse operation binding the contract event 0xbdfea73f1d581ded04f608d47cd124146963b03795e30c9f9e16d57bb1fed652.
//
// Solidity: event UpdateAssetLiquidateCollateralFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidateCF, uint64 newLiquidateCF)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAssetLiquidateCollateralFactor(log types.Log) (*ConfiguratorUpdateAssetLiquidateCollateralFactor, error) {
	event := new(ConfiguratorUpdateAssetLiquidateCollateralFactor)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAssetLiquidateCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetLiquidationFactorIterator is returned from FilterUpdateAssetLiquidationFactor and is used to iterate over the raw logs and unpacked data for UpdateAssetLiquidationFactor events raised by the Configurator contract.
type ConfiguratorUpdateAssetLiquidationFactorIterator struct {
	Event *ConfiguratorUpdateAssetLiquidationFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetLiquidationFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAssetLiquidationFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAssetLiquidationFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetLiquidationFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetLiquidationFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAssetLiquidationFactor represents a UpdateAssetLiquidationFactor event raised by the Configurator contract.
type ConfiguratorUpdateAssetLiquidationFactor struct {
	CometProxy           common.Address
	Asset                common.Address
	OldLiquidationFactor uint64
	NewLiquidationFactor uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAssetLiquidationFactor is a free log retrieval operation binding the contract event 0x4bce180f5fbcc3f04e85605e3179ceaa66041248373ced1a406d4c44d91e3797.
//
// Solidity: event UpdateAssetLiquidationFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidationFactor, uint64 newLiquidationFactor)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAssetLiquidationFactor(opts *bind.FilterOpts, cometProxy []common.Address, asset []common.Address) (*ConfiguratorUpdateAssetLiquidationFactorIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAssetLiquidationFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetLiquidationFactorIterator{contract: _Configurator.contract, event: "UpdateAssetLiquidationFactor", logs: logs, sub: sub}, nil
}

// WatchUpdateAssetLiquidationFactor is a free log subscription operation binding the contract event 0x4bce180f5fbcc3f04e85605e3179ceaa66041248373ced1a406d4c44d91e3797.
//
// Solidity: event UpdateAssetLiquidationFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidationFactor, uint64 newLiquidationFactor)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAssetLiquidationFactor(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAssetLiquidationFactor, cometProxy []common.Address, asset []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAssetLiquidationFactor", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAssetLiquidationFactor)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAssetLiquidationFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAssetLiquidationFactor is a log parse operation binding the contract event 0x4bce180f5fbcc3f04e85605e3179ceaa66041248373ced1a406d4c44d91e3797.
//
// Solidity: event UpdateAssetLiquidationFactor(address indexed cometProxy, address indexed asset, uint64 oldLiquidationFactor, uint64 newLiquidationFactor)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAssetLiquidationFactor(log types.Log) (*ConfiguratorUpdateAssetLiquidationFactor, error) {
	event := new(ConfiguratorUpdateAssetLiquidationFactor)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAssetLiquidationFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetPriceFeedIterator is returned from FilterUpdateAssetPriceFeed and is used to iterate over the raw logs and unpacked data for UpdateAssetPriceFeed events raised by the Configurator contract.
type ConfiguratorUpdateAssetPriceFeedIterator struct {
	Event *ConfiguratorUpdateAssetPriceFeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAssetPriceFeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAssetPriceFeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAssetPriceFeed represents a UpdateAssetPriceFeed event raised by the Configurator contract.
type ConfiguratorUpdateAssetPriceFeed struct {
	CometProxy   common.Address
	Asset        common.Address
	OldPriceFeed common.Address
	NewPriceFeed common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAssetPriceFeed is a free log retrieval operation binding the contract event 0x1736faf0e37344d44b8757867808a35e1d4961804de8c3dec2896bf1c1d6fc51.
//
// Solidity: event UpdateAssetPriceFeed(address indexed cometProxy, address indexed asset, address oldPriceFeed, address newPriceFeed)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAssetPriceFeed(opts *bind.FilterOpts, cometProxy []common.Address, asset []common.Address) (*ConfiguratorUpdateAssetPriceFeedIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAssetPriceFeed", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetPriceFeedIterator{contract: _Configurator.contract, event: "UpdateAssetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchUpdateAssetPriceFeed is a free log subscription operation binding the contract event 0x1736faf0e37344d44b8757867808a35e1d4961804de8c3dec2896bf1c1d6fc51.
//
// Solidity: event UpdateAssetPriceFeed(address indexed cometProxy, address indexed asset, address oldPriceFeed, address newPriceFeed)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAssetPriceFeed(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAssetPriceFeed, cometProxy []common.Address, asset []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAssetPriceFeed", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAssetPriceFeed)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAssetPriceFeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAssetPriceFeed is a log parse operation binding the contract event 0x1736faf0e37344d44b8757867808a35e1d4961804de8c3dec2896bf1c1d6fc51.
//
// Solidity: event UpdateAssetPriceFeed(address indexed cometProxy, address indexed asset, address oldPriceFeed, address newPriceFeed)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAssetPriceFeed(log types.Log) (*ConfiguratorUpdateAssetPriceFeed, error) {
	event := new(ConfiguratorUpdateAssetPriceFeed)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAssetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConfiguratorUpdateAssetSupplyCapIterator is returned from FilterUpdateAssetSupplyCap and is used to iterate over the raw logs and unpacked data for UpdateAssetSupplyCap events raised by the Configurator contract.
type ConfiguratorUpdateAssetSupplyCapIterator struct {
	Event *ConfiguratorUpdateAssetSupplyCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConfiguratorUpdateAssetSupplyCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorUpdateAssetSupplyCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConfiguratorUpdateAssetSupplyCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConfiguratorUpdateAssetSupplyCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfiguratorUpdateAssetSupplyCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfiguratorUpdateAssetSupplyCap represents a UpdateAssetSupplyCap event raised by the Configurator contract.
type ConfiguratorUpdateAssetSupplyCap struct {
	CometProxy   common.Address
	Asset        common.Address
	OldSupplyCap *big.Int
	NewSupplyCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAssetSupplyCap is a free log retrieval operation binding the contract event 0x561b42b6b7ded5005c55634f0a6bd24d306cb04b04533bab7e1a644b67245486.
//
// Solidity: event UpdateAssetSupplyCap(address indexed cometProxy, address indexed asset, uint128 oldSupplyCap, uint128 newSupplyCap)
func (_Configurator *ConfiguratorFilterer) FilterUpdateAssetSupplyCap(opts *bind.FilterOpts, cometProxy []common.Address, asset []common.Address) (*ConfiguratorUpdateAssetSupplyCapIterator, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "UpdateAssetSupplyCap", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorUpdateAssetSupplyCapIterator{contract: _Configurator.contract, event: "UpdateAssetSupplyCap", logs: logs, sub: sub}, nil
}

// WatchUpdateAssetSupplyCap is a free log subscription operation binding the contract event 0x561b42b6b7ded5005c55634f0a6bd24d306cb04b04533bab7e1a644b67245486.
//
// Solidity: event UpdateAssetSupplyCap(address indexed cometProxy, address indexed asset, uint128 oldSupplyCap, uint128 newSupplyCap)
func (_Configurator *ConfiguratorFilterer) WatchUpdateAssetSupplyCap(opts *bind.WatchOpts, sink chan<- *ConfiguratorUpdateAssetSupplyCap, cometProxy []common.Address, asset []common.Address) (event.Subscription, error) {

	var cometProxyRule []interface{}
	for _, cometProxyItem := range cometProxy {
		cometProxyRule = append(cometProxyRule, cometProxyItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "UpdateAssetSupplyCap", cometProxyRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfiguratorUpdateAssetSupplyCap)
				if err := _Configurator.contract.UnpackLog(event, "UpdateAssetSupplyCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateAssetSupplyCap is a log parse operation binding the contract event 0x561b42b6b7ded5005c55634f0a6bd24d306cb04b04533bab7e1a644b67245486.
//
// Solidity: event UpdateAssetSupplyCap(address indexed cometProxy, address indexed asset, uint128 oldSupplyCap, uint128 newSupplyCap)
func (_Configurator *ConfiguratorFilterer) ParseUpdateAssetSupplyCap(log types.Log) (*ConfiguratorUpdateAssetSupplyCap, error) {
	event := new(ConfiguratorUpdateAssetSupplyCap)
	if err := _Configurator.contract.UnpackLog(event, "UpdateAssetSupplyCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
