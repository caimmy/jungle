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

// date     : 2017/12/4 10:24
// author   : caimmy@hotmail.com

package logger

import (
	"log"
	"fmt"
	"os"
)

func NewLoggingManager (logPath string) *LoggingManager {
	logging_instance := &LoggingManager{LoggingPath: logPath + "/t.log"}
	logging_instance.StartRecord()
	return logging_instance
}

type LoggingManager struct {
	LoggingPath 			string
	LoggingSize				int
	LoggingTailLabel		string

	logger_file_prt 		*os.File

	log.Logger
}

func (this *LoggingManager) StartRecord() {
	var err error
	this.logger_file_prt, err = os.OpenFile(this.LoggingPath, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	this.Logger = *log.New(this.logger_file_prt, "", log.Ldate | log.Ltime | log.Llongfile)
}

// Close the filelog handler
func (this *LoggingManager) StopRecord() {
	this.logger_file_prt.Close()
}

func (this *LoggingManager) Writeline(txt string) {
	this.Output(2, fmt.Sprintf(txt))
	this.logger_file_prt.Sync()
}