package helper

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var rawURL string = "https://polygon-mumbai.infura.io/v3/b779a81806f04c9685f3986bf229d9b0"

func GetBalance(account string) string {
	cl, err := ethclient.DialContext(context.Background(), rawURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer cl.Close()
	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	wei, err := cl.BalanceAt(context.Background(), common.HexToAddress(account), block.Number())
	if err != nil {
		log.Fatalf("Error to get balance. \naccount: %s \nblock: %s \n", account, block.Number())
	}
	ether := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
	return ether.String()
}

// // func testingGetBalance() {
// // 	taskPtr := flag.String("block", "1234", "Block without")
// // 	flag.Parse()
// // 	account := "0xfc1c48eA080398b55F69C87a6e6386E1d77229Fc"
// // 	for {

// // 		if *taskPtr == "" {
// // 			flag.PrintDefaults()
// // 			os.Exit(1)
// // 		}

// // 		timer1 := time.NewTimer(3 * time.Second)
// // 		<-timer1.C
// // 		balance := GetBalance(account)
// // 		fmt.Printf("Account: %s ", account)
// // 		fmt.Printf("Balance: %s\n", balance)
// // 	}
// }
