package cipher_test

import (
	"enc/cipher"
	"fmt"
	"testing"
)

func TestEncryptMessage(t *testing.T){
	testTable := []struct {
		mes string
		key string
		expected string
	}{
		{
			mes: "hello",
			key: "1234567890abcdef",
			expected: "iUcGEvin7do+NEuMfiPVfTm0uSsI",
		},
		{
			mes: "",
			key: "1234567890abcdef",
			expected: "S+Ivm2f9f2pkqjwaB3In0Q==",
		},
		{
			mes: "привет",
			key: "0987654321щеш",
			expected: "Igc3cKlzy2RTqod39UyviWc5urKl1kmwsRUIcg==",
		},
		{
			mes: "",
			key: "",
			expected: "could not create new cipher",
		},
		{
			mes: "hello",
			key: "",
			expected: "could not create new cipher",
		},
		{
			mes: "hello",
			key: "1234567890abc",
			expected: "could not create new cipher",
		},
	}

	for _, testCase := range testTable {
		result, err := cipher.EncryptMessage(testCase.mes, []byte(testCase.key))
		if result != testCase.expected {
			t.Errorf("Incorerct result. Expected %s, got %s", testCase.expected, result)
		}
		if err != nil {
			fmt.Println("your cipher cannot be encrypted")
		}
	}
}
	// mes := "hello"
	// key := "1234567890abcdef"
	// res, err:= EncryptMessage(mes,[]byte(key))
	// if err != nil {
	// 	return 
	// }
	// fmt.Println(res)

	// mes2 := ""
	// key2 := "0987654321poiuyt" 
	// res2, err:= EncryptMessage(mes2,[]byte(key2))
	// if err != nil {
	// 	return 
	// }
	// fmt.Println(res2)

	// mes3 := "привет"
	// key3 := "0987654321щеш" 
	// res3 , err:= EncryptMessage(mes3,[]byte(key3))
	// if err != nil {
	// 	return 
	// }
	// fmt.Println(res3)

	// mes4 := "что-то осмысленное"
	// key4 := "нога" 
	// if len(key4) != 16 || len(key4) != 24|| len(key4) != 32{
	// 	fmt.Println("this  key-size is incorrect")

	// }
	// res4, err:= EncryptMessage(mes4,[]byte(key4))	
	// if err != nil {
	// 	return 
	// }
	// fmt.Println(res4)
	
