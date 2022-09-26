package first_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GenerateHaskell() {
	result := make([]string, 0, 9)
	for i := 1; i <= 9; i++ {
		line := ""
		for j := 1; j <= i; j++ {
			str := fmt.Sprintf("%d x %d = %d", i, j, i*j)
			if j != i {
				str = fmt.Sprintf("%v ", str)
			}
			line = fmt.Sprintf("%v%v ", line, str)
		}
		result = append(result, line+"\n")
		line = ""
	}
	writeToFile(result)
}
func writeToFile(data []string) {
	f, err := openFile()
	if err != nil {
		log.Fatal(err.Error())
	}

	w := bufio.NewWriter(f)
	for _, v := range data {
		n, err := w.WriteString(v)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(n)
	}
	w.Flush()
	f.Close()
}

func openFile() (*os.File, error) {
	filename := "ninenine.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
}
