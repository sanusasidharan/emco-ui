package app

import "net/http"

func RegisterCertCPHandlers(handle HandleFunc, bootConf MiddleendConfig) {
	// swagger:route POST /cluster-provider/{clusterprovider-name}/caRequest ClusterProviderCaRequest ClusterProvidercaRequestPOST
	// Create CA request
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	//  + name: input
	//  in: body
	//  type: CaRequest
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&clpHandler{createInstance(bootConf, r)}).caRequest(w, r) }).Methods("POST")

	// swagger:route DELETE /cluster-provider/{clusterprovider-name}/caRequest ClusterProviderCaRequest ClusterProvidercaRequestDelete
	// Delete CA request
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&clpHandler{createInstance(bootConf, r)}).caDelete(w, r) }).Methods("DELETE")

	// swagger:route GET /cluster-provider/{clusterprovider-name}/caRequest ClusterProviderCaRequest ClusterProviderGetCaCert
	// Get CA Request
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCert
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&clpHandler{createInstance(bootConf, r)}).caCert(w, r) }).Methods("GET")

	// swagger:route GET /cluster-provider/{clusterprovider-name}/caRequest/clusters ClusterProviderCaRequest ClusterProviderGetCaClusters
	// Get CA clusters
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseClusters
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/clusters", func(w http.ResponseWriter, r *http.Request) { (&clpHandler{createInstance(bootConf, r)}).caClusters(w, r) }).Methods("GET")

	// swagger:route PUT /cluster-provider/{clusterprovider-name}/caRequest/clusters ClusterProviderCaRequest ClusterProviderUpdateClusters
	// Update Cert Clusters
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	//  + name: input
	//  in: body
	//  type: UpdateClusters
	// responses:
	// 200: JsonResponseUpdateClusters
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/clusters", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caUpdateClusters(w, r)
	}).Methods("PUT")

	// swagger:route GET /cluster-provider/{clusterprovider-name}/caRequest/enrollment/status ClusterProviderEnrollment ClusterProviderenrollmentStatus
	// Get enrollment status
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCertStatus
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/enrollment/status", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caGetEnrollmentStatus(w, r)
	}).Methods("GET")

	// swagger:route GET /cluster-provider/{clusterprovider-name}/caRequest/distribution/status ClusterProviderDistribution ClusterProviderdistributionStatus
	// Get distribution status
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCertStatus
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/distribution/status", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caGetDistributionStatus(w, r)
	}).Methods("GET")

	// swagger:route POST /cluster-provider/{clusterprovider-name}/caRequest/enrollment/instantiate ClusterProviderEnrollment ClusterProviderEnrollmentInstantiate
	// Instantiate Cert Enrollment
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/enrollment/instantiate", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caEnrollmentInstantiate(w, r)
	}).Methods("POST")

	// swagger:route POST /cluster-provider/{clusterprovider-name}/caRequest/distribution/instantiate ClusterProviderDistribution ClusterProviderDistributionInstantiate
	// Instantiate Cert Distribution
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/distribution/instantiate", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caDistributionInstantiate(w, r)
	}).Methods("POST")

	// swagger:route POST /cluster-provider/{clusterprovider-name}/caRequest/enrollment/terminate ClusterProviderEnrollment ClusterProviderEnrollmentTerminate
	// Terminate Cert Enrollment
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/enrollment/terminate", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caEnrollmentTerminate(w, r)
	}).Methods("POST")

	// swagger:route POST /cluster-provider/{clusterprovider-name}/caRequest/distribution/terminate ClusterProviderDistribution ClusterProviderDistributionTerminate
	// Terminate Cert Distribution
	//  Parameters:
	//  + name: clusterprovider-name
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/cluster-provider/{clusterprovider-name}/caRequest/distribution/terminate", func(w http.ResponseWriter, r *http.Request) {
		(&clpHandler{createInstance(bootConf, r)}).caDistributionTerminate(w, r)
	}).Methods("POST")
}
