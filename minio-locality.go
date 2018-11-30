package main

import (
	_ "fmt"
	"log"
	"io/ioutil"
)

func test10MB() {
	buffer := make([]byte, 10485760)
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 0xff
	}

	data := 12

	for i := 0; i < data; i++ {
		partSize := len(buffer) / data + 1

		for j := i*partSize; j < (i+1)*partSize; j++ {
			if j < len(buffer) {
				buffer[j] = byte(i + 65)
			}
		}
	}

	err := ioutil.WriteFile("test10mb.obj", buffer, 0644)
	if err != nil {
		log.Fatalln("Error writing file", err)
	}
}

func fillBuffer(chr, multiplier int) (buffer []byte) {
	buffer = make([]byte, 10485760*multiplier)
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 0xff
	}

	data := 12

	for i := 0; i < data; i++ {
		partSize := len(buffer) / data //+ 1

		for j := i*partSize; j < (i+1)*partSize; j++ {
			if j < len(buffer) {
				buffer[j] = byte(i + chr)
			}
		}
	}
	return
}

func test20MB() {

	buffer := make([]byte, 10485760*2)

	buf1 := fillBuffer(65, 1)
	buf2 := fillBuffer(65+12, 1)

	copy(buffer, buf1)
	copy(buffer[len(buf1):], buf2)

	err := ioutil.WriteFile("test20mb.obj", buffer, 0644)
	if err != nil {
		log.Fatalln("Error writing file", err)
	}
}

func test240MB() {

	buffer := make([]byte, 10485760*6*4)

	buf1 := fillBuffer(65, 6*4)

	copy(buffer, buf1)

	err := ioutil.WriteFile("test240mb.obj", buffer, 0644)
	if err != nil {
		log.Fatalln("Error writing file", err)
	}
}


func main() {
	test240MB()
}

//$ hexdump -C test240mb.obj
//00000000  41 41 41 41 41 41 41 41  41 41 41 41 41 41 41 41  |AAAAAAAAAAAAAAAA|
//*
//01400000  42 42 42 42 42 42 42 42  42 42 42 42 42 42 42 42  |BBBBBBBBBBBBBBBB|
//*
//02800000  43 43 43 43 43 43 43 43  43 43 43 43 43 43 43 43  |CCCCCCCCCCCCCCCC|
//*
//03c00000  44 44 44 44 44 44 44 44  44 44 44 44 44 44 44 44  |DDDDDDDDDDDDDDDD|
//*
//05000000  45 45 45 45 45 45 45 45  45 45 45 45 45 45 45 45  |EEEEEEEEEEEEEEEE|
//*
//06400000  46 46 46 46 46 46 46 46  46 46 46 46 46 46 46 46  |FFFFFFFFFFFFFFFF|
//*
//07800000  47 47 47 47 47 47 47 47  47 47 47 47 47 47 47 47  |GGGGGGGGGGGGGGGG|
//*
//08c00000  48 48 48 48 48 48 48 48  48 48 48 48 48 48 48 48  |HHHHHHHHHHHHHHHH|
//*
//0a000000  49 49 49 49 49 49 49 49  49 49 49 49 49 49 49 49  |IIIIIIIIIIIIIIII|
//*
//0b400000  4a 4a 4a 4a 4a 4a 4a 4a  4a 4a 4a 4a 4a 4a 4a 4a  |JJJJJJJJJJJJJJJJ|
//*
//0c800000  4b 4b 4b 4b 4b 4b 4b 4b  4b 4b 4b 4b 4b 4b 4b 4b  |KKKKKKKKKKKKKKKK|
//*
//0dc00000  4c 4c 4c 4c 4c 4c 4c 4c  4c 4c 4c 4c 4c 4c 4c 4c  |LLLLLLLLLLLLLLLL|
//*
//0f000000
