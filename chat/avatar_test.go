package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	//値なしで確認
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL) //OnメソッドでAvatarURLを指定 Returnで値を返す。
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.AvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" +
			"ErrNoAvatarを返すべきです")
	}
	//値をセットします
	testUrl := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testChatUser
	testUser.On("AvatarURL").Return(testUrl, nil)
	url, err = authAvatar.AvatarURL(testChatUser)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.AvatarURLは" + "エラーをかえすべきではありません")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.AvatarURLは正しいURLを返すべきです")
		}
	}
}

const (
	userid    = "1c55860804895a165a6bf5eca7c3cf3e"
	avatarurl = "https://www.gravatar.com/avatar/1c55860804895a165a6bf5eca7c3cf3e"
)

func TestGravatarAvatar(t *testing.T) {
	//値なしで確認
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: userid}
	url, err := gravatarAvatar.AvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != avatarurl {
		t.Errorf("Gravatar.AvatarURLが%sという誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {

	// テスト用のアバターのファイルを生成します
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.AvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURLが%sという誤った値を返しました", url)
	}
}
