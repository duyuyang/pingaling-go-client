// Copyright © 2018 The Pingaling Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pingaling

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// ErrNotExpectedJSON is returned when the API response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}
