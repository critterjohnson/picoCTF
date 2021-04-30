package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	bmpHeader := []byte{66, 77}
	for i := 0; i < 2; i++ {
		if data[i] != bmpHeader[i] {
			fmt.Println("Stopping: BMP does not have valid signature")
			os.Exit(1)
		}
	}
	fmt.Println("BMP header validated")

	// Header
	signature := string(data[0:2])
	fileSize := binary.LittleEndian.Uint32(data[2:6])
	reserved := binary.LittleEndian.Uint32(data[6:10])
	dataOffset := binary.LittleEndian.Uint32(data[10:14])
	fmt.Printf("signature: %s\nfileSize: %d\nreserved: %b\ndataOffset: %d\n", signature, fileSize, reserved, dataOffset)

	infoHeader := make(map[string][]byte)
	infoHeaderLabels := []string{
		"size",
		"width",
		"height",
		"planes",
		"bitsPerPixel",
		"compression",
		"imageSize",
		"xPixelsPerM",
		"yPixelsPerM",
		"colorsUsed",
		"importantColors",
	}
	infoHeaderSizes := []int{
		4,
		4,
		4,
		2,
		2,
		4,
		4,
		4,
		4,
		4,
		4,
	}
	i := 14
	for j, label := range infoHeaderLabels {
		size := infoHeaderSizes[j]
		var f func([]byte) uint
		if size == 4 {
			f = func(b []byte) uint { return uint(binary.LittleEndian.Uint32(b)) }
		} else {
			f = func(b []byte) uint { return uint(binary.LittleEndian.Uint16(b)) }
		}
		infoHeader[label] = data[i : i+size]
		fmt.Println(label+":", f(infoHeader[label]))
		i += size
	}

	//i += int(math.Pow(2, float64(binary.LittleEndian.Uint16(infoHeader["bitsPerPixel"]))))

	newDataOffset := make([]byte, 4)
	binary.LittleEndian.PutUint32(newDataOffset, uint32(i))
	for i := 10; i < 14; i++ {
		data[i] = newDataOffset[i-10]
	}

	newSize := make([]byte, 4)
	binary.LittleEndian.PutUint16(newSize, 40)
	for i := 14; i < 18; i++ {
		data[i] = newSize[i-14]
	}

	for s := 100; s < 600; s += 10 {
		newHeight := make([]byte, 4)
		binary.LittleEndian.PutUint32(newHeight, uint32(306+s))
		for i := 22; i < 26; i++ {
			data[i] = newHeight[i-22]
		}
		ioutil.WriteFile("out"+strconv.Itoa(s)+".bmp", data, 0777)
	}
}
