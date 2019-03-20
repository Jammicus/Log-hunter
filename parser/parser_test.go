package parser

import (
	"log-hunter/encryption"
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

func TestGetNode(t *testing.T) {

	var expectedCases = []Node{
		// Empty
		Node{},
		// Non encrypted password
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// Encrypted Password
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default username
		Node{
			"testHost",
			"defaultUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default port
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"33",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default connection
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"defaultConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default log location
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"/tmp/defaultLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default download directory
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"home/james/downloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default log name
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"defaultLog.log",
			"true",
			"",
		},
		// default deletelog
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"false",
			"",
		},
		// default checksum
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"sha256",
		},
		// overriden username
		Node{
			"testHost",
			"overridenName",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overidden port
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"5555",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden connection
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"overridenConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden log location
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"over/ridden/directory",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overridden download directory
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/overrides/directory",
			"testLog.name",
			"true",
			"",
		},
		// overriden log name
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"overridenLog.overrides",
			"true",
			"",
		},
		// overriden deletelog
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden checksum
		Node{
			"testHost",
			"testUser",
			"testPassword",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"sha256",
		},
	}

	var nodeConfigCases = []nodeInfo{
		// Empty
		nodeInfo{},
		// Non encrypted password
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// Encrypted Password
		nodeInfo{
			"testHost",
			"testUser",
			"",
			"U2FsdGVkX1+PVwxRcRmy3OKS3XYfxr06bRUdeYHqmpw=",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default username
		nodeInfo{
			"testHost",
			"",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default port
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default connection
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default log location
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// default download directory
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"",
			"testLog.name",
			"true",
			"",
		},
		// default log name
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"",
			"true",
			"",
		},
		// default deletelog
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"",
			"",
		},
		// default checksum
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden username
		nodeInfo{
			"testHost",
			"overridenName",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overidden port
		nodeInfo{
			"testHost",
			"testUser",
			"",
			"U2FsdGVkX1+PVwxRcRmy3OKS3XYfxr06bRUdeYHqmpw=",
			"5555",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden connection
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"overridenConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden log location
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"over/ridden/directory",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overridden download directory
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/overrides/directory",
			"testLog.name",
			"true",
			"",
		},
		// overriden log name
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"overridenLog.overrides",
			"true",
			"",
		},
		// overriden deletelog
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"",
		},
		// overriden checksum
		nodeInfo{
			"testHost",
			"testUser",
			"testPassword",
			"",
			"22",
			"testConnection",
			"var/testLogLocation",
			"download/testDownloadDirectory",
			"testLog.name",
			"true",
			"sha256",
		},
	}

	var defaultConfigCases = []defaultInfo{
		// empty
		defaultInfo{},
		// non encrypted password
		defaultInfo{},
		// Encrypted Password
		defaultInfo{},
		// default username
		defaultInfo{
			"defaultUser",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
		},
		// default port
		defaultInfo{
			"",
			"33",
			"",
			"",
			"",
			"",
			"",
			"",
		},
		// default connection
		defaultInfo{
			"",
			"",
			"defaultConnection",
			"",
			"",
			"",
			"",
			"",
		},
		// default log location
		defaultInfo{
			"",
			"",
			"",
			"/tmp/defaultLogLocation",
			"",
			"",
			"",
			"",
		},
		// default download directory
		defaultInfo{
			"",
			"",
			"",
			"",
			"home/james/downloadDirectory",
			"",
			"",
			"",
		},
		// default log name
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"defaultLog.log",
			"",
			"",
		},
		// default deletelog
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"",
			"false",
			"",
		},
		// default checksum
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"sha256",
		},
		// overriden username
		defaultInfo{
			"defaultUser",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
		},
		// overidden port
		defaultInfo{
			"",
			"33",
			"",
			"",
			"",
			"",
			"",
			"",
		},
		// overriden connection
		defaultInfo{
			"",
			"",
			"defaultConnection",
			"",
			"",
			"",
			"",
			"",
		},
		// overriden log location
		defaultInfo{
			"",
			"",
			"",
			"home/james/downloadDirectory",
			"",
			"",
			"",
			"",
		},
		// overridden download directory
		defaultInfo{
			"",
			"",
			"",
			"",
			"home/james/downloadDirectory",
			"",
			"",
			"",
		},
		// overriden log name
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"defaultLog.log",
			"",
			"",
		},
		// overriden deletelog
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"",
			"false",
			"",
		},
		// overriden checksum
		defaultInfo{
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"defaultChecksum",
		},
	}

	for i := range expectedCases {
		encryption.Passphrase = "z4yH36a6zerhfE5427ZV"

		if item := getNode(defaultConfigCases[i], nodeConfigCases[i]); item != expectedCases[i] {
			t.Error("Result did not match expected", '\n',
				"default config: ", defaultConfigCases[i], '\n',
				"node config: ", nodeConfigCases[i], '\n',
				"expected config:", expectedCases[i])
		}
	}

}

func TestParseSize(t *testing.T) {
	var testcases = []struct {
		file               string
		numOfNodesExpected int
	}{
		{"testData/noDefault/example1.yml", 1},
		{"testData/noDefault/example5.yml", 5},
		{"testData/noDefault/example10.yml", 10},
		{"testData/noDefault/example100.yml", 100},
		{"testData/default/example1.yml", 1},
		{"testData/default/example5.yml", 5},
		{"testData/default/example10.yml", 10},
		{"testData/default/example100.yml", 100},
	}

	for _, test := range testcases {
		if item := Parse(test.file); len(item) != test.numOfNodesExpected {
			t.Error("Returned array of nodes is not the expected size" + "\n" +
				"Returned: " + string(len(item)) + "\n" +
				"expected: " + string(test.numOfNodesExpected))
		}
	}
}

//go test -run=XXX -bench=.
// eg go test -run=100 -bench=BenchmarkDefault100

func BenchmarkNoDefault1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/noDefault/example1.yml")
	}
}

func BenchmarkNoDefault5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/noDefault/example5.yml")
	}
}

func BenchmarkNoDefault10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/noDefault/example10.yml")
	}
}

func BenchmarkNoDefault100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/noDefault/example100.yml")
	}
}
func BenchmarkDefault1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/default/example1.yml")
	}
}

func BenchmarkDefault5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/default/example5.yml")
	}
}

func BenchmarkDefault10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/default/example10.yml")
	}
}
func BenchmarkDefault100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parse("testData/default/example100.yml")
	}
}
