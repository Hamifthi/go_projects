package jsonReader

import (
	"choose_adventure"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Reader struct {
	TempStorage map[string]adventure.Story
}

func (r Reader) ReadInput(filepath string) error {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	baseValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(baseValue, &r.TempStorage)
	if err != nil {
		return err
	}
	return nil
}
