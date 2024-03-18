import requests
import urllib.parse
from datetime import datetime
from config import Config


def new_http_connection(timeout: int, proxy: bool, proxy_str: str = None):
    proxies = None
    if proxy and proxy_str:
        parsed_proxy = urllib.parse.urlparse(proxy_str)
        proxies = {
            "http": f"{parsed_proxy.scheme}://{parsed_proxy.netloc}",
            "https": f"{parsed_proxy.scheme}://{parsed_proxy.netloc}",
        }
    session = requests.Session()
    session.proxies = proxies
    session.timeout = timeout / 1000  # Convert milliseconds to seconds

    return session


class HttpClient:
    def __init__(self, conf: Config):
        self.timeout = conf.timeout
        if conf.proxy:
            self.session_with_proxy = new_http_connection(conf.timeout, True, conf.proxy_str)
        self.session_without_proxy = new_http_connection(conf.timeout, False)

    def network_available(self, url: str, timeout: int, proxy: bool = False):
        start_time = datetime.now()
        try:
            response = self.session.get(url)
            response.raise_for_status()
            latency = int((datetime.now() - start_time).total_seconds() * 1000)  # Convert to milliseconds
            return True, latency, response.status_code
        except requests.exceptions.RequestException:
            latency = timeout
            return False, latency, 0
