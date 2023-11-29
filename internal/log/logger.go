// Copyright (c) 2023
//
// @author 贺鹏Kavin
// 微信公众号:技术岁月
// https://github.com/skyhackvip/risk_engine
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
package log

import (
	"github.com/skyhackvip/risk_engine/configs"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

const (
	flag      = log.Ldate | log.Ltime | log.Lshortfile
	calldepth = 3
)

type Level int

const (
	LevelError Level = iota
	LevelWarn
	LevelInfo
	LevelDebug

	LevelMax
)

func (l Level) String() string {
	switch l {
	case LevelError:
		return "[ERROR]"
	case LevelWarn:
		return "[WARN]"
	case LevelInfo:
		return "[INFO]"
	case LevelDebug:
		return "[DEBUG]"
	}
	return ""
}

//interface
type Logger interface {
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

//logger
var l Logger

func Error(v ...interface{}) {
	l.Error(v)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v)
}
func Warn(v ...interface{}) {
	l.Warn(v)
}

func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v)
}

func Info(v ...interface{}) {
	l.Info(v)
}

func Infof(format string, v ...interface{}) {
	l.Infof(format, v)
}

func Debug(v ...interface{}) {
	l.Debug(v)
}

func Debugf(format string, v ...interface{}) {
	l.Debugf(format, v)
}

//init logger, in file
func InitLogger(outputMethod, path string) {
	if outputMethod == configs.FILE {
		l = NewDefaultLogger(&lumberjack.Logger{
			Filename:  path,
			MaxSize:   500,
			MaxAge:    10,
			LocalTime: true,
		}, "", flag, LevelDebug)
	} else { //default
		l = NewDefaultLogger(os.Stdout, "", flag, LevelInfo)
	}
}
