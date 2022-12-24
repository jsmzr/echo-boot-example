# echo-boot-example

boot-echo 使用示例

按需声明需要使用的插件，并通过 `application.yaml` 配置文件或 env 参数设置插件

```sh
curl 'http://127.0.0.1:8080/demo'
curl 'http://127.0.0.1:8080/demo/config?key=foo.bar'
# prometheus
curl 'http://127.0.0.1:8081/mertrics'
```