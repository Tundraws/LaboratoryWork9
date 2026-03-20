package main

import (
	"bufio"
	"net"
	"testing"
)

func TestProcessIncomingConnection(t *testing.T) {
	const testAddress = "127.0.0.1:8083"

	t.Run("SuccessfulConversion", func(t *testing.T) {
		const testMessage = "hello\n"
		const expectedResponse = "HELLO\n"

		listener, _ := net.Listen("tcp", testAddress)
		defer listener.Close()

		go func() {
			conn, _ := listener.Accept()
			ProcessIncomingConnection(conn)
		}()

		client, _ := net.Dial("tcp", testAddress)
		defer client.Close()

		client.Write([]byte(testMessage))
		response, _ := bufio.NewReader(client).ReadString('\n')

		if response != expectedResponse {
			t.Errorf("Expected %s, got %s", expectedResponse, response)
		}
	})

	t.Run("ReadErrorHandling", func(t *testing.T) {
		networkListener, _ := net.Listen("tcp", "127.0.0.1:8084")
		defer networkListener.Close()

		go func() {
			conn, _ := networkListener.Accept()
			conn.Close()
			ProcessIncomingConnection(conn)
		}()

		client, _ := net.Dial("tcp", "127.0.0.1:8084")
		client.Close()
	})
}

func TestStartNetworkListener(t *testing.T) {
	listener, err := StartNetworkListener("tcp", "127.0.0.1:0") 
	if err != nil {
		t.Errorf("Failed to start listener: %v", err)
	}
	if listener != nil {
		listener.Close()
	}
}