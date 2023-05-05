### go-binance

A Golang SDK for [binance](https://www.binance.com) API.

[![Build Status](https://travis-ci.org/adshao/go-binance.svg?branch=master)](https://travis-ci.org/adshao/go-binance)
[![GoDoc](https://godoc.org/github.com/adshao/go-binance?status.svg)](https://godoc.org/github.com/adshao/go-binance)
[![Go Report Card](https://goreportcard.com/badge/github.com/adshao/go-binance)](https://goreportcard.com/report/github.com/adshao/go-binance)
[![codecov](https://codecov.io/gh/adshao/go-binance/branch/master/graph/badge.svg)](https://codecov.io/gh/adshao/go-binance)

All the REST APIs listed in [binance API document](https://github.com/binance-exchange/binance-official-api-docs) are implemented, as well as the websocket APIs.

For best compatibility, please use Go >= 1.8.

Make sure you have read binance API document before continuing.

### API List

Name | Description | Status
------------ | ------------ | ------------
[rest-api.md](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md) | Details on the Rest API (/api) | <input type="checkbox" checked> Implemented
[web-socket-streams.md](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md) | Details on available streams and payloads | <input type="checkbox" checked>  Implemented
[user-data-stream.md](https://github.com/binance/binance-spot-api-docs/blob/master/user-data-stream.md) | Details on the dedicated account stream | <input type="checkbox" checked>  Implemented
[margin-api.md](https://binance-docs.github.io/apidocs/spot/en) | Details on the Margin API (/sapi) | <input type="checkbox" checked>  Implemented
[futures-api.md](https://binance-docs.github.io/apidocs/futures/en/#general-info) | Details on the Futures API (/fapi) | <input type="checkbox" checked>  Partially Implemented
[delivery-api.md](https://binance-docs.github.io/apidocs/delivery/en/#general-info) | Details on the Coin-M Futures API (/dapi) | <input type="checkbox" checked>  Partially Implemented

### Installation

```shell
go get github.com/adshao/go-binance/v2
```

For v1 API, it has been moved to `v1` branch, please use:

```shell
go get github.com/adshao/go-binance/v1
```

### Importing

```golang
import (
    "github.com/adshao/go-binance/v2"
)
```

