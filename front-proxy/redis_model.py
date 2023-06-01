import redis

from util import load_config

config_dict = load_config()
redis_config = config_dict.get("redis")

pool = redis.ConnectionPool(host=redis_config.get("host"), password=redis_config.get("password"), port=redis_config.get(
    "port"), max_connections=redis_config.get("max_connection"))
conn = redis.Redis(connection_pool=pool)
