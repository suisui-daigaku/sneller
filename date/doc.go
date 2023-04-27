// Copyright 2023 Sneller, Inc.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package date implements optimized date-parsing routines
// specific to the date formats that we support.
//
// Currently, only RFC3339Nano dates are supported.
package date

//go:generate ragel -Z -G2 parse_date.rl
//go:generate ragel -Z -G2 parse_duration.rl

//go:generate gofmt -w .
