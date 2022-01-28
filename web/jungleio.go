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

// date     : 2017/11/26 10:19
// author   : caimmy@hotmail.com

package web

import (
	"net/http"
)


const (
	METHOD_GET 		= "GET"
	METHOD_POST		= "POST"
	METHOD_PUT		= "PUT"
	METHOD_DELETE	= "DELETE"
)



type JungleResponseWriter struct {
	http.ResponseWriter
}

type JungleRequest struct {
	http.Request
}

func NotFound(w JungleResponseWriter, r *JungleRequest) {
	http.NotFound(w, &r.Request)
}

func Redirect(w JungleResponseWriter, r *JungleRequest, url string, code int) {
	http.Redirect(w, &r.Request, url, code)
}