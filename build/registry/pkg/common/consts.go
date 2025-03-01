// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

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

package common

const (
	RulesArtifactSuffix = "-rules"
	// EngineVersionKey is the name given to all the engine requirements.
	// The same name used by Falco when outputting the engine version.
	EngineVersionKey = "engine_version"
	// PluginAPIVersion is the name givet to the plugin api version requirements.
	// The same name used by Falco when outputting the plugin api version
	PluginAPIVersion = "plugin_api_version"
)
