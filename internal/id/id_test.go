package id

import (
	"testing"

	"github.com/google/uuid"
)

func TestNew_Format(t *testing.T) {
	s := New()

	u, err := uuid.Parse(s)
	if err != nil {
		t.Fatalf("Parse(%q) 返回错误: %v", s, err)
	}
	if got := u.Version(); got != 4 {
		t.Errorf("Version() = %d, 期望 4", got)
	}
	if got := u.Variant(); got != uuid.RFC4122 {
		t.Errorf("Variant() = %v, 期望 %v", got, uuid.RFC4122)
	}
}

func TestNew_Unique(t *testing.T) {
	const n = 1000

	seen := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		s := New()
		if _, dup := seen[s]; dup {
			t.Fatalf("第 %d 次生成出现重复: %q", i, s)
		}
		seen[s] = struct{}{}
	}
	if len(seen) != n {
		t.Errorf("生成了 %d 个不同的 UUID, 期望 %d", len(seen), n)
	}
}

func TestNewRandom(t *testing.T) {
	s, err := NewRandom()
	if err != nil {
		t.Fatalf("NewRandom() 返回错误: %v", err)
	}
	if !IsValid(s) {
		t.Errorf("NewRandom() = %q, 未通过 IsValid", s)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"标准形式", "f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"urn:uuid: 前缀", "urn:uuid:f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"花括号包裹", "{f47ac10b-58cc-4372-a567-0e02b2c3d479}", true},
		{"32 位无连字符", "f47ac10b58cc4372a5670e02b2c3d479", true},
		{"空串", "", false},
		{"长度不足", "f47ac10b-58cc-4372-a567", false},
		{"非 hex 字符", "z47ac10b-58cc-4372-a567-0e02b2c3d479", false},
		{"错误的 urn 前缀", "urn:uuidx:f47ac10b-58cc-4372-a567-0e02b2c3d47", false},
		{"花括号只有一半", "{f47ac10b-58cc-4372-a567-0e02b2c3d479", false},
		{"连字符位置错误", "f47ac10b5-8cc-4372-a567-0e02b2c3d479", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.in); got != tt.want {
				t.Errorf("IsValid(%q) = %v, 期望 %v", tt.in, got, tt.want)
			}
		})
	}
}
