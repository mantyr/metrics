package metrics

import (
    "sync"
)

type Metrics struct {
    sync.RWMutex
    s map[string]int64
}