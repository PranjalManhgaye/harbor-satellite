package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRetainProcessLocalSettings(t *testing.T) {
	remote := &Config{AppConfig: AppConfig{LogLevel: "info"}}
	local := AppConfig{
		BringOwnRegistry: true,
		LocalRegistryCredentials: RegistryCredentials{
			Username: "robot",
			Password: "secret",
			URL:      URL("http://127.0.0.1:5000"),
		},
		DirectDelivery: DirectDeliveryConfig{
			Enabled:  true,
			ImageDir: "/var/lib/rancher/k3s/agent/images",
		},
	}

	RetainProcessLocalSettings(remote, local, true)

	require.Equal(t, "info", remote.AppConfig.LogLevel)
	require.True(t, remote.AppConfig.BringOwnRegistry)
	require.Equal(t, "secret", remote.AppConfig.LocalRegistryCredentials.Password)
	require.Equal(t, URL("http://127.0.0.1:5000"), remote.AppConfig.LocalRegistryCredentials.URL)
	require.True(t, remote.AppConfig.DirectDelivery.Enabled)
	require.Equal(t, "/var/lib/rancher/k3s/agent/images", remote.AppConfig.DirectDelivery.ImageDir)
	require.True(t, remote.AppConfig.UseUnsecure)
}

func TestRetainProcessLocalSettingsLeavesRemoteWhenProcessHasNoOverrides(t *testing.T) {
	remote := &Config{AppConfig: AppConfig{LogLevel: "debug", UseUnsecure: true}}

	RetainProcessLocalSettings(remote, AppConfig{}, false)

	require.Equal(t, "debug", remote.AppConfig.LogLevel)
	require.True(t, remote.AppConfig.UseUnsecure)
	require.False(t, remote.AppConfig.BringOwnRegistry)
	require.False(t, remote.AppConfig.DirectDelivery.Enabled)
}
