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

package agency

import (
	"context"
	"net/http"

	"github.com/dalet-oss/arangodb-operator/pkg/util/arangod/conn"
	"github.com/dalet-oss/arangodb-operator/pkg/util/errors"
)

func GetAgencyConfig(ctx context.Context, connection conn.Connection) (*Config, error) {
	resp, code, err := conn.NewExecutor[any, Config](connection).ExecuteGet(ctx, "/_api/agency/config")
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, errors.Newf("Unknown response code %d", code)
	}

	return resp, nil
}

type Config struct {
	LeaderId string `json:"leaderId"`

	CommitIndex uint64 `json:"commitIndex"`

	Configuration struct {
		ID string `json:"id"`
	} `json:"configuration"`
}
