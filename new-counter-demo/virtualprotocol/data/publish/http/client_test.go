package http

import (
	"testing"

	"github.com/kubeedge/mapper-framework/pkg/common"
)

func TestPushReturnsWhenHTTPPostFails(t *testing.T) {
	pushMethod := &PushMethod{HTTP: &HTTPConfig{
		HostName: "http://127.0.0.1",
		Port:     1,
	}}

	pushMethod.Push(&common.DataModel{PropertyName: "count", Value: "1"})
}