package validator

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
)

func TestValidateRegisterJSONBody(t *testing.T) {
	type args struct {
		param generated.PostRegisterJSONBody
	}
	tests := []struct {
		name    string
		args    args
		mock    func() (string, string, string)
		wantErr bool
	}{
		{
			name: "pass",
			mock: func() (string, string, string) {
				var fullName, phoneNumber, password string
				fullName = "testtest"
				phoneNumber = "+6281123123"
				password = "Password123!"
				return fullName, phoneNumber, password
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullName, phoneNumber, password := tt.mock()
			tt.args.param = generated.PostRegisterJSONBody{
				FullName:    &fullName,
				PhoneNumber: &phoneNumber,
				Password:    &password,
			}
			if err := ValidateRegisterJSONBody(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("ValidateRegisterJSONBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
