// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package config

type Config struct {
	// Full path to the config file and folder where this Config originates from
	ConfigFilePath   string
	ConfigFolderPath string

	// Namespaces where operator watches for events
	Namespaces []string `yaml:"namespaces"`

	// Paths where to look for additional ClickHouse config .xml files to be mounted into Pod
	// config.d
	// conf.d
	// users.d
	// respectively
	CommonConfigsPath     string `yaml:"commonConfigsPath"`
	DeploymentConfigsPath string `yaml:"deploymentConfigsPath"`
	UsersConfigsPath      string `yaml:"usersConfigsPath"`
	// Config files fetched by these paths
	CommonConfigs     map[string]string
	DeploymentConfigs map[string]string
	UsersConfigs      map[string]string

	// Rolling update behavior - for how long to wait for StatefulSet to reach new Generation
	StatefulSetUpdateTimeout uint64 `yaml:"statefulSetUpdateTimeout"`
	// Rolling update behavior - for how long to sleep while polling StatefulSet to reach new Generation
	StatefulSetUpdatePollPeriod uint64 `yaml:"statefulSetUpdatePollPeriod"`
	// Rolling update behavior - what to do in case StatefulSet can't reach new Generation
	OnStatefulSetUpdateFailureAction string `yaml:"onStatefulSetUpdateFailureAction"`
}

const (
	// What to do in case StatefulSet can't reach new Generation - abort rolling update
	OnStatefulSetUpdateFailureActionAbort = "abort"

	// What to do in case StatefulSet can't reach new Generation - delete Pod and revert StatefulSet to previous Generation
	OnStatefulSetUpdateFailureActionRevert = "revert"
)
