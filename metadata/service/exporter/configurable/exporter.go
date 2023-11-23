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

package configurable

import (
	"sync"
)

import (
	"github.com/skeyic/dubbo-go/common"
	"github.com/skeyic/dubbo-go/common/constant"
	"github.com/skeyic/dubbo-go/common/extension"
	"github.com/skeyic/dubbo-go/common/logger"
	"github.com/skeyic/dubbo-go/config"
	_ "github.com/skeyic/dubbo-go/metadata/mapping/metadata"
	"github.com/skeyic/dubbo-go/metadata/service"
	"github.com/skeyic/dubbo-go/metadata/service/exporter"
	_ "github.com/skeyic/dubbo-go/metadata/service/remote"
	_ "github.com/skeyic/dubbo-go/protocol/dubbo"
)

// MetadataServiceExporter is the ConfigurableMetadataServiceExporter which implement MetadataServiceExporter interface
type MetadataServiceExporter struct {
	ServiceConfig   *config.ServiceConfig
	lock            sync.RWMutex
	metadataService service.MetadataService
}

func init() {
	extension.SetMetadataServiceExporter(constant.DefaultKey, NewMetadataServiceExporter)
}

// NewMetadataServiceExporter will return a service_exporter.MetadataServiceExporter with the specified  metadata service
func NewMetadataServiceExporter(metadataService service.MetadataService) exporter.MetadataServiceExporter {
	return &MetadataServiceExporter{
		metadataService: metadataService,
	}
}

// Export will export the metadataService
func (exporter *MetadataServiceExporter) Export(url *common.URL) error {
	if !exporter.IsExported() {
		version, _ := exporter.metadataService.Version()
		exporter.lock.Lock()
		defer exporter.lock.Unlock()
		exporter.ServiceConfig = config.NewServiceConfigBuilder().
			SetServiceID(constant.SimpleMetadataServiceName).
			SetProtocolIDs(constant.DefaultProtocol).
			AddRCProtocol(constant.DefaultProtocol, config.NewProtocolConfigBuilder().
				SetName(constant.DefaultProtocol).
				Build()).
			SetRegistryIDs("N/A").
			SetInterface(constant.MetadataServiceName).
			SetGroup(config.GetApplicationConfig().Name).
			SetVersion(version).
			SetProxyFactoryKey(constant.DefaultKey).
			SetMetadataType(config.GetApplicationConfig().MetadataType).
			Build()
		exporter.ServiceConfig.Implement(exporter.metadataService)
		err := exporter.ServiceConfig.Export()

		logger.Infof("[Metadata Service] The MetadataService exports urls : %v ", exporter.ServiceConfig.GetExportedUrls())
		return err
	}
	logger.Warnf("[Metadata Service] The MetadataService has been exported : %v ", exporter.ServiceConfig.GetExportedUrls())
	return nil
}

// Unexport will unexport the metadataService
func (exporter *MetadataServiceExporter) Unexport() {
	if exporter.IsExported() {
		exporter.ServiceConfig.Unexport()
	}
}

// GetExportedURLs will return the urls that export use.
// Notice！The exported url is not same as url in registry , for example it lack the ip.
func (exporter *MetadataServiceExporter) GetExportedURLs() []*common.URL {
	return exporter.ServiceConfig.GetExportedUrls()
}

// isExported will return is metadataServiceExporter exported or not
func (exporter *MetadataServiceExporter) IsExported() bool {
	exporter.lock.RLock()
	defer exporter.lock.RUnlock()
	return exporter.ServiceConfig != nil && exporter.ServiceConfig.IsExport()
}
