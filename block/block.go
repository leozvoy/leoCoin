package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"mygo/go_project/2.ZqChain/transaction"
)

type Block struct {
	PreHash      string                    //挖矿前读取上一个区块Hash值产生
	ThisHash     string                    //挖矿（SHA256()）产生
	Verify       int                       //挖矿过程中计算得
	Transactions []transaction.Transaction //从transactionPool中获取
	Difficult    int                       //读取所在链的Difficult产生
	TimeStamp    time.Time                 //挖矿过程中戳上时间戳
}

//计算Hash值
func (this *Block) ComputeHash() {
	hash := sha256.New()
	transTemp, _ := json.Marshal(this.Transactions)
	hash.Write([]byte(this.PreHash + strconv.Itoa(this.Verify) + this.TimeStamp.String() + string(transTemp)))
	this.ThisHash = hex.EncodeToString(hash.Sum(nil))
}

//获取当前挖矿难度
func (this *Block) GetDiff() string {
	result := ""
	for i := 0; i < this.Difficult; i++ {
		result += "0"
	}
	return result
}

//挖矿（生成新的区块）
func (this *Block) Mine() {

	for {
		this.TimeStamp = time.Now()
		//this.Transactions.Timestamp = this.TimeStamp
		this.ComputeHash()
		if strings.HasPrefix(this.ThisHash, this.GetDiff()) {
			//fmt.Printf("mine successed! hash:0x%v\n", result)
			break
		}
		this.Verify++
	}

}
