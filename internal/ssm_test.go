package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestFindInstances(t *testing.T) {
	assert := assert.New(t)

	cfg, err := NewConfig(context.Background(), "", "", "", "", "")
	assert.NoError(err)

	tests := map[string]struct {
		ctx   context.Context
		cfg   aws.Config
		isErr bool
	}{
		"success": {
			ctx:   context.Background(),
			cfg:   cfg,
			isErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindInstances(tc.ctx, tc.cfg)
			assert.Equal(tc.isErr, err != nil)
			fmt.Println(len(result))
		})
	}
}
func TestFindInstanceIdsWithConnectedSSM(t *testing.T) {
	assert := assert.New(t)

	cfg, err := NewConfig(context.Background(), "", "", "", "", "")
	assert.NoError(err)

	tests := map[string]struct {
		ctx   context.Context
		cfg   aws.Config
		isErr bool
	}{
		"success": {
			ctx:   context.Background(),
			cfg:   cfg,
			isErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindInstanceIdsWithConnectedSSM(tc.ctx, tc.cfg)
			assert.Equal(tc.isErr, err != nil)
			fmt.Println(len(result))
		})
	}
}

func TestFindInstanceIdByIp(t *testing.T) {
	assert := assert.New(t)

	cfg, err := NewConfig(context.Background(), "", "", "", "", "")
	assert.NoError(err)

	tests := map[string]struct {
		ctx   context.Context
		cfg   aws.Config
		ip    string
		isErr bool
	}{
		"success": {
			ctx:   context.Background(),
			cfg:   cfg,
			ip:    "1.1.1.1",
			isErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindInstanceIdByIp(tc.ctx, tc.cfg, tc.ip)
			assert.Equal(tc.isErr, err != nil)
			fmt.Println(result)
		})
	}
}

func TestFindDomainByInstanceId(t *testing.T) {
	assert := assert.New(t)

	cfg, err := NewConfig(context.Background(), "", "", "", "", "")
	assert.NoError(err)

	tests := map[string]struct {
		ctx        context.Context
		cfg        aws.Config
		instanceId string
		isErr      bool
	}{
		"success": {
			ctx:        context.Background(),
			cfg:        cfg,
			instanceId: "i-unknown",
			isErr:      false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindDomainByInstanceId(tc.ctx, tc.cfg, tc.instanceId)
			assert.Equal(tc.isErr, err != nil)
			fmt.Println(result)
		})
	}
}

func TestAskUser(t *testing.T)                {}
func TestAskTeam(t *testing.T)                {}
func TestAskRegion(t *testing.T)              {}
func TestAskTarget(t *testing.T)              {}
func TestAskMultiTarget(t *testing.T)         {}
func TestAskPorts(t *testing.T)               {}
func TestCreateStartSession(t *testing.T)     {}
func TestDeleteStartSession(t *testing.T)     {}
func TestSendCommand(t *testing.T)            {}
func TestPrintCommandInvocation(t *testing.T) {}
func TestGenerateSSHExecCommand(t *testing.T) {}
func TestPrintReady(t *testing.T)             {}
