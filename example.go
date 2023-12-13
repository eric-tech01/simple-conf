package example

import (
	file_datasource "github.com/eric-tech01/datasource/file"
	conf "github.com/eric-tech01/simple-conf"
	toml "github.com/pelletier/go-toml"
)

func main() {

	provider := file_datasource.NewDataSource("./config.toml")
	if err := conf.Load(provider, toml.Unmarshal); err != nil {
		panic(err)
	}
}
