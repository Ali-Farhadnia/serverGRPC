package book

import "encoding/json"

//cinvert book to the json string
func (book *Book) ToString() (string, error) {
	b, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
