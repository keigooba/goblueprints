package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid") //ハッシュ化したメールアドレス
	file, header, err := req.FormFile("avatarFile") //１つ目がmultipart.File ２つ目がメタデータ
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file) //バイト列の読み込み
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename)) //filepath.Extで名前を削除してuseridを使用
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "成功")
}
