package common

import (
	"encoding/json"
	"os"
)

func LoadJson(filename string, conf interface{}) error {
	err := loadJson(filename, conf)
	if err != nil {
		err := loadJson(filename+".back", conf)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadJson(filename string, conf interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		return err
	}
	return nil
}

func SaveJson(filename string, conf interface{}) error {
	err1 := saveJson(filename, conf)
	err2 := saveJson(filename+".back", conf)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func saveJson(filename string, conf interface{}) error {
	str, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	jsonFile, err := os.OpenFile(filename,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	jsonFile.Write(str)
	jsonFile.Close()
	return nil
}
