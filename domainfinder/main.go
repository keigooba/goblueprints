package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{ //コマンドを格納
	exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout //availableの最後の出力

	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		nextCmd.Stdin = stdout //次のコマンドに出力結果を渡す
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil { //cmd.Start()でプログラムを実行できる Run()だとコマンドの終了まで待つ為今回は適していない。
			log.Panic(err) //log.Fatalln()だとdefer設定が処理されない
		} else {
			defer cmd.Process.Kill() //最後にコマンドのプロセスを終了
		}
	}

	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil { //それぞれのコマンドに対してループして終了を待つ。
			log.Panic(err)
		}
	}
}
