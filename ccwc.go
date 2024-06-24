package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var flag string = os.Args[1]
	switch flag {
	case "-c":
		var fileName string = os.Args[2]
		print(getFileSize(fileName), " ", os.Args[2])
	case "-l":
		var fileName string = os.Args[2]
		print(getNumberofLines(fileName), " ", os.Args[2])
	case "-w":
		var fileName string = os.Args[2]
		print(getNumberofWords(fileName), " ", os.Args[2])
	case "-m":
		var fileName string = os.Args[2]
		print(getNumberOfChars(fileName), " ", os.Args[2])
	default:
		var fileName string = os.Args[1]
		var fileSize int64 = getFileSize(fileName)
		var lineNumbers int = getNumberofLines(fileName)
		var wordNumbers int = getNumberofWords(fileName)
		print(strconv.Itoa(int(fileSize))+" "+strconv.Itoa(lineNumbers)+" "+strconv.Itoa(wordNumbers), " ", fileName)
	}
}
func getNumberOfChars(input_file string) int {

	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	charCount := 0
	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file:", r, err)

		}
		charCount++
	}

	// content, err := ioutil.ReadFile(input_file)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// }

	// charCount := len(content)
	return charCount
}
func getFileSize(input_file string) int64 {
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)

	}
	fileSize := fileInfo.Size()
	return fileSize
	// print(strconv.FormatInt(fileSize, 10) + " " + input_file)
}
func getNumberofLines(input_file string) int {
	dat, err := os.ReadFile(input_file)
	check(err)
	count := 0
	var r rune = '\n'
	s := string(dat)
	for _, c := range s {
		if c == r {
			count++
		}
	}
	// lines := strings.Split(string(dat), "\n")
	return count

}

func getNumberofWords(input_file string) int {
	dat, err := os.ReadFile(input_file)
	check(err)
	words := strings.Fields(string(dat))

	// words := strings.Split(string(dat), " ")
	return len(words)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
