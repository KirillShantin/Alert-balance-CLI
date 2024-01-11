package helper

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/go-toast/toast"
	"github.com/spf13/cobra"
)

func Notify(addr, balance string) toast.Notification {
	msg := fmt.Sprintf("Top up your wallet by clicking the button at the bottom.\nWallet: %s Balance: %s", addr, balance)
	return toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "Top up wallet!",
		Message: msg,
		Icon:    "C:\\Users\\ShantinKV\\Pictures\\cat.png",
		Actions: []toast.Action{
			{"protocol", "Polyscan", `https://mumbai.polygonscan.com/address/0xfc1c48eA080398b55F69C87a6e6386E1d77229Fc`},
		},
	}
}

type Opt func(cmd *cobra.Command, args []string)

func Run() Opt {
	return func(cmd *cobra.Command, args []string) {
		for {
			timer1 := time.NewTimer(10 * time.Second)
			<-timer1.C
			balance := GetBalance(args[0])
			if balance <= big.NewFloat(0.958559).String() {
				fmt.Printf("Account: %s ", args[0])
				fmt.Printf("Balance: %s eth\n", balance)
				noti := Notify(args[0], balance)
				err := noti.Push()
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}

// func Run(cmd *cobra.Command, args []string) {
// 	for {
// 		timer1 := time.NewTimer(10 * time.Second)
// 		<-timer1.C
// 		balance := helper.GetBalance(args[0])
// 		if balance <= big.NewFloat(0.958559).String() {
// 			fmt.Printf("Account: %s ", args[0])
// 			fmt.Printf("Balance: %s eth\n", balance)
// 			noti := Notify(args[0], balance)
// 			err := noti.Push()
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 		}
// 	}
// }
