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
	"html"
	"html/template"
	"strings"
	"io"
	"os"
	"bytes"
	"github.com/caimmy/jungle/context"
	"github.com/caimmy/jungle/web"
	"errors"
)

type ControllerInterface interface {
	Init(cptr *ControllerInterface, w http.ResponseWriter, r *http.Request)
	Prepare()
	Action()
	BeforeAction()
	AfterAction()

	Get()
	Post()
	Put()
	Delete()
}

type JungleController struct {
	Ctx 			context.Context

	// Templates setting
	TplPath		 	string
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

	c.instance_prt  		= cptr
	(*c.instance_prt).Prepare()
}

func (c *JungleController) SetSession(key string, value interface{}) {
	if JungleApp.SessionManager != nil {
		JungleApp.SessionManager.Set(&c.Ctx, key, value)
	}
}

func (c *JungleController) GetSession(key string) interface{} {
	if JungleApp.SessionManager != nil {
		return JungleApp.SessionManager.Get(&c.Ctx, key)
	}
	return nil
}

func (c *JungleController) Prepare() {
	// 开启全局会话管理器
	if JungleApp.SessionManager != nil {
		JungleApp.SessionManager.OpenSession(&c.Ctx)
		JungleApp.SessionManager.UpdateSession(&c.Ctx)
	}
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
	(*c.instance_prt).BeforeAction()
	switch strings.ToUpper(c.Ctx.Request.Method) {
	case METHOD_GET:
		(*c.instance_prt).Get()
	case METHOD_POST:
		(*c.instance_prt).Post()
	case METHOD_PUT:
		(*c.instance_prt).Put()
	case METHOD_DELETE:
		(*c.instance_prt).Delete()
	default:
		c.ResponseError("Method not Allowed", http.StatusMethodNotAllowed)
	}
	(*c.instance_prt).AfterAction()
}

func (c *JungleController) BeforeAction() {
	c.Ctx.Request.ParseForm()
}

func (c *JungleController) AfterAction() {

}

// Response standard Error information to client
func (c *JungleController) ResponseError(err_msg string, err_code int) {
	http.Error(c.Ctx.ResponseWriter, err_msg, err_code)
}

func (c *JungleController) Render(tplfile string, tpl_params map[string] interface{}) {
	// cached and prehot template in TplManager
	content_str := bytes.NewBufferString("")
	JungleApp.TemplateManager.RenderHtml(content_str, TemplatesPath + string(os.PathSeparator) + tplfile, tpl_params)
	layout_template := JungleApp.TemplateManager.LoadLayout(TemplatesPath + string(os.PathSeparator) + "/layout/layout.phtml")
	layout_template.Execute(c.Ctx.ResponseWriter, template.HTML(content_str.String()))
}

func (c *JungleController) RenderPartial(tplfile string, tpl_params map[string] interface{})  {
	// cached and prehot template in TplManager
	JungleApp.TemplateManager.RenderHtml(c.Ctx.ResponseWriter, TemplatesPath + string(os.PathSeparator) + tplfile, tpl_params)
}

func (c *JungleController) Echo(content string) {
	content = html.EscapeString(content)
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

func (c *JungleController) GetInstancesByName(filename string) (*web.UploadFile, error) {
	file, handler, err := c.Ctx.Request.FormFile(filename)
	if err == nil {
		return &web.UploadFile{FileHeader:handler, File:file}, nil
	} else {
		return nil, errors.New("not find file")
	}
}
