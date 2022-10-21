package timeparse

import "testing"

func TestParsTime(t *testing.T) {

	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:12", true},
		{"bad", false},
		{"1:-3:12", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23", false},
	}
	for _, date := range table {
		_, err := parsTime(date.time)
		if date.ok && err != nil {
			t.Errorf("%v:%v, error shoud be nil",date.time,err)	
		}
	}
}
