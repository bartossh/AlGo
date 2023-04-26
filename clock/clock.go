package clock

import "fmt"

const (
	hrs  = 24
	mInH = 60
)

type Clock struct {
	h int
	m int
}

func New(h int, m int) Clock {
	totalM := h*mInH + m
	hr := int(totalM / mInH)
	mLeft := int(totalM % mInH)
	hr = int(hr % hrs)
	if mLeft < 0 {
		hr = hr - 1
		mLeft = mInH + mLeft
	}
	if hr < 0 {
		hr = hrs + hr
	}
	return Clock{
		h: hr,
		m: mLeft,
	}
}

func (c Clock) String() string {
	s := ""
	if c.h < 10 {
		s = fmt.Sprintf("0%v", c.h)
	} else {
		s = fmt.Sprintf("%v", c.h)
	}
	if c.m < 10 {
		s = fmt.Sprintf("%s:0%v", s, c.m)
	} else {
		s = fmt.Sprintf("%s:%v", s, c.m)
	}
	return s
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}
