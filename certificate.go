package ycsdk

import (
	"github.com/fedorpatlin/go-sdk/gen/certificatemanager"
	certificatemanagerdata "github.com/fedorpatlin/go-sdk/gen/certificatemanagerdata"
)

const (
	CertificateManagerID     = "certificate-manager"
	CertificateManagerDataID = "certificate-manager-data"
)

func (sdk *SDK) Certificates() *certificatemanager.CertificateManager {
	return certificatemanager.NewCertificateManager(sdk.getConn(CertificateManagerID))
}

func (sdk *SDK) CertificatesData() *certificatemanagerdata.CertificateManagerData {
	return certificatemanagerdata.NewCertificateManagerData(sdk.getConn(CertificateManagerDataID))
}
