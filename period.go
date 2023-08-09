package period

import "time"

type Period struct {
	start time.Time
	end   time.Time
}

func New(start, end time.Time) *Period {
	if start.After(end) {
		panic("start must be before end")
	}
	return &Period{start: start, end: end}
}

func NewFromString(start, end string, layout string) (*Period, error) {
	s, err := time.Parse(layout, start)
	if err != nil {
		return nil, err
	}
	e, err := time.Parse(layout, end)
	if err != nil {
		return nil, err
	}

	return New(s, e), nil
}

func (p *Period) Start() time.Time {
	return p.start
}

func (p *Period) End() time.Time {
	return p.end
}

func (p *Period) IsWithinPeriod(t time.Time) bool {
	return t.After(p.start) && t.Before(p.end)
}

func (p *Period) IsOutsidePeriod(t time.Time) bool {
	return !p.IsWithinPeriod(t)
}

func (p *Period) ChunkByDateInterval(date int) []*Period {
	// TODO
	return nil
}

func (p *Period) ChunkByHourInterval(hour int) []*Period {
	// TODO
	return nil
}
