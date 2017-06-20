# Identigen

[![GoDoc](https://godoc.org/github.com/empijei/identigen?status.png)](https://godoc.org/github.com/empijei/identigen/identities)
[![Changelog](https://img.shields.io/github/release/empijei/identigen.svg)](https://github.com/empijei/identigen/releases)
[![Build Status](https://travis-ci.org/empijei/identigen.svg?branch=master)](https://travis-ci.org/empijei/identigen) 
[![Go Report Card](https://goreportcard.com/badge/github.com/empijei/identigen)](https://goreportcard.com/report/github.com/empijei/identigen)

This software is meant to create fake Italian identities to provide test accounts.

We provide [binaries](https://github.com/empijei/identigen/releases), but it is strongly suggested to obtain this package via 
```
go get github.com/empijei/identigen
```

Documentation on how to use the library can be found [here](https://godoc.org/github.com/empijei/identigen/identities) and APIs are documented [here](https://empijei.github.io/identigen/)

The usage is fairly simple:

```
Usage of identigen:
  -dt_fmt string
    	The format of the dates. Supports: 'eu','us','ja' (default "eu")
  -fields string
    	The comma separated list of fields to print. Use 'all' to print all of them (default "all")
  -format string
    	The comma separated list of formats for the output. Supports: 'json', 'csv', 'human' (default "human")
  -maxage int
    	The maximum age for random people generation. Must be positive and more than minage. (default 55)
  -minage int
    	The minimum age for random people generation. Must be positive and less than maxage. (default 25)
  -number int
    	The amount of random people to generate. Must be positive. (default 1)
```

Sample output:
```sh
$ identigen -number 2 -fields Nome,Cognome,DataDiNascita -dt_fmt ja -format json
```
```json
[
	{
		"Cognome": "Sammartano",
		"DataDiNascita": "1967/08/19",
		"Nome": "Cleo"
	},
	{
		"Cognome": "Buonfiglio",
		"DataDiNascita": "1970/09/17",
		"Nome": "Anna"
	}
]
```

For any feature request or bug report feel free to [open new issues](https://github.com/empijei/identigen/issues/new).

Want to add a feature you coded yourself? [Pull requests](https://github.com/empijei/identigen#fork-destination-box) are also welcome!

Bonus feature: all flags are six characters long!
