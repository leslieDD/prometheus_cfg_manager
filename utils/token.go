package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

// 生成一个可以解析的TOKEN

// TokenInfo TokenInfo
type TokenInfo struct {
	UserName string
	Expire   int64
}

// GetSimpleToken 生成简易的token
func GetSimpleToken() (string, error) {
	token, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

// GetSalt 生成简易的salt
func GetSalt() (string, error) {
	salt, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return salt.String(), nil
}

// GetToken 获取token - 复杂的
// GetToken，要求包含的字符不能超过64个字符，证书的长度是512
func GetToken(username string, expire int64, publicKey string) (string, error) {
	info := TokenInfo{
		UserName: username,
		Expire:   expire,
	}
	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	token, err := RsaEncrypt(bytes, []byte(publicKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// DecryptToken 对加密的token进行解码
func DecryptToken(token string, privateKey string) (*TokenInfo, error) {
	bytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	result, err := RsaDecrypt(bytes, []byte(privateKey))
	if err != nil {
		return nil, err
	}

	var info TokenInfo
	err = jsoniter.Unmarshal(result, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// 公钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

// RsaEncrypt 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

type storeRsa struct {
	rsaInfo bytes.Buffer
}

func (sr *storeRsa) Write(p []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("empty, recv nil")
	}
	sr.rsaInfo.Write(p)
	return len(p), nil
}

func (sr *storeRsa) Bytes() []byte {
	return sr.rsaInfo.Bytes()
}

func (sr *storeRsa) String() string {
	return sr.rsaInfo.String()
}

func newSR() *storeRsa {
	sr := &storeRsa{}
	sr.rsaInfo = bytes.Buffer{}
	return sr
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey(bits int) (map[string]string, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateStore := newSR()
	err = pem.Encode(privateStore, block)
	if err != nil {
		return nil, err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	publicStore := newSR()
	err = pem.Encode(publicStore, block)
	if err != nil {
		return nil, err
	}
	rsaInfo := map[string]string{}
	rsaInfo["private"] = privateStore.String()
	rsaInfo["pubilc"] = publicStore.String()
	return rsaInfo, nil
}
