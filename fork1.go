package main

import (
  "crypto/sha256"
  "encoding/hex"
  "fmt"
  "time"
)

//The Block struct represents a block in the blockchain

type Block struct {
  Index int
  Timestamp string
  Data string
  PrevHash string
  Hash string
}

//CalculateHash calculates the SHA256 hash of a block

func calculateHash(b Block) string {
  record := string(b.Index) + b.Timestamp + b.Data + b.PrevHash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)

  return hex.EncodeToString(hashed)
}

//Agora criamos a blockchain ao armazenar os blocos em um slice e implementar a função de adicionar novos blocos.

var Blockchain []Block

//createBlock creates a new block and adds it to the blockchain

func createBlock(data string, prevHash string) Block {
  var newBlock Block
  newBlock.Index = len(Blockchain)
  newBlock.Timestamp = time.Now().string()
  newBlock.Data = data
  newBlock.PrevHash = prevHash
  newBlock.Hash = calculateHash(newBlock)
  return newBlock
}

//Now we need to validate the Blockchain
//isBlockValid checks if a block is valid by verifying its hash andindex

func isBlockValid(newBlock, prevBlock Block) bool {
  if newBlock.Index != prevBlock.Index+1 {
    return false
  }

  if newBlock.PrevHash != prevBlock.Hash {
    return false
  }

  if calculateHash(newBlock) != newBlock.Hash {
    return false
  }

  return true
}

//isChainValid checks the validity of the entire blockchain

func isChainValid() bool {
  for i := 1; i < len(Blockchain); i++ {
    if !isBlockValid(Blockchain[i], Blockchain[i-1]) {
      return false
    }
  }

  return true
}
