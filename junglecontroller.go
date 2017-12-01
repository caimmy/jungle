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
	"html/template"
	"fmt"
	"github.com/caimmy/jungle/context"
	"strings"
	"github.com/caimmy/jungle/tools"
	"io"
)

type ControllerInterface interface {
	Init(cptr *ControllerInterface, w http.ResponseWriter, r *http.Request)
	Prepare()
	Action()

	Get()
	Post()
	Put()
	Delete()
}

type JungleController struct {
	Ctx 			context.Context

	// Templates setting
	TplName		 	string
	Layout 			string
	cache_layout 	template.Template

	// Runtime Instances
	instance_prt	*ControllerInterface
}

// intialize controller instance.
// cptr params ControllerInterface : receive a instance to pointer the final Implements of JungleController
func (c *JungleController) Init(cptr *ControllerInterface, w http.ResponseWriter, r *http.Request) {
	c.Ctx.ResponseWriter	= w
	c.Ctx.Request 			= r

	c.instance_prt  					= cptr
	(*c.instance_prt).Prepare()
}

func (c *JungleController) Prepare() {
	fmt.Println("Prepared function called!")
}

func (c *JungleController) Get() {
	c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Post() {
	c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Put() {
	c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
}

func (c *JungleController) Delete() {
	c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
}

func (c* JungleController) Action() {
	switch strings.ToUpper(c.Ctx.Request.Method) {
	case METHOD_GET:
		(*c.instance_prt).Get()
	case METHOD_POST:
		(*c.instance_prt).Get()
	case METHOD_PUT:
		(*c.instance_prt).Put()
	case METHOD_DELETE:
		(*c.instance_prt).Delete()
	default:
		c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
	}
}

// Response standard Error information to client
func (c *JungleController) ResponseError(err_msg string, err_code int) {
	http.Error(c.Ctx.ResponseWriter, err_msg, err_code)
}

func (c * JungleController) RenderHtml(tpl_path string, tpl_params map[string] interface{}) {
	tpl_string := tools.RenderHtml(tpl_path, tpl_params)
	c.SetLayout("templates/layout.phtml")

	c.cache_layout.Execute(c.Ctx.ResponseWriter, template.HTML(tpl_string))
}

func (c *JungleController) Echo(content string) {
	io.WriteString(c.Ctx.ResponseWriter, content)
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