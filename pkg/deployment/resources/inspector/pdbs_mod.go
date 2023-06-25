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

package inspector

import (
	policy "k8s.io/api/policy/v1"

	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/constants"
	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/definitions"
	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/generic"
	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/mods"
	policyv1 "github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/poddisruptionbudget/v1"
)

func (i *inspectorState) PodDisruptionBudgetsModInterface() mods.PodDisruptionBudgetsMods {
	return podDisruptionBudgetsMod{
		i: i,
	}
}

type podDisruptionBudgetsMod struct {
	i *inspectorState
}

func (p podDisruptionBudgetsMod) V1() policyv1.ModInterface {
	return wrapMod[*policy.PodDisruptionBudget](definitions.PodDisruptionBudget, p.i.GetThrottles, generic.WithModStatusGetter[*policy.PodDisruptionBudget](constants.PodDisruptionBudgetGKv1(), p.clientv1))
}

func (p podDisruptionBudgetsMod) clientv1() generic.ModClient[*policy.PodDisruptionBudget] {
	return p.i.Client().Kubernetes().PolicyV1().PodDisruptionBudgets(p.i.Namespace())
}
