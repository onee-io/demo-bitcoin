package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain() // 初始化区块链，创建第一个区块

	bc.AddBlock("Send 1 BTC to onee")      // 加入一个区块
	bc.AddBlock("Send 2 more BTC to onee") // 加入一个区块

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
