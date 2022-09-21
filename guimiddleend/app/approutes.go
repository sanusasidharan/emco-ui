package app

import "net/http"

func RegisterApplicationHandlers(handle HandleFunc, bootConf MiddleendConfig) {
	// APIs related to service checkout
	handle(cAppUriPattern+"/{version}/checkout", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CreateDraftCompositeApp(w, r)
	}).Methods("POST")

	handle(cAppUriPattern+"/versions", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetSvcVersions(w, r)
	}).Methods("GET")

	handle(cAppUriPattern+"/versions/", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetSvcVersions(w, r)
	}).Queries("state", "{state}")

	handle(cAppUriPattern+"/{version}/app", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).UpdateCompositeApp(w, r)
	}).Methods("POST")
	handle(cAppUriPattern+"/{version}/apps/{appName}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).RemoveApp(w, r)
	}).Methods("DELETE")

	handle(cAppUriPattern+"/{version}/update", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CreateService(w, r)
	}).Methods("POST")

	// POST, GET, DELETE composite apps
	handle("/projects/{projectName}/composite-apps", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CreateApp(w, r)
	}).Methods("POST")

	handle(cAppUriPattern+"/{version}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetSvc(w, r)
	}).Methods("GET")

	handle("/projects/{projectName}/composite-apps", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetSvc(w, r)
	}).Methods("GET")

	handle("/projects/{projectName}/composite-apps", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetSvc(w, r)
	}).Queries("filter", "{filter}")

	handle(cAppUriPattern+"/{version}", func(w http.ResponseWriter, r *http.Request) {
		_ = createInstance(bootConf, r).DelSvc(w, r)
	}).Methods("DELETE")

	// POST, GET, DELETE deployment intent groups
	handle(cAppUriPattern+"/{version}/deployment-intent-groups", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CreateDig(w, r)
	}).Methods("POST")

	handle("/projects/{projectName}/deployment-intent-groups", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetAllDigs(w, r)
	}).Methods("GET")

	handle(cAppUriPattern+"/{version}/deployment-intent-groups", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetAllDigs(w, r)
	}).Methods("GET")
}
