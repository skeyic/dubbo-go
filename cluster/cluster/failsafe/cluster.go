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

package failsafe

import (
	clusterpkg "github.com/skeyic/dubbo-go/cluster/cluster"
	"github.com/skeyic/dubbo-go/cluster/directory"
	"github.com/skeyic/dubbo-go/common/constant"
	"github.com/skeyic/dubbo-go/common/extension"
	"github.com/skeyic/dubbo-go/protocol"
)

func init() {
	extension.SetCluster(constant.ClusterKeyFailsafe, newFailsafeCluster)
}

type failsafeCluster struct{}

// newFailsafeCluster returns a failsafeCluster instance.
//
// Failure of security, anomalies, directly ignored. Usually it is
// used to write audit logs and other operations.
func newFailsafeCluster() clusterpkg.Cluster {
	return &failsafeCluster{}
}

// Join returns a baseClusterInvoker instance
func (cluster *failsafeCluster) Join(directory directory.Directory) protocol.Invoker {
	return clusterpkg.BuildInterceptorChain(newFailsafeClusterInvoker(directory))
}
