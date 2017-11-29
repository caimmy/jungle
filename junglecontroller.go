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

// date     : 2017/11/22 14:27
// author   : caimmy@hotmail.com

package jungle

import (
	"net/http"
	"github.com/caimmy/jungle/tools"
	"html/template"
	"fmt"
)

type ControllerInterface interface {
	Init(JungleResponseWriter, *JungleRequest)
	Prepare()
	Get()
	Post()
	Put()
	Delete()
}

type JungleController struct {
	ResponseWriter 	JungleResponseWriter
	Request 		*JungleRequest
	// Templates setting
	TplName		 	string
	Layout 			string
	cache_layout 	template.Template
}

func (c *JungleController) Init(w JungleResponseWriter, r *JungleRequest) {
	c.ResponseWriter 	= w
	c.Request			= r
}

func (c *JungleController) Prepare() {
	fmt.Println("Prepared function called!")
}

func (c *JungleController) Get() {
	http.Error(c.ResponseWriter, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Post() {
	http.Error(c.ResponseWriter, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Put() {
	http.Error(c.ResponseWriter, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Delete() {
	http.Error(c.ResponseWriter, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c * JungleController) RenderHtml(tpl_path string, tpl_params map[string] interface{}) {
	tpl_string := tools.RenderHtml(tpl_path, tpl_params)
	c.SetLayout("test_demo/templates/layout.phtml")
	c.cache_layout.Execute(c.ResponseWriter, template.HTML(tpl_string))
}

func (c *JungleController) SetLayout(layout string) {
	c.Layout = layout
	if (c.Layout != "") {
		_t_layout, err := template.ParseFiles(c.Layout)
		if err != nil {
			panic("layout template not found")
		} else {
			c.cache_layout = *_t_layout
		}
	}
}