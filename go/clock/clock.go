package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

func New(h, m int) Clock {
	for {
		if m < 0 {
			m = m + 60
			h -= 1
			continue
		}
		break
	}

	for {
		if h < 0 {
			h = h + 24
			continue
		}
		break
	}

	if m >= 60 {
		tmp := math.Floor(float64(m) / 60)
		m = m - (int(tmp) * 60)
		h = h + int(tmp)
	}

	h = h % 24
	return Clock{
		Hours:   h,
		Minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	var hours float64
	minutes := c.Minutes + m
	if minutes > 60 {
		hours = math.Floor(float64(minutes) / 60)
		minutes = minutes - (minutes * int(hours))
	}

	if minutes == 60 {
		minutes = 0
	}

	temp := c.Hours + int(hours)
	if temp == 24 {
		temp = 0
	}
	return Clock{
		Minutes: minutes,
		Hours:   temp,
	}
}

func (c Clock) Subtract(m int) Clock {
	hour := c.Hours
	minute := c.Minutes
	for {
		if minute-m <= 0 {
			minute += 60
			hour -= 1
			continue
		}

		minute -= m
		break
	}

	return Clock{
		Hours:   hour,
		Minutes: minute,
	}
}

func (c Clock) String() string {
	hour := fmt.Sprintf("%d", c.Hours)
	if c.Hours < 10 {
		hour = fmt.Sprintf("0%d", c.Hours)
	}

	minute := fmt.Sprintf("%d", c.Minutes)
	if c.Minutes < 10 {
		minute = fmt.Sprintf("0%d", c.Minutes)
	}

	return fmt.Sprintf("%s:%s", hour, minute)
}
