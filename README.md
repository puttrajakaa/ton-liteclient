# ton-liteclient

> ton · lite · derive

[![Go 1.22+](https://img.shields.io/badge/go-1.22+-00ADD8)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Build](https://img.shields.io/badge/build-passing-brightgreen)]()

TON liteclient shell — derive, stub seqno.

## Features

- TON derivation path m/44'/607'/0'
- Local vault JSON with XOR wrap
- SHA-256 stand-in keys — no live RPC
- stdlib CLI via flag

## Prerequisites

- Go 1.22+
- Git

## Getting Started

```bash
git clone <repo-url>
cd ton-liteclient
make build
./bin/tonlite -help
```

## CLI Usage

```bash
make test
go run ./cmd/tonlite -help
```

## Project Structure

```
cmd/tonlite/main.go
internal/config/config.go
internal/crypto/keys.go
internal/wallet/wallet.go
internal/wallet/wallet_test.go
```

## Background

TON Go notes search liteclient.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.


---

## Topics

![ton](https://img.shields.io/badge/ton-111827?style=flat-square) ![liteclient](https://img.shields.io/badge/liteclient-111827?style=flat-square) ![ton-liteclient](https://img.shields.io/badge/ton%20liteclient-111827?style=flat-square) ![cryptocurrency](https://img.shields.io/badge/cryptocurrency-111827?style=flat-square) ![wallet](https://img.shields.io/badge/wallet-111827?style=flat-square) ![blockchain](https://img.shields.io/badge/blockchain-111827?style=flat-square) ![web3](https://img.shields.io/badge/web3-111827?style=flat-square) ![bitcoin](https://img.shields.io/badge/bitcoin-111827?style=flat-square)

`ton` `liteclient` `ton-liteclient` `cryptocurrency` `wallet` `blockchain` `web3` `bitcoin` `ethereum` `hd-wallet` `open-source` `golang` `go`

Search: ton-liteclient · ton · lite · derive · TON liteclient shell — derive, stub seqno.

---

<sub>TON liteclient shell — derive, stub seqno.</sub>
