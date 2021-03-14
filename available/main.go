package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43") //ポート43で接続
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "\r\n")) //\r\n復帰と改行を表す ドメイン名を書き込む
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") { //no matchがあればドメインが使用されていない
			return false, nil
		}
	}
	return true, nil
}

var marks = map[bool]string{true: "OK", false: "NG"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain) //domainifyに渡す
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second) //WHOISサーバーの負荷を下げる為、1秒休止
	}
}
