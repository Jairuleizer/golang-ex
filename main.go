// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/Jairuleizer/golang-ex/controllers"
)

const (
	APP_VER = "0.1.1.0227"
)

var (
	servingCertFile = os.Getenv("SERVING_CERT")
	servingKeyFile  = os.Getenv("SERVING_KEY")
)

func main() {
	beego.Info(beego.AppConfig.String("appname"), APP_VER)

	// Register routers.
	beego.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// Long polling.
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	// serve securely if the certificates are present
	_, certErr := os.Stat(servingCertFile)
	_, keyErr := os.Stat(servingKeyFile)
	if certErr == nil && keyErr == nil && len(servingCertFile) > 0 && len(servingKeyFile) > 0 {
		beego.BConfig.Listen.HTTPSCertFile = servingCertFile
		beego.BConfig.Listen.HTTPSKeyFile = servingKeyFile
		beego.BConfig.Listen.EnableHTTPS = true
	}

	beego.Run()
}
