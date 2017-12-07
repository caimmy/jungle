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

// date     : 2017/11/30 15:00
// author   : caimmy@hotmail.com

package context

import "net/http"

type JungleInput struct {

}

type JungleOutput struct {
	Status 			int
	EnableGzip 		bool
}

type Context struct {
	Request 		*http.Request
	ResponseWriter 	http.ResponseWriter
}
