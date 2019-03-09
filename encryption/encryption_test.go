package encryption

import "testing"

func TestDecrypt(t *testing.T) {
	var testcases = []struct {
		input    string
		expected string
	}{
		{"U2FsdGVkX19YalVZkD9ulTLrymqTjqat8MajHbz9+go=", "password"},
		{"U2FsdGVkX1/EnuxHbWYuehpJHM0w84d8tTKVJcXYHPg=", "OpenTh3GAT3S"},
		{"U2FsdGVkX188B+Ia2gh9vADCoU3pDX16B4NgWtQplQk=", "S0M3RanD0mPASS"},
		{"U2FsdGVkX19bleVN45lxmW73QBDNYynWX5cV2kN3V78=", "JU1cy"},
		{"U2FsdGVkX1+nirWzLGZu6kmwSbfb+UjjTcS8xCbEwJk=", "HelL0W0rlD"},
		{"U2FsdGVkX185hRVkKP47ZVSe23o1eUfPld91Wp0kGbDV5lM+WH5Fc2rrb+uri6be", "INternetExplorerIsCrap"},
		{"U2FsdGVkX19GUeP+ebLSo24I7WGU+oobdJKvpncIcWUyv7xBpfmGwswOyFDis8nG", "WELCOME!@£@!!!!"},
		{"U2FsdGVkX1/UMLeIXuncaz6H6HK6J6q6st5fJ4YjbX1KH6pa2OtaraVf4H0qC0QF", "ENOUGHISEnough_PArseMe"},
	}

	for _, test := range testcases {
		Passphrase = "z4yH36a6zerhfE5427ZV"
		if item := Decrypt(test.input); item != test.expected {
			t.Errorf("encrypt(%v) = %v", test.input, item)
		}
	}
}

func TestEncrypt(t *testing.T) {
	var testcases = []struct {
		originalPassword string
	}{
		{"password"},
		{"LLCooLJ"},
		{"9$n)j7pCvB.cQC*q"},
		{"h&'y9#n7%,eHJKsGAq_5TLv_"},
		{">V6xyq5$f.,`!,GkEe$P@K`t_E:Ba2hdx{h=&P5zCm"},
		{"Et2XReL=b(Z+!qZ[sQ*8y9s(<qm@mg;*zA.YT932]_Au-C&m4g"},
		{"q5Rpzhn$?6!^W8af8crA==*^eg!xHNauxyNH4GG-naK2Hv6sEvU!UZrqX+t!7*2dQTCLSKdp6!5bE^Y!Qj7L#98vUHG@dKXF*X9nZ9mXuE9M2sX=t9X*-"},
		{"8PC??tdFbjPfMvbcpwW7kE7p$PdP-R!+aL%2huHRepyEhmuDQy#*LQq$UScy$?V_LavHQ*kG3es&!U*aB&mun%?LeGxkBRq-3R=Ca2Ch@SQ&k-A*f7$KP+z!?EgUTDK^g=mT^9^V9N5t^eb5dQTYMjMkpPYBrLCf2QTXXw*E55dNz2KzNAPrWcX*&R&n?T7Uu@a4h_PWLF6UyGHD29BZz4p*DHh5Ats6h5j@^nNHTZUDfDnCXV2MFXv9sj=2Pj6xGdkT5e^C3zPs3PNWxJC9frMnQG=m^PQL-kGNC3nne?rY&*4PLM^SmVE^SU9dQ2nD_SF%PEEhrpELv@KJG4uaCgyzkgCbChet94DvkLneRZ-577qe?jrFMcnnU$2Mvmh?4k&nm_M5jr#UyFyaUkP2#tk8KjadhpwarM8TKZ4PGmXWH?uxV4a+sy9YyBYhJ@w&^8E6Z^C3naKLu+jprV_?-uWu_X8VT+*#WRR?A+H=U_LUc#geQWmSwL4n?Z$c2N^Xd+FpJ3p^5wMSahBGb7tL_MgdA_RGFCWMceaCVd+uWm4R_+*Y4cqmZsXkh-nFWaJ5+ETz=ycC22_-kHYDQudYnQ!JJUSFPSUK2JjxVVjkuuMw2XhrYP2P@*XjPbk%qnB287bzTKK9jeNJ9u+QAP9X3_qAQ^u!3#rbH=uUM*S@6WS%VN5b3y4S?5Ad55uAjAQmD++$M73$D?nCW2-Vx$wk2qGjU8w#D%NTB!K-BzFQBmM+CELVp@Rpp2@q@yMsvzFHvmXdW-nfka6W?AWHfxk*cy#@hw@G#b&X@$R@2V_!Mz$KV$6s#&$sW3mjHvTwC$Pm7ZB8bjqT2d^d*BxFY#8U&fbjg!n*yvL@u2UQr-bun2XrYQN@2ug4_*+@kqbRZFQ*Dm=5s@jPYpCS3g+?4kJrpXAXZgqHV6!VDCM8c&Fq7*D$u3D5+%KPX?LHz=eMB*69tU*%BY=?V$B5UZdqxhSHcP+h-6?Xkj-UKX$H7wN4FF&%hJdhHTVeDCCubn"},
	}

	for _, test := range testcases {
		Passphrase = "z4yH36a6zerhfE5427ZV"
		if item := Encrypt(test.originalPassword); item == test.originalPassword {
			t.Errorf("encrypt(%v) = %v", test.originalPassword, item)
		}
	}
}
