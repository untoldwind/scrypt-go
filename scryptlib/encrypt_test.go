package scryptlib_test

import (
	"bytes"
	"testing"

	"github.com/untoldwind/scrypt-go/scryptlib"

	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	require := require.New(t)

	encryptedOut := bytes.NewBuffer(nil)
	err := scryptlib.Encrypt([]byte("12345678"), bytes.NewBufferString("Simple test message"), encryptedOut)
	require.Nil(err)

	decryptedOut := bytes.NewBuffer(nil)
	err = scryptlib.Decrypt([]byte("12345678"), encryptedOut, decryptedOut)
	require.Nil(err)

}
