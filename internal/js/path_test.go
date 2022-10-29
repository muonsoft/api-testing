package js_test

import (
	"testing"

	"github.com/muonsoft/api-testing/internal/js"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertyPath_With(t *testing.T) {
	path := js.NewPath(js.PropertyName("top"), js.ArrayIndex(0))

	path = path.With(
		js.NewPath(
			js.PropertyName("low"),
			js.ArrayIndex(1),
			js.PropertyName("property"),
		),
	)

	assert.Equal(t, "top[0].low[1].property", path.String())
}

func TestPropertyPath_Elements(t *testing.T) {
	want := []js.PathElement{
		js.PropertyName("top"),
		js.ArrayIndex(0),
		js.PropertyName("low"),
		js.ArrayIndex(1),
		js.PropertyName("property"),
	}

	got := js.NewPath(want...).Elements()

	assert.Equal(t, want, got)
}

func TestPropertyPath_String(t *testing.T) {
	tests := []struct {
		path *js.Path
		want string
	}{
		{path: nil, want: ""},
		{path: js.NewPath(), want: ""},
		{path: js.NewPath(js.PropertyName(" ")), want: "[' ']"},
		{path: js.NewPath(js.PropertyName("$")), want: "$"},
		{path: js.NewPath(js.PropertyName("_")), want: "_"},
		{path: js.NewPath(js.PropertyName("id$_")), want: "id$_"},
		{
			path: js.NewPath().WithProperty("array").WithIndex(1).WithProperty("property"),
			want: "array[1].property",
		},
		{
			path: js.NewPath().WithProperty("@foo").WithProperty("bar"),
			want: "['@foo'].bar",
		},
		{
			path: js.NewPath().WithProperty("@foo").WithIndex(0),
			want: "['@foo'][0]",
		},
		{
			path: js.NewPath().WithProperty("foo.bar").WithProperty("baz"),
			want: "['foo.bar'].baz",
		},
		{
			path: js.NewPath().WithProperty("foo.'bar'").WithProperty("baz"),
			want: `['foo.\'bar\''].baz`,
		},
		{
			path: js.NewPath().WithProperty(`0`).WithProperty("baz"),
			want: `['0'].baz`,
		},
		{
			path: js.NewPath().WithProperty(`foo[0]`).WithProperty("baz"),
			want: `['foo[0]'].baz`,
		},
		{
			path: js.NewPath().WithProperty(``).WithProperty("baz"),
			want: `[''].baz`,
		},
		{
			path: js.NewPath().WithProperty("foo").WithProperty(""),
			want: `foo['']`,
		},
		{
			path: js.NewPath().WithProperty(`'`).WithProperty("baz"),
			want: `['\''].baz`,
		},
		{
			path: js.NewPath().WithProperty(`\`).WithProperty("baz"),
			want: `['\\'].baz`,
		},
		{
			path: js.NewPath().WithProperty(`\'foo`).WithProperty("baz"),
			want: `['\\\'foo'].baz`,
		},
		{
			path: js.NewPath().WithProperty(`фу`).WithProperty("baz"),
			want: `фу.baz`,
		},
	}
	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			got := test.path.String()

			assert.Equal(t, test.want, got)
		})
	}
}

func TestPathFromAny(t *testing.T) {
	tests := []struct {
		args []interface{}
		want string
	}{
		{[]interface{}{"foo", "bar", "baz"}, "foo.bar.baz"},
		{[]interface{}{"foo", String("stringer")}, "foo.stringer"},
		{[]interface{}{"foo", 0, "bar", 1, "baz"}, "foo[0].bar[1].baz"},
		{[]interface{}{uint8(1), uint16(2), uint32(3), uint64(4), uint(5)}, "[1][2][3][4][5]"},
		{[]interface{}{int8(1), int16(2), int32(3), int64(4), int(5)}, "[1][2][3][4][5]"},
	}
	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			path, err := js.PathFromAny(test.args...)

			require.NoError(t, err)
			assert.Equal(t, test.want, path.String())
		})
	}
}

func TestPathFromAny_Error(t *testing.T) {
	path, err := js.PathFromAny(123.123)

	assert.Nil(t, path)
	assert.EqualError(t, err, "invalid path: should contain only strings and numbers")
}

type String string

func (s String) String() string {
	return string(s)
}
