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

package cluster_impl

// This package is for being compatible with older dubbo-go, please use `imports` package.
// This package may be DEPRECATED OR REMOVED in the future.

import (
	_ "github.com/skeyic/dubbo-go/cluster/cluster/adaptivesvc"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/available"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/broadcast"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/failback"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/failfast"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/failover"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/failsafe"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/forking"
	_ "github.com/skeyic/dubbo-go/cluster/cluster/zoneaware"
)
