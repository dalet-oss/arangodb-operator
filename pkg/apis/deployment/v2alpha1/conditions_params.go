//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

package v2alpha1

const (
	// ConditionParamContainerUpdatingName define parameter used during Image Runtime update
	ConditionParamContainerUpdatingName ConditionParam = "updatingContainerName"
)

// ConditionParam is a strongly typed condition parameter
type ConditionParam string

type ConditionParams map[ConditionParam]string

// Equal compare two ConditionParams objects
func (c ConditionParams) Equal(b ConditionParams) bool {
	if len(c) != len(b) {
		return false
	}

	for k, v := range c {
		if v2, ok := b[k]; !ok || v != v2 {
			return false
		}
	}

	return true
}
