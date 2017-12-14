# gongt [![License: Apache](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![release](https://img.shields.io/github/release/yahoojapan/gongt.svg)](https://github.com/yahoojapan/gongt/releases/latest) [![CircleCI](https://circleci.com/gh/yahoojapan/gongt.svg?style=shield)](https://circleci.com/gh/yahoojapan/gongt) [![codecov](https://codecov.io/gh/yahoojapan/gongt/branch/master/graph/badge.svg)](https://codecov.io/gh/yahoojapan/gongt) [![Go Report Card](https://goreportcard.com/badge/github.com/yahoojapan/gongt)](https://goreportcard.com/report/github.com/yahoojapan/gongt) [![GoDoc](http://godoc.org/github.com/yahoojapan/gongt?status.svg)](http://godoc.org/github.com/yahoojapan/gongt) [![Join the chat at https://gitter.im/yahoojapan/gongt](https://badges.gitter.im/yahoojapan/gongt.svg)](https://gitter.im/yahoojapan/gongt?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## Description
gongt provides Go API for [NGT](https://github.com/yahoojapan/NGT).

NGT is Neighborhood Graph and Tree for Indexing High-dimensional Data. If you want more information, please read [NGT repository](https://github.com/yahoojapan/NGT).

## Dependency
- [NGT](https://github.com/yahoojapan/NGT)

## Installation
```
$ go get -u github.com/yahoojapan/gongt
```

## Example
```go
package main

import (
	"fmt"

	"github.com/yahoojapan/gongt"
)

func main() {
	defer gongt.Get().SetIndexPath("assets/example").Open().Close()
	if errs := gongt.GetErrors(); len(errs) > 0 {
		panic(errs)
	}

	fmt.Printf("Dimension: %d\n", gongt.GetDim())

	query := []float64{12, 17, 21, 18, 17, 31, 33, 25, 26, 19, 42, 31, 25, 26, 49, 30, 19, 23, 29, 29, 22, 19, 28, 27, 28, 19, 13, 12, 25, 21, 25, 21, 35, 12, 44, 36, 19, 49, 104, 33, 29, 77, 43, 36, 28, 44, 90, 46, 52, 37, 65, 42, 33, 40, 104, 103, 44, 26, 50, 43, 18, 20, 48, 68, 28, 16, 104, 27, 6, 36, 98, 327, 53, 81, 40, 36, 61, 104, 44, 27, 42, 84, 55, 54, 49, 53, 28, 27, 103, 42, 27, 28, 24, 53, 60, 66, 7, 42, 14, 6, 32, 69, 15, 3, 4, 79, 27, 7, 30, 82, 26, 3, 15, 27, 18, 6, 19, 52, 21, 16, 104, 72, 30, 40, 22, 36, 19, 22}

	results, err := gongt.Search(query, 10, gongt.DefaultEpsilon)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, r := range results {
		fmt.Printf("Rank %d\n", i+1)
		fmt.Printf("  ID: %d\n", r.ID)
		fmt.Printf("  Distance: %f\n", r.Distance)
	}
}
```
### result
```
$ go run example.go
Dimension: 128
Rank 1
  ID: 2892
  Distance: 273.355072
Rank 2
  ID: 2138
  Distance: 274.874512
Rank 3
  ID: 2422
  Distance: 276.372925
Rank 4
  ID: 1586
  Distance: 277.263428
Rank 5
  ID: 679
  Distance: 277.564392
Rank 6
  ID: 1564
  Distance: 278.792023
Rank 7
  ID: 2594
  Distance: 281.176086
Rank 8
  ID: 2159
  Distance: 281.895386
Rank 9
  ID: 2738
  Distance: 282.876312
Rank 10
  ID: 318
  Distance: 283.339020
```

License
-------

Copyright (C) 2017 Yahoo Japan Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this software except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Contributor License Agreement
-----------------------------

This project requires contributors to agree to a [Contributor License Agreement (CLA)](https://gist.github.com/ydnjp/3095832f100d5c3d2592).

Note that only for contributions to the gongt repository on the GitHub (https://github.com/yahoojapan/gongt), the contibutors of them shall be deemed to have agreed to the CLA without individual written agreements.

Authors
-------

[Kosuke Morimoto](https://github.com/kou-m)  
[kpango](https://github.com/kpango)
