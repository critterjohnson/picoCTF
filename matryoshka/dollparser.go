// This file parses a PNG block by block, saving any remaining data at the end (after the IEND block) to an output file for
// further forensics.
// First argument: PNG file to parse
// Second argument: output file name
package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	pngHeader := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	for i := 0; i < 8; i++ {
		if data[i] != pngHeader[i] {
			fmt.Println("Stopping: PNG does not have valid signature")
			os.Exit(1)
		}
	}
	fmt.Println("PNG header validated")

	//var chunk [][]byte
	i := 8
	for {
		chunkSize := binary.BigEndian.Uint32(data[i : i+4])
		fmt.Println("size:", chunkSize)
		i += 4

		var chunkType string
		for j := 0; j < 4; j++ {
			chunkType += string(rune(data[i+j]))
		}
		fmt.Println("type:", chunkType)
		i += 4

		chunkData := data[i : i+int(chunkSize)]
		fmt.Println("data:", chunkData)
		i += int(chunkSize) + 4 // note that this skips CRC

		fmt.Print("\n")

		if chunkType == "IEND" {
			break
		}
	}

	if i < len(data) {
		ioutil.WriteFile(os.Args[2], data[i:], 0777)
		fmt.Println("extra data written:", os.Args[2])
	}
}
