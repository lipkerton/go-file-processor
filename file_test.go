package main
import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"testing"
)

func openResultFile(fn *os.File) bool {
	sc := bufio.NewScanner(fn)
	for tk := sc.Scan(); tk; tk = sc.Scan() {
		for _, v := range strings.Fields(sc.Text()) {
			switch v{
			case "error":
				return false
			}
		}
	}
	return true
}

func TestFile(t *testing.T) {
	input_file, err := os.OpenFile("input.txt", os.O_RDONLY, 0111)
	if err != nil {
		fmt.Println(err)
		return
	}
	output_file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input_file.Close()
	defer output_file.Close()
	parseFile(input_file, output_file)
	if !openResultFile(output_file) {
		t.Errorf("`output.txt` has `error` word in it!")
	}
}