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

package reconcile

import (
	"context"

	api "github.com/dalet-oss/arangodb-operator/pkg/apis/deployment/v1"
)

func newJWTPropagatedAction(action api.Action, actionCtx ActionContext) Action {
	a := &actionJWTPropagated{}

	a.actionImpl = newActionImplDefRef(action, actionCtx)

	return a
}

type actionJWTPropagated struct {
	actionImpl

	actionEmptyCheckProgress
}

func (a *actionJWTPropagated) Start(ctx context.Context) (bool, error) {
	_, err := ensureJWTFolderSupportFromAction(a.actionCtx)
	if err != nil {
		a.log.Err(err).Error("Action not supported")
		return true, nil
	}

	propagatedFlag, exists := a.action.Params[propagated]
	if !exists {
		a.log.Err(err).Error("Propagated flag is missing")
		return true, nil
	}

	propagatedFlagBool := propagatedFlag == conditionTrue

	if err = a.actionCtx.WithStatusUpdate(ctx, func(s *api.DeploymentStatus) bool {
		if s.Hashes.JWT.Propagated != propagatedFlagBool {
			s.Hashes.JWT.Propagated = propagatedFlagBool
			return true
		}

		return false
	}); err != nil {
		return false, err
	}

	return true, nil
}
