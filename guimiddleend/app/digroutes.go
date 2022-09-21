package app

import "net/http"

func RegisterDIGHandlers(handle HandleFunc, bootConf MiddleendConfig) {
	handle(digUriPattern, func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetAllDigs(w, r)
	}).Methods("GET")

	handle(digUriPattern, func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DelDig(w, r)
	}).Methods("DELETE")

	handle(digUriPattern+"/", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DelDig(w, r)
	}).Queries("operation", "{operation}").Methods("DELETE")

	handle(digUriPattern+"/status", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetDigStatus(w, r)
	}).Methods("GET")

	// DIG migrate/update/rollback related APIs
	handle(digUriPattern+"/checkout", func(w http.ResponseWriter, r *http.Request) {
		_ = createInstance(bootConf, r).GetDigInEdit(w, r)
	}).Methods("GET")

	handle(digUriPattern+"/checkout", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CheckoutDIG(w, r)
	}).Methods("POST")

	handle(digUriPattern+"/checkout/", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).CheckoutDIG(w, r)
	}).Queries("operation", "{operation}").Methods("POST")

	handle(digUriPattern+"/checkout/submit", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).UpgradeDIG(w, r)
	}).Methods("POST")

	handle(digUriPattern+"/checkout", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DigUpdateHandler(w, r)
	}).Methods("PUT")

	handle(digUriPattern+"/checkout/", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DigUpdateHandler(w, r)
	}).Queries("operation", "{operation}").Methods("PUT")

	handle(digUriPattern+"/scaleout", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).ScaleOutDig(w, r)
	}).Methods("POST")

	// GAC related APIs
	handle(digUriPattern+"/resources", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).GetK8sResources(w, r)
	}).Methods("GET")

	handle(digUriPattern+"/resources/{resourceName}", func(w http.ResponseWriter, r *http.Request) {
		createInstance(bootConf, r).DeleteK8sResources(w, r)
	}).Methods("DELETE")

	handle(digUriPattern+"/resources/{resourceName}/customizations/{customizationName}",
		func(w http.ResponseWriter, r *http.Request) {
			createInstance(bootConf, r).DeleteK8sResourceCustomizations(w, r)
		}).Methods("DELETE")
}
