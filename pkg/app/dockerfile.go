package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/cobra"
)

var dockerfileCmd = &cobra.Command{
	Use:  "dockerfile",
	Long: "Sanitize Dockerfile",
	Run: func(cmd *cobra.Command, args []string) {
		sanitizeDockerfile()
	},
}

func init() {
	RootCmd.AddCommand(dockerfileCmd)
}

func sanitizeDockerfile() {
	// read from stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "FROM") {
			fmt.Println(sanitizeFrom(line))
		} else {
			fmt.Println(line)
		}
	}
}

func sanitizeFrom(line string) string {
	// split line into tokens
	tokens := strings.SplitN(line, " ", 3)

	// sanitize image name
	imageName := tokens[1]
	sanitized := sanitizeImageTag(imageName)

	// keep the remainer of the line
	ret := fmt.Sprintf("FROM %s", sanitized)
	if len(tokens) == 3 {
		ret = ret + " " + tokens[2]
	}

	return ret
}

func sanitizeImageTag(imageName string) string {
	// check if image tag is a digest
	if strings.Contains(imageName, "@") {
		return imageName
	}

	digest, err := imageDigest(imageName)
	if err != nil {
		log.Printf("Error parsing reference: %v", err)
		return imageName
	}

	return fmt.Sprintf("%s@%s", imageName, digest)
}

func imageDigest(imageName string) (string, error) {
	ref, err := name.ParseReference(imageName)
	if err != nil {
		return "", err
	}

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return "", err
	}

	digest, err := img.Digest()
	if err != nil {
		return "", err
	}

	return digest.String(), nil
}
