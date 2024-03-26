package checkertype

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCheckType(t *testing.T) {
	type args struct {
		value any
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with int",
			args: args{value: 10},
			want: "int",
		},
		{
			name: "Test with float",
			args: args{value: 10.0},
			want: "float",
		},
		{
			name: "Test with string",
			args: args{value: "10.0"},
			want: "string",
		},
		{
			name: "Test with boolean",
			args: args{value: true},
			want: "boolean",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckType(tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}
