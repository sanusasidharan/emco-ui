// Package classification Middleend API.
//
//     Schemes: http, https
//     Host: localhost:3000
//     BasePath: /middleend
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package app

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type HandleFunc func(string, func(http.ResponseWriter, *http.Request)) *mux.Route

const (
	cAppUriPattern = "/projects/{projectName}/composite-apps/{compositeAppName}"
	digUriPattern  = cAppUriPattern + "/{version}/deployment-intent-groups/{deploymentIntentGroupName}"
)

func createInstance(bootConf MiddleendConfig, r *http.Request) *OrchestrationHandler {
	o := NewAppHandler()

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	if bootConf.LogLevel == "debug" {
		l.SetLevel(logrus.DebugLevel)
	}
	fields := logrus.Fields{
		"uuid":   uuid.New(),
		"from":   r.RemoteAddr,
		"method": r.Method,
		"url":    r.URL.String(),
	}
	o.Logger = l.WithFields(fields)
	o.MiddleendConf = bootConf
	return o
}

// RegisterHandlers is a helper function for net/http/HandleFunc
// This function was introduced as a workaround for a concurrency issue in the middleend code
// The receiver for handlers functions was not thread-safe.
// All ServeHTTP functions required a refactoring to instantiate a new instance of
// the receiver, OrchestrationHandler, on every request.
// Ideally, future endpoints should not be using the closure for creating instance,
// instead do it as part handler intialization
func RegisterHandlers(handle HandleFunc, bootConf MiddleendConfig) {
	rapiopts := middleware.RapiDocOpts{SpecURL: "/middleend/swagger.yaml", BasePath: "/middleend/", Path: "/rapidocs"}
	rapidoc := middleware.RapiDoc(rapiopts, nil)

	redocopts := middleware.RedocOpts{SpecURL: "/middleend/swagger.yaml", BasePath: "/middleend", Path: "/redocs"}
	redoc := middleware.Redoc(redocopts, nil)

	handle("/redocs", redoc.ServeHTTP).Methods("GET")
	handle("/rapidocs", rapidoc.ServeHTTP).Methods("GET")
	handle("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		dat, err := os.ReadFile("./swagger.yaml")
		if err != nil {
			w.WriteHeader(404)
			return
		}
		if _, err := w.Write(dat); err != nil {
			log.Error(err, PrintFunctionName())
		}
	}).Methods("GET")

	handle("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetHealth(w)
	}).Methods("GET")

	RegisterApplicationHandlers(handle, bootConf)
	RegisterDIGHandlers(handle, bootConf)

	// ClusterProvider/Cluster creation APIs
	handle("/cluster-providers", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CreateClusterProvider(w, r)
	}).Methods("POST")
	handle("/cluster-providers/{clusterProvider}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DeleteClusterProvider(w, r)
	}).Methods("DELETE")
	handle("/cluster-providers/{cluster-provider-name}/clusters", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CheckConnection(w, r)
	}).Methods("POST")
	handle("/all-clusters", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetClusters(w, r)
	}).Methods("GET")

	// GET dashboard
	handle("/projects/{projectName}/dashboard", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetDashboardData(w, r)
	}).Methods("GET")

	// Logical Cloud related APIs
	handle("/projects/{projectName}/logical-clouds", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).HandleLCCreateRequest(w, r)
	}).Methods("POST")
	handle("/projects/{projectName}/logical-clouds", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetLogicalClouds(w, r)
	}).Methods("GET")
	handle("/projects/{projectName}/logical-clouds/{logicalCloud}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DeleteLogicalCloud(w, r)
	}).Methods("DELETE")
	handle("/projects/{projectName}/logical-clouds/{logicalCloud}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).UpdateLogicalCloud(w, r)
	}).Methods("PUT")
	handle("/projects/{projectName}/logical-cloud/{logicalCloud}/status", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetLogicalCloudsStatus(w, r)
	}).Methods("GET")

	// Get cluster networks
	handle("/cluster-providers/{clusterprovider-name}/clusters/{cluster-name}/networks",
		func(w http.ResponseWriter, r *http.Request) {
			createInstance(bootConf, r).GetClusterNetworks(w, r)
		}).Methods("GET")

	RegisterCertCPHandlers(handle, bootConf)
	RegisterCertLCHandlers(handle, bootConf)
}
