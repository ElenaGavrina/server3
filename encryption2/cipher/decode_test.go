package cipher_test

import (
	"enc/cipher"
	"fmt"
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecryptMessage(t *testing.T) {
	testTable := []struct {
		name string
		cipher   string
		key      string
		expected string
	}{
		{	name: "one",
			cipher:   "DefsAroGZHaqZRQJHgYV8/dag5DB",
			key:      "1234567890abcdef",
			expected: "hello",
		},
		{
			name: "two",
			cipher:   "nqd585mIjLy0aptDSr+9Fw==",
			key:      "0987654321poiuyt",
			expected: "",
		},
		{
			name: "three",
			cipher:   "9PsV3VxrYubyB7mXtuBvSTOGzhXu5sZDA9n8Jg==",
			key:      "0987654321щеш",
			expected: "привет",
		},
		{
			name: "four",
			cipher:   "что-то осмысленное",
			key:      "нога",
			expected: "could not base64 decode",
		},
		{
			name: "five",
			cipher:   "",
			key:      "",
			expected: "could not create new cipher",
		},
		{
			name: "six",
			cipher:   "9PsV3VxrYubyB7mXtuBvSTOGzhXu5sZDA9n8Jg==",
			key:      "",
			expected: "could not create new cipher",
		},
		{
			name: "seven",
			cipher:   "",
			key:      "1234567890abcdef",
			expected: "invalid ciphertext block size",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T){
			result,err := cipher.DecryptMessage(testCase.cipher,[]byte(testCase.key))
			if err != nil {
				fmt.Println("your cipher cannot be decrypted")
			}
			assert.Equal(t, testCase.expected, result)
		})
		
		// result, err := cipher.DecryptMessage(testCase.cipher, []byte(testCase.key))
		// if result != testCase.expected {
		// 	t.Errorf("Incorerct result. Expected %s, got %s", testCase.expected, result)
		// }
		// if err != nil {
		// 	fmt.Println("your cipher cannot be decrypted")
		// }
	}
}