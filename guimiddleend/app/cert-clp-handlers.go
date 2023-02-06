package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type clpHandler struct {
	*OrchestrationHandler
}

func (h *clpHandler) jsonOK(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode == 0 {
		statusCode = 200
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonResponse{
		Data:       data,
		Errors:     make(map[string]string),
		IsSuccess:  true,
		StatusCode: statusCode,
	}.Byte()); err != nil {
		h.Logger.Error(err)
	}
}

func (h *clpHandler) jsonError(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode == 0 {
		statusCode = 500
	}
	h.Logger.Error(err)
	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonResponse{
		IsSuccess:  false,
		Errors:     make(map[string]string),
		Error:      err,
		StatusCode: statusCode,
	}.Byte()); err != nil {
		h.Logger.Error(err)
	}
}

func (h *clpHandler) caEnrollmentInstantiate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Enrollment instantiante request")
	statusCode, err := h.InstantiateEnrollment(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}
	h.jsonOK(w, "Certificate enrollment istantiated", statusCode)
}

func (h *clpHandler) caEnrollmentTerminate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Enrollment terminate request")
	statusCode, err := h.TerminateEnrollment(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}
	h.jsonOK(w, "Certificate enrollment terminated", statusCode)
}

func (h *clpHandler) caDistributionInstantiate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Distribution instantiante request")
	statusCode, err := h.InstantiateDistribution(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.jsonOK(w, "Certificate distribution istantiated", statusCode)
}

func (h *clpHandler) caDistributionTerminate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Distribution terminate request")
	statusCode, err := h.TerminateDistribution(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.jsonOK(w, "Certificate distribution terminated", statusCode)
}

func (h *clpHandler) caGetEnrollmentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Getting CA Cert enrollment status")
	enrStatus, err := h.GetCaCertEnrollmentStatus(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, enrStatus, 0)
}

func (h *clpHandler) caGetDistributionStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Getting CA Cert distribution status")
	enrStatus, err := h.GetCaCertDistributionStatus(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, enrStatus, 0)
}

func (h *clpHandler) caCert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Get CACert")
	cert, err := h.GetCaCert(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, cert, 0)
}

func (h *clpHandler) caClusters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)

	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Getting CA Clusters")

	clusters, err := h.GetCAIntentClusters(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, clusters, 0)
}

func (h *clpHandler) caDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)
	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("CA Delete")

	clusters, err := h.GetClustersByProvider(clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	statusCode, err := h.caCertTerminate(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}

	for _, cl := range clusters {
		err := h.DeleteCertCluster(caIntent, clusterProvider, cl.Metadata.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	statusCode, err = h.DeleteCaCert(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}

	successMsg := "Ca Request succesfully deleted"
	h.Logger.Info(successMsg)
	h.jsonOK(w, successMsg, 200)
}

func (h *clpHandler) caUpdateClusters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)

	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("Updating CA Clusters")
	reqClusterNames := []string{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(payload, &reqClusterNames)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Logger.WithField("payload", string(payload)).Debug("Payload")

	clusters, err := h.GetClustersByProvider(clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	caClusters, err := h.GetCAIntentClusters(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqClusters, err := findClustersByNames(clusters, reqClusterNames)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqClustersMap := make(map[string]*ClusterWithLabel)
	for _, cl := range reqClusters {
		reqClustersMap[cl.Metadata.Name] = cl
	}

	caClustersMap := make(map[string]*CertIntentCluster)
	for _, cl := range caClusters {
		caClustersMap[cl.Metadata.Name] = cl
	}

	clustersForCreate := []*ClusterWithLabel{}
	createdClusters := []string{}
	for _, cl := range reqClusters {
		if _, ok := caClustersMap[cl.Metadata.Name]; !ok {
			clustersForCreate = append(clustersForCreate, cl)
			createdClusters = append(createdClusters, cl.Metadata.Name)
		}
	}

	err = h.CreateCertClusters(caIntent, clusterProvider, clustersForCreate)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	clustersForDelete := []string{}
	for _, cl := range caClusters {
		if _, ok := reqClustersMap[cl.Metadata.Name]; !ok {
			clustersForDelete = append(clustersForDelete, cl.Metadata.Name)
		}
	}

	err = h.DeleteCertClusters(caIntent, clusterProvider, clustersForDelete)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	logMsg := "Not need for clusters update"

	if len(createdClusters) > 0 || len(clustersForDelete) > 0 {
		logMsg = "CA Clusters successfully updated"
		statusCode, err := h.caCertReInstantiate(caIntent, clusterProvider)
		if err != nil {
			h.jsonError(w, err.Error(), statusCode)
			return
		}
	}

	response := &CertUpdateClustersResponse{
		Created: createdClusters,
		Deleted: clustersForDelete,
	}

	response_json, _ := json.Marshal(response)
	h.Logger.WithField("response", string(response_json)).Info(logMsg)

	h.jsonOK(w, response, 0)
}

func (h *clpHandler) caRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clusterProvider := vars["clusterprovider-name"]
	caIntent := h.generateCaIntentName(clusterProvider)

	h.Logger = h.Logger.WithFields(logrus.Fields{"clusterProvider": clusterProvider, "caIntent": caIntent, "function": PrintFunctionName()})
	h.Logger.Info("CA Request")

	caRequest := &CaRequest{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(payload, caRequest)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	clusters, err := h.GetClustersByProvider(clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	caCert := &CaCert{
		Metadata: CaCertMetadata{
			Name: caIntent,
		},
		Spec: CaCertSpec{
			CSRInfo: CaCertSpecCSRInfo{
				KeySize: 4096,
				Version: 1,
				Algorithm: CaCertSpecCSRInfoAlgorithm{
					PublicKeyAlgorithm: "RSA",
					SignatureAlgorithm: "SHA512WithRSA",
				},
				Subject: CaCertSpecCSRInfoSubject{
					Names: CaCertSpecCSRInfoSubjectNames{CommonNamePrefix: clusterProvider},
				},
			},
			Duration: "8760h",
			IsCA:     true,
			IssuerRef: CaCertSpecIssuerRef{
				Name:  "new-istio-system",
				Kind:  "ClusterIssuer",
				Group: "cert-manager.io",
			},
			IssuingCluster: CaCertSpecIssuingCluster{
				Cluster:         caRequest.IssuingCluster,
				ClusterProvider: clusterProvider,
			},
		},
	}

	statusCode, err := h.PostCaCert(caCert, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}

	clusterList, err := findClustersByNames(clusters, caRequest.RequestingClusters)
	if err != nil {
		h.Logger.Error(err, "Rolling back")
		_, delErr := h.DeleteCaCert(caIntent, clusterProvider)
		if delErr != nil {
			h.Logger.Error(delErr)
		}
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.CreateCertClusters(caIntent, clusterProvider, clusterList)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		_, err := h.DeleteCaCert(caIntent, clusterProvider)
		if err != nil {
			h.Logger.Error(err)
		}
		return
	}

	statusCode, err = h.caCertInstantiate(caIntent, clusterProvider)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		return
	}

	successMsg := "Ca Request succesfully completed"
	h.Logger.Info(successMsg)

	h.jsonOK(w, successMsg, 0)
}
