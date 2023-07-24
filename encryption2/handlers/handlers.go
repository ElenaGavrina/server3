package handlers

import (
	"enc/cipher"
	"enc/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func EncodingString(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}

	data, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var TextRequest struct {
		Text string `json:"text"`
		Key  string `json:"key"`
	}

	err := json.Unmarshal(data, &TextRequest)
	if err != nil {
		log.Printf("Error happened in JSON unmarshal. Err: %s", err)
	}

	enc, err := cipher.EncryptMessage(TextRequest.Text, []byte(TextRequest.Key))
	if err != nil {
		log.Printf("could not encrypt: %v",err)
	}

	res, err := json.Marshal(enc)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	fmt.Println(string(res))

}

func DecodingString(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}

	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var CipherRequest struct {
		Cipher string `json:"cipher"`
		Key    string `json:"key"`
	}

	err := json.Unmarshal(data, &CipherRequest)
	if err != nil {
		log.Printf("Error happened in JSON unmarshal. Err: %s", err)
	}

	dec, err := cipher.DecryptMessage(CipherRequest.Cipher, []byte(CipherRequest.Key))
	if err != nil {
		log.Printf("%v", err.Error())
	}

	res, err := json.Marshal(dec)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}

	fmt.Println(string(res))

}

func EncodingFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}

	r.ParseMultipartForm(10 << 20)
	multForm := r.MultipartForm
	for key := range multForm.File {
		file, _, err := r.FormFile(key)
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
	
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		err = utils.CheckingFileFormat(w, fileBytes)
		if err != nil{
			return 
		}
		
		eKey := multForm.Value["key"][0]
		
		enc, err := cipher.EncryptMessage(string(fileBytes), []byte(eKey))

		tempFile, err := ioutil.TempFile("enc-file", "*.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		tempFile.Write([]byte(enc))

	}
	
}

func DecodingFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}

	r.ParseMultipartForm(10 << 20)
	multForm := r.MultipartForm
	for key := range multForm.File {
		file, _, err := r.FormFile(key)
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		err = utils.CheckingFileFormat(w, fileBytes)
		if err != nil{
			return 
		}

		dKey := multForm.Value["key"][0]

		dec, err := cipher.DecryptMessage(string(fileBytes), []byte(dKey))

		tempFile, err := ioutil.TempFile("dec-file", "decrypt-*.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		tempFile.Write([]byte(dec))
	}
}