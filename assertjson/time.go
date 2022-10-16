package assertjson

import (
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"
)

const dateLayout = "2006-01-02"

// IsTime asserts that the JSON node has a string value with time in RFC3339 format.
func (node *AssertNode) IsTime(msgAndArgs ...interface{}) *TimeAssertion {
	node.t.Helper()
	return node.IsString().WithTime(msgAndArgs...)
}

// IsTimeWithLayout asserts that the JSON node has a string value with time with the given layout.
func (node *AssertNode) IsTimeWithLayout(layout string, msgAndArgs ...interface{}) *TimeAssertion {
	node.t.Helper()
	return node.IsString().WithTimeWithLayout(layout, msgAndArgs...)
}

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
		return &TimeAssertion{t: a.t, message: a.message, path: a.path, layout: layout, value: t}
	}

	a.fail(
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is time: %s`,
			a.path,
			err.Error(),
		),
		msgAndArgs...,
	)

	return nil
}

// IsDate asserts that the JSON node has a string value with date in "YYYY-MM-DD" format.
func (node *AssertNode) IsDate(msgAndArgs ...interface{}) *TimeAssertion {
	node.t.Helper()
	return node.IsString().WithDate(msgAndArgs...)
}

// WithDate asserts that the JSON node has a string value with date in "YYYY-MM-DD" format.
func (a *StringAssertion) WithDate(msgAndArgs ...interface{}) *TimeAssertion {
	return a.WithTimeWithLayout(dateLayout, msgAndArgs...)
}

// TimeAssertion is used to build a chain of assertions for the time node.
type TimeAssertion struct {
	t       TestingT
	message string
	path    string
	layout  string
	value   time.Time
}

// EqualTo asserts that the JSON node is time equals to the given value.
func (a *TimeAssertion) EqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time equal to "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotEqualTo asserts that the JSON node is time not equals to the given value.
func (a *TimeAssertion) NotEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time not equal to "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// After asserts that the JSON node is time after the given time.
func (a *TimeAssertion) After(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.After(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time after "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// AfterOrEqualTo asserts that the JSON node is time after or equal to the given time.
func (a *TimeAssertion) AfterOrEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.After(expected) && !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time after or equal to "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Before asserts that the JSON node is time before the given time.
func (a *TimeAssertion) Before(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Before(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time before "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// BeforeOrEqualTo asserts that the JSON node is time before or equal to the given time.
func (a *TimeAssertion) BeforeOrEqualTo(expected time.Time, msgAndArgs ...interface{}) *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.value.Before(expected) && !a.value.Equal(expected) {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is time before or equal to "%s", actual is "%s"`,
				a.path,
				expected.Format(a.layout),
				a.value.Format(a.layout),
			),
			msgAndArgs...,
		)
	}

	return a
}

// EqualToDate asserts that the JSON node is time equals to the given date.
func (a *TimeAssertion) EqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.EqualTo(newDate(year, month, day), msgAndArgs...)
}

// NotEqualToDate asserts that the JSON node is time not equals to the given date.
func (a *TimeAssertion) NotEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.NotEqualTo(newDate(year, month, day), msgAndArgs...)
}

// AfterDate asserts that the JSON node is time after the given date.
func (a *TimeAssertion) AfterDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.After(newDate(year, month, day), msgAndArgs...)
}

// AfterOrEqualToDate asserts that the JSON node is time after or equal to the given date.
func (a *TimeAssertion) AfterOrEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.AfterOrEqualTo(newDate(year, month, day), msgAndArgs...)
}

// BeforeDate asserts that the JSON node is time before the given date.
func (a *TimeAssertion) BeforeDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.Before(newDate(year, month, day), msgAndArgs...)
}

// BeforeOrEqualToDate asserts that the JSON node is time before or equal to the given date.
func (a *TimeAssertion) BeforeOrEqualToDate(year int, month time.Month, day int, msgAndArgs ...interface{}) *TimeAssertion {
	return a.BeforeOrEqualTo(newDate(year, month, day), msgAndArgs...)
}

// Value returns JSON node value as time.Time. If string is not a valid time it returns empty time.
func (a *TimeAssertion) Value() time.Time {
	if a == nil {
		return time.Time{}
	}
	a.t.Helper()

	return a.value
}

// Time asserts that the JSON node is time.Time and returns its value.
// If string is not a valid time it returns empty time. It is an alias for IsTime().Value().
func (node *AssertNode) Time() time.Time {
	return node.IsTime().Value()
}

// TimeWithLayout asserts that the JSON node is time.Time with specific layout and returns its value.
// If string is not a valid time it returns empty time. It is an alias for IsTimeWithLayout().Value().
func (node *AssertNode) TimeWithLayout(layout string) time.Time {
	return node.IsTimeWithLayout(layout).Value()
}

func (a *TimeAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	if a.message != "" {
		message = a.message + ": " + message
	}
	assert.Fail(a.t, message, msgAndArgs...)
}

func newDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
