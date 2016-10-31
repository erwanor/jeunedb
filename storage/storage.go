package storage

import "bufio"
import "bytes"
import "fmt"
import "encoding/binary"

type Block struct {
	Key    []byte
	KeyL   uint32
	Value  []byte
	ValueL uint32
}

func NewBlock(k []byte, v []byte) *Block {
	lenK := uint32(len(k))
	lenV := uint32(len(v))
	b := &Block{
		Key:    k,
		KeyL:   lenK,
		Value:  v,
		ValueL: lenV,
	}
	return b
}

func (b *Block) Serialize() *bytes.Buffer {
	toBuffer := new(bytes.Buffer)
	toBuffer.Reset()
	binary.Write(toBuffer, binary.BigEndian, b.KeyL)
	binary.Write(toBuffer, binary.BigEndian, b.Key)
	binary.Write(toBuffer, binary.BigEndian, b.ValueL)
	binary.Write(toBuffer, binary.BigEndian, b.Value)
	return toBuffer
}

func fillBuffer(size uint32, r *bufio.Reader) []byte {
	buff := make([]byte, size)
	_, err := r.Read(buff)
	if err != nil {
		fmt.Println("Error: ", err) // To replace w/ proper error handling
	}
	return buff
}
