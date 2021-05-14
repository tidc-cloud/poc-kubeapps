# Copyright 2021 VMware. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Ensure the required version of cli tooling in tools/tools.go is installed.
# The versions are stored in ../go.mod as usual and this same approach is used in
# both octant and pinniped for build tooling.
# See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
cli-dependencies:
	go get \
		github.com/spf13/cobra/cobra

run:
	go run main.go

.PHONY: cli-dependencies run