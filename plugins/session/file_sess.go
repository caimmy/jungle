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
	"net/http"
	"os"
	"encoding/json"
	"io/ioutil"
	"errors"
	"time"
)

type FileSession struct {
	SessPath 			string

	SessionManager
}

func (this *FileSession) GetSessionFilePath(id string) string {
	return this.SessPath + string(os.PathSeparator) + id
}

func (this *FileSession) OpenSession(ctx context.Context) {
	this.mLock.Lock()
	defer this.mLock.Unlock()
	cookie, err := ctx.Request.Cookie(this.m_strCookieName)
	if err != nil || cookie.Value == "" {
		newSessionID, err := this.NewSessionID()
		if err == nil {
			//sess := Session{SessID: newSessionID, LastTimeAccessed: time.Now().Unix(), Values: make(map[string] interface{})}
			//this.SetSession(ctx, sess)
			cookie := http.Cookie{Name: this.m_strCookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(this.m_iMaxLifeTime)}
			http.SetCookie(ctx.ResponseWriter, &cookie)
		}
	}
}

func (this *FileSession) CloseSession(ctx context.Context) {
	panic("tobe implements.")
}

func (this *FileSession) SetSession(ctx context.Context, session Session) {
	this.mLock.Lock()
	defer this.mLock.Unlock()
	id := this.LoadCookieValue(ctx)
	serialize_json, err := json.Marshal(session)
	if err == nil {
		_ = ioutil.WriteFile(this.GetSessionFilePath(id), serialize_json, 0666)
	}
}

func (this *FileSession) Set(ctx context.Context, key string, value interface{})  {
	session, err := this.GetSession(ctx)
	if err == nil {
		session.Values[key] = value
	} else {
		session = Session{this.LoadCookieValue(ctx), time.Now().Unix(), make(map[string]interface{})}
		session.Values[key] = value
	}

	this.SetSession(ctx, session)
}

func (this *FileSession) GetSession(ctx context.Context) (Session, error) {
	this.mLock.Lock()
	defer this.mLock.Unlock()
	id := this.LoadCookieValue(ctx)
	_, se := os.Stat(this.GetSessionFilePath(id))
	if se == nil {
		session_json, err := ioutil.ReadFile(this.GetSessionFilePath(id))
		if err == nil {
			var ret_session Session
			if ue := json.Unmarshal(session_json, ret_session); ue==nil {
				return ret_session, nil
			}
		}
	}
	return Session{}, errors.New("load session failure")
}

func (this *FileSession) Get(ctx context.Context, key string) interface{} {
	session, err := this.GetSession(ctx)
	if err != nil {
		val, ok := session.Values[key]
		if ok {
			return val
		}
	}
	return nil
}

func (this *FileSession) UpdateSession() {
	panic("tobe implements.")
}

func (this *FileSession) GC() {
	panic("tobe implements.")
}