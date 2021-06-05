package util

import (
	"context"
	"fmt"
	"goldtalkAPI/pkg/thirdparty/go-trace"
	"strings"
	"time"
)

const (
	DLTAG_PREFIX = "_gogooltalk"
)

type LogType int64

const (
	_                     = iota
	LOG_TYPE_NONE LogType = iota
	LOG_TYPE_OUT
	LOG_TYPE_IN
	LOG_TYPE_SUCCESS
	LOG_TYPE_FAIL
)

func (lt LogType) String() string {
	switch lt {
	case LOG_TYPE_NONE:
		return ""
	case LOG_TYPE_OUT:
		return "out"
	case LOG_TYPE_IN:
		return "in"
	case LOG_TYPE_FAIL:
		return "fail"
	case LOG_TYPE_SUCCESS:
		return "success"
	}
	return ""
}

type LogBuild map[string]interface{}

func (lb LogBuild) Set(key string, value interface{}) LogBuild {
	lb[key] = value
	return lb
}

func (lb LogBuild) Get(key string) interface{} {
	return lb[key]
}

func (lb LogBuild) Remove(key string) LogBuild {
	delete(lb, key)
	return lb
}

func (lb LogBuild) SetWarn(msg string) LogBuild {
	return lb.Set("warnmsg", msg)
}

func (lb LogBuild) SetSuccess(msg string) LogBuild {
	return lb.Set("successmsg", msg)
}

func (lb LogBuild) SetError(err error) LogBuild {
	if err == nil{
		return lb
	}
	return lb.SetType(LOG_TYPE_FAIL).Set("errmsg", err.Error())
}

func (lb LogBuild) SetErrorNo(errno int64) LogBuild {
	return lb.Set("errorno", errno)
}

func (lb LogBuild) SetType(lt LogType) LogBuild {
	return lb.Set("_logtype", lt)
}

func (lb LogBuild) SetDescp(descp string) LogBuild {
	return lb.Set("descp", descp)
}

func (lb LogBuild) IsError() bool {
	return lb.Get("_logtype") == LOG_TYPE_FAIL
}

func (lb LogBuild) String() string {
	dltag := DLTAG_PREFIX
	logtype := fmt.Sprintf("_%v", LOG_TYPE_NONE)
	list := make([]string, 0)

	for key, value := range lb {
		switch key {
		case "_action":
			dltag += fmt.Sprintf("_%s", value)
		case "_logtype":
			logtype = fmt.Sprintf("_%s", value)
		case "_ctx":
			if value, ok := value.(context.Context); ok {
				list = append(list, fmt.Sprintf("%v", trace.FromContext(value)))
			}
		case "_startTime":
			if value, ok := value.(time.Time); ok {
				list = append(list, fmt.Sprintf("cost=%v", time.Since(value).Seconds()))
			}
		default:
			t := JsonString(value)
			list = append(list, fmt.Sprintf("%s=%+v", key, string(t)))
		}
	}
	return dltag + logtype + "||" + strings.Join(list, "||")
}

func NewLogBuild(ctx context.Context, action string) LogBuild {
	lb := make(LogBuild)
	lb.Set("_action", action)
	lb.Set("_ctx", ctx)
	lb.Set("_startTime", time.Now())

	return lb
}
