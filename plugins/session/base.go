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

// date     : 2017/12/5 15:01
// author   : caimmy@hotmail.com

package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/caimmy/jungle/context"
)

const (
	FILE_SESSION  = "file"
	REDIS_SESSION = "redis"
)

var (
	SESS_ID string
)

func init() {
	SESS_ID = "jungleid"
}

type SessionMgrInterface interface {
	OpenSession(ctx *context.Context)
	CloseSession(ctx *context.Context)
	SetSession(ctx *context.Context, session Session)
	Set(ctx *context.Context, key string, value interface{})
	GetSession(ctx *context.Context) (Session, error)
	Get(ctx *context.Context, key string) interface{}
	UpdateSession(ctx *context.Context)
	GC()
}

type Session struct {
	SessID           string                 `json:"sessid"`
	LastTimeAccessed int64                  `json:"lsttm"`
	Values           map[string]interface{} `json:"values"`
}

type SessionManager struct {
	m_strCookieName string
	m_iMaxlife      int64
	mLock           sync.RWMutex
}

func (this *SessionManager) LoadCookieValue(ctx *context.Context) string {
	cookie, err := ctx.Request.Cookie(this.m_strCookieName)
	if err != nil || cookie.Value == `` {
		n_id, _ := this.NewSessionID()
		return n_id
	} else {
		return cookie.Value
	}
}

func (this *SessionManager) OpenSession(ctx *context.Context) {
	panic("tobe implements.")
}

func (this *SessionManager) CloseSession(ctx *context.Context) {
	panic("tobe implements.")
}

func (this *SessionManager) SetSession(ctx *context.Context, session Session) {
	panic("tobe implements.")
}

func (this *SessionManager) Set(ctx *context.Context, key string, value interface{}) {
	panic("tobe implements.")
}

func (this *SessionManager) GetSession(ctx *context.Context) (Session, error) {
	panic("tobe implements.")
}

func (this *SessionManager) Get(ctx *context.Context, key string) interface{} {
	panic("tobe implements.")
}

func (this *SessionManager) UpdateSession(ctx *context.Context) {
	panic("tobe implements.")
}

func (this *SessionManager) GC() {
	panic("tobe implements.")
}

func (this *SessionManager) NewSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", errors.New("failed to generate session id")
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func NewSessionManager(sess_type string, iMaxlife int64, ext_params map[string]interface{}) SessionMgrInterface {
	var ret_session_manager SessionMgrInterface
	switch strings.ToLower(sess_type) {
	case REDIS_SESSION:
		ret_session_manager = &RedisSession{SessionManager: SessionManager{m_strCookieName: SESS_ID, m_iMaxlife: iMaxlife}}
	case FILE_SESSION:
		cache_path, ok := ext_params["path"]
		if !ok {
			return nil
		}
		ret_session_manager = &FileSession{SessPath: cache_path.(string), SessionManager: SessionManager{m_strCookieName: SESS_ID, m_iMaxlife: iMaxlife}}
	}
	go ret_session_manager.GC()
	return ret_session_manager
}

func NewSessionObject() *Session {
	return &Session{SESS_ID, time.Now().Unix(), make(map[string]interface{})}
}
