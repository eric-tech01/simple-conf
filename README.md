# conf

## 加载配置

### 从字符串中加载配置

```golang
var content = `[app] mode="dev"`
if err := conf.Load(bytes.NewBufferString(content), toml.Unmarshal); err != nil {
    panic(err)
}
```

### 从配置文件中加载配置

```golang
import (
    file_datasource "github.com/eric-tech01/datasource/file"
)

provider := file_datasource.NewDataSource(path)
if err := conf.Load(provider, toml.Unmarshal); err != nil {
    panic(err)
}
```
