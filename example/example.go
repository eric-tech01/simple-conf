package main

import (
	"fmt"

	conf "github.com/eric-tech01/simple-conf"
	file_datasource "github.com/eric-tech01/simple-conf/datasource/file"
	toml "github.com/pelletier/go-toml"
)

func main() {

	provider, err := file_datasource.NewDataSource("./config.toml", false)
	if err != nil {
		panic(err)
	}
	if err := conf.LoadFromDataSource(provider, toml.Unmarshal); err != nil {
		panic(err)
	}
	s := conf.GetString("simple.conf.default.name")
	fmt.Printf("%s", s)

	conf.Set("simple.conf.default.age", 11)
	fmt.Printf("\n %d", conf.GetInt("simple.conf.default.age"))
}
