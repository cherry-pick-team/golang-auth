package Utils

import (
	"encoding/json"
	"os"
)

/**
Данные для авторизации через соц сети
*/
type Secrets struct {
	ClientID     string
	ClientSecret string
}

/**
Читает из файла параметры Secrets в формате json
*/
func ReadSecrets(fileName string) (Secrets, error) {
	s := Secrets{}
	file, err := os.Open(fileName)
	if err != nil {
		return s, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)
	return s, err
}
