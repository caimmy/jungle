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

// date     : 2017/11/24 13:53
// author   : caimmy@hotmail.com

package web

import (
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/caimmy/jungle/plugins/blueprint"
)

type JungleHttpServerHandler struct {
	routers map[string]reflect.Type
}

func (hander *JungleHttpServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	valid_uri := r.RequestURI
	uri_end_pos := strings.Index(valid_uri, "?")
	if uri_end_pos >= 0 {
		valid_uri = r.RequestURI[0:uri_end_pos]
	}

	if len(hander.routers) == 0 {
		io.WriteString(w, "Welcome to Jungle, make up your first JungleController please!")
	} else {
		predef_controller, ok := hander.routers[valid_uri]
		if !ok {
			predef_controller, ok = hander.routers[valid_uri+"/"]
		}

		if ok {
			controller := reflect.New(predef_controller).Interface().(ControllerInterface)
			controller.Init(&controller, w, r)
			controller.Action()
		} else {
			http.NotFound(w, r)
		}
	}
}

func (hander *JungleHttpServerHandler) Add(pattern string, c ControllerInterface) {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	hander.routers[pattern] = t
}

func (hander *JungleHttpServerHandler) AddBlueprint(prefix string, bp *blueprint.Blueprint) {
	if strings.Index(prefix, "/") != 0{
		prefix = "/" + prefix
	}
	for r, v := range *bp.GetRouter() {
		bp_prefix := prefix + r
		hander.routers[bp_prefix] = v
	}
}
