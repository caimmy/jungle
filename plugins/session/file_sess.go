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

// date     : 2017/12/5 15:32
// author   : caimmy@hotmail.com

package session

import (
	"github.com/caimmy/jungle/context"
	"time"
	"net/http"
	"os"
	"fmt"
)

type FileSession struct {
	SessPath 			string

	SessionManager
}

func (this *FileSession) OpenSession(ctx context.Context) {
	this.mLock.Lock()
	defer this.mLock.Unlock()
	// TODO Load current valid session
	// continue...
	// if client has no sess cookie, generate one jungle id to identify client
	newSessionID, err := this.NewSessionID()
	if err == nil {
		sess := Session{SessID: newSessionID, LastTimeAccessed: time.Now().Unix(), Values: make(map[string] interface{})}
		this.SetSession(newSessionID, sess)
		cookie := http.Cookie{Name: this.m_strCookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(this.m_iMaxLifeTime)}
		http.SetCookie(ctx.ResponseWriter, &cookie)
	}
}

func (this *FileSession) CloseSession(ctx context.Context) {
	panic("tobe implements.")
}

func (this *FileSession) SetSession(id string, session Session) {
	sess_file_prt, err := os.OpenFile(this.SessPath + string(os.PathSeparator) + id, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprint(sess_file_prt, "asdfsadf")
	}
}

func (this *FileSession) GetSession(id string) Session {
	panic("tobe implements.")
}

func (this *FileSession) UpdateSession(id string) {
	panic("tobe implements.")
}

func (this *FileSession) GC() {
	panic("tobe implements.")
}