
Pug
===

[![Build Status](https://www.travis-ci.org/wodog/pug.svg?branch=master)](https://www.travis-ci.org/wodog/pug)
[![GoDoc](https://godoc.org/github.com/wodog/pug?status.svg)](https://godoc.org/github.com/wodog/pug)

Pug used for development or production environment config

Install
-------
```
go get github.com/wodog/pug
```

Usage
-----

replace fromat: `{{_ .XXX}}`

```
pug.Parse("config.json")
xxx := pug.GetString("XXX")
```
