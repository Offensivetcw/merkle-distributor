package main

import (
	"encoding/json"
	"fmt"
	distributor "github.com/Offensivetcw/merkle-distributor"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	balances := []distributor.Balance{
		{Account: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"), Amount: big.NewInt(20)},
		{Account: common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"), Amount: big.NewInt(40)},
	}
	info, err := distributor.ParseBalanceMap(balances)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
