package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	mathrand "math/rand"
	"net"
	"strings"
	"time"
)

var Keywords []string

func GenerateTCPClient() net.Conn {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	return conn
}

func GenerateRandPayload(c int) string {
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func PopulateKeywords() {
	buffer, err := ioutil.ReadFile("keywords.data")
	if err != nil {
		panic(err)
	}

	Keywords = strings.Split(string(buffer), "\n")
	fmt.Println(Keywords)
}

func RandomKeyword() string {
	return Keywords[mathrand.Intn(len(Keywords))]
}

func main() {
	PopulateKeywords()

	conn := GenerateTCPClient()
	for {
		payload := ""
		numKeywords := mathrand.Intn(5) + 1
		for i := 0; i < numKeywords; i++ {
			payload += RandomKeyword() + " "

			if mathrand.Intn(10) == 0 {
				payload += GenerateRandPayload(mathrand.Intn(4))
			}
		}

		if mathrand.Intn(4) == 2 {
			payload += GenerateRandPayload(mathrand.Intn(200))
		}

		log.Println(payload)

		payload += "\r\n\r\n"

		_, err := conn.Write([]byte(payload))

		if err == nil {
			buffer := make([]byte, 1024)
			n, _ := conn.Read(buffer)
			fmt.Println(n)
			log.Println("response", buffer[:10])
		}

		if err != nil {
			log.Println("[-]", payload, err.Error())
			conn.Close()

			time.Sleep(10 * time.Millisecond)
			conn = GenerateTCPClient()
		}
	}
}
