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
)

type ControllerInterface interface {
	Init()
	Prepare()
	Get(JungleResponseWriter, *JungleRequest)
	Post(JungleResponseWriter, *JungleRequest)
	Put(JungleResponseWriter, *JungleRequest)
	Delete(JungleResponseWriter, *JungleRequest)
}

type JungleController struct {
	// Templates setting
	TplName		 	string
	Layout 			string
	cache_layout 	template.Template
}

func (c *JungleController) Init() {
	c.Layout = "templates/layout.phtml"
}

func (c *JungleController) Prepare() {
	panic("need impleted!")
}

func (c *JungleController) Get(w JungleResponseWriter, r *JungleRequest) {
	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Post(w JungleResponseWriter, r *JungleRequest) {
	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Put(w JungleResponseWriter, r *JungleRequest) {
	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Delete(w JungleResponseWriter, r *JungleRequest) {
	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}

func (c * JungleController) RenderHtml(w JungleResponseWriter, tpl_path string, tpl_params map[string] interface{}) {
	tpl_string := tools.RenderHtml(tpl_path, tpl_params)
	c.SetLayout("test_demo/templates/layout.phtml")
	c.cache_layout.Execute(w, template.HTML(tpl_string))
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