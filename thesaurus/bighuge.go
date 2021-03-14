package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"` //名詞
	Verb *words `json:"verb"` //動詞
}

type words struct {
	Syn []string `json:"syn"` //名詞・動詞の類語
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: %qの類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil { //デコードして結果をdata変数にセット
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...) //名詞の類語だけを追加
	syns = append(syns, data.Verb.Syn...) //動詞の類語だけを追加
	return syns, nil
}
