package app

// nolint
// swaggerJsonResponse
// swagger:response swaggerJsonResponse
type swaggerJsonResponse struct {
	// in: body
	Body JsonResponseInterface
}

// nolint
// JsonResponseCert
// swagger:response JsonResponseCert
type swaggerJsonResponseCert struct {
	// in: body
	Body JsonResponseCert
}

// nolint
// JsonResponseCertStatus
// swagger:response JsonResponseCertStatus
type swaggerJsonResponseCertStatus struct {
	// in: body
	Body JsonResponseCertStatus
}

// nolint
// JsonResponseClusters
// swagger:response JsonResponseClusters
type swaggerJsonResponseClusters struct {
	// in: body
	Body JsonResponseCertIntentClusters
}

// nolint
// JsonResponseError
// swagger:response JsonResponseError
type swaggerJsonErrorResponse struct {
	// in: body
	Body JsonResponseInterface
}

// nolint
// JsonResponseUpdateClusters
// swagger:response JsonResponseUpdateClusters
type swaggerJsonResponseUpdateClusters struct {
	// in: body
	Body JsonResponseUpdateClusters
}

type JsonResponseInterface struct {
	Data string `json:"data"`
	jsonResponse
}

type JsonResponseCertStatus struct {
	Data *CertStatus `json:"data"`
	jsonResponse
}

type JsonResponseCertIntentClusters struct {
	Data []*CertIntentCluster `json:"data"`
	jsonResponse
}

type JsonResponseCert struct {
	Data *CaCert `json:"data"`
	jsonResponse
}

type JsonResponseUpdateClusters struct {
	Data *CertUpdateClustersResponse `json:"data"`
	jsonResponse
}

// UpdateClusters
//
// swagger:model UpdateClusters
type SwaggerUpdateClusters []string

type JsonResponseCertIntentLogicalClouds struct {
	Data []*CaCertLogicalCloud `json:"data"`
	jsonResponse
}

// nolint
// swaggerJsonResponseLogicalClouds
// swagger:response swaggerJsonResponseLogicalClouds
type swaggerJsonResponseLogicalClouds struct {
	// in: body
	Body JsonResponseCertIntentLogicalClouds
}
