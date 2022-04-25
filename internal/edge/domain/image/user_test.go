package image

import (
	"reflect"
	"testing"

	"github.com/Avielyo10/edge-api/internal/common/models"
	"github.com/bxcodec/faker/v3"
)

func TestNewUser(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"
	type args struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "valid user",
			args: args{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: User{
				username: "valid",
				sshKey:   validSSHKey,
			},
			wantErr: false,
		},
		{
			name: "invalid username",
			args: args{
				username: "",
				sshKey:   validSSHKey,
			},
			want:    User{},
			wantErr: true,
		},
		{
			name: "invalid ssh key",
			args: args{
				username: "valid",
				sshKey:   faker.Word(),
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUser(tt.args.username, tt.args.sshKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ValidSSHKey(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid ssh key",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: true,
		},
		{
			name: "invalid ssh key",
			fields: fields{
				username: "valid",
				sshKey:   faker.Word(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.ValidSSHKey(); got != tt.want {
				t.Errorf("User.ValidSSHKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ValidUsername(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid username",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: true,
		},
		{
			name: "invalid username",
			fields: fields{
				username: "",
				sshKey:   validSSHKey,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.ValidUsername(); got != tt.want {
				t.Errorf("User.ValidUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsZero(t *testing.T) {
	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero user",
			fields: fields{
				username: "",
				sshKey:   "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.IsZero(); got != tt.want {
				t.Errorf("User.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidUser(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type args struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid user",
			args: args{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: true,
		},
		{
			name: "invalid user",
			args: args{
				username: "",
				sshKey:   validSSHKey,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isValidUser(tt.args.username, tt.args.sshKey); got != tt.want {
				t.Errorf("isValidUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Username(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid user",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: "valid",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.Username(); got != tt.want {
				t.Errorf("User.Username() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SSHKey(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "valid user",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: validSSHKey,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.SSHKey(); got != tt.want {
				t.Errorf("User.SSHKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_MarshalJSON(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid user",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want:    []byte(`{"username":"valid","ssh_key":"` + validSSHKey + `"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			got, err := u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestUser_UnmarshalJSON(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid user",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			args: args{
				data: []byte(`{"username":"valid","ssh_key":"` + validSSHKey + `"}`),
			},
			wantErr: false,
		},
		{
			name: "invalid user",
			fields: fields{
				username: "",
				sshKey:   validSSHKey,
			},
			args: args{
				data: []byte(`{"username":"","ssh_key":"` + validSSHKey + `"}`),
			},
			wantErr: true,
		},
		{
			name: "invalid JSON",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			args: args{
				data: []byte(`{""}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if err := u.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("User.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_MarshalGorm(t *testing.T) {
	validSSHKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFjRxF1E73z1K9AjltDkuJyGUW3YluTEAW6PvHEZH6vnzNHI+cut716lGGRFHlYk1Fk51Q/92ZlynJ/HqByaK/MJppkQSL4x3KEm6s5ciwXbVEb3ct4waTgqxPD9gy7NN0uzbrhQMillb50yZgox6d9A/JmyRA1Dlai/esrlKfZ4wtSUl+CMsPoVxC6pIsh1YqUWE7S/dvXsQ8V+O7H0sdXAkZMg09kLUOQe3fliTMg6wppW+tb30g4MWAbHSrXksL1TpYjmP0M+stNetO2EIZ07bc8KpQhZybdM8LUhhPGuZXuKzIlwbkDI7C1yLv574wOYCjG/zk7Zu9qO7p6u8x valid@sshkey"

	type fields struct {
		username string
		sshKey   string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.User
	}{
		{
			name: "valid user",
			fields: fields{
				username: "valid",
				sshKey:   validSSHKey,
			},
			want: &models.User{
				Name:   "valid",
				SSHKey: validSSHKey,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := User{
				username: tt.fields.username,
				sshKey:   tt.fields.sshKey,
			}
			if got := u.MarshalGorm(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.MarshalGorm() = %v, want %v", got, tt.want)
			}
		})
	}
}
