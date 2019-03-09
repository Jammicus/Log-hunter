package parser

import (
	"testing"
)

func TestGetDefault(t *testing.T) {
	var testcases = []struct {
		defaultVal      string
		nonDefaultValue string
		expected        string
	}{
		{"22", "22", "22"},
		{"hello", "", "hello"},
		{"", "example", "example"},
		{"", "", ""},
	}

	for _, test := range testcases {
		if item := isDefault(test.defaultVal, test.nonDefaultValue); item != test.expected {
			t.Errorf("getDefault(%v) = %v", test.defaultVal, test.nonDefaultValue)
		}
	}
}

func TestPasswordHandler(t *testing.T) {
	var testcases = []struct {
		encrypted    string
		nonEncrypted string
		expected     string
	}{
		{"", "exAmpLePassWord", "exAmpLePassWord"},
		{"U2FsdGVkX19YalVZkD9ulTLrymqTjqat8MajHbz9+go=", "", "password"},
		{"", "", ""},
	}
	for _, test := range testcases {
		if item := passwordHandler(test.encrypted, test.nonEncrypted); item != test.expected {
			t.Errorf("getDefault(%v) = %v", test.encrypted, test.nonEncrypted)
		}
	}
}

//func getNode(defaultList defaultInfo, nonDefaultList nodeInfo) Node {
func TestParse(t *testing.T) {
	var testcases = []struct {
		encrypted    string
		nonEncrypted string
		expected     string
	}{
		{"", "exAmpLePassWord", "exAmpLePassWord"},
		{"U2FsdGVkX19YalVZkD9ulTLrymqTjqat8MajHbz9+go=", "", "password"},
		{"", "", ""},
	}

	for _, test := range testcases {
		if item := passwordHandler(test.encrypted, test.nonEncrypted); item != test.expected {
			t.Errorf("getDefault(%v) = %v", test.encrypted, test.nonEncrypted)
		}
	}
}

// func TestGetNode(t *testing.T) {
// 	var testcases = []Node {
// 		encrypted    string
// 		nonEncrypted string
// 		expected     string
// 	}{
// 		{"", "exAmpLePassWord", "exAmpLePassWord"},
// 		{"U2FsdGVkX19YalVZkD9ulTLrymqTjqat8MajHbz9+go=", "", "password"},
// 		{"", "", ""},
// 	}

// 	for _, test := range testcases {
// 		if item := passwordHandler(test.encrypted, test.nonEncrypted); item != test.expected {
// 			t.Errorf("getDefault(%v) = %v", test.encrypted, test.nonEncrypted)
// 		}
// 	}
// }
