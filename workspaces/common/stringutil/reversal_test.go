// Package strings strings
package stringutil

import (
	"testing"
)

func TestReversal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english",
			args: args{
				str: "hello world",
			},
			want: "dlrow olleh",
		},
		{
			name: "chinese",
			args: args{
				str: "你好，世界",
			},
			want: "界世，好你",
		},
		{
			name: "empty",
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reversal(tt.args.str); got != tt.want {
				t.Errorf("Reversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
