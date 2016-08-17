package providertests

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tracer0tong/gosaml2"
	"github.com/tracer0tong/goxmldsig"
)

func LoadXMLResponse(path string) string {
	xml, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(xml)
}

func LoadCertificateStore(path string) dsig.X509CertificateStore {
	encoded, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(encoded)
	if block == nil {
		panic("no certificate block found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	return &dsig.MemoryX509CertificateStore{
		Roots: []*x509.Certificate{cert},
	}
}

type ProviderTestScenario struct {
	Response        string
	ServiceProvider *saml2.SAMLServiceProvider
}

func ExerciseProviderTestScenarios(t *testing.T, scenarios []ProviderTestScenario) {
	for _, scenario := range scenarios {
		_, err := scenario.ServiceProvider.RetrieveAssertionInfo(scenario.Response)
		require.NoError(t, err)
	}
}
