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

package inspector

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/dalet-oss/arangodb-operator/pkg/util/k8sutil/inspector/constants"
)

func (p *secretsInspectorV1) GroupVersionKind() schema.GroupVersionKind {
	return constants.SecretGKv1()
}

func (p *secretsInspectorV1) GroupVersionResource() schema.GroupVersionResource {
	return constants.SecretGRv1()
}

func (p *secretsInspector) GroupKind() schema.GroupKind {
	return constants.SecretGK()
}

func (p *secretsInspector) GroupResource() schema.GroupResource {
	return constants.SecretGR()
}
