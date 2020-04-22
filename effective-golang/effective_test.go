package effective

import (
	"testing"
)

func TestSwitch(t *testing.T) {

}

func Test_switchType(t *testing.T) {
	type args struct {
		a interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				a: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switchType(tt.args.a)
		})
	}
}

func Test_aboutRange(t *testing.T) {
	aboutRange()
}
