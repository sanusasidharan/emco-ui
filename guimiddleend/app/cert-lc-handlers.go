package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type lcHandler struct {
	*OrchestrationHandler
}

func (h *lcHandler) jsonOK(w http.ResponseWriter, data interface{}, statusCode int) {
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

func (h *lcHandler) jsonError(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode == 0 {
		statusCode = 500
	}
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

func (h *lcHandler) caEnrollmentInstantiate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	statusCode, err := h.InstantiateEnrollment(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}
	h.jsonOK(w, "Certificate enrollment istantiated", statusCode)
}

func (h *lcHandler) caEnrollmentTerminate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	statusCode, err := h.TerminateEnrollment(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}
	h.jsonOK(w, "Certificate enrollment terminated", statusCode)
}

func (h *lcHandler) caDistributionInstantiate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	statusCode, err := h.InstantiateDistribution(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	h.jsonOK(w, "Certificate distribution istantiated", statusCode)
}

func (h *lcHandler) caDistributionTerminate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	statusCode, err := h.TerminateDistribution(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	h.jsonOK(w, "Certificate distribution terminated", statusCode)
}

func (h *lcHandler) caGetEnrollmentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)
	enrStatus, err := h.GetCaCertEnrollmentStatus(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, enrStatus, 0)
}

func (h *lcHandler) caGetDistributionStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)
	enrStatus, err := h.GetCaCertDistributionStatus(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, enrStatus, 0)
}

func (h *lcHandler) caCert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)
	cert, err := h.GetCaCert(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, cert, 0)
}

func (h *lcHandler) caLClouds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)
	clouds, err := h.GetCAIntentLogicalClouds(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusNotFound)
		return
	}
	h.jsonOK(w, clouds, 0)
}

func (h *lcHandler) caDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	clouds, err := h.GetLogicalCloudsByProject(project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	statusCode, err := h.caCertTerminate(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}

	for _, cl := range clouds {
		err := h.DeleteCertLogicalCloud(caIntent, project, cl.Metadata.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}

	statusCode, err = h.DeleteCaCert(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}
	h.jsonOK(w, "Ca Request succesfully deleted", 200)
}

func (h *lcHandler) caUpdateClouds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	reqCloudNames := []string{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = json.Unmarshal(payload, &reqCloudNames)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	clouds, err := h.GetLogicalCloudsByProject(project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	caClouds, err := h.GetCAIntentLogicalClouds(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	reqClusters, err := findCloudsByNames(clouds, reqCloudNames)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	reqClustersMap := make(map[string]*CaCertLogicalCloud)
	for _, cl := range reqClusters {
		reqClustersMap[cl.Metadata.Name] = cl
	}

	caClustersMap := make(map[string]*CaCertLogicalCloud)
	for _, cl := range caClouds {
		caClustersMap[cl.Metadata.Name] = cl
	}

	clustersForCreate := []*CaCertLogicalCloud{}
	createdClusters := []string{}
	for _, cl := range reqClusters {
		if _, ok := caClustersMap[cl.Metadata.Name]; !ok {
			clustersForCreate = append(clustersForCreate, cl)
			createdClusters = append(createdClusters, cl.Metadata.Name)
		}
	}

	err = h.CreateCertLogicalClouds(caIntent, project, clustersForCreate)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	clustersForDelete := []string{}
	for _, cl := range caClouds {
		if _, ok := reqClustersMap[cl.Metadata.Name]; !ok {
			clustersForDelete = append(clustersForDelete, cl.Metadata.Name)
		}
	}

	err = h.DeleteCertLogicalClouds(caIntent, project, clustersForDelete)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	statusCode, err := h.caCertReInstantiate(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}

	response := &CertUpdateClustersResponse{
		Created: createdClusters,
		Deleted: clustersForDelete,
	}

	h.jsonOK(w, response, 0)
}

func (h *lcHandler) caRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	caIntent := h.generateCaIntentName(project)

	clouds, err := h.GetLogicalCloudsByProject(project)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	caRequest := &CaRequest{}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(payload, caRequest)
	if err != nil {
		log.Println(err)
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
					Names: CaCertSpecCSRInfoSubjectNames{CommonNamePrefix: project},
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
				ClusterProvider: project,
			},
		},
	}

	statusCode, err := h.PostCaCert(caCert, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}

	clusterList, err := findCloudsByNames(clouds, caRequest.RequestingClusters)
	if err != nil {
		log.Println(err)
		log.Println("Rolling back")
		_, delErr := h.DeleteCaCert(caIntent, project)
		if delErr != nil {
			log.Println(delErr)
		}
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.CreateCertLogicalClouds(caIntent, project, clusterList)
	if err != nil {
		h.jsonError(w, err.Error(), http.StatusBadRequest)
		_, err := h.DeleteCaCert(caIntent, project)
		if err != nil {
			log.Println(err)
		}
		return
	}

	statusCode, err = h.caCertInstantiate(caIntent, project)
	if err != nil {
		h.jsonError(w, err.Error(), statusCode)
		log.Println(err)
		return
	}

	h.jsonOK(w, "Ca Request succesfully completed", 0)
}
