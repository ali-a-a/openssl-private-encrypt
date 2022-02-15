package opd_test

import (
	"strings"
	"testing"

	"github.com/ali-a-a/openssl-private/opd"
	"github.com/ali-a-a/openssl-private/ope"
	"github.com/stretchr/testify/assert"
)

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

//nolint:gochecknoglobals
var pemPrivateKey = testingKey(`-----BEGIN RSA TESTING KEY-----
MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0
fd7Ai2KW5ToIwzFofvJcS/STa6HA5gQenRUCAwEAAQJBAIq9amn00aS0h/CrjXqu
/ThglAXJmZhOMPVn4eiu7/ROixi9sex436MaVeMqSNf7Ex9a8fRNfWss7Sqd9eWu
RTUCIQDasvGASLqmjeffBNLTXV2A5g4t+kLVCpsEIZAycV5GswIhANEPLmax0ME/
EO+ZJ79TJKN5yiGBRsv5yvx5UiHxajEXAiAhAol5N4EUyq6I9w1rYdhPMGpLfk7A
IU2snfRJ6Nq2CQIgFrPsWRCkV+gOYcajD17rEqmuLrdIRexpg8N1DOSXoJ8CIGlS
tAboUGBxTDq3ZroNism3DaMIbKPyYrAqhKov1h5V
-----END RSA TESTING KEY-----
`)

//nolint:gochecknoglobals
var invalidTypeKey = testingKey(`-----BEGIN TESTING KEY-----
MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0
fd7Ai2KW5ToIwzFofvJcS/STa6HA5gQenRUCAwEAAQJBAIq9amn00aS0h/CrjXqu
/ThglAXJmZhOMPVn4eiu7/ROixi9sex436MaVeMqSNf7Ex9a8fRNfWss7Sqd9eWu
RTUCIQDasvGASLqmjeffBNLTXV2A5g4t+kLVCpsEIZAycV5GswIhANEPLmax0ME/
EO+ZJ79TJKN5yiGBRsv5yvx5UiHxajEXAiAhAol5N4EUyq6I9w1rYdhPMGpLfk7A
IU2snfRJ6Nq2CQIgFrPsWRCkV+gOYcajD17rEqmuLrdIRexpg8N1DOSXoJ8CIGlS
tAboUGBxTDq3ZroNism3DaMIbKPyYrAqhKov1h5V
-----END TESTING KEY-----
`)

//nolint:gochecknoglobals
var invalidBlockKey = testingKey(`-----BEGIN RSA TESTING KEY-----
MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0
-----END RSA TESTING KEY-----
`)

func TestOpensslPrivateEncrypt(t *testing.T) {
	t.Parallel()

	testData := "ali"

	encryptedData, err := ope.OpensslPrivateEncrypt(testData, pemPrivateKey)
	assert.NoError(t, err)

	cases := []struct {
		name          string
		privateKey    string
		encryptedData string
		shouldFail    bool
	}{
		{
			name:          "successful",
			privateKey:    pemPrivateKey,
			encryptedData: encryptedData,
			shouldFail:    false,
		},
		{
			name:          "invalid type",
			privateKey:    invalidTypeKey,
			encryptedData: encryptedData,
			shouldFail:    true,
		},
		{
			name:          "invalid block",
			privateKey:    invalidBlockKey,
			encryptedData: encryptedData,
			shouldFail:    true,
		},
		{
			name:          "invalid encrypted data",
			privateKey:    pemPrivateKey,
			encryptedData: "invalid",
			shouldFail:    true,
		},
	}

	for _, testcase := range cases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			data, err := opd.OpensslPrivateDecrypt(testcase.encryptedData, testcase.privateKey)
			if testcase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testData, data)
			}
		})
	}
}
