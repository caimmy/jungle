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

// date     : 2017/12/13 11:21
// author   : caimmy@hotmail.com

package web

import (
	"mime/multipart"
	"os"
	"io"
)

type UploadFile struct {
	FileHeader 				*multipart.FileHeader
	File 					multipart.File
}

func (uf *UploadFile) SaveAs(descpath string) error{
	f, err := os.OpenFile(descpath, os.O_CREATE, 0666)
	defer f.Close()
	if err == nil {
		_, cperr := io.Copy(f, uf.File)
		return cperr
	} else {
		return err
	}
}

func (uf *UploadFile) Release() {
	uf.File.Close()
}

func (uf *UploadFile) GetContentType() string {
	type_string, ok := "", true//uf.FileHeader.Header["Content-Type"]
	if ok {
		return string(type_string)
	} else {
		return ""
	}
}