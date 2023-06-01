# IAST Front Proxy

`IAST Front Proxy`是一款基于[mitmproxy](https://mitmproxy.org/) 的 Web 流量获取工具，可以根据配置文件对流量进行前置处理：黑白名单判断、状态码过滤、动态链接过滤、参数值标准化、流量去重等，具体可以看功能特性。

最主要功能是作为 IAST 的 Web 前置代理，来获取目标流量。

## 功能特性

1. 异步处理流量，用户端不受影响；
2. 通过域名/IP 判断目标；
3. 对请求进行标准化处理去重；
4. 根据静态资源后缀、content type 过滤流量；
5. 根据状态码、http 方法、接口、指定保存的流量；
6. 流量过滤使用 Redis，流量保存到 Redis 中；
7. 可指定代理账号密码。

## 依赖

1. 通过[官方 Advanced Installation](https://docs.mitmproxy.org/stable/overview-installation/#advanced-installation)（推荐）安装 mitmproxy
2. 安装 docker、docker compose，需要通过 docker compose 启动 Redis
3. `Python >= 3.10` ，其它版本 Python3 应该也行，没试

## 运行方式

1. 修改`docker-compose.yaml`和`config.toml`文件中的 Redis 密码（建议修改）；如果已经有 Redis 服务了，可以直接修改配置文件，然后就不用启动 Redis 服务了
2. `docker compose up` 启动 Redis 服务
3. 向 mitmporxy 注入依赖，`pipx inject mitmproxy rtoml`
4. 调整`config.toml`中的配置参数，如：`request.target`
5. 后续通过无账号密码或带账号密码启动后，再配置浏览器代理即可

### 无账号密码启动

`mitmdump -s addons.py`

### 带账号密码启动（用户名 test、密码 test）

`mitmdump -s addons.py --proxyauth test:test`

### 浏览器配置代理并访问目标

启动后，默认代理 IP 端口：`*:8080`，浏览器配置对应代理即可（可能需要输入密码），访问目标站点，可在 Redis 中查看结果。

## 运行结果保存在 Redis 中

去重的流量保存的 key，`target:`开头

标准化流量保存的 key，`flow:` 开头

## 注意事项

开发过程中，如果代码有报错，运行时不会输出报错，比较难排查。

## 后续操作

目前已经能够通过去重方式获取 Web 流量，后续可以从 Redis 中获取标准化后的流量进行处理，如安全扫描等。
