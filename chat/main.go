package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// 現在のアクティブなAvatarの実装
var avatars Avatar = TryAvatars{
	UserFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar}

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host, //HostとUserDataが含まれる
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈する
	//Gomniauthのセットアップ
	gomniauth.SetSecurityKey("セキュリティキー")
	gomniauth.WithProviders(
		google.New("587125882573-qc7jp6pmuvps29kd541qi0qds97480mt.apps.googleusercontent.com", "j409W6R1jm6kB8NNrEOWfFjk", "http://localhost:8080/auth/callback/google"),
	)
	http.Handle("/avatars/",http.StripPrefix("/avatars/",http.FileServer(http.Dir("./avatars"))))
	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/room", r)
	// チャットルームを開始します
	go r.run()
	// Webサーバーを起動します
	log.Println("Webサーバーを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
