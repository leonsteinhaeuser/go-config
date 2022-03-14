package config

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type ExampleConfigA struct {
	Name     string
	Age      int
	Size     float64
	IsActive bool
	Uint     int64
	Hosts    []string
	Children ExampleConfigB
}

type ExampleConfigB struct {
	Name     string
	Age      int
	Size     float64
	IsActive bool
}

func TestAutoloadAndEnrichConfigWithEnvPrefix(t *testing.T) {
	type args struct {
		filePath string
		receiver interface{}
		prefix   string
	}
	tests := []struct {
		name     string
		args     args
		preFunc  func() error
		postFunc func() error
		want     *ExampleConfigA
		wantErr  bool
	}{
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			want:    &ExampleConfigA{},
			wantErr: true,
		},

		// with env vars
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			preFunc: func() error {
				os.Setenv("AAA_NAME", "emil")
				os.Setenv("AAA_AGE", "100")
				os.Setenv("AAA_SIZE", "10.1")
				os.Setenv("AAA_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("AAA_NAME")
				os.Unsetenv("AAA_AGE")
				os.Unsetenv("AAA_SIZE")
				os.Unsetenv("AAA_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			preFunc: func() error {
				os.Setenv("AAA_NAME", "emil")
				os.Setenv("AAA_AGE", "100")
				os.Setenv("AAA_SIZE", "10.1")
				os.Setenv("AAA_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("AAA_NAME")
				os.Unsetenv("AAA_AGE")
				os.Unsetenv("AAA_SIZE")
				os.Unsetenv("AAA_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			preFunc: func() error {
				os.Setenv("AAA_NAME", "emil")
				os.Setenv("AAA_AGE", "100")
				os.Setenv("AAA_SIZE", "10.1")
				os.Setenv("AAA_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("AAA_NAME")
				os.Unsetenv("AAA_AGE")
				os.Unsetenv("AAA_SIZE")
				os.Unsetenv("AAA_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
				receiver: &ExampleConfigA{},
				prefix:   "aaa",
			},
			preFunc: func() error {
				os.Setenv("AAA_NAME", "emil")
				os.Setenv("AAA_AGE", "100")
				os.Setenv("AAA_SIZE", "10.1")
				os.Setenv("AAA_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("AAA_NAME")
				os.Unsetenv("AAA_AGE")
				os.Unsetenv("AAA_SIZE")
				os.Unsetenv("AAA_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
				receiver: &ExampleConfigA{},
			},
			want:    &ExampleConfigA{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AutoloadAndEnrichConfigWithEnvPrefix(tt.args.filePath, tt.args.prefix, tt.args.receiver); (err != nil) != tt.wantErr {
				t.Errorf("AutoloadAndEnrichConfigWithEnvPrefix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAutoloadAndEnrichConfig(t *testing.T) {
	type args struct {
		filePath string
		receiver interface{}
	}
	tests := []struct {
		name     string
		args     args
		preFunc  func() error
		postFunc func() error
		want     *ExampleConfigA
		wantErr  bool
	}{
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
				receiver: &ExampleConfigA{},
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
				receiver: &ExampleConfigA{},
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
				receiver: &ExampleConfigA{},
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
				receiver: &ExampleConfigA{},
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
				receiver: &ExampleConfigA{},
			},
			want:    &ExampleConfigA{},
			wantErr: true,
		},

		// with env vars
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
				receiver: &ExampleConfigA{},
			},
			preFunc: func() error {
				os.Setenv("CFG_NAME", "emil")
				os.Setenv("CFG_AGE", "100")
				os.Setenv("CFG_SIZE", "10.1")
				os.Setenv("CFG_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("CFG_NAME")
				os.Unsetenv("CFG_AGE")
				os.Unsetenv("CFG_SIZE")
				os.Unsetenv("CFG_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
				receiver: &ExampleConfigA{},
			},
			preFunc: func() error {
				os.Setenv("CFG_NAME", "emil")
				os.Setenv("CFG_AGE", "100")
				os.Setenv("CFG_SIZE", "10.1")
				os.Setenv("CFG_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("CFG_NAME")
				os.Unsetenv("CFG_AGE")
				os.Unsetenv("CFG_SIZE")
				os.Unsetenv("CFG_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
				receiver: &ExampleConfigA{},
			},
			preFunc: func() error {
				os.Setenv("CFG_NAME", "emil")
				os.Setenv("CFG_AGE", "100")
				os.Setenv("CFG_SIZE", "10.1")
				os.Setenv("CFG_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("CFG_NAME")
				os.Unsetenv("CFG_AGE")
				os.Unsetenv("CFG_SIZE")
				os.Unsetenv("CFG_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
				receiver: &ExampleConfigA{},
			},
			preFunc: func() error {
				os.Setenv("CFG_NAME", "emil")
				os.Setenv("CFG_AGE", "100")
				os.Setenv("CFG_SIZE", "10.1")
				os.Setenv("CFG_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("CFG_NAME")
				os.Unsetenv("CFG_AGE")
				os.Unsetenv("CFG_SIZE")
				os.Unsetenv("CFG_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "emil",
				Age:      100,
				Size:     10.1,
				IsActive: false,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
				receiver: &ExampleConfigA{},
			},
			want:    &ExampleConfigA{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preFunc != nil {
				tt.preFunc()
			}

			if err := AutoloadAndEnrichConfig(tt.args.filePath, tt.args.receiver); (err != nil) != tt.wantErr {
				t.Errorf("AutoloadAndEnrichConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.args.receiver, tt.want)
			if diff != "" {
				t.Errorf("AutoloadAndEnrichConfig() diff = %v", diff)
			}

			if tt.postFunc != nil {
				tt.postFunc()
			}
		})
	}
}

func Test_detectFormat(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want format
	}{
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
			},
			want: YAML,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
			},
			want: JSON,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
			},
			want: TOML,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
			},
			want: HCL,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectFormat(tt.args.filePath); got != tt.want {
				t.Errorf("detectFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadAndParseFile(t *testing.T) {
	type args struct {
		filePath string
		receiver interface{}
		f        format
	}
	tests := []struct {
		name    string
		args    args
		want    *ExampleConfigA
		wantErr bool
	}{
		{
			name: "yaml",
			args: args{
				filePath: ".file/simple.yml",
				receiver: &ExampleConfigA{},
				f:        YAML,
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "json",
			args: args{
				filePath: ".file/simple.json",
				receiver: &ExampleConfigA{},
				f:        JSON,
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "toml",
			args: args{
				filePath: ".file/simple.toml",
				receiver: &ExampleConfigA{},
				f:        TOML,
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "hcl",
			args: args{
				filePath: ".file/simple.hcl",
				receiver: &ExampleConfigA{},
				f:        HCL,
			},
			want: &ExampleConfigA{
				Name:     "Simple Sam",
				Age:      25,
				Size:     1.87,
				IsActive: true,
				Uint:     8,
				Hosts:    []string{"localhost", "127.0.0.1"},
				Children: ExampleConfigB{
					Name:     "Chris Sam",
					Age:      3,
					Size:     0.87,
					IsActive: true,
				},
			},
			wantErr: false,
		},
		{
			name: "unsupported type",
			args: args{
				filePath: ".file/simple.usu",
				receiver: &ExampleConfigA{},
				f:        "usu",
			},
			want:    &ExampleConfigA{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := loadAndParseFile(tt.args.filePath, tt.args.receiver, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadAndParseFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			diff := cmp.Diff(tt.args.receiver, tt.want)
			if diff != "" {
				t.Errorf("loadAndParseFile() diff = %s\n", diff)
			}
		})
	}
}

func Test_prefixString(t *testing.T) {
	type args struct {
		prefix    string
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty prefix",
			args: args{
				prefix:    "",
				fieldName: "field",
			},
			want: "FIELD",
		},
		{
			name: "with prefix",
			args: args{
				prefix:    "second",
				fieldName: "field",
			},
			want: "SECOND_FIELD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixString(tt.args.prefix, tt.args.fieldName); got != tt.want {
				t.Errorf("prefixString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readStructAndEnrichWithEnv(t *testing.T) {
	type args struct {
		st     interface{}
		prefix string
	}
	tests := []struct {
		name     string
		args     args
		preFunc  func() error
		postFunc func() error
		want     *ExampleConfigA
	}{
		{
			name: "without env settings",
			args: args{
				st: &ExampleConfigA{
					Name:     "John",
					Age:      30,
					Size:     1.78,
					IsActive: true,
					Children: ExampleConfigB{
						Name:     "John",
						Age:      30,
						Size:     1.78,
						IsActive: true,
					},
				},
				prefix: "envprefix",
			},
			preFunc: func() error {
				return nil
			},
			postFunc: func() error {
				return nil
			},
			want: &ExampleConfigA{
				Name:     "John",
				Age:      30,
				Size:     1.78,
				IsActive: true,
				Children: ExampleConfigB{
					Name:     "John",
					Age:      30,
					Size:     1.78,
					IsActive: true,
				},
			},
		},
		{
			name: "with env settings",
			args: args{
				st: &ExampleConfigA{
					Name:     "John",
					Age:      30,
					Size:     1.78,
					IsActive: true,
					Children: ExampleConfigB{
						Name:     "John",
						Age:      30,
						Size:     1.78,
						IsActive: true,
					},
				},
				prefix: "envprefix",
			},
			preFunc: func() error {
				os.Setenv("ENVPREFIX_NAME", "Johnathan")
				os.Setenv("ENVPREFIX_AGE", "25")
				os.Setenv("ENVPREFIX_SIZE", "1.55")
				os.Setenv("ENVPREFIX_ISACTIVE", "false")
				os.Setenv("ENVPREFIX_CHILDREN_NAME", "Marge")
				os.Setenv("ENVPREFIX_CHILDREN_AGE", "11")
				os.Setenv("ENVPREFIX_CHILDREN_SIZE", "1.22")
				os.Setenv("ENVPREFIX_CHILDREN_ISACTIVE", "false")
				return nil
			},
			postFunc: func() error {
				os.Unsetenv("ENVPREFIX_NAME")
				os.Unsetenv("ENVPREFIX_AGE")
				os.Unsetenv("ENVPREFIX_SIZE")
				os.Unsetenv("ENVPREFIX_ISACTIVE")
				os.Unsetenv("ENVPREFIX_CHILDREN_NAME")
				os.Unsetenv("ENVPREFIX_CHILDREN_AGE")
				os.Unsetenv("ENVPREFIX_CHILDREN_SIZE")
				os.Unsetenv("ENVPREFIX_CHILDREN_ISACTIVE")
				return nil
			},
			want: &ExampleConfigA{
				Name:     "Johnathan",
				Age:      25,
				Size:     1.55,
				IsActive: false,
				Children: ExampleConfigB{
					Name:     "Marge",
					Age:      11,
					Size:     1.22,
					IsActive: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preFunc != nil {
				tt.preFunc()
			}

			readStructAndEnrichWithEnv(tt.args.st, tt.args.prefix)
			diff := cmp.Diff(tt.args.st, tt.want)
			if diff != "" {
				t.Errorf("readStructAndEnrichWithEnv() diff = %v\n", diff)
			}

			if tt.postFunc != nil {
				tt.postFunc()
			}
		})
	}
}
