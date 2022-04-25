package image

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/Avielyo10/edge-api/internal/edge/domain/common"
)

func TestImage_AddTag(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		uuid         string
		name         common.Name
		description  string
		timing       common.Time
		status       Status
		version      Version
		distribution Distribution
		user         User
		packages     Packages
		repos        Repos
		installer    Installer
		outputType   []OutputType
		tags         common.Tags
	}
	type args struct {
		tag []common.Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add tag",
			fields: fields{
				tags: common.NewTags("tag1", "tag2"),
			},
			args: args{
				tag: []common.Tag{
					common.NewTag("tag3"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &Image{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				uuid:         tt.fields.uuid,
				name:         tt.fields.name,
				description:  tt.fields.description,
				timing:       tt.fields.timing,
				status:       tt.fields.status,
				version:      tt.fields.version,
				distribution: tt.fields.distribution,
				user:         tt.fields.user,
				packages:     tt.fields.packages,
				repos:        tt.fields.repos,
				installer:    tt.fields.installer,
				outputType:   tt.fields.outputType,
				tags:         tt.fields.tags,
			}
			i.AddTag(tt.args.tag...)
			if !reflect.DeepEqual(append(tt.fields.tags.Tags(), tt.args.tag...), i.tags.Tags()) {
				t.Errorf("AddTag() = %v, want %v", i.tags.Tags(), append(tt.fields.tags.Tags(), tt.args.tag...))
			}
		})
	}
}

func TestImage_RemoveTag(t *testing.T) {
	type fields struct {
		ctx          context.Context
		cancel       context.CancelFunc
		uuid         string
		name         common.Name
		description  string
		timing       common.Time
		status       Status
		version      Version
		distribution Distribution
		user         User
		packages     Packages
		repos        Repos
		installer    Installer
		outputType   []OutputType
		tags         common.Tags
	}
	type args struct {
		tag []common.Tag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "remove tag",
			fields: fields{
				tags: common.NewTags("tag1", "tag2"),
			},
			args: args{
				tag: []common.Tag{
					common.NewTag("tag1"),
				},
			},
		},
		{
			name: "remove tag, none exist",
			fields: fields{
				tags: common.NewTags("tag1", "tag2"),
			},
			args: args{
				tag: []common.Tag{
					common.NewTag("tag3"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &Image{
				ctx:          tt.fields.ctx,
				cancel:       tt.fields.cancel,
				uuid:         tt.fields.uuid,
				name:         tt.fields.name,
				description:  tt.fields.description,
				timing:       tt.fields.timing,
				status:       tt.fields.status,
				version:      tt.fields.version,
				distribution: tt.fields.distribution,
				user:         tt.fields.user,
				packages:     tt.fields.packages,
				repos:        tt.fields.repos,
				installer:    tt.fields.installer,
				outputType:   tt.fields.outputType,
				tags:         tt.fields.tags,
			}
			i.RemoveTag(tt.args.tag...)
			tt.fields.tags.Remove(tt.args.tag...)
			if !string_array_sorted_equal(tt.fields.tags.StringArray(), i.tags.StringArray()) {
				t.Errorf("RemoveTag() = %v, want %v", i.tags.StringArray(), tt.fields.tags.StringArray())
			}
		})
	}
}

func string_array_sorted_equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]string, len(a))
	b_copy := make([]string, len(b))

	copy(a_copy, a)
	copy(b_copy, b)

	sort.Strings(a_copy)
	sort.Strings(b_copy)

	return reflect.DeepEqual(a_copy, b_copy)
}
