package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pearshop_backend/app/domain/entity"
)

func TestPagingRequest_StructName(t *testing.T) {
	type fields struct {
		Size *int
		Page *int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Size: nil,
				Page: nil,
			},
			want: "PagingRequest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PagingRequest{
				Size: tt.fields.Size,
				Page: tt.fields.Page,
			}.StructName()

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPagingRequest_Form(t *testing.T) {
	type fields struct {
		Size *int
		Page *int
	}

	tests := []struct {
		name   string
		fields fields
		want   entity.PagingRequest
	}{
		{
			name: "no input",
			fields: fields{
				Size: nil,
				Page: nil,
			},
			want: entity.PagingRequest{
				Size: 30,
				Page: 1,
			},
		},
		{
			name: "has size only",
			fields: fields{
				Size: func() *int {
					i := 100

					return &i
				}(),
				Page: nil,
			},
			want: entity.PagingRequest{
				Size: 100,
				Page: 1,
			},
		},
		{
			name: "has page only",
			fields: fields{
				Size: nil,
				Page: func() *int {
					i := 20

					return &i
				}(),
			},
			want: entity.PagingRequest{
				Size: 30,
				Page: 20,
			},
		},
		{
			name: "has both",
			fields: fields{
				Size: func() *int {
					i := 100

					return &i
				}(),
				Page: func() *int {
					i := 20

					return &i
				}(),
			},
			want: entity.PagingRequest{
				Size: 100,
				Page: 20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PagingRequest{
				Size: tt.fields.Size,
				Page: tt.fields.Page,
			}
			got := p.Form()

			assert.Equal(t, tt.want, got)
		})
	}
}
