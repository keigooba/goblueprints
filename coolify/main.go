package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0 //0か1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() { //変換の確率は1/2
			var vI int = -1
			for i, char := range word {
				switch char { //文字列の値
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() { //変換の確率は1/2
						vI = i //母音の位置をセット
					}
				}
			}
			if vI >= 0 {
				switch randBool() { //1/2で重ねるか削除
				case duplicateVowel: //母音を重ねる
					word = append(word[:vI+1], word[vI:]...)
				case removeVowel: //母音を削除する
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
