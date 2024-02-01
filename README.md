# go-Word-of-Wisdom
Test task for Faraway. 

Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used. 
- The choice of the POW algorithm should be explained. 
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes. 
- Docker file should be provided both for the server and for the client that solves the POW challenge


## How to run
## Docker compose
1. Clone the repository
2. Run `docker-compose up` in the root directory
3. Run `docker-compose run gwow-client` in the root directory
4. Follow the instructions in the client container

## without docker
1. Clone the repository
2. Run `go run ./server/src/main.go` in the root directory
3. Run `go run ./client/src/main.go` in the root directory
4. Follow the instructions in the client container