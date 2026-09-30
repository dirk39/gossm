package internal

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		ctx     context.Context
		key     string
		secret  string
		token   string
		region  string
		roleArn string
		isErr   bool
	}{
		"fail":    {isErr: true},
		"success": {ctx: context.Background(), key: mockAwsKey, secret: mockAwsSecret, region: mockRegion, isErr: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewConfig(tc.ctx, tc.key, tc.secret, tc.token, tc.region, tc.roleArn)
			assert.Equal(tc.isErr, err != nil)
		})
	}
}

func TestNewSharedConfig(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		ctx               context.Context
		profile           string
		sharedCredentials []string
		sharedConfigs     []string
		isErr             bool
	}{
		"fail": {isErr: true},
		"success": {
			ctx:               context.Background(),
			profile:           mockProfile,
			sharedConfigs:     []string{config.DefaultSharedConfigFilename()},
			sharedCredentials: []string{config.DefaultSharedCredentialsFilename()},
			isErr:             false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := NewSharedConfig(tc.ctx, tc.profile, tc.sharedConfigs, tc.sharedCredentials)
			assert.Equal(tc.isErr, err != nil)
		})
	}
}
