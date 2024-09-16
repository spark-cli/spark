package cmd

import (
	"errors"
	"fmt"
	"strings"

	"net/url"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

var newCmd = &cobra.Command{
	Use:   "new [flags] source [directory]",
	Short: "Spark new project",
	Long:  `Sparks a new project from a template.`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]
		source, err := ResolveSource(src)
		if err != nil {
			fmt.Println(err)
			return
		}

		directory := source.name
		if len(args) > 1 {
			directory = args[1]
		}

		directory, err = filepath.Abs(directory)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = git.PlainClone(directory, false, &git.CloneOptions{
			URL:   source.Primary(),
			Depth: 1,
		})
		if err != nil {
			_, err = git.PlainClone(directory, false, &git.CloneOptions{
				URL:   source.url,
				Depth: 1,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		base, err := filepath.Abs(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		relDir, err := filepath.Rel(base, directory)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Cloned", src, "to", relDir)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ResolveSource(source string) (Source, error) {
	err := validate.Var(source, "dirpath")
	if err == nil {
		return Source{source, pathLeaf(source)}, nil
	}

	err = validate.Var(source, "url")
	if err == nil {
		url, err := url.Parse(source)
		if err == nil {
			return ResolveSourceURL(url)
		}
	}

	return Source{}, errors.New("Source is neither a directory path nor a supported URL protocol")
}

func ResolveSourceURL(source *url.URL) (Source, error) {
	switch source.Scheme {
	case "file":
		return Source{source.Path, pathLeaf(source.Path)}, nil
	case "http":
		source.Scheme = "https"
		fallthrough
	case "https":
		return Source{source.String(), pathLeaf(source.String())}, nil
	case "std":
		source.Opaque = "spark-cli/" + source.Opaque
		fallthrough
	case "github":
		return Source{"https://github.com/" + source.Opaque + ".git", pathLeaf(source.Opaque)}, nil
	default:
		return Source{}, errors.New("Unsupported URL protocol")
	}
}

type Source struct {
	url  string
	name string
}

func pathLeaf(path string) string {
	nodes := strings.Split(path, "/")
	return nodes[len(nodes)-1]
}

func (t Source) Primary() string {
	nodes := strings.Split(t.url, "/")
	return strings.Join(nodes[0:len(nodes)-1], "/") + "/-" + nodes[len(nodes)-1]
}
