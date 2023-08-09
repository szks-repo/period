package period

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeriod_New(t *testing.T) {
	tests := []struct {
		start time.Time
		end   time.Time
	}{
		{
			start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
		},
		{
			start: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("New#%02d", i+1), func(t *testing.T) {
			p := New(tt.start, tt.end)

			assert.True(t, tt.start.Equal(p.Start()))
			assert.True(t, tt.end.Equal(p.End()))
		})
	}
}

func TestPeriod_IsWitninPeriod(t *testing.T) {
	tests := []struct {
		start   time.Time
		end     time.Time
		argTime time.Time
		want    bool
	}{
		{
			start:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:     time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			argTime: time.Date(2023, 1, 10, 23, 59, 59, 0, time.Local),
			want:    true,
		},
		{
			start:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:     time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			argTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			want:    true,
		},
		{
			start:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:     time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			argTime: time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			want:    true,
		},
		{
			start:   time.Date(2023, 1, 1, 0, 0, 0, 1, time.Local),
			end:     time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			argTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			want:    false,
		},
		{
			start:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			end:     time.Date(2023, 3, 31, 23, 59, 59, 0, time.Local),
			argTime: time.Date(2023, 3, 31, 23, 59, 59, 1, time.Local),
			want:    false,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("New#%02d", i+1), func(t *testing.T) {
			p := New(tt.start, tt.end)
			assert.Equal(t, tt.want, p.IsWithinPeriod(tt.argTime))
		})
	}
}
