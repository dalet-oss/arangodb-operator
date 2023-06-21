//
// DISCLAIMER
//
// Copyright 2016-2023 ArangoDB GmbH, Cologne, Germany
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

package rotation

import (
	"reflect"

	core "k8s.io/api/core/v1"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
)

func comparePodTolerations(_ api.DeploymentSpec, _ api.ServerGroup, spec, status *core.PodSpec) comparePodFunc {
	return func(builder api.ActionBuilder) (mode Mode, plan api.Plan, err error) {
		if !reflect.DeepEqual(spec.Tolerations, status.Tolerations) {
			plan = append(plan, builder.NewAction(api.ActionTypeRuntimeContainerSyncTolerations))

			status.Tolerations = spec.Tolerations
			mode = mode.And(InPlaceRotation)

			return
		}

		return
	}
}
