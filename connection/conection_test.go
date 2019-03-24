package connection

import (
	"strings"
	"testing"

	"github.com/Jammicus/log-hunter/parser"
)

var node = parser.Node{
	"testHost",
	"testUser",
	"testPassword",
	"/home/example/.ssh/id_rsa",
	"22",
	"testConnection",
	"var/testLogLocation",
	"download/testDownloadDirectory",
	"testLog.name",
	"true",
	"sha256",
}

func TestValidChecksumAlgo(t *testing.T) {

	var testcases = []struct {
		algo string
	}{
		{"sha512"},
		{"sha256"},
		{"sha1"},
		{"md5"},
	}

	for _, test := range testcases {
		content := strings.NewReader("wEsMk7RReDciYKxXfxMN1ZyQINuq0laMqtjO0MhZabdVpCiuEKeMT3jOTaXaPyssyEACbh9SJHJ8TgtKosvOy61kRzSAhkoJetEDhTxotf72CFjLVVal87ecxoATEevDID8RvQ==")
		if err := verifyFileIntegrity(content, content, test.algo); (err != nil) != false {
			t.Error(err)
			t.Error("String:", content, "Algo", test.algo)
			t.Errorf(test.algo)
		}
	}
}

func TestInvalidChecksumAlgo(t *testing.T) {
	var testcases = []struct {
		algo string
	}{
		{"random"},
		{"sha726"},
		{"md42"},
		{"sha512md5"},
	}

	for _, test := range testcases {
		content := strings.NewReader("wEsMk7RReDciYKxXfxMN1ZyQINuq0laMqtjO0MhZabdVpCiuEKeMT3jOTaXaPyssyEACbh9SJHJ8TgtKosvOy61kRzSAhkoJetEDhTxotf72CFjLVVal87ecxoATEevDID8RvQ==")
		if err := verifyFileIntegrity(content, content, test.algo); (err == nil) != false {
			t.Error("Test should have thrown an error!")
			t.Error("String:", content, "Algo", test.algo)
			t.Errorf(test.algo)
		}
	}
}
