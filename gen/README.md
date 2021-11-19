# 项目 Go 代码自动生成

1. 复制 config.example.json 重命名为 config.json
2. 编辑 config.json 配置
3. 根据不同的 OS 平台执行对应的可执行文件：
   1. MacOS：`./gen.macos [gen][new][del] <param>`
   2. Linux：`./gen.linux [gen][new][del] <param>`
   3. Windows：`.\gen.exe [gen][new][del] <param>`
   
   以下文档，将以 linux 平台为例。
4. 操作说明：
   - `./gen.linux gen`：代码生成操作 
   - `./gen.linux new`：新项目创建操作 
   - `./gen.linux del`：删除项目操作

## 代码生成操作：`./gen.linux gen`

- 执行文件后，也可以带上配置文件名，若不指定，则为 `config.json`，例如：

```shell
# 加载同目录下的 config-xxx.json 配置文件
 ./gen.linux gen config-xxx.json
 # 加载同目录下的 config.json 配置文件
 ./gen.linux gen
```

- 也可指定环境变量：`config.file.path`，指定加载其他路径的配置

## 新项目创建操作：`./gen.linux new`

## 删除项目操作：`./gen.linux new`



