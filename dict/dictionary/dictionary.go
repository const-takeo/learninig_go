package dictionary

import (
	"errors"
)

var (
	errNotFound   error = errors.New("見つかる事ができません")
	errCantAdd    error = errors.New("追加出来ません")
	errCantUpdate error = errors.New("更新できません")
	errCantDelete error = errors.New("削除できません")
)

//Dictionaryはただのmapの別名
//typeにもメソッドを追加出来る
//Dictionary type
//Dictonary is a hashmap. And by default a hashmap already has the * included.
type Dictionary map[string]string

//Search method of Dictionary
func (d Dictionary) Search(word string) (string, error) {
	// mapのkeyを呼び出すとmapはvalueとokを返す
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add method of Dictionary
func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)
	if err != nil {
		d[key] = word
		return nil
	}
	return errCantAdd
}

//Update method of Dictionary
func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	if err == nil {
		d[key] = word
		return nil
	}
	return errCantUpdate
}

//Delete method of Dictionary
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return errCantDelete
	}
	delete(d, key)
	return nil
}
