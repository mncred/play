package system

import "github.com/shirou/gopsutil/cpu"

type System struct{}

// CpuTotalLoad returns current CPU load percent as a single value.
// -1 indicates that invalid value got.
func (s *System) CpuTotalLoad() float64 {
	loads, err := cpu.Percent(0, true)
	if err != nil {

		return -1
	}
	total := float64(0)
	for _, load := range loads {
		total += load
	}
	total /= float64(len(loads))
	return total
}

func (s *System) CpuName() string {
	info, _ := cpu.Info()
	return info[0].ModelName
}
