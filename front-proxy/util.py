import json
import re

import rtoml


# 将url中的参数值标准化
def params_standardised_default_v(param):
    if '=' in param:
        key = param.split('=')[0]
        value = param.split('=')[1]
        if isinstance(value, int):
            value = 1
        elif isinstance(value, str):
            value = 'a'
        else:
            value = 'b'
        param = key + '=' + value
    return param


def params_split(param):
    if '=' in param:
        key = param.split('=')[0]
        value = param.split('=')[1]
        param = key + '=' + value
    return param


# 将data中的常规参数值标准化
def url_data_standardised_default_v(data):
    if '=' in data:
        params = data.split('&')
        result = list(map(params_standardised_default_v, params))
        return '&'.join(result)
    else:
        return data


# url参数
def url_data_standardised(data):
    if '=' in data:
        params = data.split('&')
        result = list(map(params_split, params))
        return '&'.join(result)
    else:
        return data


# 将data中的json参数值标准化
def json_data_standardised_default_v(data):
    try:
        # 不是json会直接返回
        dict_data = json.loads(data)
        dict_result = {}
        for k, v in dict_data.items():
            if isinstance(v, int):
                temp_v = 1
            elif isinstance(v, str):
                temp_v = 'a'
            else:
                temp_v = 'b'
            dict_result.update({k: temp_v})
        return json.dumps(dict_result)
    except ValueError:
        return data


def json_data_standardised(data):
    try:
        # 不是json会直接返回
        dict_data = json.loads(data)
        dict_result = {}
        for k, v in dict_data.items():
            dict_result.update({k: v})
        return json.dumps(dict_result)
    except ValueError:
        return data


def load_config():
    with open('config.toml', 'r', encoding='utf-8') as f:
        file_content = f.read()
    return rtoml.loads(file_content)


def check_not_contains(data: [], key: str):
    if len(data) != 0:
        if key not in data:
            return True
    return False


def check_contains(data: [], key: str):
    if len(data) != 0:
        if key in data:
            return True
    return False


def check_not_contains_re(data: [], key: str):
    if len(data) == 0:
        return False
    for target in data:
        pattern = re.compile(target, re.M | re.I)
        if pattern.findall(key):
            return False
    return True


def check_contains_re(data: [], key: str):
    if len(data) == 0:
        return False
    for target in data:
        pattern = re.compile(target, re.M | re.I)
        if pattern.findall(key):
            return True
    return False
