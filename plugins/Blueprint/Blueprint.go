// Copyright 2017 jungle Author. All Rights Reserved.
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

// date     : 2017/12/7 16:21
// author   : caimmy@hotmail.com

package Blueprint

import (
	"reflect"
)

type Blueprint struct {
	routers					map[string] reflect.Type
}

func (b *Blueprint) AddRouter(pattern string, c interface{}) {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	b.routers[pattern]	= t
}

func (b *Blueprint) GetRouter() *map[string] reflect.Type {
	return &b.routers
}

func NewBlueprint() *Blueprint {
	return &Blueprint{routers: make(map[string] reflect.Type)}
}
