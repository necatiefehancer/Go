package data

import "fmt"

var isConnected bool = false

func connect() {
	isConnected = true
	fmt.Println("Connected to database")

	//added mongo db
}

func disconnect() {
	isConnected = false
	fmt.Println("disconnected to database")
}

func databaseProcessing() {
	connect()
	fmt.Println("deferring disconnect")
	defer disconnect()
	fmt.Printf("connection open %v \n", isConnected)
	fmt.Println("Doing something")
}

func Start() {
	fmt.Printf("connection state %v \n", isConnected)
	databaseProcessing()
	fmt.Printf("connection state %v \n", isConnected)

}
