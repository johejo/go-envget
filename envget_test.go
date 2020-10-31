package envget_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/johejo/go-envget"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name, key, value, fallback, want string
	}{
		{name: "get value", key: "TEST_KEY", value: "TEST_VALUE", fallback: "TEST_FALLBACK", want: "TEST_VALUE"},
		{name: "get fallback", key: "TEST_KEY", value: "", fallback: "TEST_FALLBACK", want: "TEST_FALLBACK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.key, tt.value)
			got := envget.GetString(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("want=%v, but got %v", tt.want, got)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name, key, value string
		fallback, want   int
	}{
		{name: "get 0", key: "TEST_KEY", value: "0", fallback: 99, want: 0},
		{name: "get 1", key: "TEST_KEY", value: "1", fallback: 99, want: 1},
		{name: "get fallback", key: "TEST_KEY", value: "", fallback: 99, want: 99},
		{name: "invalid value", key: "TEST_KEY", value: "xyz", fallback: 99, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.key, tt.value)
			got := envget.GetInt(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("want=%v, but got %v", tt.want, got)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name, key, value string
		fallback, want   bool
	}{
		{name: "get 0", key: "TEST_KEY", value: "0", fallback: true, want: false},
		{name: "get 1", key: "TEST_KEY", value: "1", fallback: false, want: true},
		{name: "get true", key: "TEST_KEY", value: "true", fallback: false, want: true},
		{name: "get True", key: "TEST_KEY", value: "True", fallback: false, want: true},
		{name: "get TRUE", key: "TEST_KEY", value: "TRUE", fallback: false, want: true},
		{name: "get fallback", key: "TEST_KEY", value: "", fallback: true, want: true},
		{name: "invalid value", key: "TEST_KEY", value: "xyz", fallback: true, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.key, tt.value)
			got := envget.GetBool(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("want=%v, but got %v", tt.want, got)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	tests := []struct {
		name, key, value string
		fallback, want   time.Duration
	}{
		{name: "get 0", key: "TEST_KEY", value: "0ns", fallback: 99 * time.Second, want: 0 * time.Nanosecond},
		{name: "get 1", key: "TEST_KEY", value: "1ms", fallback: 99 * time.Second, want: 1 * time.Millisecond},
		{name: "get fallback", key: "TEST_KEY", value: "", fallback: 99 * time.Second, want: 99 * time.Second},
		{name: "invalid format", key: "TEST_KEY", value: "xyx", fallback: 99 * time.Second, want: 99 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.key, tt.value)
			got := envget.GetDuration(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("want=%v, but got %v", tt.want, got)
			}
		})
	}
}

func TestGetStringSlice(t *testing.T) {
	tests := []struct {
		name, key, value string
		fallback, want   []string
	}{
		{name: "get one value", key: "TEST_KEY", value: "A", fallback: []string{"F"}, want: []string{"A"}},
		{name: "get one value", key: "TEST_KEY", value: "A,", fallback: []string{"F"}, want: []string{"A"}},
		{name: "get one value", key: "TEST_KEY", value: ",A,", fallback: []string{"F"}, want: []string{"A"}},
		{name: "get one value", key: "TEST_KEY", value: ",A", fallback: []string{"F"}, want: []string{"A"}},
		{name: "get multi value", key: "TEST_KEY", value: "A,B", fallback: []string{"F"}, want: []string{"A", "B"}},
		{name: "get multi value", key: "TEST_KEY", value: "A, B", fallback: []string{"F"}, want: []string{"A", "B"}},
		{name: "get multi value", key: "TEST_KEY", value: ",A, B", fallback: []string{"F"}, want: []string{"A", "B"}},
		{name: "get multi value", key: "TEST_KEY", value: "A,B,", fallback: []string{"F"}, want: []string{"A", "B"}},
		{name: "get fallback value", key: "TEST_KEY", value: "", fallback: []string{"F"}, want: []string{"F"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.key, tt.value)
			got := envget.GetStringSlice(tt.key, tt.fallback)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, but got %v", tt.want, got)
			}
		})
	}
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()
	if err := os.Setenv(key, value); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Unsetenv(key); err != nil {
			t.Fatal(err)
		}
	})
}
