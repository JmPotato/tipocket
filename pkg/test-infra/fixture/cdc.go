// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package fixture

// CDCConfig for binlog component
type CDCConfig struct {
	CDCVersion       string
	DockerRepository string
	HubAddress       string
	LogPath          string
	Upstream         TiDBImageConfig
	Downstream       TiDBImageConfig
}

// TiDBImageConfig defines the tidb cluster image
type TiDBImageConfig struct {
	TiDBImage string
	TiKVImage string
	PDImage   string
}