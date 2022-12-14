package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  // 区块创建时间戳
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 前一个区块的哈希值
	Hash          []byte // 区块自身的哈希值，用于校验区块数据有效
	Nonce         int    // 工作量证明
}

// 创建一个新区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// 设置哈希值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// 创建一个创世纪区块（区块链的第一个区块）
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
