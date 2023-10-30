package ascii

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func Display(a []string, t [][]int) string {
	var output bytes.Buffer
	for _, b := range t {
		if b[0] == -1 {
			output.WriteString("\n")
		} else {
			for i := 1; i < 9; i++ {
				for _, j := range b {
					output.WriteString(a[j+i])
				}
				output.WriteString("\n")
			}
		}
	}
	return output.String()
}

func Open_File(s string) (t []string) {
	file, err := os.Open("static/banners/" + s + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	scn := bufio.NewScanner(file)
	for scn.Scan() {
		t = append(t, scn.Text())
	}
	return
}
func IsPrintable(s string) bool {
	for _, v := range s {
		if (v < 32 || v > 127) && (v != 10 && v != 13) {
			return false
		}
	}
	return true
}

func Finder(str string) (nn [][]int) {
	t := strings.Split(str, "\r\n")
	start := 0
	if t[0] == "" && t[1] == "" {
		start = 1
	}

	var n []int
	for i := start; i < len(t); i++ {
		if t[i] != "" {
			for _, v := range t[i] {
				n = append(n, (int(v-32) * 9))
			}
		} else {
			n = append(n, -1)
		}
		nn = append(nn, n)
		n = nil
	}
	return
}
