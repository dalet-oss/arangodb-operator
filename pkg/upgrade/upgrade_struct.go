//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package upgrade

import (
	api "github.com/dalet-oss/arangodb-operator/pkg/apis/deployment/v1"
	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/interfaces"
)

func newUpgrade(version api.Version, u func(obj api.ArangoDeployment, status *api.DeploymentStatus, cache interfaces.Inspector) error) Upgrade {
	return &upgrade{
		version: version,
		upgrade: u,
	}
}

type upgrade struct {
	version api.Version

	upgrade func(obj api.ArangoDeployment, status *api.DeploymentStatus, cache interfaces.Inspector) error
}

func (u upgrade) Version() api.Version {
	return u.version
}

func (u upgrade) ArangoDeployment(obj api.ArangoDeployment, status *api.DeploymentStatus, cache interfaces.Inspector) error {
	if q := u.upgrade; q != nil {
		return q(obj, status, cache)
	}

	return nil
}
