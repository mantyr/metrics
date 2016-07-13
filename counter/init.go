package counter

var (
    Counter *Metrics
)

func init() {
    Counter = NewMetrics()
}

func Add(name string, count int64) {
    Counter.Add(name, count)
}

func GetAvgSpeed(name string) float64 {
    return Counter.GetAvgSpeed(name)
}

func GetAvgSpeedString(name string) string {
    return Counter.GetAvgSpeedString(name)
}

func GetHourCount(name string) int64 {
    return Counter.GetHourCount(name)
}

