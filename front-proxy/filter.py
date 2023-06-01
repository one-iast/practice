import mitmproxy.http
import asyncio
from threading import Thread

from response_filter import response_filter
from util import load_config


def thread_loop(loop):
    asyncio.set_event_loop(loop)
    loop.run_forever()


parser_loop = asyncio.new_event_loop()
t = Thread(target=thread_loop, args=(parser_loop,))
t.daemon = True
t.start()

config_dict = load_config()


class Filter:

    def response(self, flow: mitmproxy.http.HTTPFlow):
        asyncio.run_coroutine_threadsafe(response_filter(flow, config_dict), parser_loop)
