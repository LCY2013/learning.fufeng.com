package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nbeginning read file")
	printFileContents(file)
	fmt.Println("ended read file\n")
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

//func main() {
//	fmt.Println("convertToBin results:")
//	fmt.Println(
//		convertToBin(5),  // 101
//		convertToBin(13), // 1101
//		convertToBin(72387885),
//		convertToBin(0),
//	)
//
//	fmt.Println("abc.txt contents:")
//	printFile("/Users/magicLuoMacBook/software/go/projects/reptile/learning.fufeng.com/project/reptile/basic/branch/abc.txt")
//
//	fmt.Println("printing a string:")
//	s := `abc"d"
//	kkkk
//	123
//
//	p`
//	printFileContents(strings.NewReader(s))
//
//	// Uncomment to see it runs forever
//	// forever()
//}
