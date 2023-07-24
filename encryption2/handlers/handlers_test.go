package handlers_test

import (
	"enc/handlers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodingString(t *testing.T){
	testCases := []struct{
		name string
		cipher string
		key string
		expected []byte
	}{
		{	name: "one",
			cipher:   "DefsAroGZHaqZRQJHgYV8/dag5DB",
			key:      "1234567890abcdef",
			expected: []byte("hello"),
		},
		{
			name: "two",
			cipher:   "nqd585mIjLy0aptDSr+9Fw==",
			key:      "0987654321poiuyt",
			expected: []byte(""),
		},
		{
			name: "three",
			cipher:   "9PsV3VxrYubyB7mXtuBvSTOGzhXu5sZDA9n8Jg==",
			key:      "0987654321щеш",
			expected: []byte("привет"),
		},
		{
			name: "four",
			cipher:   "что-то осмысленное",
			key:      "нога",
			expected: []byte("could not base64 decode"),
		},
		{
			name: "five",
			cipher:   "",
			key:      "",
			expected: []byte("could not create new cipher"),
		},
		{
			name: "six",
			cipher:   "9PsV3VxrYubyB7mXtuBvSTOGzhXu5sZDA9n8Jg==",
			key:      "",
			expected: []byte("could not create new cipher"),
		},
		{
			name: "seven",
			cipher:   "",
			key:      "1234567890abcdef",
			expected: []byte("invalid ciphertext block size"),
		},
	}

	handler := http.HandlerFunc(handlers.DecodingString)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rCipher,err := json.Marshal(tc.cipher)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s", err)
			}
			rKey,err := json.Marshal(tc.key)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s", err)
			}
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/ds?cipher=%s,?key=%s", string(rCipher),string(rKey)), rec.Body)
			handler.ServeHTTP(rec,req)
			assert.Equal(t, tc.expected, rec.Body.Bytes())
		})
	}
}