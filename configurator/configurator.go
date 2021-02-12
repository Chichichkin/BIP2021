package configurator

import (
	"polybay/configurator/file"
	"polybay/model"
)

type config struct {
	file model.IConfigReader
	env  model.IConfigReader
}

func Read(path string) (*model.Config, error) {
	ret := config{file: file.New(path)}
	return ret.file.Read()
}
