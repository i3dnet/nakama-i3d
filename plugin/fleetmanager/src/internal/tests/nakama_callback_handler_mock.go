package tests

import (
	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CallbackHandlerMock struct {
	t              *testing.T
	id             string
	fn             runtime.FmCreateCallbackFn
	expectedStatus *runtime.FmCreateStatus
}

func NewCallbackHandlerMock(t *testing.T) *CallbackHandlerMock {
	return &CallbackHandlerMock{
		id: uuid.New().String(),
		t:  t,
	}
}

func (c *CallbackHandlerMock) ExpectedStatus(status runtime.FmCreateStatus) {
	c.expectedStatus = &status
}

func (c *CallbackHandlerMock) GenerateCallbackId() string {
	return c.id
}

func (c *CallbackHandlerMock) SetCallback(callbackId string, fn runtime.FmCreateCallbackFn) {
	assert.Equal(c.t, c.id, callbackId)
	assert.NotNil(c.t, fn)

	c.fn = fn
}

func (c *CallbackHandlerMock) InvokeCallback(callbackId string, status runtime.FmCreateStatus, instanceInfo *runtime.InstanceInfo, sessionInfo []*runtime.SessionInfo, metadata map[string]any, err error) {
	assert.Equal(c.t, c.id, callbackId)
	if c.expectedStatus != nil {
		assert.Equal(c.t, c.expectedStatus, status)
	}

	if c.fn == nil {
		return
	}

	c.fn(status, instanceInfo, sessionInfo, metadata, err)
}
