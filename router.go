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

package jungle

import (
	"io"
	"net/http"
	"reflect"
	"github.com/caimmy/jungle/plugins/blueprint"
	"strings"
	"os"
)

type JungleHttpServerHandler struct {
	routers			map[string] reflect.Type
	ws_routers		map[string] func(w http.ResponseWriter, r *http.Request)
}

func (handler *JungleHttpServerHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if !handler.filterRunWebsocketRequest(w, r) {
		handler.filterRunWebRequest(w, r)
	}
}

func (handler *JungleHttpServerHandler)filterRunStaticFile(w http.ResponseWriter, r *http.Request, name string) bool {
	http.ServeFile(w, r, name)
	return true
}

func (handler *JungleHttpServerHandler)filterRunWebsocketRequest(w http.ResponseWriter, r *http.Request) bool {
	prefix_matched := false

	predef_handler, ok 		:= handler.ws_routers[r.RequestURI]
	if ok {
		predef_handler(w, r)
		prefix_matched = true
	}

	return prefix_matched
}

func (handler *JungleHttpServerHandler)filterRunWebRequest(w http.ResponseWriter, r *http.Request) bool {
	prefix_matched := false
	valid_uri := r.RequestURI
	uri_end_pos := strings.Index(valid_uri, "?")
	if uri_end_pos >= 0 {
		valid_uri = r.RequestURI[0 : uri_end_pos]
	}

	if (len(handler.routers) == 0) {
		io.WriteString(w, "Welcome to Jungle, make up your first JungleController please!")
	} else {
		predef_controller, ok 			:= handler.routers[valid_uri]
		if !ok {
			predef_controller, ok = handler.routers[valid_uri + "/"]
		}

		if ok {
			controller := reflect.New(predef_controller).Interface().(ControllerInterface)
			controller.Init(&controller, w, r)
			controller.Action()
			prefix_matched = true
		} else {
			handler.filterRunStaticFile(w, r, StaticFilePath + string(os.PathSeparator) + strings.Replace(valid_uri, "/", string(os.PathSeparator), -1))
			http.NotFound(w, r)
		}
	}
	return prefix_matched
}

func (handler *JungleHttpServerHandler)AddWsHannler(pattern string, callback func(w http.ResponseWriter, r *http.Request)) {
	handler.ws_routers[pattern] = callback
}

func (handler *JungleHttpServerHandler)Add(pattern string, c ControllerInterface)  {
	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	handler.routers[pattern] = t
}

func (handler *JungleHttpServerHandler)AddBlueprint(prefix string, bp *blueprint.Blueprint) {
	if 0 != strings.Index(prefix, "/") {
		prefix = "/" + prefix
	}
	for r, v := range *bp.GetRouter() {
		bp_prefix := prefix + r
		handler.routers[bp_prefix] = v
	}
}