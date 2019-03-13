package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the diagonalDifference function below.
func diagonalDifference(ar [][]int32) int32 {
	pSum := getPrincipalDiagonalSum(ar)
	tSum := getOppositeDiagonalSum(ar)
	return int32(math.Abs(float64(pSum - tSum)))
}

func getPrincipalDiagonalSum(ar [][]int32) int32 {
	var pSum int32
	for i := range ar {
		for j := range ar[i] {
			if i == j {
				fmt.Println(ar[i][j])
				pSum += ar[i][j]
			}
		}
	}
	return pSum
}

func getOppositeDiagonalSum(ar [][]int32) int32 {
	var oSum int32
	s := len(ar) - 1
	for i := range ar {
		for j := range ar[i] {
			if i+j == s {
				oSum += ar[i][j]
			}
		}
	}
	return oSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
