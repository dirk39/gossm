package internal

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAsset(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		input string
		isErr bool
	}{
		"fail": {input: "fail", isErr: true},
		"plugin/darwin_amd64/session-manager-plugin":      {input: "plugin/darwin_amd64/session-manager-plugin", isErr: false},
		"plugin/darwin_arm64/session-manager-plugin":      {input: "plugin/darwin_arm64/session-manager-plugin", isErr: false},
		"plugin/linux_amd64/session-manager-plugin":       {input: "plugin/linux_amd64/session-manager-plugin", isErr: false},
		"plugin/linux_arm64/session-manager-plugin":       {input: "plugin/linux_arm64/session-manager-plugin", isErr: false},
		"plugin/windows_amd64/session-manager-plugin.exe": {input: "plugin/windows_amd64/session-manager-plugin.exe", isErr: false},
	}

	for _, t := range tests {
		_, err := GetAsset(t.input)
		assert.Equal(t.isErr, err != nil)
	}

}

func TestGetSsmPluginName(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		output string
	}{
		"success": {output: "session-manager-plugin"},
	}

	for _, t := range tests {
		assert.Equal(t.output, GetSsmPluginName())
	}
}

func TestGetSsmPlugin(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		isErr bool
	}{
		"success": {isErr: false},
	}

	for _, t := range tests {
		_, err := GetSsmPlugin()
		assert.Equal(t.isErr, err != nil)
	}
}

func TestAssetChecksums(t *testing.T) {
	assert := assert.New(t)

	checksumData, err := GetAsset("plugin/checksums.sha256")
	assert.NoError(err)

	lines := strings.Split(strings.TrimSpace(string(checksumData)), "\n")
	assert.NotEmpty(lines)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		assert.Len(parts, 2)
		expectedHash := parts[0]
		relPath := parts[1]

		data, err := GetAsset("plugin/" + relPath)
		assert.NoError(err, "failed to get asset plugin/%s", relPath)
		actualHash := fmt.Sprintf("%x", sha256.Sum256(data))
		assert.Equal(expectedHash, actualHash, "checksum mismatch for %s", relPath)
	}
}
