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
	"sync"
	"strings"
	"errors"
	"github.com/caimmy/jungle/context"
	"io"
	"crypto/rand"
	"encoding/base64"
)

var (
	SESS_ID				string
)

func init() {
	SESS_ID = "jungleid"
}

type SessionMgrInterface interface {
	OpenSession(ctx context.Context)
	CloseSession(ctx context.Context)
	SetSession(id string, session Session)
	GetSession(id string) Session
	UpdateSession(id string)
	GC()
}

type Session struct {
	SessID				string						`json:"sessid"`
	LastTimeAccessed	int64						`json:"lsttm"`
	Values 				map[string] interface{}		`json:"values"`
}

type SessionManager struct {
	m_strCookieName			string
	m_iMaxLifeTime			int

	mLock					sync.RWMutex
}

func (this *SessionManager) OpenSession(ctx context.Context) {
	panic("tobe implements.")
}

func (this *SessionManager) CloseSession(ctx context.Context) {
	panic("tobe implements.")
}

func (this *SessionManager) SetSession(id string, session Session) {
	panic("tobe implements.")
}

func (this *SessionManager) GetSession(id string) Session {
	panic("tobe implements.")
}

func (this *SessionManager) UpdateSession(id string) {
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

func NewSessionManager(sess_type string, max_life int, ext_params map[string] interface{}) (SessionMgrInterface, error)  {
	switch strings.ToLower(sess_type) {
	case "redis":
		return &RedisSession{SessionManager: SessionManager{m_strCookieName: SESS_ID, m_iMaxLifeTime:max_life}}, nil
	case "file":
		cache_path, ok := ext_params["path"]
		if ok {
			return nil, errors.New("not set cache file path for filesession")
		}
		return &FileSession{SessPath: cache_path.(string), SessionManager: SessionManager{m_strCookieName: SESS_ID, m_iMaxLifeTime:max_life}}, nil
	default:
		return nil, errors.New("unkown session type")
	}
}