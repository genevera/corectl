// Copyright (c) 2016 by António Meireles  <antonio.meireles@reformi.st>.
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
//

package coreos

import "github.com/TheNewNormal/corectl/components/common/assets"

const (
	defaultChannel = "alpha"
	defaultVersion = "latest"
)

var (
	// CoreOS default Channels
	Channels = []string{"alpha", "beta", "stable"}

	GPGLongID          = "50E0885593D2DCB4"
	GPGKey             = assets.Contents("target/coreos/CoreOSkey.public")
	CoreOSIgnitionTmpl = assets.
				Contents("target/coreos/corectl.ignition.yaml")
)
