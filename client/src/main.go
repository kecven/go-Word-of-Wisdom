package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	serverAddr = getEnvString("GWOW-SERVER", "localhost:8080")  // TCP server listening address
	difficulty = getEnvInt("GWOW-DIFFICULTY", 6)                // Number of leading zeros required in the hash
	difficultyString = strings.Repeat("0", difficulty)          // String of leading zeros (e.g. "000000")
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Read challenge from server
	challenge, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	challenge = strings.TrimSpace(challenge)

	fmt.Println("Challenge received:", challenge)

	nonce := solveChallenge(challenge)
	fmt.Println("Sending nonce:", nonce)

	// Send nonce back to server
	_, err = conn.Write([]byte(nonce + "\n"))

	if err != nil {
		panic(err)
	}

	// Wait for and print the quote
	quote, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("Quote received:", quote)
}

func solveChallenge(challenge string) string {
	var nonce int64
	for {
		attempt := fmt.Sprintf("%s%d", challenge, nonce)
		hash := sha256.Sum256([]byte(attempt))
		hexHash := hex.EncodeToString(hash[:])

		if hexHash[:difficulty] == difficultyString {
			return fmt.Sprintf("%d", nonce)
		}
		nonce++
	}
}

func getEnvString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := getEnvString(key, "")
	if value == "" {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		fmt.Printf("Error parse int: %s\n", err)
		return defaultValue
	}

	return int(result)
}
