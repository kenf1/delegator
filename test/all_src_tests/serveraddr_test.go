package testauth

import (
	"os"
	"testing"

	"github.com/kenf1/delegator/src/configs"
	"github.com/kenf1/delegator/src/models"
)

func TestImportServerAddrWrapper(t *testing.T) {
	const envFile = ".env.test"

	//create temp .env, rm after each test
	writeEnvFile := func(contents string) {
		err := os.WriteFile(envFile, []byte(contents), 0644)
		if err != nil {
			t.Fatalf("failed to write env file: %v", err)
		}
	}
	defer os.Remove(envFile)

	tests := []struct {
		name         string
		remoteEnv    map[string]string
		localEnvFile string
		want         models.ServerAddr
		wantErr      bool
	}{
		{
			name:      "remote env vars present",
			remoteEnv: map[string]string{"HOST": "remotehost", "PORT": "8080"},
			want:      models.ServerAddr{Host: "remotehost", Port: "8080"},
			wantErr:   false,
		},
		{
			name:         "remote env missing, local env file loads",
			remoteEnv:    nil,
			localEnvFile: "HOST=localhost\nPORT=9090\n",
			want:         models.ServerAddr{Host: "localhost", Port: "9090"},
			wantErr:      false,
		},
		{
			name:         "remote env missing, local env file missing HOST",
			remoteEnv:    nil,
			localEnvFile: "PORT=9090\n",
			wantErr:      true,
		},
		{
			name:         "remote env missing HOST, local env present",
			remoteEnv:    map[string]string{"PORT": "8080"},
			localEnvFile: "HOST=localhost\nPORT=9090\n",
			want:         models.ServerAddr{Host: "localhost", Port: "9090"},
			wantErr:      false,
		},
		{
			name:         "remote env missing PORT, local env missing",
			remoteEnv:    map[string]string{"HOST": "remotehost"},
			localEnvFile: "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//clear env before each test
			os.Unsetenv("HOST")
			os.Unsetenv("PORT")

			//set remote env vars
			for k, v := range tt.remoteEnv {
				os.Setenv(k, v)
			}

			//ensure only nec .env present
			if tt.localEnvFile != "" {
				writeEnvFile(tt.localEnvFile)
			} else {
				os.Remove(envFile)
			}

			got, err := configs.ImportServerAddrWrapper(envFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("ImportServerAddrWrapper() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got != tt.want {
				t.Errorf("ImportServerAddrWrapper() = %+v, want %+v", got, tt.want)
			}

			os.Unsetenv("HOST")
			os.Unsetenv("PORT")
			os.Remove(envFile)
		})
	}
}
