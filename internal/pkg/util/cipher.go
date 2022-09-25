package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

//CBC mode encrypt
//Use base64 encoding to show the string after ase encrypt
func DesEncrypt(plain, key []byte) ([]byte, error) {
	if len(key) != des.BlockSize {
		return nil, errors.New("key length invalid")
	}

	//create a des block
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//data padding
	paddingPlain := padding(plain, block.BlockSize())

	//cbc mode
	iv := bytes.Repeat([]byte{1}, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)

	//encrypt blocks
	crypt := make([]byte, len(paddingPlain))
	blockMode.CryptBlocks(crypt, paddingPlain)

	return crypt, nil
}

//CBC mode decrypt
func DesDecrypt(crypt []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := bytes.Repeat([]byte{1}, block.BlockSize())
	blockMode := cipher.NewCBCDecrypter(block, iv)

	//decrypt block
	plain := make([]byte, len(crypt))
	blockMode.CryptBlocks(plain, crypt)

	//unpadding and return
	return unpadding(plain), nil
}

//PKCS5 padding
func padding(src []byte, blockSize int) []byte {
	remain := len(src) % blockSize
	pnum := blockSize - remain
	pcontent := bytes.Repeat([]byte{byte(pnum)}, pnum)
	return append(src, pcontent...)
}

//PKCS5 unpadding
func unpadding(plain []byte) []byte {
	length := len(plain)
	if length == 0 {
		return []byte{}
	}
	last := plain[length-1]
	unpaddingNum := int(last)
	return plain[:length-unpaddingNum]
}

//generate RSA public key and private key
func GenRsaKeyPair(privateKeyWriter, publicKeyWriter io.Writer, bits int) error {
	//generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	//x509 marshal private key
	privateDer := x509.MarshalPKCS1PrivateKey(privateKey)
	//create a private key pem block
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateDer,
	}
	//pem encode private key
	err = pem.Encode(privateKeyWriter, block)
	if err != nil {
		return err
	}

	//x509 marshal public key
	publicDer := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	//create a public key pem block
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicDer,
	}
	err = pem.Encode(publicKeyWriter, block)
	if err != nil {
		return err
	}

	return nil
}

type Rsa struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func NewRsa(publicKey, privateKey []byte) (*Rsa, error) {
	if publicKey == nil && privateKey == nil {
		return nil, errors.New("public key and private key is nil")
	}

	r := new(Rsa)

	if publicKey != nil && len(publicKey) > 0 {
		//decode public key
		block, _ := pem.Decode(publicKey)
		if block == nil {
			return nil, errors.New("pem decode public key failed")
		}
		//x509 parse
		pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		r.publicKey = pub
	}

	if privateKey != nil && len(privateKey) > 0 {
		//decode private key
		block, _ := pem.Decode(privateKey)
		if block == nil {
			return nil, errors.New("pem decode private key failed")
		}
		//x509 parse
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		r.privateKey = priv
	}

	return r, nil
}

func (r *Rsa) Encrypt(data []byte) ([]byte, error) {
	if r.publicKey == nil {
		return nil, errors.New("no public key")
	}
	bs, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, data)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (r *Rsa) Decrypt() {

}
