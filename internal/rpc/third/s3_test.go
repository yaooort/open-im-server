package third

import (
	"testing"

	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
	"github.com/stretchr/testify/require"
)

func TestDownloadURL(t *testing.T) {
	rawURL := "https://origin-ims1.example.com/openim/openim/temp/a.png?X-Amz-Signature=abc&response-content-type=image%2Fpng"

	t.Run("empty download address keeps raw url", func(t *testing.T) {
		server := &thirdServer{config: &Config{}}
		require.Equal(t, rawURL, server.downloadURL(rawURL))
	})

	t.Run("rewrite host and keep signed query", func(t *testing.T) {
		server := &thirdServer{config: &Config{
			MinioConfig: config.Minio{DownloadAddress: "https://ims1.example.com"},
		}}

		require.Equal(t,
			"https://ims1.example.com/openim/openim/temp/a.png?X-Amz-Signature=abc&response-content-type=image%2Fpng",
			server.downloadURL(rawURL),
		)
	})

	t.Run("download address path works as public prefix", func(t *testing.T) {
		server := &thirdServer{config: &Config{
			MinioConfig: config.Minio{DownloadAddress: "https://cdn.example.com/files"},
		}}

		require.Equal(t,
			"https://cdn.example.com/files/openim/openim/temp/a.png?X-Amz-Signature=abc&response-content-type=image%2Fpng",
			server.downloadURL(rawURL),
		)
	})
}
