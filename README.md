<h1 align="center"><a href="https://github.com/Dentrax/obscure-go">obscure-go</a></h1>

<div align="center">
 <strong>
   [W.I.P] In-memory secure types framework for Go.
 </strong>
</div>

<br />

<p align="center">
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square" alt="MIT"></a>
  <a href="https://goreportcard.com/report/github.com/Dentrax/obscure-go"><img src="https://goreportcard.com/badge/github.com/Dentrax/obscure-go?style=flat-square" alt="Go Report"></a>
  <a href="https://github.com/Dentrax/obscure-go/actions?workflow=test"><img src="https://img.shields.io/github/workflow/status/Dentrax/obscure-go/Test?label=build&logo=github&style=flat-square" alt="Build Status"></a>
  <a href="https://coveralls.io/repos/github/jandelgado/golang-ci-template-github-actions/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/Dentrax/obscure-go/master?style=flat-square" alt="Build Status"></a>
  <a href="https://github.com/Dentrax/obscure-go/releases/latest"><img src="https://img.shields.io/github/release/Dentrax/obscure-go.svg?style=flat-square" alt="GitHub release"></a>
</p>

<br />

# Supports

> * SecureInteger
> * SecureString

# Usage

## Importing

```go
import (
	secure "github.com/Dentrax/obscure-go/types"
)
```

## Creating

```go
secInt := secure.NewInt(15)
secStr := secure.NewString("foo")
```

## Hack Detecting

```go
// Importing
import (
	secure "github.com/Dentrax/obscure-go/types"
)

// Creating
w := obs.CreateWatcher("watcher")

// Attaching
secInt.AddWatcher(w)
secStr.AddWatcher(w)
```

# Function Interfaces

## Integer

```go
type ISecureInt interface {
	Apply() ISecureInt
	AddWatcher(obs obs.Observer)
	SetKey(int)
	Inc() ISecureInt
	Dec() ISecureInt
	Set(int) ISecureInt
	Get() int
	GetSelf() *SecureInt
	Decrypt() int
	RandomizeKey()
	IsEquals(ISecureInt) bool
}
```

## String

```go
type ISecureString interface {
	Apply() ISecureString
	AddWatcher(obs obs.Observer)
	SetKey(int)
	Set(string) ISecureString
	Get() string
	GetSelf() *SecureString
	Decrypt() []rune
	RandomizeKey()
	IsEquals(ISecureString) bool
}
```

# Example

* Non-Secure

```bash
$ go run ./examples/games/nonsecure/game_nonsecure.go
```

* Secure

```bash
$ go run ./examples/games/secure/game_secure.go 
```

# License

The base project code is licensed under _MIT License_ unless otherwise specified. Please see the **[LICENSE](https://github.com/Dentrax/obscure-go/blob/master/LICENSE)** file for more information.

# Copyright

_obscure-go_ was created by Furkan ([Dentrax](https://github.com/Dentrax))

<kbd>obscure-go</kbd>