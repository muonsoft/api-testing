package assertjson

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// IsStrings asserts that the JSON node is an array of strings.
// It returns StringsAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsStrings(msgAndArgs ...interface{}) *StringsAssertion {
	node.t.Helper()
	if node.exists() {
		arr, ok := node.value.([]interface{})
		if !ok {
			node.fail(
				fmt.Sprintf(`failed asserting that JSON node "%s" is array of strings`, node.path.String()),
				msgAndArgs...,
			)
			return nil
		}
		values := make([]string, 0, len(arr))
		for i, v := range arr {
			s, ok := v.(string)
			if !ok {
				pathStr := node.path.WithIndex(i).String()
				node.fail(
					fmt.Sprintf(`failed asserting that JSON node "%s" is array of strings, element at index %d is not string`, pathStr, i),
					msgAndArgs...,
				)
				return nil
			}
			values = append(values, s)
		}
		return &StringsAssertion{
			t:       node.t,
			message: fmt.Sprintf(`%sfailed asserting that JSON node "%s": `, node.message, node.path.String()),
			path:    node.path.String(),
			value:   values,
		}
	}

	return nil
}

// StringsAssertion is used to build a chain of assertions for the array of strings node.
type StringsAssertion struct {
	t       TestingT
	message string
	path    string
	value   []string
}

// EqualTo asserts that the array of strings equals the given values (length and elements match).
func (a *StringsAssertion) EqualTo(expected ...string) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !areStringsEqual(a.value, expected) {
		a.fail(
			fmt.Sprintf(`equal to [%s], actual is [%s]`, formatStrings(expected), formatStrings(a.value)),
		)
	}

	return a
}

// WithUniqueValues asserts that the array of strings has unique values.
func (a *StringsAssertion) WithUniqueValues(msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	uniques := make(map[string][]int, len(a.value))
	keys := make([]string, 0, len(a.value))

	for i, s := range a.value {
		if _, exist := uniques[s]; !exist {
			keys = append(keys, s)
		}
		uniques[s] = append(uniques[s], i)
	}

	duplicates := make([]string, 0)
	for _, key := range keys {
		if len(uniques[key]) > 1 {
			duplicates = append(duplicates, fmt.Sprintf(
				"value %s is duplicated at %s",
				strconv.Quote(key),
				strings.Join(intsToStrings(uniques[key]), ", "),
			))
		}
	}

	if len(duplicates) > 0 {
		a.fail(
			fmt.Sprintf(
				"failed asserting that JSON node \"%s\" is array of strings with unique values, duplicated elements:\n%s",
				a.path,
				strings.Join(duplicates, ";\n"),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Contains asserts that the array of strings contains the given value.
func (a *StringsAssertion) Contains(value string, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	for _, s := range a.value {
		if s == value {
			return a
		}
	}

	a.fail(
		fmt.Sprintf(`array contains %s, actual elements are [%s]`, strconv.Quote(value), formatStrings(a.value)),
		msgAndArgs...,
	)

	return a
}

// WithLength asserts that the array of strings has length equal to the given value.
func (a *StringsAssertion) WithLength(expected int, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) != expected {
		a.fail(
			fmt.Sprintf(
				`is array of strings with length is %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThan asserts that the array of strings has length greater than the value.
func (a *StringsAssertion) WithLengthGreaterThan(expected int, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) <= expected {
		a.fail(
			fmt.Sprintf(
				`is array of strings with length greater than %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThanOrEqual asserts that the array of strings has length
// greater than or equal to the value.
func (a *StringsAssertion) WithLengthGreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) < expected {
		a.fail(
			fmt.Sprintf(
				`is array of strings with length greater than or equal to %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThan asserts that the array of strings has length less than the value.
func (a *StringsAssertion) WithLengthLessThan(expected int, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) >= expected {
		a.fail(
			fmt.Sprintf(
				`is array of strings with length less than %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThanOrEqual asserts that the array of strings has length
// less than or equal to the value.
func (a *StringsAssertion) WithLengthLessThanOrEqual(expected int, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) > expected {
		a.fail(
			fmt.Sprintf(
				`is array of strings with length less than or equal to %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// That asserts that the array of strings is satisfied by the callback function.
func (a *StringsAssertion) That(f func(values []string) error, msgAndArgs ...interface{}) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if err := f(a.value); err != nil {
		a.fail(
			fmt.Sprintf(
				`failed asserting JSON node "%s": %s`,
				a.path,
				err.Error(),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Assert asserts that the array of strings is satisfied by the user function assertFunc.
func (a *StringsAssertion) Assert(assertFunc func(tb testing.TB, values []string)) *StringsAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	assertFunc(a.t.(testing.TB), a.value)

	return a
}

// Length returns the array of strings length.
func (a *StringsAssertion) Length() int {
	if a == nil {
		return 0
	}
	a.t.Helper()

	return len(a.value)
}

func (a *StringsAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}
