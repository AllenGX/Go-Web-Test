package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 证书结构体
	template := x509.Certificate{
		SerialNumber: serialNumber,                                                 //证书序列号
		Subject:      subject,                                                      //证书名称
		NotBefore:    time.Now(),                                                   //生效时间
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),                         //结束时间
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, //证书用于服务器身份验证
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               //证书用于服务器身份验证
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},                           //启用地址
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) //创建私钥

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	//生成证书 cert.pem
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	//生成私钥 key.pem
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
