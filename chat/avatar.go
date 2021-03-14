package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatarはユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	AvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar //下記3つの型を保持できる

type AuthAvatar struct{}

type GravatarAvatar struct{}

type FileSystemAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) AvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) AvatarURL(u ChatUser) (string, error) {
	return "https://www.gravatar.com/avatar/" + u.UniqueID(), nil
}

var UserFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) AvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil { //avatarsのファイルのリストを取得
		for _, file := range files {
			if file.IsDir() { //ディレクトリを除外
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}

func (a TryAvatars) AvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.AvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}
