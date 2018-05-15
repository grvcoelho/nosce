package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	t.Run("NewMetadata", func(t *testing.T) {
		metadata := NewMetadata("https://169254.now.sh")

		assert.Equal(t, metadata.Endpoint, "https://169254.now.sh")
	})

	t.Run("Get", func(t *testing.T) {
		metadata := NewMetadata("https://169254.now.sh")

		ami, err := metadata.Get("ami-id")
		assert.NoError(t, err)
		assert.Equal(t, ami, "ami-1f3ca179")

		az, err := metadata.Get("availability-zone")
		assert.NoError(t, err)
		assert.Equal(t, az, "us-east-1b")
	})
}
