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

// // CometConfigurationAssetConfig is an auto generated low-level Go binding around an user-defined struct.
// type CometConfigurationAssetConfig struct {
// 	Asset                     common.Address
// 	PriceFeed                 common.Address
// 	Decimals                  uint8
// 	BorrowCollateralFactor    uint64
// 	LiquidateCollateralFactor uint64
// 	LiquidationFactor         uint64
// 	SupplyCap                 *big.Int
// }

// // CometConfigurationConfiguration is an auto generated low-level Go binding around an user-defined struct.
// type CometConfigurationConfiguration struct {
// 	Governor                           common.Address
// 	PauseGuardian                      common.Address
// 	BaseToken                          common.Address
// 	BaseTokenPriceFeed                 common.Address
// 	ExtensionDelegate                  common.Address
// 	SupplyKink                         uint64
// 	SupplyPerYearInterestRateSlopeLow  uint64
// 	SupplyPerYearInterestRateSlopeHigh uint64
// 	SupplyPerYearInterestRateBase      uint64
// 	BorrowKink                         uint64
// 	BorrowPerYearInterestRateSlopeLow  uint64
// 	BorrowPerYearInterestRateSlopeHigh uint64
// 	BorrowPerYearInterestRateBase      uint64
// 	StoreFrontPriceFactor              uint64
// 	TrackingIndexScale                 uint64
// 	BaseTrackingSupplySpeed            uint64
// 	BaseTrackingBorrowSpeed            uint64
// 	BaseMinForRewards                  *big.Int
// 	BaseBorrowMin                      *big.Int
// 	TargetReserves                     *big.Int
// 	AssetConfigs                       []CometConfigurationAssetConfig
// }

// CometCoreAssetInfo is an auto generated low-level Go binding around an user-defined struct.
type CometCoreAssetInfo struct {
	Offset                    uint8
	Asset                     common.Address
	PriceFeed                 common.Address
	Scale                     uint64
	BorrowCollateralFactor    uint64
	LiquidateCollateralFactor uint64
	LiquidationFactor         uint64
	SupplyCap                 *big.Int
}

// CometMetaData contains all meta data concerning the Comet contract.
var CometMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structCometConfiguration.Configuration\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Absurd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInt104\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt104\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NegativeNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferInFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferOutFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAbsorbed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basePaidOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"BuyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"PauseAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SupplyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawReserves\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"absorb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accrueAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveThis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseBorrowMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseMinForRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingBorrowSpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingSupplySpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"buyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometCore.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetInfoByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometCore.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getCollateralReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAbsorbPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBorrowCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBuyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSupplyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWithdrawPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidatorPoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"numAbsorbs\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"numAbsorbed\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"approxSpend\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_reserved\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"quoteCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storeFrontPriceFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalsCollateral\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"totalSupplyAsset\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_reserved\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackingIndexScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAssetFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBasic\",\"outputs\":[{\"internalType\":\"int104\",\"name\":\"principal\",\"type\":\"int104\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingAccrued\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"assetsIn\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_reserved\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCollateral\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_reserved\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CometABI is the input ABI used to generate the binding from.
// Deprecated: Use CometMetaData.ABI instead.
var CometABI = CometMetaData.ABI

// Comet is an auto generated Go binding around an Ethereum contract.
type Comet struct {
	CometCaller     // Read-only binding to the contract
	CometTransactor // Write-only binding to the contract
	CometFilterer   // Log filterer for contract events
}

// CometCaller is an auto generated read-only Go binding around an Ethereum contract.
type CometCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CometTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CometFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CometSession struct {
	Contract     *Comet            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CometCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CometCallerSession struct {
	Contract *CometCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CometTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CometTransactorSession struct {
	Contract     *CometTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CometRaw is an auto generated low-level Go binding around an Ethereum contract.
type CometRaw struct {
	Contract *Comet // Generic contract binding to access the raw methods on
}

// CometCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CometCallerRaw struct {
	Contract *CometCaller // Generic read-only contract binding to access the raw methods on
}

// CometTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CometTransactorRaw struct {
	Contract *CometTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComet creates a new instance of Comet, bound to a specific deployed contract.
func NewComet(address common.Address, backend bind.ContractBackend) (*Comet, error) {
	contract, err := bindComet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comet{CometCaller: CometCaller{contract: contract}, CometTransactor: CometTransactor{contract: contract}, CometFilterer: CometFilterer{contract: contract}}, nil
}

// NewCometCaller creates a new read-only instance of Comet, bound to a specific deployed contract.
func NewCometCaller(address common.Address, caller bind.ContractCaller) (*CometCaller, error) {
	contract, err := bindComet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CometCaller{contract: contract}, nil
}

// NewCometTransactor creates a new write-only instance of Comet, bound to a specific deployed contract.
func NewCometTransactor(address common.Address, transactor bind.ContractTransactor) (*CometTransactor, error) {
	contract, err := bindComet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CometTransactor{contract: contract}, nil
}

// NewCometFilterer creates a new log filterer instance of Comet, bound to a specific deployed contract.
func NewCometFilterer(address common.Address, filterer bind.ContractFilterer) (*CometFilterer, error) {
	contract, err := bindComet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CometFilterer{contract: contract}, nil
}

// bindComet binds a generic wrapper to an already deployed contract.
func bindComet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comet *CometRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comet.Contract.CometCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comet *CometRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comet.Contract.CometTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comet *CometRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comet.Contract.CometTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comet *CometCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comet *CometTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comet *CometTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comet.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Comet *CometCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Comet *CometSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Comet.Contract.BalanceOf(&_Comet.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Comet *CometCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Comet.Contract.BalanceOf(&_Comet.CallOpts, account)
}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_Comet *CometCaller) BaseBorrowMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseBorrowMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_Comet *CometSession) BaseBorrowMin() (*big.Int, error) {
	return _Comet.Contract.BaseBorrowMin(&_Comet.CallOpts)
}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_Comet *CometCallerSession) BaseBorrowMin() (*big.Int, error) {
	return _Comet.Contract.BaseBorrowMin(&_Comet.CallOpts)
}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_Comet *CometCaller) BaseMinForRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseMinForRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_Comet *CometSession) BaseMinForRewards() (*big.Int, error) {
	return _Comet.Contract.BaseMinForRewards(&_Comet.CallOpts)
}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_Comet *CometCallerSession) BaseMinForRewards() (*big.Int, error) {
	return _Comet.Contract.BaseMinForRewards(&_Comet.CallOpts)
}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_Comet *CometCaller) BaseScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_Comet *CometSession) BaseScale() (*big.Int, error) {
	return _Comet.Contract.BaseScale(&_Comet.CallOpts)
}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_Comet *CometCallerSession) BaseScale() (*big.Int, error) {
	return _Comet.Contract.BaseScale(&_Comet.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Comet *CometCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Comet *CometSession) BaseToken() (common.Address, error) {
	return _Comet.Contract.BaseToken(&_Comet.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Comet *CometCallerSession) BaseToken() (common.Address, error) {
	return _Comet.Contract.BaseToken(&_Comet.CallOpts)
}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_Comet *CometCaller) BaseTokenPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseTokenPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_Comet *CometSession) BaseTokenPriceFeed() (common.Address, error) {
	return _Comet.Contract.BaseTokenPriceFeed(&_Comet.CallOpts)
}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_Comet *CometCallerSession) BaseTokenPriceFeed() (common.Address, error) {
	return _Comet.Contract.BaseTokenPriceFeed(&_Comet.CallOpts)
}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_Comet *CometCaller) BaseTrackingBorrowSpeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseTrackingBorrowSpeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_Comet *CometSession) BaseTrackingBorrowSpeed() (*big.Int, error) {
	return _Comet.Contract.BaseTrackingBorrowSpeed(&_Comet.CallOpts)
}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_Comet *CometCallerSession) BaseTrackingBorrowSpeed() (*big.Int, error) {
	return _Comet.Contract.BaseTrackingBorrowSpeed(&_Comet.CallOpts)
}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_Comet *CometCaller) BaseTrackingSupplySpeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "baseTrackingSupplySpeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_Comet *CometSession) BaseTrackingSupplySpeed() (*big.Int, error) {
	return _Comet.Contract.BaseTrackingSupplySpeed(&_Comet.CallOpts)
}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_Comet *CometCallerSession) BaseTrackingSupplySpeed() (*big.Int, error) {
	return _Comet.Contract.BaseTrackingSupplySpeed(&_Comet.CallOpts)
}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_Comet *CometCaller) BorrowBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "borrowBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_Comet *CometSession) BorrowBalanceOf(account common.Address) (*big.Int, error) {
	return _Comet.Contract.BorrowBalanceOf(&_Comet.CallOpts, account)
}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_Comet *CometCallerSession) BorrowBalanceOf(account common.Address) (*big.Int, error) {
	return _Comet.Contract.BorrowBalanceOf(&_Comet.CallOpts, account)
}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_Comet *CometCaller) BorrowKink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "borrowKink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_Comet *CometSession) BorrowKink() (*big.Int, error) {
	return _Comet.Contract.BorrowKink(&_Comet.CallOpts)
}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_Comet *CometCallerSession) BorrowKink() (*big.Int, error) {
	return _Comet.Contract.BorrowKink(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometCaller) BorrowPerSecondInterestRateBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "borrowPerSecondInterestRateBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometSession) BorrowPerSecondInterestRateBase() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateBase(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometCallerSession) BorrowPerSecondInterestRateBase() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateBase(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometCaller) BorrowPerSecondInterestRateSlopeHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "borrowPerSecondInterestRateSlopeHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometSession) BorrowPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateSlopeHigh(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometCallerSession) BorrowPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateSlopeHigh(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometCaller) BorrowPerSecondInterestRateSlopeLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "borrowPerSecondInterestRateSlopeLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometSession) BorrowPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateSlopeLow(&_Comet.CallOpts)
}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometCallerSession) BorrowPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _Comet.Contract.BorrowPerSecondInterestRateSlopeLow(&_Comet.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Comet *CometCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Comet *CometSession) Decimals() (uint8, error) {
	return _Comet.Contract.Decimals(&_Comet.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Comet *CometCallerSession) Decimals() (uint8, error) {
	return _Comet.Contract.Decimals(&_Comet.CallOpts)
}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_Comet *CometCaller) ExtensionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "extensionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_Comet *CometSession) ExtensionDelegate() (common.Address, error) {
	return _Comet.Contract.ExtensionDelegate(&_Comet.CallOpts)
}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_Comet *CometCallerSession) ExtensionDelegate() (common.Address, error) {
	return _Comet.Contract.ExtensionDelegate(&_Comet.CallOpts)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometCaller) GetAssetInfo(opts *bind.CallOpts, i uint8) (CometCoreAssetInfo, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getAssetInfo", i)

	if err != nil {
		return *new(CometCoreAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CometCoreAssetInfo)).(*CometCoreAssetInfo)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometSession) GetAssetInfo(i uint8) (CometCoreAssetInfo, error) {
	return _Comet.Contract.GetAssetInfo(&_Comet.CallOpts, i)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometCallerSession) GetAssetInfo(i uint8) (CometCoreAssetInfo, error) {
	return _Comet.Contract.GetAssetInfo(&_Comet.CallOpts, i)
}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometCaller) GetAssetInfoByAddress(opts *bind.CallOpts, asset common.Address) (CometCoreAssetInfo, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getAssetInfoByAddress", asset)

	if err != nil {
		return *new(CometCoreAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CometCoreAssetInfo)).(*CometCoreAssetInfo)

	return out0, err

}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometSession) GetAssetInfoByAddress(asset common.Address) (CometCoreAssetInfo, error) {
	return _Comet.Contract.GetAssetInfoByAddress(&_Comet.CallOpts, asset)
}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_Comet *CometCallerSession) GetAssetInfoByAddress(asset common.Address) (CometCoreAssetInfo, error) {
	return _Comet.Contract.GetAssetInfoByAddress(&_Comet.CallOpts, asset)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_Comet *CometCaller) GetBorrowRate(opts *bind.CallOpts, utilization *big.Int) (uint64, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getBorrowRate", utilization)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_Comet *CometSession) GetBorrowRate(utilization *big.Int) (uint64, error) {
	return _Comet.Contract.GetBorrowRate(&_Comet.CallOpts, utilization)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_Comet *CometCallerSession) GetBorrowRate(utilization *big.Int) (uint64, error) {
	return _Comet.Contract.GetBorrowRate(&_Comet.CallOpts, utilization)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_Comet *CometCaller) GetCollateralReserves(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getCollateralReserves", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_Comet *CometSession) GetCollateralReserves(asset common.Address) (*big.Int, error) {
	return _Comet.Contract.GetCollateralReserves(&_Comet.CallOpts, asset)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_Comet *CometCallerSession) GetCollateralReserves(asset common.Address) (*big.Int, error) {
	return _Comet.Contract.GetCollateralReserves(&_Comet.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_Comet *CometCaller) GetPrice(opts *bind.CallOpts, priceFeed common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getPrice", priceFeed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_Comet *CometSession) GetPrice(priceFeed common.Address) (*big.Int, error) {
	return _Comet.Contract.GetPrice(&_Comet.CallOpts, priceFeed)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_Comet *CometCallerSession) GetPrice(priceFeed common.Address) (*big.Int, error) {
	return _Comet.Contract.GetPrice(&_Comet.CallOpts, priceFeed)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_Comet *CometCaller) GetReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_Comet *CometSession) GetReserves() (*big.Int, error) {
	return _Comet.Contract.GetReserves(&_Comet.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_Comet *CometCallerSession) GetReserves() (*big.Int, error) {
	return _Comet.Contract.GetReserves(&_Comet.CallOpts)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_Comet *CometCaller) GetSupplyRate(opts *bind.CallOpts, utilization *big.Int) (uint64, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getSupplyRate", utilization)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_Comet *CometSession) GetSupplyRate(utilization *big.Int) (uint64, error) {
	return _Comet.Contract.GetSupplyRate(&_Comet.CallOpts, utilization)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_Comet *CometCallerSession) GetSupplyRate(utilization *big.Int) (uint64, error) {
	return _Comet.Contract.GetSupplyRate(&_Comet.CallOpts, utilization)
}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_Comet *CometCaller) GetUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "getUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_Comet *CometSession) GetUtilization() (*big.Int, error) {
	return _Comet.Contract.GetUtilization(&_Comet.CallOpts)
}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_Comet *CometCallerSession) GetUtilization() (*big.Int, error) {
	return _Comet.Contract.GetUtilization(&_Comet.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Comet *CometCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Comet *CometSession) Governor() (common.Address, error) {
	return _Comet.Contract.Governor(&_Comet.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Comet *CometCallerSession) Governor() (common.Address, error) {
	return _Comet.Contract.Governor(&_Comet.CallOpts)
}

// HasPermission is a free data retrieval call binding the contract method 0xcde68041.
//
// Solidity: function hasPermission(address owner, address manager) view returns(bool)
func (_Comet *CometCaller) HasPermission(opts *bind.CallOpts, owner common.Address, manager common.Address) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "hasPermission", owner, manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0xcde68041.
//
// Solidity: function hasPermission(address owner, address manager) view returns(bool)
func (_Comet *CometSession) HasPermission(owner common.Address, manager common.Address) (bool, error) {
	return _Comet.Contract.HasPermission(&_Comet.CallOpts, owner, manager)
}

// HasPermission is a free data retrieval call binding the contract method 0xcde68041.
//
// Solidity: function hasPermission(address owner, address manager) view returns(bool)
func (_Comet *CometCallerSession) HasPermission(owner common.Address, manager common.Address) (bool, error) {
	return _Comet.Contract.HasPermission(&_Comet.CallOpts, owner, manager)
}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_Comet *CometCaller) IsAbsorbPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isAbsorbPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_Comet *CometSession) IsAbsorbPaused() (bool, error) {
	return _Comet.Contract.IsAbsorbPaused(&_Comet.CallOpts)
}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_Comet *CometCallerSession) IsAbsorbPaused() (bool, error) {
	return _Comet.Contract.IsAbsorbPaused(&_Comet.CallOpts)
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1654379.
//
// Solidity: function isAllowed(address , address ) view returns(bool)
func (_Comet *CometCaller) IsAllowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isAllowed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xa1654379.
//
// Solidity: function isAllowed(address , address ) view returns(bool)
func (_Comet *CometSession) IsAllowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Comet.Contract.IsAllowed(&_Comet.CallOpts, arg0, arg1)
}

// IsAllowed is a free data retrieval call binding the contract method 0xa1654379.
//
// Solidity: function isAllowed(address , address ) view returns(bool)
func (_Comet *CometCallerSession) IsAllowed(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Comet.Contract.IsAllowed(&_Comet.CallOpts, arg0, arg1)
}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_Comet *CometCaller) IsBorrowCollateralized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isBorrowCollateralized", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_Comet *CometSession) IsBorrowCollateralized(account common.Address) (bool, error) {
	return _Comet.Contract.IsBorrowCollateralized(&_Comet.CallOpts, account)
}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_Comet *CometCallerSession) IsBorrowCollateralized(account common.Address) (bool, error) {
	return _Comet.Contract.IsBorrowCollateralized(&_Comet.CallOpts, account)
}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_Comet *CometCaller) IsBuyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isBuyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_Comet *CometSession) IsBuyPaused() (bool, error) {
	return _Comet.Contract.IsBuyPaused(&_Comet.CallOpts)
}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_Comet *CometCallerSession) IsBuyPaused() (bool, error) {
	return _Comet.Contract.IsBuyPaused(&_Comet.CallOpts)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Comet *CometCaller) IsLiquidatable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isLiquidatable", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Comet *CometSession) IsLiquidatable(account common.Address) (bool, error) {
	return _Comet.Contract.IsLiquidatable(&_Comet.CallOpts, account)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_Comet *CometCallerSession) IsLiquidatable(account common.Address) (bool, error) {
	return _Comet.Contract.IsLiquidatable(&_Comet.CallOpts, account)
}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_Comet *CometCaller) IsSupplyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isSupplyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_Comet *CometSession) IsSupplyPaused() (bool, error) {
	return _Comet.Contract.IsSupplyPaused(&_Comet.CallOpts)
}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_Comet *CometCallerSession) IsSupplyPaused() (bool, error) {
	return _Comet.Contract.IsSupplyPaused(&_Comet.CallOpts)
}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_Comet *CometCaller) IsTransferPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isTransferPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_Comet *CometSession) IsTransferPaused() (bool, error) {
	return _Comet.Contract.IsTransferPaused(&_Comet.CallOpts)
}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_Comet *CometCallerSession) IsTransferPaused() (bool, error) {
	return _Comet.Contract.IsTransferPaused(&_Comet.CallOpts)
}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_Comet *CometCaller) IsWithdrawPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "isWithdrawPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_Comet *CometSession) IsWithdrawPaused() (bool, error) {
	return _Comet.Contract.IsWithdrawPaused(&_Comet.CallOpts)
}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_Comet *CometCallerSession) IsWithdrawPaused() (bool, error) {
	return _Comet.Contract.IsWithdrawPaused(&_Comet.CallOpts)
}

// LiquidatorPoints is a free data retrieval call binding the contract method 0xc5fa15cf.
//
// Solidity: function liquidatorPoints(address ) view returns(uint32 numAbsorbs, uint64 numAbsorbed, uint128 approxSpend, uint32 _reserved)
func (_Comet *CometCaller) LiquidatorPoints(opts *bind.CallOpts, arg0 common.Address) (struct {
	NumAbsorbs  uint32
	NumAbsorbed uint64
	ApproxSpend *big.Int
	Reserved    uint32
}, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "liquidatorPoints", arg0)

	outstruct := new(struct {
		NumAbsorbs  uint32
		NumAbsorbed uint64
		ApproxSpend *big.Int
		Reserved    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NumAbsorbs = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.NumAbsorbed = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ApproxSpend = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Reserved = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// LiquidatorPoints is a free data retrieval call binding the contract method 0xc5fa15cf.
//
// Solidity: function liquidatorPoints(address ) view returns(uint32 numAbsorbs, uint64 numAbsorbed, uint128 approxSpend, uint32 _reserved)
func (_Comet *CometSession) LiquidatorPoints(arg0 common.Address) (struct {
	NumAbsorbs  uint32
	NumAbsorbed uint64
	ApproxSpend *big.Int
	Reserved    uint32
}, error) {
	return _Comet.Contract.LiquidatorPoints(&_Comet.CallOpts, arg0)
}

// LiquidatorPoints is a free data retrieval call binding the contract method 0xc5fa15cf.
//
// Solidity: function liquidatorPoints(address ) view returns(uint32 numAbsorbs, uint64 numAbsorbed, uint128 approxSpend, uint32 _reserved)
func (_Comet *CometCallerSession) LiquidatorPoints(arg0 common.Address) (struct {
	NumAbsorbs  uint32
	NumAbsorbed uint64
	ApproxSpend *big.Int
	Reserved    uint32
}, error) {
	return _Comet.Contract.LiquidatorPoints(&_Comet.CallOpts, arg0)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_Comet *CometCaller) NumAssets(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_Comet *CometSession) NumAssets() (uint8, error) {
	return _Comet.Contract.NumAssets(&_Comet.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_Comet *CometCallerSession) NumAssets() (uint8, error) {
	return _Comet.Contract.NumAssets(&_Comet.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comet *CometCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comet *CometSession) PauseGuardian() (common.Address, error) {
	return _Comet.Contract.PauseGuardian(&_Comet.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comet *CometCallerSession) PauseGuardian() (common.Address, error) {
	return _Comet.Contract.PauseGuardian(&_Comet.CallOpts)
}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_Comet *CometCaller) QuoteCollateral(opts *bind.CallOpts, asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "quoteCollateral", asset, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_Comet *CometSession) QuoteCollateral(asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _Comet.Contract.QuoteCollateral(&_Comet.CallOpts, asset, baseAmount)
}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_Comet *CometCallerSession) QuoteCollateral(asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _Comet.Contract.QuoteCollateral(&_Comet.CallOpts, asset, baseAmount)
}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_Comet *CometCaller) StoreFrontPriceFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "storeFrontPriceFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_Comet *CometSession) StoreFrontPriceFactor() (*big.Int, error) {
	return _Comet.Contract.StoreFrontPriceFactor(&_Comet.CallOpts)
}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_Comet *CometCallerSession) StoreFrontPriceFactor() (*big.Int, error) {
	return _Comet.Contract.StoreFrontPriceFactor(&_Comet.CallOpts)
}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_Comet *CometCaller) SupplyKink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "supplyKink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_Comet *CometSession) SupplyKink() (*big.Int, error) {
	return _Comet.Contract.SupplyKink(&_Comet.CallOpts)
}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_Comet *CometCallerSession) SupplyKink() (*big.Int, error) {
	return _Comet.Contract.SupplyKink(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometCaller) SupplyPerSecondInterestRateBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "supplyPerSecondInterestRateBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometSession) SupplyPerSecondInterestRateBase() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateBase(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_Comet *CometCallerSession) SupplyPerSecondInterestRateBase() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateBase(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometCaller) SupplyPerSecondInterestRateSlopeHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "supplyPerSecondInterestRateSlopeHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometSession) SupplyPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateSlopeHigh(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_Comet *CometCallerSession) SupplyPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateSlopeHigh(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometCaller) SupplyPerSecondInterestRateSlopeLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "supplyPerSecondInterestRateSlopeLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometSession) SupplyPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateSlopeLow(&_Comet.CallOpts)
}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_Comet *CometCallerSession) SupplyPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _Comet.Contract.SupplyPerSecondInterestRateSlopeLow(&_Comet.CallOpts)
}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_Comet *CometCaller) TargetReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "targetReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_Comet *CometSession) TargetReserves() (*big.Int, error) {
	return _Comet.Contract.TargetReserves(&_Comet.CallOpts)
}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_Comet *CometCallerSession) TargetReserves() (*big.Int, error) {
	return _Comet.Contract.TargetReserves(&_Comet.CallOpts)
}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_Comet *CometCaller) TotalBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "totalBorrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_Comet *CometSession) TotalBorrow() (*big.Int, error) {
	return _Comet.Contract.TotalBorrow(&_Comet.CallOpts)
}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_Comet *CometCallerSession) TotalBorrow() (*big.Int, error) {
	return _Comet.Contract.TotalBorrow(&_Comet.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Comet *CometCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Comet *CometSession) TotalSupply() (*big.Int, error) {
	return _Comet.Contract.TotalSupply(&_Comet.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Comet *CometCallerSession) TotalSupply() (*big.Int, error) {
	return _Comet.Contract.TotalSupply(&_Comet.CallOpts)
}

// TotalsCollateral is a free data retrieval call binding the contract method 0x59e017bd.
//
// Solidity: function totalsCollateral(address ) view returns(uint128 totalSupplyAsset, uint128 _reserved)
func (_Comet *CometCaller) TotalsCollateral(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalSupplyAsset *big.Int
	Reserved         *big.Int
}, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "totalsCollateral", arg0)

	outstruct := new(struct {
		TotalSupplyAsset *big.Int
		Reserved         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSupplyAsset = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserved = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalsCollateral is a free data retrieval call binding the contract method 0x59e017bd.
//
// Solidity: function totalsCollateral(address ) view returns(uint128 totalSupplyAsset, uint128 _reserved)
func (_Comet *CometSession) TotalsCollateral(arg0 common.Address) (struct {
	TotalSupplyAsset *big.Int
	Reserved         *big.Int
}, error) {
	return _Comet.Contract.TotalsCollateral(&_Comet.CallOpts, arg0)
}

// TotalsCollateral is a free data retrieval call binding the contract method 0x59e017bd.
//
// Solidity: function totalsCollateral(address ) view returns(uint128 totalSupplyAsset, uint128 _reserved)
func (_Comet *CometCallerSession) TotalsCollateral(arg0 common.Address) (struct {
	TotalSupplyAsset *big.Int
	Reserved         *big.Int
}, error) {
	return _Comet.Contract.TotalsCollateral(&_Comet.CallOpts, arg0)
}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_Comet *CometCaller) TrackingIndexScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "trackingIndexScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_Comet *CometSession) TrackingIndexScale() (*big.Int, error) {
	return _Comet.Contract.TrackingIndexScale(&_Comet.CallOpts)
}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_Comet *CometCallerSession) TrackingIndexScale() (*big.Int, error) {
	return _Comet.Contract.TrackingIndexScale(&_Comet.CallOpts)
}

// UserBasic is a free data retrieval call binding the contract method 0xdc4abafd.
//
// Solidity: function userBasic(address ) view returns(int104 principal, uint64 baseTrackingIndex, uint64 baseTrackingAccrued, uint16 assetsIn, uint8 _reserved)
func (_Comet *CometCaller) UserBasic(opts *bind.CallOpts, arg0 common.Address) (struct {
	Principal           *big.Int
	BaseTrackingIndex   uint64
	BaseTrackingAccrued uint64
	AssetsIn            uint16
	Reserved            uint8
}, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "userBasic", arg0)

	outstruct := new(struct {
		Principal           *big.Int
		BaseTrackingIndex   uint64
		BaseTrackingAccrued uint64
		AssetsIn            uint16
		Reserved            uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Principal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseTrackingIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.BaseTrackingAccrued = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.AssetsIn = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Reserved = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserBasic is a free data retrieval call binding the contract method 0xdc4abafd.
//
// Solidity: function userBasic(address ) view returns(int104 principal, uint64 baseTrackingIndex, uint64 baseTrackingAccrued, uint16 assetsIn, uint8 _reserved)
func (_Comet *CometSession) UserBasic(arg0 common.Address) (struct {
	Principal           *big.Int
	BaseTrackingIndex   uint64
	BaseTrackingAccrued uint64
	AssetsIn            uint16
	Reserved            uint8
}, error) {
	return _Comet.Contract.UserBasic(&_Comet.CallOpts, arg0)
}

// UserBasic is a free data retrieval call binding the contract method 0xdc4abafd.
//
// Solidity: function userBasic(address ) view returns(int104 principal, uint64 baseTrackingIndex, uint64 baseTrackingAccrued, uint16 assetsIn, uint8 _reserved)
func (_Comet *CometCallerSession) UserBasic(arg0 common.Address) (struct {
	Principal           *big.Int
	BaseTrackingIndex   uint64
	BaseTrackingAccrued uint64
	AssetsIn            uint16
	Reserved            uint8
}, error) {
	return _Comet.Contract.UserBasic(&_Comet.CallOpts, arg0)
}

// UserCollateral is a free data retrieval call binding the contract method 0x2b92a07d.
//
// Solidity: function userCollateral(address , address ) view returns(uint128 balance, uint128 _reserved)
func (_Comet *CometCaller) UserCollateral(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Balance  *big.Int
	Reserved *big.Int
}, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "userCollateral", arg0, arg1)

	outstruct := new(struct {
		Balance  *big.Int
		Reserved *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserved = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserCollateral is a free data retrieval call binding the contract method 0x2b92a07d.
//
// Solidity: function userCollateral(address , address ) view returns(uint128 balance, uint128 _reserved)
func (_Comet *CometSession) UserCollateral(arg0 common.Address, arg1 common.Address) (struct {
	Balance  *big.Int
	Reserved *big.Int
}, error) {
	return _Comet.Contract.UserCollateral(&_Comet.CallOpts, arg0, arg1)
}

// UserCollateral is a free data retrieval call binding the contract method 0x2b92a07d.
//
// Solidity: function userCollateral(address , address ) view returns(uint128 balance, uint128 _reserved)
func (_Comet *CometCallerSession) UserCollateral(arg0 common.Address, arg1 common.Address) (struct {
	Balance  *big.Int
	Reserved *big.Int
}, error) {
	return _Comet.Contract.UserCollateral(&_Comet.CallOpts, arg0, arg1)
}

// UserNonce is a free data retrieval call binding the contract method 0x2e04b8e7.
//
// Solidity: function userNonce(address ) view returns(uint256)
func (_Comet *CometCaller) UserNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comet.contract.Call(opts, &out, "userNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNonce is a free data retrieval call binding the contract method 0x2e04b8e7.
//
// Solidity: function userNonce(address ) view returns(uint256)
func (_Comet *CometSession) UserNonce(arg0 common.Address) (*big.Int, error) {
	return _Comet.Contract.UserNonce(&_Comet.CallOpts, arg0)
}

// UserNonce is a free data retrieval call binding the contract method 0x2e04b8e7.
//
// Solidity: function userNonce(address ) view returns(uint256)
func (_Comet *CometCallerSession) UserNonce(arg0 common.Address) (*big.Int, error) {
	return _Comet.Contract.UserNonce(&_Comet.CallOpts, arg0)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_Comet *CometTransactor) Absorb(opts *bind.TransactOpts, absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "absorb", absorber, accounts)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_Comet *CometSession) Absorb(absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _Comet.Contract.Absorb(&_Comet.TransactOpts, absorber, accounts)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_Comet *CometTransactorSession) Absorb(absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _Comet.Contract.Absorb(&_Comet.TransactOpts, absorber, accounts)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_Comet *CometTransactor) AccrueAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "accrueAccount", account)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_Comet *CometSession) AccrueAccount(account common.Address) (*types.Transaction, error) {
	return _Comet.Contract.AccrueAccount(&_Comet.TransactOpts, account)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_Comet *CometTransactorSession) AccrueAccount(account common.Address) (*types.Transaction, error) {
	return _Comet.Contract.AccrueAccount(&_Comet.TransactOpts, account)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) ApproveThis(opts *bind.TransactOpts, manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "approveThis", manager, asset, amount)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_Comet *CometSession) ApproveThis(manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.ApproveThis(&_Comet.TransactOpts, manager, asset, amount)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) ApproveThis(manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.ApproveThis(&_Comet.TransactOpts, manager, asset, amount)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_Comet *CometTransactor) BuyCollateral(opts *bind.TransactOpts, asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "buyCollateral", asset, minAmount, baseAmount, recipient)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_Comet *CometSession) BuyCollateral(asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Comet.Contract.BuyCollateral(&_Comet.TransactOpts, asset, minAmount, baseAmount, recipient)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_Comet *CometTransactorSession) BuyCollateral(asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Comet.Contract.BuyCollateral(&_Comet.TransactOpts, asset, minAmount, baseAmount, recipient)
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_Comet *CometTransactor) InitializeStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "initializeStorage")
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_Comet *CometSession) InitializeStorage() (*types.Transaction, error) {
	return _Comet.Contract.InitializeStorage(&_Comet.TransactOpts)
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_Comet *CometTransactorSession) InitializeStorage() (*types.Transaction, error) {
	return _Comet.Contract.InitializeStorage(&_Comet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_Comet *CometTransactor) Pause(opts *bind.TransactOpts, supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "pause", supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_Comet *CometSession) Pause(supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _Comet.Contract.Pause(&_Comet.TransactOpts, supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_Comet *CometTransactorSession) Pause(supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _Comet.Contract.Pause(&_Comet.TransactOpts, supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_Comet *CometTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "supply", asset, amount)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_Comet *CometSession) Supply(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Supply(&_Comet.TransactOpts, asset, amount)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) Supply(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Supply(&_Comet.TransactOpts, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) SupplyFrom(opts *bind.TransactOpts, from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "supplyFrom", from, dst, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_Comet *CometSession) SupplyFrom(from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.SupplyFrom(&_Comet.TransactOpts, from, dst, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) SupplyFrom(from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.SupplyFrom(&_Comet.TransactOpts, from, dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) SupplyTo(opts *bind.TransactOpts, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "supplyTo", dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_Comet *CometSession) SupplyTo(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.SupplyTo(&_Comet.TransactOpts, dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) SupplyTo(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.SupplyTo(&_Comet.TransactOpts, dst, asset, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Comet *CometTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Comet *CometSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Transfer(&_Comet.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Comet *CometTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Transfer(&_Comet.TransactOpts, dst, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) TransferAsset(opts *bind.TransactOpts, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "transferAsset", dst, asset, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_Comet *CometSession) TransferAsset(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferAsset(&_Comet.TransactOpts, dst, asset, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) TransferAsset(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferAsset(&_Comet.TransactOpts, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) TransferAssetFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "transferAssetFrom", src, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_Comet *CometSession) TransferAssetFrom(src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferAssetFrom(&_Comet.TransactOpts, src, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) TransferAssetFrom(src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferAssetFrom(&_Comet.TransactOpts, src, dst, asset, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Comet *CometTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Comet *CometSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferFrom(&_Comet.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Comet *CometTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.TransferFrom(&_Comet.TransactOpts, src, dst, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_Comet *CometTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "withdraw", asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_Comet *CometSession) Withdraw(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Withdraw(&_Comet.TransactOpts, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) Withdraw(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.Withdraw(&_Comet.TransactOpts, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) WithdrawFrom(opts *bind.TransactOpts, src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "withdrawFrom", src, to, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_Comet *CometSession) WithdrawFrom(src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawFrom(&_Comet.TransactOpts, src, to, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) WithdrawFrom(src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawFrom(&_Comet.TransactOpts, src, to, asset, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_Comet *CometTransactor) WithdrawReserves(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "withdrawReserves", to, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_Comet *CometSession) WithdrawReserves(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawReserves(&_Comet.TransactOpts, to, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_Comet *CometTransactorSession) WithdrawReserves(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawReserves(&_Comet.TransactOpts, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_Comet *CometTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.contract.Transact(opts, "withdrawTo", to, asset, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_Comet *CometSession) WithdrawTo(to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawTo(&_Comet.TransactOpts, to, asset, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_Comet *CometTransactorSession) WithdrawTo(to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comet.Contract.WithdrawTo(&_Comet.TransactOpts, to, asset, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Comet *CometTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Comet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Comet *CometSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Comet.Contract.Fallback(&_Comet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Comet *CometTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Comet.Contract.Fallback(&_Comet.TransactOpts, calldata)
}

// CometAbsorbCollateralIterator is returned from FilterAbsorbCollateral and is used to iterate over the raw logs and unpacked data for AbsorbCollateral events raised by the Comet contract.
type CometAbsorbCollateralIterator struct {
	Event *CometAbsorbCollateral // Event containing the contract specifics and raw log

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
func (it *CometAbsorbCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometAbsorbCollateral)
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
		it.Event = new(CometAbsorbCollateral)
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
func (it *CometAbsorbCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometAbsorbCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometAbsorbCollateral represents a AbsorbCollateral event raised by the Comet contract.
type CometAbsorbCollateral struct {
	Absorber           common.Address
	Borrower           common.Address
	Asset              common.Address
	CollateralAbsorbed *big.Int
	UsdValue           *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAbsorbCollateral is a free log retrieval operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_Comet *CometFilterer) FilterAbsorbCollateral(opts *bind.FilterOpts, absorber []common.Address, borrower []common.Address, asset []common.Address) (*CometAbsorbCollateralIterator, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "AbsorbCollateral", absorberRule, borrowerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometAbsorbCollateralIterator{contract: _Comet.contract, event: "AbsorbCollateral", logs: logs, sub: sub}, nil
}

// WatchAbsorbCollateral is a free log subscription operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_Comet *CometFilterer) WatchAbsorbCollateral(opts *bind.WatchOpts, sink chan<- *CometAbsorbCollateral, absorber []common.Address, borrower []common.Address, asset []common.Address) (event.Subscription, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "AbsorbCollateral", absorberRule, borrowerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometAbsorbCollateral)
				if err := _Comet.contract.UnpackLog(event, "AbsorbCollateral", log); err != nil {
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

// ParseAbsorbCollateral is a log parse operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_Comet *CometFilterer) ParseAbsorbCollateral(log types.Log) (*CometAbsorbCollateral, error) {
	event := new(CometAbsorbCollateral)
	if err := _Comet.contract.UnpackLog(event, "AbsorbCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometAbsorbDebtIterator is returned from FilterAbsorbDebt and is used to iterate over the raw logs and unpacked data for AbsorbDebt events raised by the Comet contract.
type CometAbsorbDebtIterator struct {
	Event *CometAbsorbDebt // Event containing the contract specifics and raw log

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
func (it *CometAbsorbDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometAbsorbDebt)
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
		it.Event = new(CometAbsorbDebt)
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
func (it *CometAbsorbDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometAbsorbDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometAbsorbDebt represents a AbsorbDebt event raised by the Comet contract.
type CometAbsorbDebt struct {
	Absorber    common.Address
	Borrower    common.Address
	BasePaidOut *big.Int
	UsdValue    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAbsorbDebt is a free log retrieval operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_Comet *CometFilterer) FilterAbsorbDebt(opts *bind.FilterOpts, absorber []common.Address, borrower []common.Address) (*CometAbsorbDebtIterator, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "AbsorbDebt", absorberRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CometAbsorbDebtIterator{contract: _Comet.contract, event: "AbsorbDebt", logs: logs, sub: sub}, nil
}

// WatchAbsorbDebt is a free log subscription operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_Comet *CometFilterer) WatchAbsorbDebt(opts *bind.WatchOpts, sink chan<- *CometAbsorbDebt, absorber []common.Address, borrower []common.Address) (event.Subscription, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "AbsorbDebt", absorberRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometAbsorbDebt)
				if err := _Comet.contract.UnpackLog(event, "AbsorbDebt", log); err != nil {
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

// ParseAbsorbDebt is a log parse operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_Comet *CometFilterer) ParseAbsorbDebt(log types.Log) (*CometAbsorbDebt, error) {
	event := new(CometAbsorbDebt)
	if err := _Comet.contract.UnpackLog(event, "AbsorbDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometBuyCollateralIterator is returned from FilterBuyCollateral and is used to iterate over the raw logs and unpacked data for BuyCollateral events raised by the Comet contract.
type CometBuyCollateralIterator struct {
	Event *CometBuyCollateral // Event containing the contract specifics and raw log

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
func (it *CometBuyCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometBuyCollateral)
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
		it.Event = new(CometBuyCollateral)
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
func (it *CometBuyCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometBuyCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometBuyCollateral represents a BuyCollateral event raised by the Comet contract.
type CometBuyCollateral struct {
	Buyer            common.Address
	Asset            common.Address
	BaseAmount       *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBuyCollateral is a free log retrieval operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_Comet *CometFilterer) FilterBuyCollateral(opts *bind.FilterOpts, buyer []common.Address, asset []common.Address) (*CometBuyCollateralIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "BuyCollateral", buyerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometBuyCollateralIterator{contract: _Comet.contract, event: "BuyCollateral", logs: logs, sub: sub}, nil
}

// WatchBuyCollateral is a free log subscription operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_Comet *CometFilterer) WatchBuyCollateral(opts *bind.WatchOpts, sink chan<- *CometBuyCollateral, buyer []common.Address, asset []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "BuyCollateral", buyerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometBuyCollateral)
				if err := _Comet.contract.UnpackLog(event, "BuyCollateral", log); err != nil {
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

// ParseBuyCollateral is a log parse operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_Comet *CometFilterer) ParseBuyCollateral(log types.Log) (*CometBuyCollateral, error) {
	event := new(CometBuyCollateral)
	if err := _Comet.contract.UnpackLog(event, "BuyCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometPauseActionIterator is returned from FilterPauseAction and is used to iterate over the raw logs and unpacked data for PauseAction events raised by the Comet contract.
type CometPauseActionIterator struct {
	Event *CometPauseAction // Event containing the contract specifics and raw log

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
func (it *CometPauseActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometPauseAction)
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
		it.Event = new(CometPauseAction)
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
func (it *CometPauseActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometPauseActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometPauseAction represents a PauseAction event raised by the Comet contract.
type CometPauseAction struct {
	SupplyPaused   bool
	TransferPaused bool
	WithdrawPaused bool
	AbsorbPaused   bool
	BuyPaused      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauseAction is a free log retrieval operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_Comet *CometFilterer) FilterPauseAction(opts *bind.FilterOpts) (*CometPauseActionIterator, error) {

	logs, sub, err := _Comet.contract.FilterLogs(opts, "PauseAction")
	if err != nil {
		return nil, err
	}
	return &CometPauseActionIterator{contract: _Comet.contract, event: "PauseAction", logs: logs, sub: sub}, nil
}

// WatchPauseAction is a free log subscription operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_Comet *CometFilterer) WatchPauseAction(opts *bind.WatchOpts, sink chan<- *CometPauseAction) (event.Subscription, error) {

	logs, sub, err := _Comet.contract.WatchLogs(opts, "PauseAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometPauseAction)
				if err := _Comet.contract.UnpackLog(event, "PauseAction", log); err != nil {
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

// ParsePauseAction is a log parse operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_Comet *CometFilterer) ParsePauseAction(log types.Log) (*CometPauseAction, error) {
	event := new(CometPauseAction)
	if err := _Comet.contract.UnpackLog(event, "PauseAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Comet contract.
type CometSupplyIterator struct {
	Event *CometSupply // Event containing the contract specifics and raw log

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
func (it *CometSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometSupply)
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
		it.Event = new(CometSupply)
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
func (it *CometSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometSupply represents a Supply event raised by the Comet contract.
type CometSupply struct {
	From   common.Address
	Dst    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_Comet *CometFilterer) FilterSupply(opts *bind.FilterOpts, from []common.Address, dst []common.Address) (*CometSupplyIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "Supply", fromRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &CometSupplyIterator{contract: _Comet.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_Comet *CometFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *CometSupply, from []common.Address, dst []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "Supply", fromRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometSupply)
				if err := _Comet.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_Comet *CometFilterer) ParseSupply(log types.Log) (*CometSupply, error) {
	event := new(CometSupply)
	if err := _Comet.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometSupplyCollateralIterator is returned from FilterSupplyCollateral and is used to iterate over the raw logs and unpacked data for SupplyCollateral events raised by the Comet contract.
type CometSupplyCollateralIterator struct {
	Event *CometSupplyCollateral // Event containing the contract specifics and raw log

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
func (it *CometSupplyCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometSupplyCollateral)
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
		it.Event = new(CometSupplyCollateral)
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
func (it *CometSupplyCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometSupplyCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometSupplyCollateral represents a SupplyCollateral event raised by the Comet contract.
type CometSupplyCollateral struct {
	From   common.Address
	Dst    common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupplyCollateral is a free log retrieval operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) FilterSupplyCollateral(opts *bind.FilterOpts, from []common.Address, dst []common.Address, asset []common.Address) (*CometSupplyCollateralIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "SupplyCollateral", fromRule, dstRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometSupplyCollateralIterator{contract: _Comet.contract, event: "SupplyCollateral", logs: logs, sub: sub}, nil
}

// WatchSupplyCollateral is a free log subscription operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) WatchSupplyCollateral(opts *bind.WatchOpts, sink chan<- *CometSupplyCollateral, from []common.Address, dst []common.Address, asset []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "SupplyCollateral", fromRule, dstRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometSupplyCollateral)
				if err := _Comet.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
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

// ParseSupplyCollateral is a log parse operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) ParseSupplyCollateral(log types.Log) (*CometSupplyCollateral, error) {
	event := new(CometSupplyCollateral)
	if err := _Comet.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Comet contract.
type CometTransferIterator struct {
	Event *CometTransfer // Event containing the contract specifics and raw log

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
func (it *CometTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometTransfer)
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
		it.Event = new(CometTransfer)
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
func (it *CometTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometTransfer represents a Transfer event raised by the Comet contract.
type CometTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Comet *CometFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CometTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CometTransferIterator{contract: _Comet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Comet *CometFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CometTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometTransfer)
				if err := _Comet.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Comet *CometFilterer) ParseTransfer(log types.Log) (*CometTransfer, error) {
	event := new(CometTransfer)
	if err := _Comet.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometTransferCollateralIterator is returned from FilterTransferCollateral and is used to iterate over the raw logs and unpacked data for TransferCollateral events raised by the Comet contract.
type CometTransferCollateralIterator struct {
	Event *CometTransferCollateral // Event containing the contract specifics and raw log

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
func (it *CometTransferCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometTransferCollateral)
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
		it.Event = new(CometTransferCollateral)
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
func (it *CometTransferCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometTransferCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometTransferCollateral represents a TransferCollateral event raised by the Comet contract.
type CometTransferCollateral struct {
	From   common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferCollateral is a free log retrieval operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) FilterTransferCollateral(opts *bind.FilterOpts, from []common.Address, to []common.Address, asset []common.Address) (*CometTransferCollateralIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "TransferCollateral", fromRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometTransferCollateralIterator{contract: _Comet.contract, event: "TransferCollateral", logs: logs, sub: sub}, nil
}

// WatchTransferCollateral is a free log subscription operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) WatchTransferCollateral(opts *bind.WatchOpts, sink chan<- *CometTransferCollateral, from []common.Address, to []common.Address, asset []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "TransferCollateral", fromRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometTransferCollateral)
				if err := _Comet.contract.UnpackLog(event, "TransferCollateral", log); err != nil {
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

// ParseTransferCollateral is a log parse operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) ParseTransferCollateral(log types.Log) (*CometTransferCollateral, error) {
	event := new(CometTransferCollateral)
	if err := _Comet.contract.UnpackLog(event, "TransferCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Comet contract.
type CometWithdrawIterator struct {
	Event *CometWithdraw // Event containing the contract specifics and raw log

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
func (it *CometWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometWithdraw)
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
		it.Event = new(CometWithdraw)
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
func (it *CometWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometWithdraw represents a Withdraw event raised by the Comet contract.
type CometWithdraw struct {
	Src    common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_Comet *CometFilterer) FilterWithdraw(opts *bind.FilterOpts, src []common.Address, to []common.Address) (*CometWithdrawIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "Withdraw", srcRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CometWithdrawIterator{contract: _Comet.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_Comet *CometFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CometWithdraw, src []common.Address, to []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "Withdraw", srcRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometWithdraw)
				if err := _Comet.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_Comet *CometFilterer) ParseWithdraw(log types.Log) (*CometWithdraw, error) {
	event := new(CometWithdraw)
	if err := _Comet.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the Comet contract.
type CometWithdrawCollateralIterator struct {
	Event *CometWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *CometWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometWithdrawCollateral)
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
		it.Event = new(CometWithdrawCollateral)
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
func (it *CometWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometWithdrawCollateral represents a WithdrawCollateral event raised by the Comet contract.
type CometWithdrawCollateral struct {
	Src    common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts, src []common.Address, to []common.Address, asset []common.Address) (*CometWithdrawCollateralIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "WithdrawCollateral", srcRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometWithdrawCollateralIterator{contract: _Comet.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *CometWithdrawCollateral, src []common.Address, to []common.Address, asset []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "WithdrawCollateral", srcRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometWithdrawCollateral)
				if err := _Comet.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_Comet *CometFilterer) ParseWithdrawCollateral(log types.Log) (*CometWithdrawCollateral, error) {
	event := new(CometWithdrawCollateral)
	if err := _Comet.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometWithdrawReservesIterator is returned from FilterWithdrawReserves and is used to iterate over the raw logs and unpacked data for WithdrawReserves events raised by the Comet contract.
type CometWithdrawReservesIterator struct {
	Event *CometWithdrawReserves // Event containing the contract specifics and raw log

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
func (it *CometWithdrawReservesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometWithdrawReserves)
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
		it.Event = new(CometWithdrawReserves)
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
func (it *CometWithdrawReservesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometWithdrawReservesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometWithdrawReserves represents a WithdrawReserves event raised by the Comet contract.
type CometWithdrawReserves struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawReserves is a free log retrieval operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_Comet *CometFilterer) FilterWithdrawReserves(opts *bind.FilterOpts, to []common.Address) (*CometWithdrawReservesIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.FilterLogs(opts, "WithdrawReserves", toRule)
	if err != nil {
		return nil, err
	}
	return &CometWithdrawReservesIterator{contract: _Comet.contract, event: "WithdrawReserves", logs: logs, sub: sub}, nil
}

// WatchWithdrawReserves is a free log subscription operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_Comet *CometFilterer) WatchWithdrawReserves(opts *bind.WatchOpts, sink chan<- *CometWithdrawReserves, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Comet.contract.WatchLogs(opts, "WithdrawReserves", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometWithdrawReserves)
				if err := _Comet.contract.UnpackLog(event, "WithdrawReserves", log); err != nil {
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

// ParseWithdrawReserves is a log parse operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_Comet *CometFilterer) ParseWithdrawReserves(log types.Log) (*CometWithdrawReserves, error) {
	event := new(CometWithdrawReserves)
	if err := _Comet.contract.UnpackLog(event, "WithdrawReserves", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
