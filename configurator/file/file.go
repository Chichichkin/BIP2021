package file

import (
	"encoding/json"
	"io/ioutil"
	"polybay/model"
)

type file struct {
	path string
}

func New(path string) model.IConfigReader {
	return &file{path: path}
}

func (f *file) Read() (*model.Config, error) {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	ret := &model.Config{}
	if err = json.Unmarshal(data, ret); err != nil {
		return nil, err
	}

	return ret, nil
}
