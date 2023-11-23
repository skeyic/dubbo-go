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

package synthesizer

import (
	"github.com/skeyic/dubbo-go/common"
	"github.com/skeyic/dubbo-go/registry"
)

// SubscribedURLsSynthesizer is used to synthesize the subscribed url
type SubscribedURLsSynthesizer interface {
	// Supports the synthesis of the subscribed url or not
	Support(subscribedURL *common.URL) bool
	// synthesize the subscribed url
	Synthesize(subscribedURL *common.URL, serviceInstances []registry.ServiceInstance) []*common.URL
}
