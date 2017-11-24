// Copyright 2014 jungle Author. All Rights Reserved.
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
	"net/http"
	"io"
)

type JungleHttpServerHandler struct {
	routers			map[string] ControllerInterface
}

func (hander *JungleHttpServerHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if (len(hander.routers) == 0) {
		io.WriteString(w, "Welcome to Jungle, make up your first JungleController please!")
	} else {
		controller, ok := hander.routers[r.RequestURI]
		if ok {
			switch r.Method {
			case "GET":
				controller.Get(w, r)
			default:
				io.WriteString(w, "Hello, Jungle!")
			}
		}
	}
}

func (hander *JungleHttpServerHandler)Add(pattern string, c ControllerInterface)  {
	hander.routers[pattern] = c
}
