/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package graceful_shutdown

import (
	"context"
	"net/url"
	"testing"
)

import (
	perrors "github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

import (
	"github.com/skeyic/dubbo-go/common"
	"github.com/skeyic/dubbo-go/common/constant"
	"github.com/skeyic/dubbo-go/common/extension"
	"github.com/skeyic/dubbo-go/config"
	"github.com/skeyic/dubbo-go/filter"
	"github.com/skeyic/dubbo-go/protocol"
	"github.com/skeyic/dubbo-go/protocol/invocation"
)

func TestProviderFilterInvoke(t *testing.T) {
	url := common.NewURLWithOptions(common.WithParams(url.Values{}))
	invocation := invocation.NewRPCInvocation("GetUser", []interface{}{"OK"}, make(map[string]interface{}))

	extension.SetRejectedExecutionHandler("test", func() filter.RejectedExecutionHandler {
		return &TestRejectedExecutionHandler{}
	})

	rootConfig := config.NewRootConfigBuilder().
		SetShutDown(config.NewShutDownConfigBuilder().
			SetTimeout("60s").
			SetStepTimeout("3s").
			SetRejectRequestHandler("test").
			Build()).Build()

	config.SetRootConfig(*rootConfig)

	filterValue, _ := extension.GetFilter(constant.GracefulShutdownProviderFilterKey)
	filter := filterValue.(*providerGracefulShutdownFilter)
	filter.Set(constant.GracefulShutdownFilterShutdownConfig, config.GetShutDown())
	assert.Equal(t, filter.shutdownConfig, config.GetShutDown())

	result := filter.Invoke(context.Background(), protocol.NewBaseInvoker(url), invocation)
	assert.NotNil(t, result)
	assert.Nil(t, result.Error())

	config.GetShutDown().RejectRequest.Store(true)
	result = filter.Invoke(context.Background(), protocol.NewBaseInvoker(url), invocation)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Error().Error(), "Rejected")

}

type TestRejectedExecutionHandler struct{}

// RejectedExecution will do nothing, it only log the invocation.
func (handler *TestRejectedExecutionHandler) RejectedExecution(url *common.URL,
	_ protocol.Invocation) protocol.Result {

	return &protocol.RPCResult{
		Err: perrors.New("Rejected"),
	}
}
