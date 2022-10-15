package assertjson

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// IsUUID asserts that the JSON node has a string value with UUID.
func (node *AssertNode) IsUUID(msgAndArgs ...interface{}) *UUIDAssertion {
	node.t.Helper()
	return node.IsString().WithUUID(msgAndArgs...)
}

// WithUUID asserts that the JSON node has a string value with UUID.
func (a *StringAssertion) WithUUID(msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	id, err := uuid.FromString(a.value)
	if err == nil {
		return &UUIDAssertion{t: a.t, message: a.message, path: a.path, value: id}
	}

	a.fail(
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is UUID, actual is "%s"`,
			a.path,
			a.value,
		),
		msgAndArgs...,
	)

	return nil
}

// UUIDAssertion is used to build a chain of assertions for the UUID node.
type UUIDAssertion struct {
	t       TestingT
	message string
	path    string
	value   uuid.UUID
}

// Nil asserts that the JSON node has a string value equals to nil UUID.
func (a *UUIDAssertion) Nil(msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !a.value.IsNil() {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is nil UUID, actual is "%s"`,
				a.path,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotNil asserts that the JSON node has a string value equals to not nil UUID.
func (a *UUIDAssertion) NotNil(msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value.IsNil() {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is not nil UUID, actual is "%s"`,
				a.path,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// Version asserts that the JSON node has a string value equals to UUID with the given version.
func (a *UUIDAssertion) Version(version byte, msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value.Version() != version {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is UUID of version %d, actual is %d`,
				a.path,
				version,
				a.value.Version(),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Variant asserts that the JSON node has a string value equals to UUID with the given variant.
func (a *UUIDAssertion) Variant(variant byte, msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value.Variant() != variant {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is UUID of variant %d, actual is %d`,
				a.path,
				variant,
				a.value.Variant(),
			),
			msgAndArgs...,
		)
	}

	return a
}

// EqualTo asserts that the JSON node is UUID equals to the given value.
func (a *UUIDAssertion) EqualTo(expected uuid.UUID, msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value != expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is UUID equal to "%s", actual is "%s"`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotEqualTo asserts that the JSON node is UUID not equals to the given value.
func (a *UUIDAssertion) NotEqualTo(expected uuid.UUID, msgAndArgs ...interface{}) *UUIDAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value == expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is UUID not equal to "%s", actual is "%s"`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// Value returns JSON node value as UUID. If string is not a valid UUID it returns nil UUID.
func (a *UUIDAssertion) Value() uuid.UUID {
	if a == nil {
		return uuid.Nil
	}
	a.t.Helper()

	return a.value
}

// UUID asserts that the JSON node is UUID and returns its value. If value is not a valid UUID,
// then it will return nil UUID. It is an alias for IsUUID().Value().
func (node *AssertNode) UUID() uuid.UUID {
	return node.IsUUID().Value()
}

func (a *UUIDAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	if a.message != "" {
		message = a.message + ": " + message
	}
	assert.Fail(a.t, message, msgAndArgs...)
}
