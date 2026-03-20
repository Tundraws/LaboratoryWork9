package main

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func ProcessIncomingConnection(connection net.Conn) {
	defer connection.Close()

	clientReader := bufio.NewReader(connection)
	receivedData, readError := clientReader.ReadString('\n')
	if readError != nil {
		return
	}

	upperCaseResult := strings.ToUpper(receivedData)
	connection.Write([]byte(upperCaseResult))
}

func StartNetworkListener(protocol string, address string) (net.Listener, error) {
	return net.Listen(protocol, address)
}

func main() {
	const serverNetwork = "tcp"
	const serverAddress = ":8080"
	const fatalExitCode = 1

	networkListener, listenError := StartNetworkListener(serverNetwork, serverAddress)
	if listenError != nil {
		os.Exit(fatalExitCode)
	}
	defer networkListener.Close()

	for {
		clientConnection, acceptError := networkListener.Accept()
		if acceptError != nil {
			continue
		}
		go ProcessIncomingConnection(clientConnection)
	}
}