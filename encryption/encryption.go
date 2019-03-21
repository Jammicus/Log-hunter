package encryption

// Read https://dequeue.blogspot.com/2014/11/decrypting-something-encrypted-with.html

import (
	openssl "github.com/Luzifer/go-openssl"
	log "github.com/sirupsen/logrus"
)

var Passphrase = "z4yH36a6zerhfE5427ZV"

// Encrypt (string).
// Takes a string and returns a encrypted password based off the passphrase provided.
// This is equivilent to the following command:
// echo -n "example" | openssl aes-256-cbc -pass pass:<yourPassPhase> -md sha256 -a -salt
func Encrypt(password string) string {
	o := openssl.New()
	enc, err := o.EncryptBytes(Passphrase, []byte(password))

	if err != nil {
		log.Fatal("Could not decrypt password", password)
	}

	return string(enc)
}

// Decrypt(string)
// Takes an encrypted string and decrypts it using the passphase.
func Decrypt(encryptedPass string) string {
	o := openssl.New()

	dec, err := o.DecryptBytes(Passphrase, []byte(encryptedPass))
	if err != nil {
		log.Fatal("Unable to decrypt password", encryptedPass, err)
	}
	return string(dec)
}
