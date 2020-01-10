package main

import (
	"fmt"

	"mygo/go_project/2.ZqChain/block"
	"mygo/go_project/2.ZqChain/chain"
)

func init() {

}

func main() {
	zqChain := chain.BlockChain{}
	zqChain.BigBang()

	newBlock := block.Block{
		Data:      "转账0.76BTC",
		Difficult: 1,
	}
	zqChain.AddBlock(newBlock)

	newBlock = block.Block{
		Data:      "转账2ETH",
		Difficult: 2,
	}
	zqChain.AddBlock(newBlock)

	newBlock = block.Block{
		Data:      "转账300EOS",
		Difficult: 3,
	}
	zqChain.AddBlock(newBlock)

	zqChain.ShowChain()
	fmt.Println(zqChain.Validate())

	// //测试篡改中间节点的数据
	// zqChain.Chain[2].Data = "转账200Eth"
	// zqChain.Chain[2].ComputeHash()
	// //zqChain.Chain[2].Mine()
	// fmt.Println(zqChain.Validate())

	//测试篡改最后一个节点的数据
	zqChain.Chain[3].Data = "转账99999999EOS"
	zqChain.Chain[3].ComputeHash()
	fmt.Println(zqChain.Validate())
	// zqChain.ShowChain()

}
