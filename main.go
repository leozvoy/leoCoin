package main

import (
	"fmt"

	"mygo/go_project/2.ZqChain/chain"
	"mygo/go_project/2.ZqChain/transaction"
)

func init() {

}

func main() {
	zqChain := chain.BlockChain{}
	zqChain.BigBang()

	trans1 := transaction.Transaction{
		From:   "user1",
		To:     "user2",
		Amount: 11,
	}
	zqChain.AddTransaction(trans1)
	trans2 := transaction.Transaction{
		From:   "user2",
		To:     "user3",
		Amount: 22,
	}
	zqChain.AddTransaction(trans2)
	trans3 := transaction.Transaction{
		From:   "user3",
		To:     "user4",
		Amount: 33,
	}
	zqChain.AddTransaction(trans3)

	zqChain.ShowChain()
	zqChain.MinerJob("miner1")
	zqChain.ShowChain()

	fmt.Println(zqChain.Validate())

	// 测试篡改中间节点的数据
	zqChain.Chain[1].Transactions[0].Amount = 999999
	//zqChain.Chain[2].ComputeHash()
	//zqChain.Chain[2].Mine()
	fmt.Println(zqChain.Validate())

	//测试篡改最后一个节点的数据
	// zqChain.Chain[3].Transactions.Amount = 999999
	// zqChain.Chain[3].ComputeHash()
	// fmt.Println(zqChain.Validate())
	// zqChain.ShowChain()

}
