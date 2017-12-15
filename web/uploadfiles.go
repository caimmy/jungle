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
	type_collection, ok := uf.FileHeader.Header["Content-Type"]
	if ok && len(type_collection) > 0 {
		return type_collection[0]
	} else {
		return ""
	}
}

/**
check uploadfile is specialize content-type or not
@return bool
 */
func (uf *UploadFile) HasContentType(content_type string) bool {
	ret_existed_check := false
	type_collection, ok := uf.FileHeader.Header["Content-Type"]
	if ok {
		for _, val := range type_collection {
			if content_type == val {
				ret_existed_check = true
				break
			}
		}
	}
	return ret_existed_check
}