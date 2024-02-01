package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// use var because const didn't want to work with call functions
var (
	listenAddr = ":" + getEnvString("GWOW-PORT", "8080")    // TCP server listening address
	difficulty = getEnvInt("GWOW-DIFFICULTY", 6)            // Number of leading zeros required in the hash
	difficultyString = strings.Repeat("0", difficulty)      // String of leading zeros (e.g. "000000")
)

var quotes = []string{
	"Knowing yourself is the beginning of all wisdom. – Aristotle",
	"The only true wisdom is in knowing you know nothing. – Socrates",
	"The wise man does not lay up his own treasures. The more he gives to others, the more he has for his own. – Lao Tzu",
	"The saddest aspect of life right now is that science gathers knowledge faster than society gathers wisdom. – Isaac Asimov",
	"Count your age by friends, not years. Count your life by smiles, not tears. – John Lennon",
	"In three words I can sum up everything I've learned about life: it goes on. – Robert Frost",
	"Life is like riding a bicycle. To keep your balance, you must keep moving. – Albert Einstein",
	"Life is really simple, but we insist on making it complicated. – Confucius",
	"Life is a succession of lessons which must be lived to be understood. – Helen Keller",
	"Life is what happens when you're busy making other plans. – John Lennon",
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := generateChallenge()
	_, err := conn.Write([]byte(challenge + "\n"))
	if err != nil {
		fmt.Printf("Error sending challenge: %s\n", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}
	response = strings.TrimSpace(response) // Trim space and newline

	if verifyResponse(challenge, response) {
		quote := quotes[rand.Intn(len(quotes))]
		conn.Write([]byte(quote + "\n"))
	} else {
		conn.Write([]byte("Failed to verify Proof of Work\n"))
	}
}

func generateChallenge() string {
	rand.Seed(time.Now().UnixNano())
	return "Faraway-" + strconv.Itoa(rand.Intn(1000000))
}

func verifyResponse(challenge, response string) bool {
	hash := sha256.Sum256([]byte(challenge + response))
	hexHash := hex.EncodeToString(hash[:])
	return hexHash[:difficulty] == difficultyString
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
