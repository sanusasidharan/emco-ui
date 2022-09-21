package app

import "net/http"

func RegisterCertLCHandlers(handle HandleFunc, bootConf MiddleendConfig) {

	// Cert Logical cloud handlers

	// swagger:route POST /projects/{project}/caRequest LogicalCloudCaRequest LogicalCloudcaRequestPOST
	// Create CA request
	//  Parameters:
	//  + name: project
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
	handle("/projects/{project}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&lcHandler{createInstance(bootConf, r)}).caRequest(w, r) }).Methods("POST")

	// swagger:route DELETE /projects/{project}/caRequest LogicalCloudCaRequest LogicalCloudcaRequestDelete
	// Delete CA request
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/projects/{project}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&lcHandler{createInstance(bootConf, r)}).caDelete(w, r) }).Methods("DELETE")

	// swagger:route GET /projects/{project}/caRequest LogicalCloudCaRequest LogicalCloudGetCaCert
	// Get CA Request
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCert
	// default: JsonResponseError
	handle("/projects/{project}/caRequest", func(w http.ResponseWriter, r *http.Request) { (&lcHandler{createInstance(bootConf, r)}).caCert(w, r) }).Methods("GET")

	// swagger:route GET /projects/{project}/caRequest/logical-clouds LogicalCloudCaRequest LogicalCloudGetCaLogicalClouds
	// Get CA Logical Clouds
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponseLogicalClouds
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/logical-clouds", func(w http.ResponseWriter, r *http.Request) { (&lcHandler{createInstance(bootConf, r)}).caLClouds(w, r) }).Methods("GET")

	// swagger:route PUT /projects/{project}/caRequest/logical-clouds LogicalCloudCaRequest LogicalCloudUpdateClusters
	// Update Cert Logical Clouds
	//  Parameters:
	//  + name: project
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
	handle("/projects/{project}/caRequest/clusters", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caUpdateClouds(w, r)
	}).Methods("PUT")

	// swagger:route GET /cluster-provider/{project}/caRequest/enrollment/status LogicalCloudEnrollment LogicalCloudenrollmentStatus
	// Get enrollment status
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCertStatus
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/enrollment/status", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caGetEnrollmentStatus(w, r)
	}).Methods("GET")

	// swagger:route GET /cluster-provider/{project}/caRequest/distribution/status LogicalCloudDistribution LogicalClouddistributionStatus
	// Get distribution status
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: JsonResponseCertStatus
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/distribution/status", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caGetDistributionStatus(w, r)
	}).Methods("GET")

	// swagger:route POST /cluster-provider/{project}/caRequest/enrollment/instantiate LogicalCloudEnrollment LogicalCloudEnrollmentInstantiate
	// Instantiate Cert Enrollment
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/enrollment/instantiate", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caEnrollmentInstantiate(w, r)
	}).Methods("POST")

	// swagger:route POST /cluster-provider/{project}/caRequest/distribution/instantiate LogicalCloudDistribution LogicalCloudDistributionInstantiate
	// Instantiate Cert Distribution
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/distribution/instantiate", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caDistributionInstantiate(w, r)
	}).Methods("POST")

	// swagger:route POST /projects/{project}/caRequest/enrollment/terminate LogicalCloudEnrollment LogicalCloudEnrollmentTerminate
	// Terminate Cert Enrollment
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/enrollment/terminate", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caEnrollmentTerminate(w, r)
	}).Methods("POST")

	// swagger:route POST /projects/{project}/caRequest/distribution/terminate LogicalCloudDistribution LogicalCloudDistributionTerminate
	// Terminate Cert Distribution
	//  Parameters:
	//  + name: project
	//  in: path
	//  description: Cluster provider name
	//  required: true
	//  type: string
	// responses:
	// 200: swaggerJsonResponse
	// default: JsonResponseError
	handle("/projects/{project}/caRequest/distribution/terminate", func(w http.ResponseWriter, r *http.Request) {
		(&lcHandler{createInstance(bootConf, r)}).caDistributionTerminate(w, r)
	}).Methods("POST")
}