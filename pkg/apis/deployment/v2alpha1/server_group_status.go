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

package v2alpha1

import "github.com/dalet-oss/arangodb-operator/pkg/util"

type ServerGroupStatus struct {
	Index *int `json:"index,omitempty"`
}

func (s *ServerGroupStatus) Equal(b *ServerGroupStatus) bool {
	if s == nil && b == nil {
		return true
	}

	if s == nil || b == nil {
		return false
	}

	return util.CompareIntp(s.Index, b.Index)
}
