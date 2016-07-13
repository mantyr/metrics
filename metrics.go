package metrics

func New() (m *Metrics) {
    m = new(Metrics)
    m.s = make(map[string]int64)
    return
}

func (m *Metrics) Set(key string, value int64) *Metrics {
    m.Lock()
    defer m.Unlock()
    m.s[key] = value

    return m
}

// Example:
//  Add(key) // default Add(key, 1)
//  Add(key, plus_value)
func (m *Metrics) Add(key string, params ...int64) *Metrics {
    m.Lock()
    defer m.Unlock()

    var value int64 = 1
    if len(params) > 0 {
        value = params[0]
    }
    m.s[key] += value

    return m
}

func (m *Metrics) Get(key string) int64 {
    m.RLock()
    defer m.RUnlock()
    v, ok := m.s[key]
    if ok {
        return v
    }
    return 0
}