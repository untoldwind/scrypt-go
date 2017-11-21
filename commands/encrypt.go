package commands

import (
	"io"
	"os"

	"github.com/untoldwind/scrypt-go/scryptlib"
	cli "gopkg.in/urfave/cli.v2"
)

var EncryptCommand = &cli.Command{
	Name:    "encrypt",
	Usage:   "<file>",
	Aliases: []string{"enc"},
	Action:  encrypt,
}

func encrypt(ctx *cli.Context) error {
	var in io.Reader = os.Stdin
	if ctx.Args().Len() > 0 {
		inFile, err := os.Open(ctx.Args().First())
		if err != nil {
			return err
		}
		defer inFile.Close()
		in = inFile
	}

	var out io.Writer = os.Stdout

	if ctx.Args().Len() > 1 {
		outFile, err := os.OpenFile(ctx.Args().Get(1), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer outFile.Close()
		out = outFile
	}

	return scryptlib.Encrypt([]byte("12345678"), in, out)
}
