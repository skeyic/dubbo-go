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

package imports

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
	_ "github.com/skeyic/dubbo-go/cluster/loadbalance/consistenthashing"
	_ "github.com/skeyic/dubbo-go/cluster/loadbalance/leastactive"
	_ "github.com/skeyic/dubbo-go/cluster/loadbalance/p2c"
	_ "github.com/skeyic/dubbo-go/cluster/loadbalance/random"
	_ "github.com/skeyic/dubbo-go/cluster/loadbalance/roundrobin"
	_ "github.com/skeyic/dubbo-go/cluster/router/v3router"
	_ "github.com/skeyic/dubbo-go/common/proxy/proxy_factory"
	_ "github.com/skeyic/dubbo-go/config_center/apollo"
	_ "github.com/skeyic/dubbo-go/config_center/nacos"
	_ "github.com/skeyic/dubbo-go/config_center/zookeeper"
	_ "github.com/skeyic/dubbo-go/filter/accesslog"
	_ "github.com/skeyic/dubbo-go/filter/active"
	_ "github.com/skeyic/dubbo-go/filter/adaptivesvc"
	_ "github.com/skeyic/dubbo-go/filter/auth"
	_ "github.com/skeyic/dubbo-go/filter/echo"
	_ "github.com/skeyic/dubbo-go/filter/exec_limit"
	_ "github.com/skeyic/dubbo-go/filter/generic"
	_ "github.com/skeyic/dubbo-go/filter/graceful_shutdown"
	_ "github.com/skeyic/dubbo-go/filter/hystrix"
	_ "github.com/skeyic/dubbo-go/filter/metrics"
	_ "github.com/skeyic/dubbo-go/filter/seata"
	_ "github.com/skeyic/dubbo-go/filter/sentinel"
	_ "github.com/skeyic/dubbo-go/filter/token"
	_ "github.com/skeyic/dubbo-go/filter/tps"
	_ "github.com/skeyic/dubbo-go/filter/tracing"
	_ "github.com/skeyic/dubbo-go/metadata/mapping/metadata"
	_ "github.com/skeyic/dubbo-go/metadata/report/etcd"
	_ "github.com/skeyic/dubbo-go/metadata/report/nacos"
	_ "github.com/skeyic/dubbo-go/metadata/report/zookeeper"
	_ "github.com/skeyic/dubbo-go/metadata/service/exporter/configurable"
	_ "github.com/skeyic/dubbo-go/metadata/service/local"
	_ "github.com/skeyic/dubbo-go/metadata/service/remote"
	_ "github.com/skeyic/dubbo-go/metrics/prometheus"
	_ "github.com/skeyic/dubbo-go/protocol/dubbo"
	_ "github.com/skeyic/dubbo-go/protocol/dubbo3"
	_ "github.com/skeyic/dubbo-go/protocol/dubbo3/reflection"
	_ "github.com/skeyic/dubbo-go/protocol/grpc"
	_ "github.com/skeyic/dubbo-go/protocol/jsonrpc"
	_ "github.com/skeyic/dubbo-go/protocol/rest"
	_ "github.com/skeyic/dubbo-go/registry/etcdv3"
	_ "github.com/skeyic/dubbo-go/registry/nacos"
	_ "github.com/skeyic/dubbo-go/registry/polaris"
	_ "github.com/skeyic/dubbo-go/registry/protocol"
	_ "github.com/skeyic/dubbo-go/registry/servicediscovery"
	_ "github.com/skeyic/dubbo-go/registry/zookeeper"
)
