/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "compose only",
			args: []string{"up"},
			want: []string{"compose", "up"},
		},
		{
			name: "with context",
			args: []string{"--context", "foo", "-f", "compose.yaml", "up"},
			want: []string{"--context", "foo", "compose", "-f", "compose.yaml", "up"},
		},
		{
			name: "with host",
			args: []string{"--host", "tcp://1.2.3.4", "up"},
			want: []string{"--host", "tcp://1.2.3.4", "compose", "up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.args)
			assert.DeepEqual(t, tt.want, got)
		})
	}
}
