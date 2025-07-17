package model

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   PersonalInfo
		wantErr bool
	}{
		{
			name: "valid data",
			input: PersonalInfo{
				Firstname: "firstname",
				Lastname:  "lastname",
				Email:     "firstname@example.com",
			},
			wantErr: false,
		},
		{
			name: "missing firstname",
			input: PersonalInfo{
				Firstname: "",
				Lastname:  "lastname",
				Email:     "firstname@example.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			input: PersonalInfo{
				Firstname: "firstname",
				Lastname:  "lastname",
				Email:     "notanemail",
			},
			wantErr: true,
		},
		{
			name: "email is empty",
			input: PersonalInfo{
				Firstname: "firstname",
				Lastname:  "lastname",
				Email:     "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
