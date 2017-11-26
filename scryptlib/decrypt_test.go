package scryptlib_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/untoldwind/scrypt-go/scryptlib"
)

const scryptTestBase64 = `
c2NyeXB0ABIAAAAIAAAAAQ+QiTNyEWthcJ/qY3sTYUS3Ytbvu6f6IsFRpyaDg5x7+62ferHxlLr3
XkE3t3FHstri7MHk8ECU7Q2iJwMFoLoxVZpVVDxwROccmraqFihxSu59lIOp0aeDF5wqJb2cLLaX
O6FkhUT36iJELbn0UIc5UT5dRdywv/c/WBrGXY3Z
`

func TestDecrypt(t *testing.T) {
	require := require.New(t)
	data, err := base64.StdEncoding.DecodeString(scryptTestBase64)
	require.Nil(err)

	out := bytes.NewBuffer(nil)
	err = scryptlib.Decrypt([]byte("12345678"), bytes.NewReader(data), out)
	require.Nil(err)

	require.Equal("This is a test.\n", out.String())
}
