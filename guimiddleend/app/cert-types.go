package app

import "encoding/json"

// CaRequest
//
// swagger:model CaRequest
type CaRequest struct {
	// The issuing cluster name
	// required: true
	// example: issuer1
	IssuingCluster string `json:"issuingCluster"`

	// The list of requesting clusters
	// required: true
	// items.example: cluster1
	RequestingClusters []string `json:"requestingClusters" example:"sample"`
}

type CaCert struct {
	Metadata CaCertMetadata `json:"metadata"`
	Spec     CaCertSpec     `json:"spec"`
}

type CaCertMetadata struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	UserData1   string `json:"userData1"`
	UserData2   string `json:"userData2"`
}

type CaCertSpec struct {
	CSRInfo        CaCertSpecCSRInfo        `json:"csrInfo"`
	Duration       string                   `json:"duration"`
	IsCA           bool                     `json:"isCA"`
	IssuerRef      CaCertSpecIssuerRef      `json:"issuerRef"`
	IssuingCluster CaCertSpecIssuingCluster `json:"issuingCluster"`
}

type CaCertSpecIssuerRef struct {
	Group string `json:"group"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
}

type CaCertSpecIssuingCluster struct {
	Cluster         string `json:"cluster"`
	ClusterProvider string `json:"clusterProvider"`
}

type CaCertSpecCSRInfo struct {
	Algorithm CaCertSpecCSRInfoAlgorithm `json:"algorithm"`
	KeySize   int                        `json:"keySize"`
	Subject   CaCertSpecCSRInfoSubject   `json:"subject"`
	Version   int                        `json:"version"`
}

type CaCertSpecCSRInfoAlgorithm struct {
	PublicKeyAlgorithm string `json:"publicKeyAlgorithm"`
	SignatureAlgorithm string `json:"signatureAlgorithm"`
}

type CaCertSpecCSRInfoSubject struct {
	Locale       struct{}                      `json:"locale"`
	Names        CaCertSpecCSRInfoSubjectNames `json:"names"`
	Organization struct{}                      `json:"organization"`
}

type CaCertSpecCSRInfoSubjectNames struct {
	CommonName       string `json:"CommonName"`
	CommonNamePrefix string `json:"commonNamePrefix"`
}

type CertStatus struct {
	ClusterProvider string                `json:"clusterProvider"`
	Clusters        []CertStatusCluster   `json:"clusters"`
	DeployedStatus  string                `json:"deployedStatus"`
	ReadyCounts     CertStatusReadyCounts `json:"readyCounts"`
	ReadyStatus     string                `json:"readyStatus"`
	States          CertStatusStates      `json:"states"`
}

type CertStatusStates struct {
	Actions     []CertStatusStatesAction `json:"actions"`
	Statusctxid string                   `json:"statusctxid"`
}

type CertStatusStatesAction struct {
	Instance string `json:"instance"`
	Revision int    `json:"revision"`
	State    string `json:"state"`
	Time     string `json:"time"`
}

type CertStatusReadyCounts struct {
	NotPresent int `json:"NotPresent"`
}

type CertStatusCluster struct {
	Cluster         string                     `json:"cluster"`
	ClusterProvider string                     `json:"clusterProvider"`
	Connectivity    string                     `json:"connectivity"`
	Resources       []CerStatusClusterResource `json:"resources"`
}

type CerStatusClusterResource struct {
	Gvk         CerStatusClusterResourceGvk `json:"GVK"`
	Name        string                      `json:"name"`
	ReadyStatus string                      `json:"readyStatus"`
}

type CerStatusClusterResourceGvk struct {
	Group   string `json:"Group"`
	Kind    string `json:"Kind"`
	Version string `json:"Version"`
}

type ClusterWithLabel struct {
	Labels []struct {
		ClusterLabel string `json:"clusterLabel"`
	} `json:"labels"`
	Metadata apiMetaData `json:"metadata"`
}

type jsonResponse struct {
	Data       interface{}       `json:"data"`
	Errors     map[string]string `json:"errors"`
	Error      string            `json:"Error"`
	IsSuccess  bool              `json:"isSuccess"`
	StatusCode int               `json:"statusCode"`
}

func (j jsonResponse) Byte() []byte {
	js, _ := json.Marshal(j)
	return js
}

type CertCluster struct {
	Metadata CertClusterMetadata `json:"metadata"`
	Spec     ClusterSpec         `json:"spec"`
}

type CertClusterMetadata struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	UserData1   string `json:"userData1"`
	UserData2   string `json:"userData2"`
}

type CertClusterSpec struct {
	GitOps CertClusterSpecGitOps `json:"gitOps"`
}

type CertClusterSpecGitOps struct {
	GitOpsReferenceObject string `json:"gitOpsReferenceObject"`
	GitOpsResourceObject  string `json:"gitOpsResourceObject"`
	GitOpsType            string `json:"gitOpsType"`
}

type CertIntentCluster struct {
	Metadata apiMetaData           `json:"metadata"`
	Spec     CertIntentClusterSpec `json:"spec"`
}

type CertIntentClusterSpec struct {
	Cluster         string `json:"cluster"`
	ClusterProvider string `json:"clusterProvider"`
	Scope           string `json:"scope"`
}

type CertUpdateClustersResponse struct {
	Created []string `json:"created"`
	Deleted []string `json:"deleted"`
}

// CaCertLogicalCloud holds the caCert logicalCloud details
type CaCertLogicalCloud struct {
	Metadata apiMetaData            `json:"metadata"`
	Spec     CaCertLogicalCloudSpec `json:"spec"`
}

// CaCertLogicalCloudSpec holds the logicalCloud details
type CaCertLogicalCloudSpec struct {
	LogicalCloud string `json:"logicalCloud"` // name of the logicalCloud
}
