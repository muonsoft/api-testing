package mock

import (
	"fmt"
	"strings"
	"testing"
)

type Tester struct {
	messages []string
}

func (tester *Tester) Helper() {}

func (tester *Tester) Error(args ...interface{}) {
	tester.messages = append(tester.messages, fmt.Sprint(args...))
}

func (tester *Tester) Errorf(format string, args ...interface{}) {
	tester.messages = append(tester.messages, fmt.Sprintf(format, args...))
}

func (tester *Tester) Fatal(args ...interface{}) {
	tester.messages = append(tester.messages, fmt.Sprint(args...))
}

func (tester *Tester) AssertContains(t *testing.T, messages []string) {
	t.Helper()
	if len(tester.messages) != len(messages) {
		t.Errorf(
			"failed asserting that tester has messages count %d, actual count is %d:\n%s",
			len(messages),
			len(tester.messages),
			strings.Join(tester.messages, "\n"),
		)
	}
	for i, message := range messages {
		if len(tester.messages) <= i {
			break
		}
		if !strings.Contains(tester.messages[i], message) {
			t.Errorf("failed asserting that tester message %d contains \"%s\", actual: \n%s", i, message, tester.messages[i])
		}
	}
}
