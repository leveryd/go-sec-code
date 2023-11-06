package research

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
	"unicode/utf8"
)

// 两字节的UTF-8编码，特殊字符\在末尾
func TestSpecialUTF82(t *testing.T) {
	testChar := '\''
	for i := 0; i < 256; i++ {
		input := []byte{byte(i), byte(testChar)}
		output := escape(string(input))

		count := 0
		for _, ch := range output {
			if ch == '\\' {
				count++
			}
		}
		if count == 0 {
			t.Errorf("%v -> %v", input, output)
		}
	}
}

// 三字节的UTF-8编码，特殊字符\在末尾
func TestSpecialUTF8(t *testing.T) {
	testChar := '\''
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			input := []byte{byte(i), byte(j), byte(testChar)}
			output := escape(string(input))

			count := 0
			for _, ch := range output {
				if ch == '\\' {
					count++
				}
			}
			if count == 0 {
				t.Errorf("%v -> %v", input, output)
			}
		}
	}
}

// 四字节的UTF-8编码，特殊字符\在末尾
func TestSpecialUTF8mb4(t *testing.T) {
	testChar := '\''
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for n := 0; j < 256; j++ {
				input := []byte{byte(i), byte(j), byte(n), byte(testChar)}
				output := escape(string(input))

				count := 0
				for _, ch := range output {
					if ch == '\\' {
						count++
					}
				}
				if count == 0 {
					t.Errorf("%v -> %v", input, output)
				}
			}
		}
	}
}

func build2(x, y byte, char byte) string {
	i := []byte{x, y, char} // 用户输入是byte数组

	var buf bytes.Buffer
	buf.WriteString(string(i)) // byte数组转string
	return buf.String()
}

func build1(x byte, char byte) string {
	i := []byte{x, char} // 用户输入是byte数组

	var buf bytes.Buffer
	buf.WriteString(string(i)) // byte数组转string
	return buf.String()
}

func escape(input string) string {
	var b strings.Builder

	for _, ch := range input {
		switch ch {
		case '\x00':
			b.WriteString(`\x00`)
		case '\r':
			b.WriteString(`\r`)
		case '\n':
			b.WriteString(`\n`)
		case '\\':
			b.WriteString(`\\`)
		case '\'':
			b.WriteString(`\'`)
		case '"':
			b.WriteString(`\"`)
		case '\x1a':
			b.WriteString(`\x1a`)
		default:
			b.WriteRune(ch)
		}
	}

	return b.String()
}

func TestRune(t *testing.T) {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			x := i*256*256 + j*256 + '\\'
			r := rune(x)
			if utf8.ValidRune(r) {
				log.Println(i, j, x, r)
			}
		}
	}
}

func TestInvalidUtf8(t *testing.T) {
	x := []byte{'\xce', '\x28'}
	r := string(x)
	log.Println(len(r))
	for _, ch := range r {
		if utf8.ValidRune(ch) {
			t.Errorf("%v is valid rune", ch)
		}
	}
}

func TestInvalidUtf8Another(t *testing.T) {
	x := []byte{'\xce', '\x28'}
	nihongo := string(x)

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
