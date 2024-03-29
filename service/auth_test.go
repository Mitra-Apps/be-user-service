package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/Mitra-Apps/be-user-service/domain/user/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestNewAuthClient(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name string
		args args
		want Authentication
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthClient(tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authClient_GenerateToken(t *testing.T) {
	type args struct {
		ctx           context.Context
		user          *entity.User
		expiredMinute int
	}
	tests := []struct {
		name      string
		c         *authClient
		args      args
		wantToken string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := tt.c.GenerateToken(tt.args.ctx, tt.args.user, tt.args.expiredMinute)
			if (err != nil) != tt.wantErr {
				t.Errorf("authClient.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("authClient.GenerateToken() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func Test_authClient_ValidateToken(t *testing.T) {
	auth := NewAuthClient("secret")
	user := &entity.User{
		Id: uuid.MustParse("b70a2a5e-bbd2-4000-96c0-aaa533b8236f"),
		Roles: []entity.Role{
			{
				RoleName: "merchant",
			},
			{
				RoleName: "customer",
			},
			{
				RoleName: "admin",
			},
		},
	}
	token, err := auth.GenerateToken(context.Background(), user, 60)
	if err != nil {
		panic(err.Error())
	}
	type args struct {
		ctx          context.Context
		requestToken string
	}
	tests := []struct {
		name    string
		c       *authClient
		args    args
		want    *JwtCustomClaim
		wantErr bool
	}{
		{
			name: "success",
			c: &authClient{
				secret: "secret",
			},
			args: args{
				ctx:          context.Background(),
				requestToken: token,
			},
			want: &JwtCustomClaim{
				Roles: []string{
					"merchant",
					"customer",
					"admin",
				},
				RegisteredClaims: jwt.RegisteredClaims{
					Subject: "b70a2a5e-bbd2-4000-96c0-aaa533b8236f",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ValidateToken(tt.args.ctx, tt.args.requestToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("authClient.ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if !reflect.DeepEqual(got.Roles, tt.want.Roles) {
					t.Errorf("authClient.ValidateToken() = %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(got.Subject, tt.want.Subject) {
					t.Errorf("authClient.ValidateToken() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
