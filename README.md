# go-socks5-proxy
[![CircleCI](https://circleci.com/gh/serjs/socks5-server.svg?style=shield)](https://circleci.com/gh/serjs/socks5-server)

Simple socks5 server using go-socks5 with auth

# Start container with proxy
```docker run -d --name socks5-proxy -p 1080:1080 -e PROXY_AUTH='{"<PROXY_USER_1>": "<PROXY_PASSWORD_1>", "<PROXY_USER_2>": "<PROXY_PASSWORD_2>"} serjs/go-socks5-proxy```

where `PROXY_AUTH` is a JSON map of authorized user names and the respective passwords.

# Test running service
```curl --socks5 <docker machine ip>:1080 -U <PROXY_USER>:<PROXY_PASSWORD> https://ifcfg.me``` - result must show docker host ip (for bridged network)
