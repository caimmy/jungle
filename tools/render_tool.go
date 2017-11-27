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

// date     : 2017/11/27 22:48
// author   : caimmy@hotmail.com

package tools

import (
	"html/template"
	"bytes"
)

func RenderHtml(tpl_path string, tpl_vars map[string] interface{}) string {
	t, err := template.ParseFiles(tpl_path)
	if err != nil {
		panic("template file not found!")
	}
	w := bytes.NewBufferString("")
	t.Execute(w, tpl_vars)
	return w.String()
}