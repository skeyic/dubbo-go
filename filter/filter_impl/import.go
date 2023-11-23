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

package filter_impl

// This package is for being compatible with older dubbo-go, please use `imports` package.
// This package may be DEPRECATED OR REMOVED in the future.

import (
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
)
