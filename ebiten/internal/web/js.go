// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build js

package web

import (
	"strings"
	"sync"

	"github.com/hajimehoshi/gopherwasm/js"
)

var (
	isNodeJSOnce sync.Once
	isNodeJS     = false
)

func IsNodeJS() bool {
	isNodeJSOnce.Do(func() {
		isNodeJS = js.Global.Get("process") != js.Undefined
	})
	return isNodeJS
}

func IsBrowser() bool {
	return !IsNodeJS()
}

func IsIOSSafari() bool {
	if IsNodeJS() {
		return false
	}
	ua := js.Global.Get("navigator").Get("userAgent").String()
	if !strings.Contains(ua, "iPhone") {
		return false
	}
	return true
}

func IsAndroidChrome() bool {
	if IsNodeJS() {
		return false
	}
	ua := js.Global.Get("navigator").Get("userAgent").String()
	if !strings.Contains(ua, "Android") {
		return false
	}
	if !strings.Contains(ua, "Chrome") {
		return false
	}
	return true
}

func IsMobileBrowser() bool {
	return IsIOSSafari() || IsAndroidChrome()
}
