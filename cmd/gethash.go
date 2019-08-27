package cmd

import (
	"errors"
	"fmt"
	"github.com/hitman99/hashing/internal/hashes"
	"github.com/spf13/cobra"
	"strings"
)

var getHash = &cobra.Command{
	Use:   "get",
	Short: "gets hash from string. Made for openfaas",
	Run: func(cmd *cobra.Command, args []string) {
		server(debug)
	},
}

var sha256 = &cobra.Command{
	Use:   "sha256",
	Short: "makes sha256 hash",
	Run: func(cmd *cobra.Command, args []string) {
		hashIt(strings.Title(cmd.Name()), args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires string to hash")
		}
		return nil
	},
}
var sha512 = &cobra.Command{
	Use:   "sha512",
	Short: "makes sha512 hash",
	Run: func(cmd *cobra.Command, args []string) {
		hashIt(strings.Title(cmd.Name()), args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires string to hash")
		}
		return nil
	},
}
var md5 = &cobra.Command{
	Use:   "md5",
	Short: "makes md5 hash",
	Run: func(cmd *cobra.Command, args []string) {
		hashIt(strings.Title(cmd.Name()), args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires string to hash")
		}
		return nil
	},
}
var sha1 = &cobra.Command{
	Use:   "sha1",
	Short: "makes sha1 hash",
	Run: func(cmd *cobra.Command, args []string) {
		hashIt(strings.Title(cmd.Name()), args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires string to hash")
		}
		return nil
	},
}
var sha384 = &cobra.Command{
	Use:   "sha384",
	Short: "makes sha384 hash",
	Run: func(cmd *cobra.Command, args []string) {
		hashIt(strings.Title(cmd.Name()), args[0])
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires string to hash")
		}
		return nil
	},
}

func hashIt(kind, text string) {
	hash, err := hashes.GetHash(kind, text)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	fmt.Print(*hash)
}

func init() {
	getHash.AddCommand(sha1)
	getHash.AddCommand(sha256)
	getHash.AddCommand(sha384)
	getHash.AddCommand(sha512)
	getHash.AddCommand(md5)
}
