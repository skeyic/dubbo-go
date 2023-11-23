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

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/skeyic/dubbo-go/common/file"
)

func TestCheckGenre(t *testing.T) {

	err := checkFileSuffix("abc")
	assert.NotNil(t, err)

	err = checkFileSuffix("zc")
	assert.NotNil(t, err)

	err = checkFileSuffix("json")
	assert.Nil(t, err)
}

func TestFileGenre(t *testing.T) {
	conf := NewLoaderConf(WithPath("../config/testdata/config/properties/application.properties"))
	assert.Equal(t, conf.suffix, "properties")
}

func TestRootConfig(t *testing.T) {
	rc := NewRootConfigBuilder().SetApplication(NewApplicationConfigBuilder().SetName("test-app").Build()).Build()
	conf := NewLoaderConf(WithRootConfig(rc))
	assert.Equal(t, conf.rc.Application.Name, "test-app")
}

func TestNewLoaderConf_WithBytes(t *testing.T) {
	str := `dubbo.application.name=dubbo-go
dubbo.application.module=local
dubbo.services.HelloService.registry=nacos,zk`

	conf := NewLoaderConf(WithBytes([]byte(str)), WithGenre("properties"))

	assert.NotNil(t, conf)
	assert.NotNil(t, conf.bytes)
}

func TestNewLoaderConf_WithSuffix(t *testing.T) {
	conf := NewLoaderConf(
		WithSuffix(file.JSON),
		WithPath("../config/testdata/config/properties/application.properties"),
	)

	assert.Equal(t, conf.suffix, string(file.PROPERTIES))
}
