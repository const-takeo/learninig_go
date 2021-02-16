package dictionary

import "errors"

var errNotFound error = errors.New("見つかる事ができません")
var errCantAdd error = errors.New("追加出来ません")

//Dictionaryはただのmapの別名
//typeにもメソッドを追加出来る
//Dictionary type
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
