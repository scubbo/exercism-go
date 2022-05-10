package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	// IMO negative hours/minutes should be an error, but :shrug:
	// Also what the hell is up with GoLang's modulus not forcing to positive!?
	var actual_hours int
	var actual_minutes int
	if m < 0 {
		h -= (-m / 60) + 1
		actual_minutes = 60 + (m % 60)
	} else {
		h += m / 60
		actual_minutes = m % 60
	}

	if h < 0 {
		actual_hours = 24 + (h % 24)
	} else {
		actual_hours = h % 24
	}
	if actual_minutes == 60 {
		actual_hours = (h + 1) % 24
		actual_minutes = 0
	}
	if actual_hours == 24 {
		actual_hours = 0
	}

	return Clock{
		hour:   actual_hours,
		minute: actual_minutes,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
