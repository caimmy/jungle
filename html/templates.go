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

// date     : 2017/12/2 21:31
// author   : caimmy@hotmail.com

package html

import (
	"html/template"
	"fmt"
	"io"
)

func NewTemplatesManager() *TemplatesManager {
	tplManager := &TemplatesManager{
		list_Layout_Templates: make(map[string] *template.Template),
		list_Page_Templates: make(map[string] *template.Template),
	}
	tplManager.None_Layout = template.Must(template.New("none_layout").Parse("{{ . }}"))
	return tplManager
}

type TemplatesManager struct {
	list_Layout_Templates 			map[string] *template.Template
	list_Page_Templates 			map[string] *template.Template

	None_Layout						*template.Template
}

func (t *TemplatesManager) LoadLayout(tpl_path string) *template.Template {
	cached_tpl, ok := t.list_Layout_Templates[tpl_path]
	if ok {
		fmt.Println("load tpl from cache")
		return cached_tpl
	} else {
		tpl_parsed, err := template.ParseFiles(tpl_path)
		if err == nil {
			t.list_Layout_Templates[tpl_path] = tpl_parsed
			fmt.Println("load tpl from file")
			return tpl_parsed
		} else {
			return t.None_Layout
		}
	}
}

func (t *TemplatesManager) RenderHtml(w io.Writer, tpl_path string, tpl_vars map[string] interface{}) {
	t.LoadLayout(tpl_path).Execute(w, tpl_vars)
}



