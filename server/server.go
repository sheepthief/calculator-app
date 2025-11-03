package main

import (
	"encoding/json"

	"fmt"

	"net"

	"os"

	"strconv"

	"example.com/calculator"
)

type Message struct {
	TypeOfReq  string
	Formula    string
	Result     string
	MsgHistory []Message
}

var messages []Message

func handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)

	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println(" We get into err")
			return
		} else {
			calc_req := "calculation"
			calc_his := "history"
			formula := msg.Formula
			if msg.TypeOfReq == calc_req {
				//To solve equation we call calculator app
				fmt.Println("Type of request: ")
				fmt.Println(msg.TypeOfReq)
				result := calculator.Solve(formula)

				//We package the result, so we can send it to client
				var hist []Message
				message := Message{"result", formula, result, hist}
				fmt.Println("The message sent: ")
				fmt.Println(message)
				messages = append(messages, message)
				sendMessage(message)
			}

			if msg.TypeOfReq == calc_his {
				message := Message{"history", "", "", messages}
				sendMessage(message)

			}
		}

	}
}

func sendMessage(message Message) {
	fmt.Println(" We send the message")
	fmt.Println(message)
	name, _ := os.Hostname()
	ip_address_list, _ := net.LookupHost(name)
	ip_address := ip_address_list[len(ip_address_list)-1]

	full_addr := ip_address + ":" + strconv.Itoa(8081)

	conn, _ := net.Dial("tcp", full_addr)
	encoder := json.NewEncoder(conn)
	encoder.Encode(message)

}

func main() {
	fmt.Println("Server is alive")
	name, _ := os.Hostname()
	ip_address_list, _ := net.LookupHost(name)
	ip_address := ip_address_list[len(ip_address_list)-1]

	full_addr := ip_address + ":" + strconv.Itoa(8080)
	full_addr_client := ip_address + ":" + strconv.Itoa(8081)

	fmt.Println("Full Address: ")
	fmt.Println(full_addr)

	listen, err := net.Listen("tcp", full_addr) // starts listening on port
	if err != nil {
		fmt.Println("Error in listen: ")
		fmt.Println(err)
	}
	defer listen.Close()

	for {

		_, err := net.Dial("tcp", full_addr_client)
		if err == nil {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			go handleConnection(conn)
		}

	}
}
