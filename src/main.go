package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("in", "", "path to input csv file.")
	outputFile := flag.String("out", "", "path to output csv file (cannot be the same as input).")
	landingPageURL := flag.String("url", "", "landing page base url.")
	key := flag.String("key", "", "project encryption key.")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" || *landingPageURL == "" || *key == "" {
		fmt.Println("Invalid parameters. Usage:\n  -in [path to input csv]\n  -out [path to output csv]\n  -url [landing page base url]\n  -encryptionKey [project encryption key]")
		os.Exit(1)
	}

	if inputFile == outputFile {
		log.Fatal("Error: input filename must be different to output filename")
	}

	k, err := hex.DecodeString(*key)
	if err != nil {
		log.Fatalf("Error: key [%s] is not a valid hexadecimal string", *key)
	}

	*landingPageURL = strings.Trim(*landingPageURL, "/")

	// setup reader
	csvIn, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)
	defer csvIn.Close()

	// setup writer
	csvOut, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	var fields []string
	// handle header
	rec, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for i, fieldName := range rec {
		if i == 0 {
			// Remove any byte order marker
			fieldName = string(bytes.Trim([]byte(fieldName), "\xef\xbb\xbf"))
		}
		fields = append(fields, fieldName)
	}
	rec = append(rec, "smartUrl")
	if err = w.Write(rec); err != nil {
		_ = csvOut.Close()
		_ = os.Remove(*outputFile)
		log.Fatal(err)
	}

	records := 0
	for {
		rec, err = r.Read()
		if err != nil {
			if err == io.EOF {
				log.Printf("Generation complete: %d records processed and written to %s", records, *outputFile)
				break
			}
			_ = csvOut.Close()
			_ = os.Remove(*outputFile)
			log.Fatal(err)
		}
		if len(rec) != len(fields) {
			_ = csvOut.Close()
			_ = os.Remove(*outputFile)
			log.Fatalf("Error: all rows must be the same length")
		}
		record := map[string]string{}
		for i, fieldValue := range rec {
			record[fields[i]] = strings.TrimSpace(fieldValue)
		}
		records++
		jsonBytes, err := json.Marshal(record)
		if err != nil {
			_ = csvOut.Close()
			_ = os.Remove(*outputFile)
			log.Fatalf("Error: could not create record JSON: %v", err)
		}

		var iv [aes.BlockSize]byte
		_, err = io.ReadFull(rand.Reader, iv[:])
		if err != nil {
			_ = csvOut.Close()
			_ = os.Remove(*outputFile)
			log.Fatalf("Error: could not generate iv: %v", err)
		}

		encrypted, err := encrypt(k, iv[:], jsonBytes)
		if err != nil {
			log.Fatalf("Error: could not encrypt data: %v", err)
		}
		b64 := base64.RawURLEncoding.EncodeToString(encrypted)
		url := fmt.Sprintf("%s?data=%s&iv=%x", *landingPageURL, b64, iv)
		rec = append(rec, url)
		if err = w.Write(rec); err != nil {
			_ = csvOut.Close()
			_ = os.Remove(*outputFile)
			log.Fatal(err)
		}
		w.Flush()
	}
}

func encrypt(key, iv, data []byte) ([]byte, error) {
	padded, err := pkcs7Pad(data)
	if err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(c, iv)
	cbc.CryptBlocks(padded, padded)
	return padded, nil
}

// pkcs7Pad appends padding.
func pkcs7Pad(data []byte) ([]byte, error) {
	padlen := 1
	for ((len(data) + padlen) % aes.BlockSize) != 0 {
		padlen = padlen + 1
	}
	pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
	return append(data, pad...), nil
}