package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytes2String(t *testing.T) {
	b := []byte{115, 112, 97, 114, 114, 111, 119}
	t.Log("result:", Bytes2String(b))
}

func TestString2Bytes(t *testing.T) {
	s := "sparrow"
	t.Log("result:", String2Bytes(s))
}

func TestDesCBCEncrypt(t *testing.T) {
	type args struct {
		src []byte
		key []byte
	}
	tests := []struct {
		name string
		args args
	}{
		//{"case1", args{src: []byte("12345678"), key: []byte("111")}},
		{"case2", args{src: []byte("12345678"), key: []byte("11111111")}},
		{"case3", args{src: []byte("1234"), key: []byte("12222222")}},
		{"case4", args{src: []byte("sparrow测试"), key: []byte("12222222")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DesEncrypt(tt.args.src, tt.args.key)
			assert.Nil(t, err)
			t.Log(string(got))
		})
	}

}

func TestGenRsaKeyPair(t *testing.T) {
	privateKey := bytes.NewBufferString("")
	publicKey := bytes.NewBufferString("")
	err := GenRsaKeyPair(privateKey, publicKey, 1024)
	t.Log(err)
	t.Log(privateKey.String())
	t.Log(publicKey.String())
}

func TestRsa_Encrypt(t *testing.T) {
	publicKey := `-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMqvl1CQ1qr054fmlFNPfX3a+cM9plmQDETIr+Be5+9ADY9grDHFKMoO
6qqHufJbVhpFma6i9misqk+j1jWgmgPXsvwknK8yiKufM5qBes2hnRYrwba0wWFX
BtQJGjTZ3wItKR4kUl8/X83mH9QZQ8b2Zvuu+6Ux7Z1Q8a8RbrejAgMBAAE=
-----END RSA PUBLIC KEY-----`
	rsa, err := NewRsa([]byte(publicKey), nil)
	t.Error(err)
	bs, err := rsa.Encrypt([]byte("i love yeji"))
	t.Error(err)
	t.Logf("%x", bs)
}
