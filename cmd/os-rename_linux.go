//go:build linux
// +build linux

// Copyright (c) 2015-2023 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"golang.org/x/sys/unix"
)

// Rename2 captures time taken to call os.Rename
func Rename2(src, dst string) (err error) {
	defer updateOSMetrics(osMetricRename2, src, dst)(err)
	return unix.Renameat2(unix.AT_FDCWD, src, unix.AT_FDCWD, dst, uint(2)) // RENAME_EXCHANGE from 'man renameat2'
}
