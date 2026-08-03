// Package id 提供 UUID 的生成与校验工具。
package id

import "github.com/google/uuid"

// New 返回一个随机 (v4) UUID 的标准字符串形式。
// 熵源不可用时会 panic —— 此时程序已无法安全继续。
func New() string {
	return uuid.NewString()
}

// NewRandom 与 New 一样生成随机 (v4) UUID，但把熵源错误返回给调用方，
// 供不接受 panic 的路径使用。
func NewRandom() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// IsValid 报告 s 是否为合法的 UUID 字符串。
// 除标准的 36 字符形式外，urn:uuid: 前缀、花括号包裹以及 32 位无连字符形式同样视为合法。
func IsValid(s string) bool {
	return uuid.Validate(s) == nil
}
