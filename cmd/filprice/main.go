package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/deixis/money"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	client *rpc.Client
	ETH    *ethclient.Client

	ten  = money.MustParseDecimal("10")
	nine = money.MustParseDecimal("9")

	endpoint = "http://localhost:8545"
)

var (
	SoloMarginAddress  = common.HexToAddress("0x1E0447b19BB6EcFdAe1e4AE1694b0C3659614e4e")
	PriceOracleAddress = common.HexToAddress("0x4250A6D3BD57455d7C6821eECb6206F507576cD2")

	DaiAddress  = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	HfilAddress = common.HexToAddress("0x9AFb950948c2370975fb91a441F36FDC02737cD4")
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "http://localhost:8545", "")

	flag.Parse()

	daiusd := getDAIUSDPrice()
	ethusd := getETHUSDPrice()
	fileth := getFILETHPrice()

	fmt.Println()

	fmt.Println("DAI / USD: $", daiusd.Round(2))
	fmt.Println("ETH / USD: $", ethusd.Round(2))
	fmt.Println("FIL / ETH: ", fileth.Round(4))
	fmt.Println("FIL / USD: $", fileth.Mul(ethusd).Round(4))
}

func getDAIUSDPrice() money.Decimal {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	ethClient := ethclient.NewClient(client)

	sm, err := NewSoloMargin(SoloMarginAddress, ethClient)
	if err != nil {
		panic(err)
	}

	res, _ := sm.GetMarketPrice(nil, big.NewInt(3))

	b := money.MustParseDecimal(res.Value.String())
	b = b.Div(ten.Pow(nine))
	b = b.Div(ten.Pow(nine))

	return b
}

func getETHUSDPrice() money.Decimal {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	ethClient := ethclient.NewClient(client)

	sm, err := NewSoloMargin(SoloMarginAddress, ethClient)
	if err != nil {
		panic(err)
	}

	res, _ := sm.GetMarketPrice(nil, big.NewInt(0))

	b := money.MustParseDecimal(res.Value.String())
	b = b.Div(ten.Pow(nine))
	b = b.Div(ten.Pow(nine))

	return b
}

func getFILETHPrice() money.Decimal {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	ethClient := ethclient.NewClient(client)

	po, err := NewPriceOracle(PriceOracleAddress, ethClient)
	if err != nil {
		panic(err)
	}

	res, _ := po.GetPrice(nil, HfilAddress)

	b := money.MustParseDecimal(res.String())
	b = b.Div(ten.Pow(nine))
	b = b.Div(ten.Pow(nine))

	return b
}
