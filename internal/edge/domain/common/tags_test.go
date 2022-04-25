package common

import (
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

func TestNewTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		{
			name: "valid tag",
			args: args{
				tag: "tag",
			},
			want: Tag{
				tag: "tag",
			},
		},
		{
			name: "invalid tag",
			args: args{
				tag: "",
			},
			want: Tag{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewTag(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTags(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want Tags
	}{
		{
			name: "valid tags",
			args: args{
				tags: []string{"tag1", "tag2"},
			},
			want: Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
		},
		{
			name: "invalid tags",
			args: args{
				tags: []string{},
			},
			want: Tags{},
		},
		{
			name: "invalid tags, empty strings",
			args: args{
				tags: []string{"", ""},
			},
			want: Tags{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewTags(tt.args.tags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_IsZero(t *testing.T) {
	tests := []struct {
		name string
		tr   Tag
		want bool
	}{
		{
			name: "empty tag",
			tr:   Tag{},
			want: true,
		},
		{
			name: "valid tag",
			tr:   Tag{tag: "tag"},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.IsZero(); got != tt.want {
				t.Errorf("Tag.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidTag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid tag",
			args: args{
				name: "tag",
			},
			want: true,
		},
		{
			name: "invalid tag",
			args: args{
				name: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidTag(tt.args.name); got != tt.want {
				t.Errorf("isValidTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_Add(t *testing.T) {
	type args struct {
		tags []Tag
	}
	tests := []struct {
		name string
		tr   *Tags
		args args
	}{
		{
			name: "add tags",
			tr:   &Tags{},
			args: args{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.tr.Add(tt.args.tags...)
			if !reflect.DeepEqual(tt.tr.tags, tt.args.tags) {
				t.Errorf("Tags.Add() = %v, want %v", len(tt.tr.tags), len(tt.args.tags))
			}
		})
	}
}

func TestTags_Remove(t *testing.T) {
	type args struct {
		tags []Tag
	}
	tests := []struct {
		name string
		tr   *Tags
		args args
	}{
		{
			name: "remove tags",
			tr: &Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			args: args{
				tags: []Tag{
					{
						tag: "tag1",
					},
				},
			},
		},
		{
			name: "remove tags, with empty tag",
			tr: &Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			args: args{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.tr.Remove(tt.args.tags...)
			if !reflect.DeepEqual(tt.tr.tags, []Tag{{"tag2"}}) {
				t.Errorf("Tags.Remove() = %v, want %v", tt.tr.tags, tt.args.tags)
			}
		})
	}
}

func TestTag_String(t *testing.T) {
	tests := []struct {
		name string
		tr   Tag
		want string
	}{
		{
			name: "valid tag",
			tr:   Tag{tag: "tag"},
			want: "tag",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("Tag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_StringArray(t *testing.T) {
	tests := []struct {
		name string
		tr   Tags
		want []string
	}{
		{
			name: "valid tags",
			tr: Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			want: []string{"tag1", "tag2"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.StringArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags.StringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_Tags(t *testing.T) {
	tests := []struct {
		name string
		tr   Tags
		want []Tag
	}{
		{
			name: "valid tags",
			tr: Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			want: []Tag{
				{
					tag: "tag1",
				},
				{
					tag: "tag2",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.Tags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags.Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tr      Tags
		want    []byte
		wantErr bool
	}{
		{
			name: "valid tags",
			tr: Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			want:    []byte(`["tag1","tag2"]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tags.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		tr      *Tags
		args    args
		wantErr bool
	}{
		{
			name: "valid tags",
			tr: &Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
				},
			},
			args: args{
				data: []byte(`["tag1","tag2"]`),
			},
			wantErr: false,
		},
		{
			name: "valid tags, with empty",
			tr: &Tags{
				tags: []Tag{
					{
						tag: "tag1",
					},
					{
						tag: "tag2",
					},
					{
						tag: "",
					},
				},
			},
			args: args{
				data: []byte(`["tag1","tag2"]`),
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			tr: &Tags{
				tags: []Tag{},
			},
			args: args{
				data: []byte(`""`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.tr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Tags.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTags_MarshalGorm(t *testing.T) {
	type args struct {
		account string
	}
	tests := []struct {
		name string
		tr   Tags
		args args
		want []models.Tag
	}{
		{
			name: "valid tags",
			tr: Tags{
				tags: []Tag{
					{"tag1"},
					{"tag2"},
				},
			},
			args: args{
				account: "test",
			},
			want: []models.Tag{
				{Account: "test", Name: "tag1"},
				{Account: "test", Name: "tag2"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.MarshalGorm(tt.args.account); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}
