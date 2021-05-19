package installer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetInstaller(t *testing.T) {
	t.Run("get installer", func(t *testing.T) {
		installer, err := getInstaller(testEventBus, nil, "", "", zap.NewExample().Sugar())
		assert.NoError(t, err)
		assert.NotNil(t, installer)
		_, ok := installer.(*natsInstaller)
		assert.True(t, ok)

		installer, err = getInstaller(testExoticBus, nil, "", "", zap.NewExample().Sugar())
		assert.NoError(t, err)
		assert.NotNil(t, installer)
		_, ok = installer.(*exoticNATSInstaller)
		assert.True(t, ok)
	})
}
