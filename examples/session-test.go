package main

import (
	"fmt"
	"github.com/shanling2016/requests"
)

func main() {
	session := requests.NewSession()
	session.Proxies = "socks5://127.0.0.1:8889"
	session.Ja3 = "771,49196-49195-49200-157-49198-49202-159-163-49199-156-49197-49201-158-162-49188-49192-61-49190-49194-107-106-49162-49172-53-49157-49167-57-56-49187-49191-60-49189-49193-103-64-49161-49171-47-49156-49166-51-50-49160-49170-10-49155-49165-22-19-255,0-5-10-11-13-50-17-23-43,23-24-25-256-257-258-259-260,0"
	response, err := session.Get("https://ipinfo.io", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Text)

	response, err = session.Get("https://ipinfo.io", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Text)

	response, err = session.Get("https://tls.peet.ws/api/all", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Text)

	response, err = session.Get("https://tls.peet.ws/api/all", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Text)
}
