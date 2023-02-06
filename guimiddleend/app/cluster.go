package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Create Cluster Provider and cluster-sync-objects required for GitOps
func (h *OrchestrationHandler) CreateClusterProvider(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.Vars = vars
	var jsonData ClusterProvider
	h.InitializeResponseMap()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&jsonData)
	if err != nil {
		log.Errorf("Failed to parse json: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create Cluster Provider
	cp := ClusterProvider{
		Metadata: apiMetaData{
			Name:        jsonData.Metadata.Name,
			Description: jsonData.Metadata.Description,
		},
	}

	jsonLoad, _ := json.Marshal(cp)
	url := "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers"
	resp, err := h.apiPost(jsonLoad, url, jsonData.Metadata.Name+"_cp")
	if err != nil {
		log.Errorf("Encountered error while creating cluster provider: %s", jsonData.Metadata.Name)
		w.WriteHeader(resp.(int))
		return
	}

	clusterProvider := jsonData.Metadata.Name

	// Create cluster-sync-object, if required payload available
	if jsonData.Spec.GitEnabled && len(jsonData.Spec.Kv) > 0 {
		var kvinfo []map[string]interface{}
		for _, kvpair := range jsonData.Spec.Kv {
			log.Info("kvpair", log.Fields{"kvpair": kvpair})
			v, ok := kvpair["gitType"]
			if ok {
				gitType := map[string]interface{}{
					"gitType": fmt.Sprintf("%v", v),
				}
				kvinfo = append(kvinfo, gitType)
			}
			v, ok = kvpair["gitToken"]
			if ok {
				gitToken := map[string]interface{}{
					"gitToken": fmt.Sprintf("%v", v),
				}
				kvinfo = append(kvinfo, gitToken)
			}
			v, ok = kvpair["repoName"]
			if ok {
				repoName := map[string]interface{}{
					"repoName": fmt.Sprintf("%v", v),
				}
				kvinfo = append(kvinfo, repoName)
			}
			v, ok = kvpair["userName"]
			if ok {
				userName := map[string]interface{}{
					"userName": fmt.Sprintf("%v", v),
				}
				kvinfo = append(kvinfo, userName)
			}
			v, ok = kvpair["branch"]
			if ok {
				branch := map[string]interface{}{
					"branch": fmt.Sprintf("%v", v),
				}
				kvinfo = append(kvinfo, branch)
			}
		}
		jsonData.Spec.Kv = kvinfo
		jsonData.Metadata.Name = "GitObjectMyRepo"
		jsonLoad, _ := json.Marshal(jsonData)
		url := "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers/" + clusterProvider + "/cluster-sync-objects"
		resp, err := h.apiPost(jsonLoad, url, jsonData.Metadata.Name+"_cp")
		if err != nil {
			log.Errorf("Encountered error while creating cluster sync object for clusterprovider: %s", jsonData.Metadata.Name)
			w.WriteHeader(resp.(int))
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(h.response.payload[clusterProvider+"_cp"]); err != nil {
		log.Error(err)
	}
}

// GetClusters get an a array of all the cluster providers and the clusters within them
func (h *OrchestrationHandler) GetClusters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.Vars = vars
	h.InitializeResponseMap()
	dashboardClient := DashboardClient{h}
	retcode := dashboardClient.getAllClusters()
	if retcode != nil {
		if intval, ok := retcode.(int); ok {
			log.Infof("Failed to get clusterdata : %d", intval)
			w.WriteHeader(intval)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errMsg := string(h.response.payload[h.response.lastKey]) + h.response.lastKey
			if _, err := w.Write([]byte(errMsg)); err != nil {
				log.Error(err)
			}
		}
		return
	}

	var retval []byte
	retval, err := json.Marshal(h.ClusterProviders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(retval); err != nil {
		log.Error(err)
	}
}

// Delete Cluster Provider and cluster-sync-objects
func (h *OrchestrationHandler) DeleteClusterProvider(w http.ResponseWriter, r *http.Request) {
	h.Vars = mux.Vars(r)
	h.InitializeResponseMap()
	// Fetch cluster sync object
	url := "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers/" + h.Vars["clusterProvider"] + "/cluster-sync-objects"
	reply, err := h.apiGet(url, h.Vars["clusterProvider"])
	if err != nil {
		log.Errorf("Encountered error while fetching cluster sync object for clusterprovider: %s", h.Vars["clusterProvider"])
		w.WriteHeader(reply.StatusCode)
		return
	}
	var jsonData []ClusterProvider
	if err := json.Unmarshal(reply.Data, &jsonData); err != nil {
		log.Error(err, PrintFunctionName())
	}
	log.Infof("clustersyncobjects: %+v", jsonData)

	// Delete cluster sync object
	if len(jsonData) > 0 && len(jsonData[0].Spec.Kv) != 0 {
		url := "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers/" + h.Vars["clusterProvider"] + "/cluster-sync-objects/" + "GitObjectMyRepo"
		resp, err := h.apiDel(url, h.Vars["clusterProvider"])
		if err != nil {
			log.Errorf("Encountered error while deleting cluster sync object for clusterprovider: %s", h.Vars["clusterProvider"])
			w.WriteHeader(resp.(int))
			return
		}
		if resp != nil && resp.(int) != http.StatusNoContent {
			log.Errorf("Encountered error while deleting cluster sync object for clusterprovider: %s", h.Vars["clusterProvider"])
			w.WriteHeader(resp.(int))
			return
		}
	}

	// Delete cluster provider
	url = "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers/" + h.Vars["clusterProvider"]
	resp, err := h.apiDel(url, h.Vars["clusterProvider"])
	if err != nil {
		log.Errorf("Encountered error while deleting clusterprovider: %s", h.Vars["clusterProvider"])
		w.WriteHeader(resp.(int))
		return
	}
	if resp != nil && resp.(int) != http.StatusNoContent {
		log.Errorf("Encountered error while deleting clusterprovider: %s", h.Vars["clusterProvider"])
		w.WriteHeader(resp.(int))
		return
	}
	w.WriteHeader(resp.(int))
}
