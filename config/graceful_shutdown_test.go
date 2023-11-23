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

package config

//
//import (
//	"github.com/skeyic/dubbo-go/config"
//	"github.com/skeyic/dubbo-go/config/consumer"
//	protocol2 "github.com/skeyic/dubbo-go/config/protocol"
//	"github.com/skeyic/dubbo-go/config/provider"
//	"github.com/skeyic/dubbo-go/config/reference"
//	"testing"
//)
//
//import (
//	"github.com/skeyic/dubbo-go/common/constant"
//	"github.com/skeyic/dubbo-go/common/extension"
//	"github.com/skeyic/dubbo-go/filter"
//	"github.com/skeyic/dubbo-go/protocol"
//)
//
//func TestGracefulShutdownInit(t *testing.T) {
//	extension.SetFilter(constant.GracefulShutdownConsumerFilterKey, func() filter.Filter {
//		return &config.mockGracefulShutdownFilter{}
//	})
//	extension.SetFilter(constant.GracefulShutdownProviderFilterKey, func() filter.Filter {
//		return &config.mockGracefulShutdownFilter{}
//	})
//	GracefulShutdownInit()
//}
//
//func TestBeforeShutdown(t *testing.T) {
//	extension.SetProtocol("registry", func() protocol.Protocol {
//		return &config.mockRegistryProtocol{}
//	})
//	extension.SetProtocol(constant.DUBBO, func() protocol.Protocol {
//		return &config.mockRegistryProtocol{}
//	})
//
//	extension.SetProtocol("mock", func() protocol.Protocol {
//		return &config.mockRegistryProtocol{}
//	})
//
//	consumerReferences := map[string]*reference.ReferenceConfig{}
//	consumerReferences[constant.DUBBO] = &reference.ReferenceConfig{
//		Protocol: constant.DUBBO,
//	}
//
//	// without configuration
//	config.consumerConfig = nil
//	config.referenceConfig = nil
//	BeforeShutdown()
//
//	config.consumerConfig = &consumer.ShutdownConfig{
//		References: consumerReferences,
//		ShutdownConfig: &ShutdownConfig{
//			Timeout:     "1",
//			StepTimeout: "1s",
//		},
//	}
//
//	providerProtocols := map[string]*protocol2.ProtocolConfig{}
//	providerProtocols[constant.DUBBO] = &protocol2.ProtocolConfig{
//		Name: constant.DUBBO,
//	}
//
//	providerProtocols["mock"] = &protocol2.ProtocolConfig{
//		Name: "mock",
//	}
//
//	config.referenceConfig = &provider.ProviderConfig{
//		ShutdownConfig: &ShutdownConfig{
//			Timeout:     "1",
//			StepTimeout: "1s",
//		},
//		Protocols: providerProtocols,
//	}
//	// test destroy protocol
//	BeforeShutdown()
//
//	config.referenceConfig = &provider.ProviderConfig{
//		ShutdownConfig: &ShutdownConfig{
//			Timeout:     "1",
//			StepTimeout: "-1s",
//		},
//		Protocols: providerProtocols,
//	}
//
//	config.consumerConfig = &consumer.ShutdownConfig{
//		References: consumerReferences,
//		ShutdownConfig: &ShutdownConfig{
//			Timeout:     "1",
//			StepTimeout: "-1s",
//		},
//	}
//
//	// test ignore steps
//	BeforeShutdown()
//}
