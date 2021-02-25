package parser

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	// Base64Suffix is appended to a label name in the request URL path to
	// mark the following label value as base64 encoded.
	Base64Suffix = "@base64"
)

//放在这里, 省得增加 rebase 时发生文件夹改动的风险

//ParseStream -
func ParseStream(req *http.Request, callback func(labels map[string]string) error) error {
	return nil
}

// decodeBase64 decodes the provided string using the “Base 64 Encoding with URL
// and Filename Safe Alphabet” (RFC 4648). Padding characters (i.e. trailing
// '=') are ignored.
func decodeBase64(s string) (string, error) {
	b, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(s, "="))
	return string(b), err
}