package presenter

import (
	"reflect"
	"testing"

	appErrors "pearshop_backend/app/errors"
)

func Test_fromSystemErrors(t *testing.T) {
	type args struct {
		errs appErrors.SystemErrors
	}

	tests := []struct {
		name string
		args args
		want []responseError
	}{
		{
			name: "input is nil",
			args: args{
				errs: nil,
			},
			want: nil,
		},
		{
			name: "input len is 0",
			args: args{
				errs: []appErrors.SystemError{},
			},
			want: nil,
		},
		{
			name: "one error",
			args: args{
				errs: []appErrors.SystemError{
					appErrors.NewInvalidArgumentErr(appErrors.CodeInvalidPayload, "invalid payload", nil),
				},
			},
			want: []responseError{
				{
					Type:    string(appErrors.TypeInvalidArgument),
					Code:    string(appErrors.CodeInvalidPayload),
					Message: "invalid payload",
					Param:   nil,
				},
			},
		},
		{
			name: "multiple errors",
			args: args{
				errs: []appErrors.SystemError{
					appErrors.NewInvalidArgumentErr(appErrors.CodeInvalidPayload, "invalid payload", nil),
					appErrors.NewInvalidArgumentErr(appErrors.CodeInvalidPayload, "invalid payload", nil),
				},
			},
			want: []responseError{
				{
					Type:    string(appErrors.TypeInvalidArgument),
					Code:    string(appErrors.CodeInvalidPayload),
					Message: "invalid payload",
					Param:   nil,
				},
				{
					Type:    string(appErrors.TypeInvalidArgument),
					Code:    string(appErrors.CodeInvalidPayload),
					Message: "invalid payload",
					Param:   nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fromSystemErrors(tt.args.errs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromSystemErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
