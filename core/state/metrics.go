// Copyright 2021 The go-subaru Authors
// This file is part of the go-subaru library.
//
// The go-subaru library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-subaru library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-subaru library. If not, see <http://www.gnu.org/licenses/>.

package state

import "github.com/subaru-network/go-subaru/metrics"

var (
	accountUpdatedMeter        = metrics.NewRegisteredMeter("state/update/account", nil)
	storageUpdatedMeter        = metrics.NewRegisteredMeter("state/update/storage", nil)
	accountDeletedMeter        = metrics.NewRegisteredMeter("state/delete/account", nil)
	storageDeletedMeter        = metrics.NewRegisteredMeter("state/delete/storage", nil)
	accountTrieCommittedMeter  = metrics.NewRegisteredMeter("state/commit/accountnodes", nil)
	storageTriesCommittedMeter = metrics.NewRegisteredMeter("state/commit/storagenodes", nil)
)
