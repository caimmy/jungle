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

// date     : 2017/11/24 14:18
// author   : caimmy@hotmail.com

package main

import (
	"fmt"

	"github.com/caimmy/jungle/web"
)

type CaimmyController struct {
	web.JungleController
}

func (c *CaimmyController) Get() {
	param := make(map[string]interface{})
	param["name"] = "caimmy"
	param["days"] = 10
	c.Render("test.phtml", param)
}

func (c *CaimmyController) Post() {
	param := make(map[string]interface{})
	param["name"] = "alsdfj撒旦浪费就"
	param["days"] = 90
	c.RenderPartial("test.phtml", param)
}

func (c *CaimmyController) Prepare() {
	c.JungleController.Prepare()
	fmt.Println("caimmycontroller's prepare function called")
}

/*
func (c *CaimmyController)Post(w jungle.JungleResponseWriter, r *jungle.JungleRequest) {
	io.WriteString(w, "abcdefg by POST")
}
*/

func main() {
	caimmyController := CaimmyController{}
	web.Router("/", &caimmyController)
	web.TemplatesPath = "templates"
	web.Run(":8082", "")
	/*
		m := reflect.TypeOf(CaimmyController{})
		p := reflect.New(m)
		f := p.MethodByName("Prepare")
		fmt.Printf("a : %v\n", f)
		params := make([]reflect.Value, 0)
		f.Call(params)
		fmt.Println("asdfasdf")
	*/
}
