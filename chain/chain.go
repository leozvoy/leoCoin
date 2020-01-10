package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"mygo/go_project/2.ZqChain/block"
)

type BlockChain struct {
	Chain []block.Block
}

func (this *BlockChain) BigBang() {
	theFirst := block.Block{
		Data:      "this is the first block of ZqChain",
		Difficult: 0,
	}
	theFirst.ComputeHash()
	this.Chain = make([]block.Block, 1)
	this.Chain[0] = theFirst
}

func (this *BlockChain) ShowChain() {
	for k, v := range this.Chain {
		fmt.Printf("---%d---\n{\n", k)
		fmt.Println("preHash\t", v.PreHash)
		fmt.Println("thisHash", v.ThisHash)
		fmt.Println("data\t", v.Data)
		fmt.Println("Verify\t", v.Verify)
		fmt.Println("Diff\t", v.Difficult)
		fmt.Println("}\n")
	}
}

func (this *BlockChain) AddBlock(newBlock block.Block) {
	newBlock.PreHash = this.Chain[len(this.Chain)-1].ThisHash
	//newBlock.ComputeHash()
	newBlock.Mine()

	this.Chain = append(this.Chain, newBlock)
}

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
		hash := sha256.New()
		hash.Write([]byte(v.PreHash + v.Data + v.Verify))
		if v.ThisHash != hex.EncodeToString(hash.Sum(nil)) {
			err := errors.New("Data has been altered!!!")
			return err
		}
	}
	err := errors.New("No error in this chain")
	return err
}
