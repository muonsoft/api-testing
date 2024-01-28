package assertions

import (
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
)

const DateLayout = "2006-01-02"

// WithTime asserts that the JSON node has a string value with time in RFC3339 format.
func (a *StringAssertion) WithTime(msgAndArgs ...interface{}) *TimeAssertion {
	return a.WithTimeWithLayout(time.RFC3339, msgAndArgs...)
}

// WithTimeWithLayout asserts that the JSON node has a string value with time with the given layout.
func (a *StringAssertion) WithTimeWithLayout(layout string, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	t, err := time.Parse(layout, a.value)
	if err == nil {
		return &TimeAssertion{t: a.t, layout: layout, value: t}
	}

	a.fail(
		fmt.Sprintf(`is time: %s`, err.Error()),
		msgAndArgs...,
	)

	return nil
}

// WithDate asserts that the JSON node has a string value with date in "YYYY-MM-DD" format.
func (a *StringAssertion) WithDate(msgAndArgs ...interface{}) *TimeAssertion {
	return a.WithTimeWithLayout(DateLayout, msgAndArgs...)
}

// TimeAssertion is used to build a chain of assertions for the time node.
type TimeAssertion struct {
	t       TestingT
	message string
	layout  string
	value   time.Time
}

func NewTimeAssertion(t TestingT, message string, value time.Time, layout string) *TimeAssertion {
	return &TimeAssertion{t: t, message: message, layout: layout, value: value}
}

// EqualTo asserts that the time equals to the given value.
func (a *TimeAssertion) EqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`equal to "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotEqualTo asserts that the time not equals to the given value.
func (a *TimeAssertion) NotEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`not equal to "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// After asserts that the time after the given time.
func (a *TimeAssertion) After(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.After(expected) {
		a.fail(
			fmt.Sprintf(
				`after "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// AfterOrEqualTo asserts that the time after or equal to the given time.
func (a *TimeAssertion) AfterOrEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.After(expected) && !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`after or equal to "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Before asserts that the time before the given time.
func (a *TimeAssertion) Before(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Before(expected) {
		a.fail(
			fmt.Sprintf(
				`before "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// BeforeOrEqualTo asserts that the time before or equal to the given time.
func (a *TimeAssertion) BeforeOrEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Before(expected) && !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`before or equal to "%s", actual is "%s"`,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// EqualToDate asserts that the time equals to the given date.
func (a *TimeAssertion) EqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.EqualTo(newDate(year, month, day), msgAndArgs...)
}

// NotEqualToDate asserts that the time not equals to the given date.
func (a *TimeAssertion) NotEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.NotEqualTo(newDate(year, month, day), msgAndArgs...)
}

// AfterDate asserts that the time after the given date.
func (a *TimeAssertion) AfterDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.After(newDate(year, month, day), msgAndArgs...)
}

// AfterOrEqualToDate asserts that the time after or equal to the given date.
func (a *TimeAssertion) AfterOrEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.AfterOrEqualTo(newDate(year, month, day), msgAndArgs...)
}

// BeforeDate asserts that the time before the given date.
func (a *TimeAssertion) BeforeDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.Before(newDate(year, month, day), msgAndArgs...)
}

// BeforeOrEqualToDate asserts that the time before or equal to the given date.
func (a *TimeAssertion) BeforeOrEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.BeforeOrEqualTo(newDate(year, month, day), msgAndArgs...)
}

// AtDate asserts that the time between the beginning and the end of the given date.
func (a *TimeAssertion) AtDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	begin := newDate(year, month, day)
	end := begin.Add(24 * time.Hour)

	if !((a.value.After(begin) || a.value.Equal(begin)) && a.value.Before(end)) {
		a.fail(
			fmt.Sprintf(
				`at date "%s", actual is "%s"`,
				begin.Format(DateLayout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

func (a *TimeAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}

func newDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
