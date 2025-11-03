package main

import (
	"encoding/json"
	"time"

	"fmt"

	"net"

	"os"

	"strconv"

	"bufio"

	"strings"

	"runtime"
)

type Message struct {
	TypeOfReq  string
	Formula    string
	Result     string
	MsgHistory []Message
}

var messages []Message
var WaitingForRespond bool

func handleConnection(conn net.Conn, wait bool) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	for {
		calc_res := "result"
		calc_his := "history"

		if !wait {
			// No message has come in from server
			// So we need to request from server

			fmt.Println("Either enter a formula or type 'history' to see past results: ")
			fmt.Println("Legal operations '+','-','*','/' and '^'")
			reader := bufio.NewReader(os.Stdin)
			calc, _ := reader.ReadString('\n')

			if runtime.GOOS == "windows" {
				calc = strings.TrimRight(calc, "\r\n")
			} else {
				calc = strings.TrimRight(calc, "\n")
			}

			if calc == calc_his {
				var hist []Message
				message := Message{"history", calc, "0", hist}
				sendMessage(message)
				WaitingForRespond = true
				return
			}
			if calc != "history" {
				var hist []Message
				message := Message{"calculation", calc, "0", hist}
				sendMessage(message)
				WaitingForRespond = true
				return
			}
		}
		var msg Message
		errjson := decoder.Decode(&msg)

		if errjson != nil {
			fmt.Println(errjson)
			return
		} else {

			WaitingForRespond = false
			if msg.TypeOfReq == calc_res {
				fmt.Println("This is the result of " + msg.Formula)
				fmt.Println(msg.Result)
				fmt.Println("")
				return
			}

			if msg.TypeOfReq == calc_his {
				messages = msg.MsgHistory
				for _, res := range messages {
					fmt.Println("The formula: ")
					fmt.Println(res.Formula)
					fmt.Println("The result: ")
					fmt.Println(res.Result)
					fmt.Println("")

				}

				return
			}

		}
	}
}

func sendMessage(message Message) {
	fmt.Println("We send the message")
	name, _ := os.Hostname()
	ip_address_list, _ := net.LookupHost(name)
	ip_address := ip_address_list[len(ip_address_list)-1]

	full_addr := ip_address + ":" + strconv.Itoa(8080)

	conn, _ := net.Dial("tcp", full_addr)
	encoder := json.NewEncoder(conn)
	encoder.Encode(message)

}

func main() {
	fmt.Println("Client is alive")
	name, _ := os.Hostname()
	ip_address_list, _ := net.LookupHost(name)
	ip_address := ip_address_list[len(ip_address_list)-1]

	full_addr_server := ip_address + ":" + strconv.Itoa(8080)
	full_addr := ip_address + ":" + strconv.Itoa(8081)

	fmt.Println("Full Address: ")
	fmt.Println(full_addr)

	listen, err := net.Listen("tcp", full_addr) // starts listening on port
	if err != nil {
		fmt.Println("Error in listen: ")
		fmt.Println(err)
	}
	defer listen.Close()
	WaitingForRespond = false

	for {

		_, err := net.Dial("tcp", full_addr_server)
		if err == nil {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			if WaitingForRespond {
				go handleConnection(conn, WaitingForRespond)

				time.Sleep(time.Second)

			} else {
				handleConnection(conn, WaitingForRespond)
			}
		}

	}
}
