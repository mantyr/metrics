package counter

func NewMetrics() (m *Metrics) {
    m = new(Metrics)
    m.d = make(map[string]*Metric)
    return
}

func (m *Metrics) Add(name string, count int64) {
    metric := m.Get(name)
    metric.Add(count)
}

func (m *Metrics) Get(name string) *Metric {
    m.RLock()
    metric, ok := m.d[name]
    m.RUnlock()

    if !ok {
        metric = NewMetric()

        m.Lock()
        m.d[name] = metric
        m.Unlock()

        go metric.Run()
    }
    return metric
}

func (m *Metrics) GetAvgSpeed(name string) float64 {
    m.RLock()
    defer m.RUnlock()

    metric, ok := m.d[name]
    if ok {
        return metric.GetAvgSpeed()
    }
    return 0.0
}

func (m *Metrics) GetAvgSpeedString(name string) string {
    m.RLock()
    defer m.RUnlock()

    metric, ok := m.d[name]
    if ok {
        return metric.GetAvgSpeedString()
    }
    return "0"
}

func (m *Metrics) GetHourCount(name string) int64 {
    m.RLock()
    defer m.RUnlock()

    metric, ok := m.d[name]
    if ok {
        return metric.GetHourCount()
    }
    return 0
}
