package scryptlib_test

import (
	"bytes"
	"crypto/rand"
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

	require.Equal("Simple test message", decryptedOut.String())

	longMessage := make([]byte, 20000)
	_, err = rand.Read(longMessage)
	require.Nil(err)

	encryptedLongOut := bytes.NewBuffer(nil)
	err = scryptlib.Encrypt([]byte("12345678"), bytes.NewBuffer(longMessage), encryptedLongOut)
	require.Nil(err)

	decryptedLongOut := bytes.NewBuffer(nil)
	err = scryptlib.Decrypt([]byte("12345678"), encryptedLongOut, decryptedLongOut)
	require.Nil(err)

	require.Equal(longMessage, decryptedLongOut.Bytes())
}
