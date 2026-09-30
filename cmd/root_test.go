package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootCommand(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(rootCmd)
	assert.Equal("gossm", rootCmd.Use)
	assert.NotEmpty(rootCmd.Short)
	assert.NotEmpty(rootCmd.Long)

	// Check flags
	profileFlag := rootCmd.PersistentFlags().Lookup("profile")
	if profileFlag == nil {
		profileFlag = rootCmd.Flags().Lookup("profile")
	}
	require.NotNil(t, profileFlag, "profile flag should be registered")
	assert.Equal("p", profileFlag.Shorthand)

	regionFlag := rootCmd.PersistentFlags().Lookup("region")
	if regionFlag == nil {
		regionFlag = rootCmd.Flags().Lookup("region")
	}
	require.NotNil(t, regionFlag, "region flag should be registered")
	assert.Equal("r", regionFlag.Shorthand)

	// Check registered subcommands
	expectedSubcommands := []string{
		"cmd",
		"fwd",
		"fwdrem",
		"mfa",
		"scp",
		"ssh",
		"start",
	}

	commandNames := make(map[string]bool)
	for _, c := range rootCmd.Commands() {
		commandNames[c.Name()] = true
	}

	for _, expected := range expectedSubcommands {
		assert.True(commandNames[expected], "expected subcommand %s to be registered", expected)
	}
}

func TestRootCommandHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "gossm [command]")
}
