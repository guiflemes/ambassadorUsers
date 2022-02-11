package logic

import (
	"reflect"
	"testing"
	"users/src/domain"
	repo "users/src/repositories"
)

func Test_loginService_Authenticate(t *testing.T) {
	type fields struct {
		loginRepo repo.UserRepository
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		want1   *domain.User
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginService{
				loginRepo: tt.fields.loginRepo,
			}
			got, got1, err := l.Authenticate(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginService.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loginService.Authenticate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("loginService.Authenticate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
