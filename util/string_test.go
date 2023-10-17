// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"testing"
)

func TestGetParamFromDataSourceName(t *testing.T) {
	dataSourceName := "postgresql://localhost:5432/casvisor?user=system&password=pass"
	key := "password"
	expected := "pass"

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetParamFromDataSourceName_NoMatch(t *testing.T) {
	dataSourceName := "postgresql://localhost:5432/casvisor?user=system"
	key := "password"
	expected := ""

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetParamFromDataSourceName_InvalidURL(t *testing.T) {
	dataSourceName := "invalid-url"
	key := "password"
	expected := ""

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetParamFromDataSourceName_EmptyInput(t *testing.T) {
	dataSourceName := ""
	key := "password"
	expected := ""

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetParamFromDataSourceName_EmptyKey(t *testing.T) {
	dataSourceName := "postgresql://localhost:5432/casvisor?user=system&password=pass"
	key := ""
	expected := ""

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetParamFromDataSourceName_KeyValueFormat(t *testing.T) {
	dataSourceName := "user=system password='with spaces' search_path=casvisor"
	key := "search_path"
	expected := "casvisor"

	result := GetParamFromDataSourceName(dataSourceName, key)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
