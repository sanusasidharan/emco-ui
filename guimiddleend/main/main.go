//=======================================================================
// Copyright (c) 2017-2020 Aarna Networks, Inc.
// All rights reserved.
// ======================================================================
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//           http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ========================================================================

package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/middleend/app"
	"example.com/middleend/authproxy"
	"example.com/middleend/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
}

/* This is the main package of the middleend. This package
 * implements the http server which exposes service ar 9891.
 * It also intialises an API router which handles the APIs with
 * subpath /v1.
 */
func main() {
	authProxyHandler := authproxy.NewAppHandler()

	configFile, err := os.Open("/opt/emco/config/middleend.conf")
	if err != nil {
		log.WithError(err).Errorf("%s(): Failed to read middleend configuration", app.PrintFunctionName())
		return
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.WithError(err).Errorf("%s(): Failed to close read file handler", app.PrintFunctionName())
			return
		}
	}(configFile)

	// Read the configuration json
	byteValue, _ := ioutil.ReadAll(configFile)
	bootConf := &app.MiddleendConfig{}
	json.Unmarshal(byteValue, bootConf)
	json.Unmarshal(byteValue, &authProxyHandler.AuthProxyConf)

	// parse string, this is built-in feature of logrus
	logLevel, err := log.ParseLevel(bootConf.LogLevel)
	if err != nil {
		logLevel = log.DebugLevel
	}

	// set global log level
	log.SetLevel(logLevel)

	// Connect to the DB
	err = db.CreateDBClient("mongo", "middleend", bootConf.Mongo)
	if err != nil {
		log.Error("Failed to connect to DB")
		return
	}

	bootConf.StoreName = "middleend"
	// Get an instance of the OrchestrationHandler, this type implements
	// the APIs i.e CreateApp, ShowApp, DeleteApp.
	httpRouter := mux.NewRouter().PathPrefix("/middleend").Subrouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, httpRouter)
	log.Infof("%s(): Starting middle end service", app.PrintFunctionName())
	log.WithFields(log.Fields{
		"ownport":        bootConf.OwnPort,
		"orchestrator":   bootConf.OrchService,
		"clm":            bootConf.Clm,
		"cert":           bootConf.Cert,
		"dcm":            bootConf.Dcm,
		"ncm":            bootConf.Ncm,
		"gac":            bootConf.Gac,
		"dtc":            bootConf.Dtc,
		"its":            bootConf.Its,
		"ovnaction":      bootConf.OvnService,
		"configSvc":      bootConf.CfgService,
		"mongo":          bootConf.Mongo,
		"logLevel":       bootConf.LogLevel,
		"storeName":      bootConf.StoreName,
		"appInstantiate": bootConf.AppInstantiate,
	}).Infof("Middle End Configuration")

	httpServer := &http.Server{
		Handler:      loggedRouter,
		Addr:         ":" + bootConf.OwnPort,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	// Package level Handlers
	app.RegisterHandlers(httpRouter.HandleFunc, *bootConf)
	// Start server in a go routine.
	go func() {
		log.Fatal(httpServer.ListenAndServe())
	}()

	// Graceful shutdown of the server,
	// create a channel and wait for SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	log.Info("wait for signal")
	<-c
	log.Info("Bye Bye")
	httpServer.Shutdown(context.Background())
}
