package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

type Block struct {
	PreHash   string
	ThisHash  string
	Data      string
	Verify    string
	Difficult int
}

func (this *Block) ComputeHash() {
	hash := sha256.New()
	hash.Write([]byte(this.PreHash + this.Data + this.Verify))
	this.ThisHash = hex.EncodeToString(hash.Sum(nil))
}

func (this *Block) GetDiff() string {
	result := ""
	for i := 0; i < this.Difficult; i++ {
		result += "0"
	}
	return result
}

func (this *Block) Mine() {
	var result string = ""
	var nonce int = 0
	for {
		hash := sha256.New()
		hash.Write([]byte(this.PreHash + this.Data + strconv.Itoa(nonce)))
		result = hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(result, this.GetDiff()) {
			//fmt.Printf("mine successed! hash:0x%v\n", result)
			break
		}
		nonce++
	}
	this.Verify = strconv.Itoa(nonce)
	this.ThisHash = result
}
