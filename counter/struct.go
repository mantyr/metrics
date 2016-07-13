package counter

import (
    "sync"
)

type Metrics struct {
    sync.RWMutex
    d map[string]*Metric
}

type Metric struct {
    sync.RWMutex

    count int64
    times []int64
    avg   float64

    time_hour int64
}
