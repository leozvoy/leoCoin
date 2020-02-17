package chain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"mygo/go_project/2.ZqChain/block"
	"mygo/go_project/2.ZqChain/transaction"
)

type BlockChain struct {
	Chain           []block.Block
	Difficult       int
	MinerReward     int
	TransactionPool []transaction.Transaction
}

//产生创世区块
func (this *BlockChain) BigBang() {
	this.Difficult = 4
	this.MinerReward = 50
	theFirst := block.Block{
		TimeStamp: time.Now(),
		Difficult: 64,
		Verify:    int(0x7FFFFFFFFFFFFFFF),
	}
	firstTrans := transaction.Transaction{
		From:   "",
		To:     "Leo",
		Amount: uint64(this.MinerReward),
	}
	theFirst.Transactions = append(theFirst.Transactions, firstTrans)

	theFirst.ComputeHash()
	this.Chain = make([]block.Block, 1)
	this.Chain[0] = theFirst

	this.TransactionPool = make([]transaction.Transaction, 0, 0)
	this.TransactionPool = append(this.TransactionPool, theFirst.Transactions...)
}

//显示所有区块
func (this *BlockChain) ShowChain() {
	for k, v := range this.Chain {
		fmt.Printf("---%d---\n【\n", k)
		fmt.Println("PreHash\t", v.PreHash)
		fmt.Println("ThisHash", v.ThisHash)
		fmt.Println("Verify\t", v.Verify)
		fmt.Println("Transactions\t", v.Transactions)
		fmt.Println("Difficult\t", v.Difficult)
		fmt.Println("TimeStamp\t", v.TimeStamp)
		fmt.Println("】\n")
	}

	fmt.Println("---TransactionPool---")
	for _, v := range this.TransactionPool {
		v.Show()
	}
	fmt.Println("---END---")
	fmt.Println()
}

//获取上一个区块
func (this *BlockChain) GetLastBlock() *block.Block {
	return &this.Chain[len(this.Chain)-1]
}

//将转账记录添加到TransactionPool中
func (this *BlockChain) AddTransaction(trans transaction.Transaction) {
	this.TransactionPool = append(this.TransactionPool, trans)
}

//矿工挖矿（从Transaction池中打包交易挖矿，并获取矿工奖励）
func (this *BlockChain) MinerJob(miner string) {
	minerRewward := transaction.Transaction{
		From:   "",
		To:     miner,
		Amount: uint64(this.MinerReward),
	}
	this.AddTransaction(minerRewward)

	newBlock := block.Block{}
	newBlock.Transactions = append(newBlock.Transactions, this.TransactionPool...)

	newBlock.PreHash = this.GetLastBlock().ThisHash
	newBlock.Difficult = this.Difficult

	newBlock.Mine()

	this.Chain = append(this.Chain, newBlock)
	this.TransactionPool = this.TransactionPool[:0]
}

//验证当前链上的所有区块
func (this *BlockChain) Validate() error {
	for k, v := range this.Chain {
		if k == 0 {
			continue
		}
		//非法连接的链
		if v.PreHash != this.Chain[k-1].ThisHash {
			err := errors.New("This is a broken chain!!!")
			return err
		}
		//Hash难度不符合
		if !strings.HasPrefix(v.ThisHash, v.GetDiff()) {
			err := errors.New("Difficult error in this chain")
			return err
		}
		//数据篡改
		testBlock := v
		testBlock.ComputeHash()

		if v.ThisHash != testBlock.ThisHash {
			err := errors.New("Data has been altered!!!")
			return err
		}
	}
	err := errors.New("No error in this chain")
	return err
}

//向区块链中添加区块（从外部添加区块，不再使用）
func (this *BlockChain) AddBlock(newBlock block.Block) {
	newBlock.PreHash = this.GetLastBlock().ThisHash
	newBlock.Difficult = this.Difficult
	//newBlock.ComputeHash()
	newBlock.Mine()

	this.Chain = append(this.Chain, newBlock)
}
