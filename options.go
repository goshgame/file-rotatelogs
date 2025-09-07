package rotatelogs

import (
	"time"

	"github.com/Bensonzjy/file-rotatelogs/internal/option"
)

const (
	optkeyClock         = "clock"
	optkeyLinkName      = "link-name"
	optkeyMaxAge        = "max-age"
	optkeyRotationSize  = "rotation-size"
	optkeyRotationCount = "rotation-count"
)

// WithClock creates a new Option that sets a clock
// that the RotateLogs object will use to determine
// the current time.
//
// By default rotatelogs.Local, which returns the
// current time in the local time zone, is used. If you
// would rather use UTC, use rotatelogs.UTC as the argument
// to this option, and pass it to the constructor.
func WithClock(c Clock) Option {
	return option.New(optkeyClock, c)
}

// WithLocation creates a new Option that sets up a
// "Clock" interface that the RotateLogs object will use
// to determine the current time.
//
// This optin works by always returning the in the given
// location.
func WithLocation(loc *time.Location) Option {
	return option.New(optkeyClock, clockFn(func() time.Time {
		return time.Now().In(loc)
	}))
}

// WithLinkName creates a new Option that sets the
// symbolic link name that gets linked to the current
// file name being used.
func WithLinkName(s string) Option {
	return option.New(optkeyLinkName, s)
}

// WithMaxAge creates a new Option that sets the
// max age of a log file before it gets purged from
// the file system.
func WithMaxAge(d time.Duration) Option {
	return option.New(optkeyMaxAge, d)
}

// WithRotationSize creates a new Option that sets the
// log file size between rotation.
func WithRotationSize(s int64) Option {
	return option.New(optkeyRotationSize, s)
}

// WithRotationCount creates a new Option that sets the
// number of files should be kept before it gets
// purged from the file system.
func WithRotationCount(n uint) Option {
	return option.New(optkeyRotationCount, n)
}
