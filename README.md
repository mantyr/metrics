# Golang Metrics and Counter

[![Build Status](https://travis-ci.org/mantyr/metrics.svg?branch=master)](https://travis-ci.org/mantyr/metrics)
[![GoDoc](https://godoc.org/github.com/mantyr/metrics?status.png)](http://godoc.org/github.com/mantyr/metrics)
[![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This don't stable version (beta-version)

## Installation

    $ go get github.com/mantyr/metrics

## Example
```GO
package main

import (
    "github.com/mantyr/metrics"
    "github.com/mantyr/metrics/counter"
)
func main() {
    m := metrics.New()

    for i := 0; i < 100; i++ {
        m.Add("key", 10)          // plus 10 int64
        counter.Add("key", 10)
    }
    m.Get("key")                  // 10 * 100 = 1000

    counter.GetAvgSpeed("key")
    counter.GetAvgSpeedString("key")
    counter.GetHourCount("key")
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
