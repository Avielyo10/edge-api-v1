package image

import (
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
)

func TestNewRepo(t *testing.T) {
	type args struct {
		name string
		url  string
	}
	tests := []struct {
		name string
		args args
		want *Repo
	}{
		{
			name: "valid",
			args: args{
				name: "test",
				url:  "http://test.com",
			},
			want: &Repo{
				name: "test",
				url:  "http://test.com",
			},
		},
		{
			name: "invalid, no name",
			args: args{
				name: "",
				url:  "http://test.com",
			},
			want: nil,
		},
		{
			name: "invalid, invalid url",
			args: args{
				name: "test",
				url:  "test",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewRepo(tt.args.name, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRepos(t *testing.T) {
	type args struct {
		repos []*Repo
	}
	tests := []struct {
		name string
		args args
		want Repos
	}{
		{
			name: "valid",
			args: args{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
				},
			},
			want: Repos{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
				},
			},
		},
		{
			name: "invalid, no repos",
			args: args{
				repos: []*Repo{},
			},
			want: Repos{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewRepos(tt.args.repos...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidRepo(t *testing.T) {
	type args struct {
		name string
		url  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				name: "test",
				url:  "http://test.com",
			},
			want: true,
		},
		{
			name: "invalid, no name",
			args: args{
				name: "",
				url:  "http://test.com",
			},
			want: false,
		},
		{
			name: "invalid, invalid url",
			args: args{
				name: "test",
				url:  "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidRepo(tt.args.name, tt.args.url); got != tt.want {
				t.Errorf("isValidRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_IsZero(t *testing.T) {
	type fields struct {
		name string
		url  string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid",
			fields: fields{
				name: "test",
				url:  "http://test.com",
			},
			want: false,
		},
		{
			name: "invalid, no name",
			fields: fields{
				name: "",
				url:  "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repo{
				name: tt.fields.name,
				url:  tt.fields.url,
			}
			if got := r.IsZero(); got != tt.want {
				t.Errorf("Repo.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepos_Add(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	type args struct {
		repos []*Repo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Repos
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
				},
			},
			args: args{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
				},
			},
			want: &Repos{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test", "http://test.com"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &Repos{
				repos: tt.fields.repos,
			}
			r.Add(tt.args.repos...)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Repos.Add() = %v, want %v", r, tt.want)
			}
		})
	}
}

func TestRepo_String(t *testing.T) {
	type fields struct {
		name string
		url  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid",
			fields: fields{
				name: "test",
				url:  "http://test.com",
			},
			want: `{"name":"test","url":"http://test.com"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repo{
				name: tt.fields.name,
				url:  tt.fields.url,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Repo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepos_StringArray(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
				},
			},
			want: []string{`{"name":"test","url":"http://test.com"}`},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repos{
				repos: tt.fields.repos,
			}
			if got := r.StringArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repos.StringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_MarshalJSON(t *testing.T) {
	type fields struct {
		name string
		url  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				name: "test",
				url:  "http://test.com",
			},
			want:    []byte(`{"name":"test","url":"http://test.com"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repo{
				name: tt.fields.name,
				url:  tt.fields.url,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_UnmarshalJSON(t *testing.T) {
	type fields struct {
		name string
		url  string
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				name: "test",
				url:  "http://test.com",
			},
			args: args{
				b: []byte(`{"name":"test","url":"http://test.com"}`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				name: "test",
				url:  "http://test.com",
			},
			args: args{
				b: []byte(`{""}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &Repo{
				name: tt.fields.name,
				url:  tt.fields.url,
			}
			if err := r.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Repo.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepos_MarshalJSON(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			want:    []byte(`[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repos{
				repos: tt.fields.repos,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repos.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repos.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestRepos_UnmarshalJSON(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			args: args{
				b: []byte(`[{"name":"test","url":"http://test.com"},{"name":"test2","url":"http://test2.com"}]`),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			args: args{
				b: []byte(`[{""}]`),
			},
			wantErr: true,
		},
		{
			name: "empty",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			args: args{
				b: []byte(`[]`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &Repos{
				repos: tt.fields.repos,
			}
			if err := r.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Repos.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReposFromArray(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	type args struct {
		repos []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			args: args{
				repos: []interface{}{
					map[string]interface{}{
						"name": "test",
						"url":  "http://test.com",
					},
					map[string]interface{}{
						"name": "test2",
						"url":  "http://test2.com",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid, unmarshal error",
			fields: fields{
				repos: nil,
			},
			args: args{
				repos: []interface{}{""},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := &Repos{
				repos: tt.fields.repos,
			}
			repos, err := ReposFromArray(tt.args.repos)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repos.FromArray() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(repos, r.repos) {
				t.Errorf("Repos.FromArray() = %v, want %v", repos, r.repos)
			}
		})
	}
}

func TestRepos_MarshalGorm(t *testing.T) {
	type fields struct {
		repos []*Repo
	}
	type args struct {
		account string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Repo
	}{
		{
			name: "valid",
			fields: fields{
				repos: []*Repo{
					NewRepo("test", "http://test.com"),
					NewRepo("test2", "http://test2.com"),
				},
			},
			args: args{
				account: "test",
			},
			want: []models.Repo{
				{
					Account: "test",
					Name:    "test",
					URL:     "http://test.com",
				},
				{
					Account: "test",
					Name:    "test2",
					URL:     "http://test2.com",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Repos{
				repos: tt.fields.repos,
			}
			if got := r.MarshalGorm(tt.args.account); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repos.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}
