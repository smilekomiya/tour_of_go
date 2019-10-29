package main

import (
	"io"
	"os"
	"strings"
)

var alphabets = map[int]rune{
	0:  'a',
	1:  'b',
	2:  'c',
	3:  'd',
	4:  'e',
	5:  'f',
	6:  'g',
	7:  'h',
	8:  'i',
	9:  'j',
	10: 'k',
	11: 'l',
	12: 'm',
	13: 'n',
	14: 'o',
	15: 'p',
	16: 'q',
	17: 'r',
	18: 's',
	19: 't',
	20: 'u',
	21: 'v',
	22: 'w',
	23: 'x',
	24: 'y',
	25: 'z',
}

type rot13Reader struct {
	r io.Reader
}

func (MyReader rot13Reader) Read(b []byte) (int, error) {
	n, err := MyReader.r.Read(b)
	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	return n, err

	// for i, value := range b {
	// 	newAlphabet := rune(value)
	// 	for num, alphabet := range alphabets {
	// 		if alphabet == rune(value) {
	// 			newAlphabet = alphabets[(num+13)%25]
	// 		}
	// 	}
	// 	b[i] = byte(newAlphabet)
	// }
	// return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
