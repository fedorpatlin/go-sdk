package ycsdk

import (
	"github.com/fedorpatlin/go-sdk/gen/kms"
	kmscrypto "github.com/fedorpatlin/go-sdk/gen/kmscrypto"
)

const (
	KMSServiceID       = "kms"
	KMSCryptoServiceID = "kms-crypto"
)

func (sdk *SDK) KMS() *kms.KMS {
	return kms.NewKMS(sdk.getConn(KMSServiceID))
}

func (sdk *SDK) KMSCrypto() *kmscrypto.KMSCrypto {
	return kmscrypto.NewKMSCrypto(sdk.getConn(KMSCryptoServiceID))
}
