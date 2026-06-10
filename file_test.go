package main
import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"testing"
)

func readError(fn *os.File) bool {
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

func readEmpty(fn *os.File) bool {
	sc := bufio.NewScanner(fn)
	return sc.Text() == ""
}

func TestFileContent(t *testing.T) {
	input_file, err := os.CreateTemp("", "input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	output_file, err := os.CreateTemp("", "output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input_file.Close()
	defer output_file.Close()
	parseFile(input_file, output_file)
	if !readError(output_file) {
		t.Errorf("`output.txt` has `error` word in it!")
	}
}

func TestFileEmpty(t *testing.T) {
	input_file, err := os.CreateTemp("", "input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	output_file, err := os.CreateTemp("", "output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input_file.Close()
	defer output_file.Close()
	parseFile(input_file, output_file)
	if !readEmpty(output_file) {
		t.Errorf("`output.txt` is empty!")
	}
}