package types

type Transaction struct {
	Hash []byte
	From []byte
	To []byte
	CallData []byte
}
