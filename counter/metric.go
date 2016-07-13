package counter

import (
    "time"
    "github.com/mantyr/debug"
    "fmt"
)

func NewMetric() (m *Metric) {
    m = new(Metric)
    return
}

func (m *Metric) Add(count int64) {
    m.Lock()
    defer m.Unlock()
    m.count += count
}

func (m *Metric) Run() {
    for {
        <- time.After(1 * time.Minute)
        m.Compress()
    }
}

func (m *Metric) Compress() {
    m.Lock()
        count := m.count
        times := m.times
        time_hour := m.time_hour

        m.count = 0
    m.Unlock()

    m.compress(count, times, time_hour)
}

func (m *Metric) compress(count int64, times []int64, time_hour int64) {
    var max_times int = 60 // максимальное количество хранимых минут
    var max_avg   int = 10 // максимальное количество минут по которым вычесляется скорость

    times = append(times, count)

    debug.LevelPrintln("counter", "Counter times in compress", times)

    times_count := len(times)

    var i int = 1
    var avg float64 = 0.0

    time_start := 0
    if times_count > max_avg {
        time_start = times_count - max_avg
    }

    for _, x := range times[time_start:] {
        avg += (float64(x) - avg) / float64(i)
        i++
    }

    if times_count > max_times {
        time_hour -= times[times_count-max_times-1]          // отсекаем лишнюю минуту в часе
        times = times[times_count-max_times:]
    }
    time_hour += count                                       // количество за последний час

    // save

    m.Lock()
    defer m.Unlock()

    m.time_hour = time_hour
    m.times = times
    m.avg = avg
    return
}

// количество за последний час с шагом в минуту
func (m *Metric) GetHourCount() int64 {
    m.RLock()
    defer m.RUnlock()
    return m.time_hour
}

// средняя скорость за последнее время (до 10 минут)
func (m *Metric) GetAvgSpeed() float64 {
    m.RLock()
    defer m.RUnlock()
    return m.avg
}

func (m *Metric) GetAvgSpeedString() string {
    s := fmt.Sprintf("%.2f", m.GetAvgSpeed())
    if s[len(s)-3:] == ".00" {
        return s[:len(s)-3]
    }
    return s
}