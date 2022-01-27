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
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func NewLoggingManager(logPath string) *LoggingManager {
	logging_instance := &LoggingManager{
		LoggingPath:   logPath,
		LoggingCached: make([]interface{}, 0, 1000),
	}
	logging_instance.StartRecord()
	return logging_instance
}

type LoggingManager struct {
	LoggingPath      string
	LoggingSize      int
	LoggingTailLabel string
	LoggingCached    []interface{}

	logger_file_prt *os.File
	logger_file_day time.Time
	dump_log_ticker *time.Ticker

	rotateMutex sync.Mutex

	log.Logger
}

func (this *LoggingManager) StartRecord() {
	this.CreateLogger()
	go this.SchedualLogit()
}

// Close the filelog handler
func (this *LoggingManager) StopRecord() {
	this.dump_log_ticker.Stop()
	this.Dumplogs()

	this.logger_file_prt.Close()
}

func (this *LoggingManager) Writeline(v interface{}) {
	this.LoggingCached = append(this.LoggingCached, v)
}

func (this *LoggingManager) Dumplogs() {
	if len(this.LoggingCached) > 0 {
		this.Rotation()
		for _, val := range this.LoggingCached {
			this.Output(3, fmt.Sprintf("%v", val))
		}
		this.LoggingCached = append([]interface{}{})
		this.logger_file_prt.Sync()
	}
}

func (this *LoggingManager) CreateLogger() {
	fmt.Println("create logger")
	this.logger_file_day = time.Now()
	var err error
	this.logger_file_prt, err = os.OpenFile(this.LoggingPath+"/"+this.logger_file_day.Format("2006_01_02")+".log",
		os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	this.Logger = *log.New(this.logger_file_prt, "", log.Ldate|log.Ltime|log.Llongfile)
}

func (this *LoggingManager) Rotation() {
	cur_time := time.Now()
	this.rotateMutex.Lock()
	defer this.rotateMutex.Unlock()
	if cur_time.Day() != this.logger_file_day.Day() {
		// rotation log and rename prev logfile
		this.logger_file_prt.Close()
		this.CreateLogger()
	}
}

func (this *LoggingManager) SchedualLogit() {
	this.dump_log_ticker = time.NewTicker(time.Second * 5)

	for _ = range this.dump_log_ticker.C {
		this.Dumplogs()
	}
}
