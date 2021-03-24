# App

[使用文档 https://docs.73zls.com/zls-go/#/](https://docs.73zls.com/zls-go/#/)

## 开发

```bash
go mod tidy # 下载依赖
go build -o tmpApp && ./tmpApp # 先编译再执行，首次执行会自动生成配置文件，配置文件和执行文件处于同一个目录
```

## 测试

### 集成测试

先建立测试专门使用的配置文件 `test/web/conf.yml` （具体配置参考根目录配置文件）,再执行测试命令。

```bash
go test ./test/web
```

