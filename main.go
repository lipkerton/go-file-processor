package main
import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func parseFile(input_file *os.File, output_file *os.File) {
	sc := bufio.NewScanner(input_file)
	wr := bufio.NewWriter(output_file)
	defer wr.Flush()
	for f, tk := true, sc.Scan(); tk; f, tk = true, sc.Scan() {
		for _, v := range strings.Fields(sc.Text()) {
			switch v {
			case "error":
				f = false
				break
			}
		}
		if f {
			_, err := wr.Write(sc.Bytes())
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {
	input_file, err := os.OpenFile("input.txt", os.O_RDONLY, 0111)
	output_file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0644)
	defer input_file.Close()
	defer output_file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	parseFile(input_file, output_file)
}