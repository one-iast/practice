import hashlib
import json

import mitmproxy.http
import redis_model
from util import (
    check_contains,
    check_contains_re,
    check_not_contains,
    check_not_contains_re,
    json_data_standardised,
    json_data_standardised_default_v,
    url_data_standardised,
    url_data_standardised_default_v,
)


async def response_filter(flow: mitmproxy.http.HTTPFlow, config_dict: {}):
    redis_con = redis_model.conn
    data = {}
    username = None
    if "proxyauth" in flow.metadata:
        username = flow.metadata["proxyauth"][0]
        data["username"] = username

    # Request
    request = flow.request
    request_url = request.url
    request_host = request.host

    request_method = request.method
    request_port = request.port
    request_headers = request.headers
    request_scheme = request.scheme
    http_version = request.http_version

    # Response
    response = flow.response
    response_status_code = response.status_code
    response_headers = response.headers

    # 处理配置文件
    proxy_config = config_dict.get("proxy")
    project = proxy_config.get("project")
    ttl = proxy_config.get("ttl")

    # 请求
    req_config = config_dict.get("request")
    # 白名单
    req_include = req_config.get("include")
    req_include_target = req_include.get("target")
    req_include_method = req_include.get("method")
    # 黑名单
    req_exclude = req_config.get("exclude")
    req_exclude_suffix = req_exclude.get("suffix")
    req_exclude_target = req_exclude.get("target")
    req_exclude_api = req_exclude.get("api")

    # 响应
    resp_config = config_dict.get("response")
    # 白名单
    resp_include = resp_config.get("include")
    resp_include_status_code = resp_include.get("statusCode")
    # 黑名单
    resp_exclude = resp_config.get("exclude")
    resp_exclude_content_type = resp_exclude.get("contentType")

    # request_host不在目标中则不做后续处理
    if check_not_contains_re(req_include_target, request_host):
        return
    # request_host在排除的目标中则不做后续处理
    if check_contains_re(req_exclude_target, request_host):
        return

    # HTTP方法白名单
    if check_not_contains(req_include_method, request_method):
        return

    # 响应
    # 仅记录状态码
    if check_not_contains(resp_include_status_code, response_status_code):
        return

    # 排除Content-Type
    if "Content-Type" in response_headers:
        if check_contains(resp_exclude_content_type, response_headers["Content-Type"]):
            return
    # 排除request api
    if check_contains_re(req_exclude_api, request_url):
        return

    # 封装数据
    data["raw_url"] = request_url
    data["scheme"] = request_scheme
    data["http_version"] = http_version
    data["method"] = request_method
    data["port"] = request_port
    data["request_host"] = request_host
    data["project"] = project
    data["response_code"] = response_status_code
    header = {}

    # 封装header
    header_field = request_headers.fields
    cookie_list = []
    for k, v in header_field:
        k_decode = k.decode()
        v_decode = v.decode()
        if k_decode.lower() != "cookie":
            header[k_decode] = v_decode
        else:
            cookie_list.append(v_decode)
    if len(cookie_list) == 1:
        header["cookie"] = cookie_list[0]
    else:
        header["cookie"] = ";".join(cookie_list)

    data["header"] = header

    standard_value_url = None
    split_url = request_url
    # 处理带?的参数
    if "?" in request_url:
        url = request_url.split("?")[0]
        split_url = url
        param = request_url.split("?")[1]
        data["url_parameter"] = param
        # 排除动态链接的静态资源的url
        req_suffix = url.rsplit(".")[-1]
        if check_contains(req_exclude_suffix, req_suffix):
            return

        # url中参数值标准化
        data_standard_v = url_data_standardised_default_v(param)
        standard_value_url = url + "?" + data_standard_v
        # 处理标准的get；非get方法，url带?的在处理Body时处理
        if request_method == "GET":
            redis_url_parameter_key = (
                f"{request_method} {standard_value_url} {request_port}"
            )
            if username is not None:
                redis_url_parameter_key += " " + username
            await data_to_redis(data, redis_con, ttl, redis_url_parameter_key)
    data["split_url"] = split_url

    # 处理带Body
    redis_body_parameter_key = ""
    if username is not None:
        redis_body_parameter_key = username + " "
    if len(request.content) and len(request.text) and request.text != "{}":
        if "=" in request.text:
            # 常规参数
            body_standard_data = url_data_standardised_default_v(request.text)
            data["body_parameter"] = url_data_standardised(request.text)
        else:
            # json参数
            body_standard_data = json_data_standardised_default_v(request.text)
            data["body_parameter"] = json_data_standardised(request.text)
        redis_body_parameter_key += (
            f"{request_method} {split_url}?{body_standard_data} {request_port}"
        )
    else:
        if "?" in request_url and request_method == "GET":
            return
            # 排除动态链接的静态资源的url
        if len(req_exclude_suffix) != 0:
            if split_url.rsplit(".")[-1] in req_exclude_suffix:
                return

        # post data为空，非GET方法，需要按get的再判断一次，不能直接在get那里入库，防止post的url带?又有data
        if "?" in request_url:
            redis_body_parameter_key += (
                f"{request_method} {standard_value_url} {request_port}"
            )
        else:
            # url 不带?，可能是restful的请求；也可能是普通请求，但是没有?参数；或者是其他不标准的请求，如PUT带参数，不带body data
            # 不管是哪种，都无法标准化，此时能记录api，但无法去重
            redis_body_parameter_key += f"{request_method} {request_url} {request_port}"
    await data_to_redis(data, redis_con, ttl, redis_body_parameter_key)


async def data_to_redis(data, redis_con, ttl, standard_set_value):
    request_host = data["request_host"]
    redis_key = "target:" + request_host
    redis_host_key = "flow:" + request_host
    if "username" in data:
        redis_key += "_" + data["username"]
        redis_host_key += "_" + data["username"]
    standard_md5 = hashlib.md5(standard_set_value.encode(encoding="utf-8")).hexdigest()
    target_exist = redis_con.sadd(redis_key, standard_md5)
    redis_con.expire(redis_key, ttl)
    if target_exist == 1:
        redis_con.lpush(redis_host_key, json.dumps(data))
