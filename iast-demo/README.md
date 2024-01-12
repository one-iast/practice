# IAST Demo

`IAST Demo`是一款基于 Go 开发的 IAST Demo，使用[前置代理](https://github.com/one-iast/practice/tree/main/front-proxy)
保存流量到 redis，通过[go cron](https://github.com/go-co-op/gocron)
定时任务自动获取量，并封装数据发送到[sqlmapapi](https://github.com/sqlmapproject/sqlmap/wiki/Usage#api-rest-json))
和[burp-rest-api](https://github.com/vmware/burp-rest-api)进行 web 漏洞扫描，同时打印结果到终端并输出到文件。

最主要功能是自动化获取前置代理的流量进行 web 漏洞扫描，打印结果到终端并输出到文件。

## 功能特性

1. 自动获取前置代理流量并处理；
2. 自动封装 sqlmapapi 和 burpsuite 数据；
3. 自动提交 sqlmap 和 burpsuite 扫描任务；
4. 自动解析结果到终端并高亮是否存在漏洞，同时输出结果到文件。

## 依赖

1. [前置代理](https://github.com/one-iast/practice/tree/main/front-proxy)
2. 安装 docker、docker compose，需要通过 docker compose 启动 Redis
3. `Python >= 3.10` ，其它版本 Python3 应该也行，没试
4. `go >= 1.20`，[task](https://github.com/go-task/task)编译及运行
5. sqlmap
6. burpsuite pro 及 burp-rest-api，同时配置好对应的 java 环境；确保可以通过 rest-api 启动
   burpsuite，[参考](https://github.com/one-iast/practice/blob/main/course/2%20%E6%89%AB%E6%8F%8F%E5%B7%A5%E5%85%B7%E9%80%89%E5%8F%96.md#%E5%AF%B9%E9%9D%B6%E5%9C%BA%E8%BF%9B%E8%A1%8Cweb%E6%BC%8F%E6%B4%9E%E6%A3%80%E6%B5%8B-burpsuite)

## 编译及运行方式

> 大概流程：配置好环境-前置代理-修改配置文件-启动（可执行文件或源码）

### 通过可执行文件运行

1. 启动前置代理：参考[前置代理](https://github.com/one-iast/practice/tree/main/front-proxy)
2. 通过可执行文件运行：根据自己的系统，下载可 `build`目录下的可执行文件和项目根目录的配置文件 `config.toml`

   **可执行文件 + config.toml**
3. 修改配置文件：修改 `config.toml`中的 `cache.redis.password`，`stage.pythonPath`、`sqlmap.path`、`burpsuite.exeCmd`
4. 运行可执行文件即可

### 通过源码运行

1. 启动前置代理
2. `git clone` IAST-Demo项目
3. 修改配置文件：修改 `config.toml`中的 `cache.redis.password`，`stage.pythonPath`、`sqlmap.path`、`burpsuite.exeCmd`后运行即可

   > sqlmap或burpsuite启动失败，可以设置 showLog=true，查看报错
   >

**运行方式[二选一]** ：

* 使用 `task`运行：`task run`
* 或者使用 `go`运行：`go run main.go`

### 编译

* mac可执行文件： `task mac -- 0.0.1` 或者 `task build build-darwin -- 0.0.1`
* windows可执行文件：`task win -- 0.0.1` 或者 `task build build-windows -- 0.0.1`
* linux可执行文件：`task linux -- 0.0.1`  或者 `task build build-linux -- 0.0.1`
* 清理编译文件：`task clean`

## 通过靶场/漏洞环境验证

1. 启动前置代理：参考[前置代理](https://github.com/one-iast/practice/tree/main/front-proxy)

   ![image](https://github.com/one-iast/practice/assets/30471543/c91f87a6-4824-4944-a859-9fb94180f2bc)

2. 启动靶场：[参考](https://github.com/one-iast/practice/blob/main/course/2%20%E6%89%AB%E6%8F%8F%E5%B7%A5%E5%85%B7%E9%80%89%E5%8F%96.md#%E6%90%AD%E5%BB%BA%E6%BC%8F%E6%B4%9E%E9%9D%B6%E5%9C%BA%E7%8E%AF%E5%A2%83)

   【截图】略
3. 通过通过源码/可执行文件运行IAST Demo

   ![image](https://github.com/one-iast/practice/assets/30471543/a4982695-d50b-4b7b-84a2-39c7a75918b3)

   ![image](https://github.com/one-iast/practice/assets/30471543/b4686223-8285-44c0-8361-ff99124de6fd)

5. 通过前置代理访问靶场并获取结果

   ![image](https://github.com/one-iast/practice/assets/30471543/01d749d9-6dbd-4c01-a86b-0bf45bd39a2a)

7. 结果输出到终端及文件

   sqlmap结果：
   ![image](https://github.com/one-iast/practice/assets/30471543/30831cff-fc3e-4a34-9af7-cd65aaa4915d)
   ![image](https://github.com/one-iast/practice/assets/30471543/a641ef23-f9b6-4ece-89b9-0657898922e5)
   ![image](https://github.com/one-iast/practice/assets/30471543/a678806a-e8a8-44c2-9119-af58e239328e)

   burpsuite结果：
   ![image](https://github.com/one-iast/practice/assets/30471543/c5bfa147-6467-48d9-ab74-0fd02d5039f1)
   ![image](https://github.com/one-iast/practice/assets/30471543/675bde24-8d4a-4066-a156-3dc2bef6f19c)
   ![image](https://github.com/one-iast/practice/assets/30471543/9f4e7740-60e1-4080-a931-9a989901ccbd)
   ![image](https://github.com/one-iast/practice/assets/30471543/1a1ff41b-ab35-4431-80d8-735dd485b5ad)

## 注意事项

如果sqlmap或burpsuite启动失败，可以设置 showLog=true，查看具体的报错

## 后续操作

目前已经能够通过去重方式获取 Web 流量，从 Redis 中获取标准化后的流量进行处理，自动化进行 SQL 注入漏洞和 Web
漏洞扫描，输出结果到终端及文件中；至此，已经完成了 IAST Demo，安全人员可用于日常安全测试。
后续需要调整任务调度，优化任务调度分发、结果保存和展示等，让 IAST 可用于生产。
