package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*" //入力された元の語の目印

var transforms = []string{ //stringのスライス
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) //最初にrand.Seedしないと乱数の結果が固定化される（デフォルト）ため時刻を取得する。
	s := bufio.NewScanner(os.Stdin)        //標準入力(今回はターミナル）からデータを読み込む
	for s.Scan() {                         //バイト列を区切り文字から一つずつ読み込む
		t := transforms[rand.Intn(len(transforms))]              //rand.Intoより項目をランダムに決定
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1)) //バイト列を文字列に変換 第２引数に置き換えされる文字列を指定
	}
	if s.Err() != nil { //エラーがあった場合強制終了
		log.Fatalln(s.Err())
	}
}
