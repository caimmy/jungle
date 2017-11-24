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

// date     : 2017/11/24 14:18
// author   : caimmy@hotmail.com

package main

import (
	"github.com/caimmy/jungle"
	"net/http"
	"io"
)

type CaimmyController struct {
	jungle.JungleController
}

func (c *CaimmyController)Get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sadfsdfsdfasdfsdfa caimmy controller")
}

func main() {
	jungle.Router("/", &CaimmyController{})
	jungle.Run()
}
