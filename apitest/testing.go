package apitest

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Helper()
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Log(args ...interface{})
}
