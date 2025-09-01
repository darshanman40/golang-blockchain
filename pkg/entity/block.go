package entity

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	// hash value of data
	Hash []byte

	// data
	Data []byte

	// hash of previous data
	PrevHash []byte

	Nonce int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

// CreateBlock creates the block for data and using previous hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}

	// create Proof of work for block and run to get nonce and hash
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock adds the block for data on existing chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
