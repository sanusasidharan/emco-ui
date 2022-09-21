//=======================================================================
// Copyright (c) 2017-2020 Aarna Networks, Inc.
// All rights reserved.
// ======================================================================
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//           http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ========================================================================

package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type logicalCloudData struct {
	Metadata apiMetaData      `json:"metadata"`
	Spec     logicalCloudSpec `json:"spec"`
}

// UserData contains the parameters needed for user
type UserData struct {
	UserName string `json:"userName"`
	Type     string `json:"type"`
}

// UserPermission contains the parameters needed for a user permission
type UserPermission struct {
	MetaData      UPMetaData `json:"metadata"`
	Specification UPSpec     `json:"spec"`
}

// UPMetaData contains the parameters needed for a user permission metadata
type UPMetaData struct {
	UserPermissionName string `json:"name"`
	Description        string `json:"description"`
	UserData1          string `json:"userData1"`
	UserData2          string `json:"userData2"`
}

// UPSpec contains the parameters needed for a user permission spec
type UPSpec struct {
	Namespace string   `json:"namespace"`
	APIGroups []string `json:"apiGroups"`
	Resources []string `json:"resources"`
	Verbs     []string `json:"verbs"`
}

// Quota contains the parameters needed for a Quota
type Quota struct {
	MetaData QMetaData `json:"metadata"`
	// Specification QSpec         `json:"spec"`
	Specification map[string]string `json:"spec"`
}

// QMetaData MetaData contains the parameters needed for metadata
type QMetaData struct {
	QuotaName   string `json:"name"`
	Description string `json:"description"`
	UserData1   string `json:"userData1"`
	UserData2   string `json:"userData2"`
}
type LogicalCloudStatus struct {
	Clusters []struct {
		Cluster         string `json:"cluster"`
		ClusterProvider string `json:"clusterProvider"`
		Connectivity    string `json:"connectivity"`
		Resources       []struct {
			Gvk struct {
				Group   string `json:"Group"`
				Kind    string `json:"Kind"`
				Version string `json:"Version"`
			} `json:"GVK"`
			Name        string `json:"name"`
			ReadyStatus string `json:"readyStatus"`
		} `json:"resources"`
	} `json:"clusters"`
	DeployedStatus string `json:"deployedStatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	ReadyCounts    struct {
		NotPresent int `json:"NotPresent"`
		Ready      int `json:"Ready"`
	} `json:"readyCounts"`
	ReadyStatus string `json:"readyStatus"`
	States      struct {
		Actions []struct {
			Instance string    `json:"instance"`
			Revision int       `json:"revision"`
			State    string    `json:"state"`
			Time     time.Time `json:"time"`
		} `json:"actions"`
		Statusctxid string `json:"statusctxid"`
	} `json:"states"`
}
type QuotaInfo struct {
	LimitsCPU                   string `json:"limits.cpu"`
	LimitsMemory                string `json:"limits.memory"`
	RequestsCPU                 string `json:"requests.cpu"`
	RequestsMemory              string `json:"requests.memory"`
	RequestsStorage             string `json:"requests.storage"`
	LimitsEphemeralStorage      string `json:"limits.ephemeral.storage"`
	PersistentVolumeClaims      string `json:"persistentvolumeclaims"`
	Pods                        string `json:"pods"`
	ConfigMaps                  string `json:"configmaps"`
	ReplicationControllers      string `json:"replicationcontrollers"`
	ResourceQuotas              string `json:"resourcequotas"`
	Services                    string `json:"services"`
	ServicesLoadBalancers       string `json:"services.loadbalancers"`
	ServicesNodePorts           string `json:"services.nodeports"`
	Secrets                     string `json:"secrets"`
	CountReplicationControllers string `json:"count/replicationcontrollers"`
	CountDeploymentsApps        string `json:"count/deployments.apps"`
	CountReplicasetsApps        string `json:"count/replicasets.apps"`
	CountStatefulSets           string `json:"count/statefulsets.apps"`
	CountJobsBatch              string `json:"count/jobs.batch"`
	CountCronJobsBatch          string `json:"count/cronjobs.batch"`
	CountDeploymentsExtensions  string `json:"count/deployments.extensions"`
}

// Logical cloud spec
type logicalCloudSpec struct {
	NameSpace string   `json:"namespace"`
	Level     string   `json:"level"`
	UserData  UserData `json:"user"`
}

type ClusterLabels struct {
	Metadata apiMetaData `json:"metadata"`
	Labels   []Labels    `json:"labels"`
}

type clusterReferenceFlat struct {
	Metadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Userdata1   string `json:"userData1"`
		Userdata2   string `json:"userData2"`
	} `json:"metadata"`
	Spec struct {
		ClusterProvider string   `json:"clusterProvider"`
		ClusterName     string   `json:"cluster"`
		LoadbalancerIP  string   `json:"loadBalancerIP"`
		Certificate     string   `json:"certificate,omitempty"`
		LabelList       []Labels `json:"labels,omitempty"`
	} `json:"spec"`
}

type clusterReferenceNested struct {
	Metadata struct {
		Name                string   `json:"name"` // Logical cloud Name
		Description         string   `json:"description"`
		ClusterRefenceNames []string `json:"clusterReferencesNames,omitempty"`
	} `json:"metadata"`
	Spec struct {
		ClusterProvidersList []ClusterProviders `json:"clusterProviders"`
	} `json:"spec"`
}

type ClusterProviders struct {
	Metadata struct {
		Name        string `json:"name"` // Cluster Provider Name
		Description string `json:"description"`
	} `json:"metadata"`
	Spec struct {
		ClustersList []Clusters `json:"clusters"`
	} `json:"spec"`
}

type Clusters struct {
	Metadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Operation   string `json:"operation,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Labels []Labels `json:"labels"`
	} `json:"spec"`
}

type Labels struct {
	LabelName string `json:"clusterLabel"`
}

type LogicalClouds struct {
	Metadata struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Userdata1   string `json:"userData1"`
		Userdata2   string `json:"userData2"`
	} `json:"metadata"`
	Spec struct {
		Namespace string `json:"namespace"`
		Level     string `json:"level"`
		User      struct {
			UserName string `json:"userName"`
			Type     string `json:"type"`
		} `json:"user"`
		UserQuota               map[string]string      `json:"userQuota"`
		UserQuotaMetadata       QMetaData              `json:"userQuotaMetadata,omitempty"`
		UserPermissionSpec      []UPSpec               `json:"userPermissions"`
		UserPerminssionMetadata []UPMetaData           `json:"userPermissionsMeta,omitempty"`
		ClusterReferences       clusterReferenceNested `json:"clusterReferences,omitempty"`
		Status                  string                 `json:"status,omitempty"`
	} `json:"spec"`
}

type userPermissions struct {
	APIGroups []string `json:"apiGroups"`
	Resources []string `json:"resources"`
	Verbs     []string `json:"verbs"`
}

type logicalCloudsPayload struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	CloudType              string `json:"cloudType"`
	Namespace              string `json:"namespace"`
	EnableServiceDiscovery bool   `json:"enableServiceDiscovery"`
	Spec                   LogicalCloudSpec
}

type LogicalCloudSpec struct {
	Namespace            string             `json:"namespace"`
	User                 *UserData          `json:"user,omitempty"`
	Permissions          *userPermissions   `json:"permissions,omitempty"`
	Quotas               *QuotaInfo         `json:"quotas,omitempty"`
	ClusterProvidersList []ClusterProviders `json:"clusterProviders"`
}

type logicalCloudUpdatePayload struct {
	CloudType            string             `json:"cloudType"`
	Namespace            string             `json:"namespace"`
	Permissions          *userPermissions   `json:"permissions,omitempty"`
	Quotas               *QuotaInfo         `json:"quotas,omitempty"`
	ClusterProvidersList []ClusterProviders `json:"clusterProviders"`
}

// logicalCloudHandler implements the orchworkflow interface
type logicalCloudHandler struct {
	orchInstance *OrchestrationHandler
}

func (h *logicalCloudHandler) getLogicalClouds() (lc []LogicalClouds, err error) {
	orch := h.orchInstance
	lc = []LogicalClouds{}
	projectName := orch.Vars["projectName"]
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projectName + "/logical-clouds"
	reply, err := orch.apiGet(url, projectName)
	if err != nil {
		log.Errorf("%s(): Failed to GET LC for %s", PrintFunctionName(), projectName)
		return lc, err
	}
	log.Infof("%s(): Get LC status: %d", PrintFunctionName(), reply.StatusCode)
	if err := json.Unmarshal(reply.Data, &lc); err != nil {
		log.Error(err, PrintFunctionName())
	}
	return lc, err
}

func (h *logicalCloudHandler) getLogicalCloud(lcName string) (lc LogicalClouds, err error) {
	orch := h.orchInstance
	projectName := orch.Vars["projectName"]
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projectName + "/logical-clouds/" + lcName
	reply, err := orch.apiGet(url, projectName)
	if err != nil {
		log.Errorf("%s(): Failed to GET LC %s", PrintFunctionName(), lcName)
		return lc, err
	}
	log.Infof("%s(): Get LC status: %d", PrintFunctionName(), reply.StatusCode)
	if err := json.Unmarshal(reply.Data, &lc); err != nil {
		log.Error(err, PrintFunctionName())
	}
	return lc, err
}

func (h *logicalCloudHandler) getLogicalCloudsStatus(ProjectName string, LogicalCloudName string) (lcStatus LogicalCloudStatus, err error) {
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + ProjectName + "/logical-clouds/" + LogicalCloudName + "/status"
	reply, err := orch.apiGet(url, LogicalCloudName)
	if err != nil {
		err = fmt.Errorf("%s(): Failed to get LC status for project %s and logical cloud %s with error %s", PrintFunctionName(),
			ProjectName, LogicalCloudName, err.Error())
		return lcStatus, err
	}
	err = json.Unmarshal(reply.Data, &lcStatus)
	return lcStatus, err
}

func (h *logicalCloudHandler) fetchLCReferencesFlat(lcName string) (lcRefList []clusterReferenceFlat, err error) {
	lcRefList = []clusterReferenceFlat{}
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		orch.Vars["projectName"] + "/logical-clouds/" + lcName + "/cluster-references"
	reply, err := orch.apiGet(url, lcName)
	if err != nil {
		log.Errorf("%s(): Failed to LC reference for %s", PrintFunctionName(), lcName)
		return lcRefList, err
	}
	err = json.Unmarshal(reply.Data, &lcRefList)
	log.Debugf("lc references: %+v", lcRefList)
	return lcRefList, err
}

// w.WriteHeader(resp.(int))
// errMsg := string(h.response.payload[h.response.lastKey]) + h.response.lastKey
// w.Write([]byte(errMsg))
// return

func (h *logicalCloudHandler) getLogicalCloudReferences(lcName string) (nestedRef clusterReferenceNested, err error) {
	orch := h.orchInstance
	lcRefList, err := h.fetchLCReferencesFlat(lcName)
	if err != nil {
		return nestedRef, err
	}
	// Create reference name array
	nestedRef.Metadata.ClusterRefenceNames = make([]string, len(lcRefList))
	for k, cluRef := range lcRefList {
		nestedRef.Metadata.ClusterRefenceNames[k] = cluRef.Metadata.Name
	}

	// Fetch label information of all clusters belonging to cluster provider part of logical cloud
	clusterProviders := make(map[string]bool)
	for _, cluRef := range lcRefList {
		clusterProviders[cluRef.Spec.ClusterProvider] = true
	}

	// Build a map of cluster providers to clusters list
	clusterProviderMap := make(map[string][]Clusters, len(lcRefList))

	var wg sync.WaitGroup
	ERR := &globalErr{}
	for clusterProvider := range clusterProviders {
		clusterProvider := clusterProvider
		wg.Add(1)
		go func(clusterProvider string) {
			defer wg.Done()
			url := "http://" + orch.MiddleendConf.Clm + "/v2/cluster-providers/" +
				clusterProvider + "/clusters?withLabels=true"
			reply, err := orch.apiGet(url, clusterProvider)
			if err != nil {
				log.Errorf("%s(): Encountered error while fetching labels for cluster provider %s",
					PrintFunctionName(), clusterProvider)
				ERR.Error(err)
				return
			}
			var clusterLabels []ClusterLabels
			if err := json.Unmarshal(reply.Data, &clusterLabels); err != nil {
				ERR.Error(err)
				return
			}

			for _, ref := range lcRefList {
				var cluster Clusters
				cluster.Metadata.Name = ref.Spec.ClusterName
				cluster.Metadata.Description = "Cluster" + ref.Spec.ClusterName
				for _, cinfo := range clusterLabels {
					if ref.Spec.ClusterProvider == clusterProvider && ref.Spec.ClusterName == cinfo.Metadata.Name {
						cluster.Spec.Labels = cinfo.Labels
					}
				}
				if clusterProvider == ref.Spec.ClusterProvider {
					clusterProviderMap[clusterProvider] = append(clusterProviderMap[clusterProvider],
						cluster)
				}
			}
		}(clusterProvider)
	}
	wg.Wait()
	// parse through the output and fill int he reference nested structure
	// that is to be returned to the GUI
	nestedRef.Metadata.Name = lcName
	nestedRef.Metadata.Description = "Cluster references for" + lcName

	for k, v := range clusterProviderMap {
		l := ClusterProviders{}
		l.Metadata.Name = k
		l.Metadata.Description = "cluster provider : " + k
		l.Spec.ClustersList = make([]Clusters, len(v))
		l.Spec.ClustersList = v
		nestedRef.Spec.ClusterProvidersList = append(nestedRef.Spec.ClusterProvidersList, l)
	}
	return nestedRef, ERR.Errors()
}

func (h *logicalCloudHandler) createLogicalCloud(lcData logicalCloudsPayload,
	lcDataRetPayload *LogicalClouds,
) int {
	orch := h.orchInstance
	if lcData.CloudType == "admin" {
		resp, err := h.createAdminLogicalCloud(lcData)
		if err != nil || resp != http.StatusCreated {
			log.Errorf("Error encountered during creation of Admin Logical Cloud: %s", err)
			return resp
		}
		// Prepare ret payload
		if err := json.Unmarshal(h.orchInstance.response.payload[lcData.Name], lcDataRetPayload); err != nil {
			log.Error(err, PrintFunctionName())
			return resp
		}
	} else if lcData.CloudType == "user" || lcData.CloudType == "privileged" {
		resp, err := h.createStandardLogicalCloud(lcData, lcDataRetPayload)
		if err != nil || resp != http.StatusCreated {
			log.Errorf("Error encountered during creation of User Logical Cloud: %s", err)
			return resp
		}
	} else {
		log.Errorf("%s(): Invalid cloud type for creation of logical cloud: %s", PrintFunctionName(), lcData.CloudType)
		return http.StatusInternalServerError
	}

	// Now Create the reference for each cluster in the logical cloud
	cretVal := clusterReferenceNested{}
	for _, clusterProvider := range lcData.Spec.ClusterProvidersList {
		cpp := ClusterProviders{}
		cretVal.Metadata.Name = lcData.Name
		cretVal.Metadata.Description = lcData.Description
		for _, cluster := range clusterProvider.Spec.ClustersList {
			resp := h.createClusterReference(orch.Vars["projectName"], lcData.Name, clusterProvider.Metadata.Name, cluster.Metadata.Name)
			if resp != http.StatusCreated {
				log.Errorf("%s(): Failed to add Cluster referecens for cloud %s", PrintFunctionName(), lcData.Name)
				lcDataRetPayload.Spec.ClusterReferences = cretVal
				return resp
			}
			// Prep ret payload
			cpayload := clusterReferenceFlat{}
			clp := Clusters{}
			if err := json.Unmarshal(h.orchInstance.response.payload[lcData.Name+"-"+cluster.Metadata.Name], &cpayload); err != nil {
				log.Error(err, PrintFunctionName())
			}
			cpp.Metadata.Name = cpayload.Spec.ClusterProvider
			clp.Metadata.Name = cpayload.Spec.ClusterName
			cpp.Spec.ClustersList = append(cpp.Spec.ClustersList, clp)
			cretVal.Metadata.ClusterRefenceNames = append(cretVal.Metadata.ClusterRefenceNames, cpayload.Metadata.Name)
		}
		cretVal.Spec.ClusterProvidersList = append(cretVal.Spec.ClusterProvidersList, cpp)
		lcDataRetPayload.Spec.ClusterReferences = cretVal
	}

	// Instantiate the cluster.
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		orch.Vars["projectName"] + "/logical-clouds/" + lcData.Name + "/instantiate"
	var jsonLoad []byte
	resp, err := orch.apiPost(jsonLoad, url, lcData.Name+"-instantiate")
	if err != nil || resp != http.StatusAccepted {
		log.Errorf("%s(): Failed to instantiate logical cloud %s", PrintFunctionName(), lcData.Name)
		return resp.(int)
	}
	lcDataRetPayload.Spec.Status = "Instantiated"
	return http.StatusCreated
}

func (h *logicalCloudHandler) createAdminLogicalCloud(lcData logicalCloudsPayload) (int, error) {
	orch := h.orchInstance
	vars := orch.Vars
	// Create the logical cloud
	apiPayload := logicalCloudData{
		Metadata: apiMetaData{
			Name:        lcData.Name,
			Description: lcData.Description,
			UserData1:   "data 1",
			UserData2:   "data 2",
		},
		Spec: logicalCloudSpec{
			Level: "0",
		},
	}
	jsonLoad, _ := json.Marshal(apiPayload)
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + vars["projectName"] + "/logical-clouds"
	resp, err := orch.apiPost(jsonLoad, url, lcData.Name)
	return resp.(int), err
}

func (h *logicalCloudHandler) addPrivilegedPermisions(lcData logicalCloudsPayload, projectName string, lcDataRetPayload *LogicalClouds) (int, error) {
	// Create User Permissions for Privileged Logical Cloud
	lcData.Spec.Permissions = &userPermissions{
		APIGroups: []string{"*"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	}
	retCode, err := h.createUserPermissions(projectName, lcData.Name, lcData.Spec.Namespace, lcData.Spec.Permissions, lcDataRetPayload)
	if retCode != http.StatusCreated {
		log.Errorf("Creating user permissions failed for logical cloud: %s", lcData.Name)
		return retCode, err
	}

	retCode, err = h.createUserPermissions("kube", lcData.Name, "kube-system", lcData.Spec.Permissions, lcDataRetPayload)
	if retCode != http.StatusCreated {
		log.Errorf("Kube-system NS : Creating user permissions failed for logical cloud: %s", lcData.Name)
		return retCode, err
	}
	// Create Cluster Wide User Permissions for Privileged Logical Cloud
	retCode, err = h.createUserPermissions("cluster", lcData.Name, "", lcData.Spec.Permissions, lcDataRetPayload)
	if retCode != http.StatusCreated {
		log.Errorf("Cluster-Wide: Creating user permissions failed for logical cloud: %s", lcData.Name)
		return retCode, err
	}
	return http.StatusCreated, nil
}

func (h *logicalCloudHandler) createStandardLogicalCloud(lcData logicalCloudsPayload, lcDataRetPayload *LogicalClouds) (int, error) {
	orch := h.orchInstance
	vars := orch.Vars

	// Create the logical cloud
	apiPayload := logicalCloudData{
		Metadata: apiMetaData{
			Name:        lcData.Name,
			Description: lcData.Description,
			UserData1:   "data 1",
			UserData2:   "data 2",
		},
		Spec: logicalCloudSpec{
			UserData: UserData{
				UserName: vars["projectName"],
				Type:     "certificate",
			},
			NameSpace: lcData.Spec.Namespace,
		},
	}

	if lcData.CloudType == "user" {
		lcData.CloudType = "standard"
	}
	jsonLoad, _ := json.Marshal(apiPayload)
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + vars["projectName"] + "/logical-clouds"
	resp, err := orch.apiPost(jsonLoad, url, lcData.Name)
	if err != nil || resp != http.StatusCreated {
		log.Errorf("%s(): Failed to crreate cloud type %s name %s", PrintFunctionName(), lcData.CloudType, lcData.Name)
		return resp.(int), err
	}
	log.Infof("%s(): Created %s logical-cloud %s retcode %d  ", PrintFunctionName(), lcData.CloudType, lcData.Name, resp.(int))
	if err := json.Unmarshal(h.orchInstance.response.payload[lcData.Name], lcDataRetPayload); err != nil {
		log.Error(err, PrintFunctionName())
		return resp.(int), err
	}

	if lcData.CloudType == "standard" {
		// Default user permissions
		defaultUserPerm := &userPermissions{
			APIGroups: []string{
				"", "apps", "k8splugin.io", "networking.k8s.io",
				"admissionregistration.k8s.io", "apiextensions.k8s.io", "apiregistration.k8s.io",
				"authentication.k8s.io", "authorization.k8s.io", "autoscaling", "batch",
				"certificates.k8s.io", "coordination.k8s.io", "discovery.k8s.io", "events.k8s.io", "flowcontrol.apiserver.k8s.io",
				"internal.apiserver.k8s.io", "node.k8s.io", "policy", "rbac.authorization.k8s.io", "scheduling.k8s.io",
				"storage.k8s.io", "networking.istio.io", "authentication.istio.io", "rbac.istio.io", "config.istio.io", "security.istio.io",
			},
			Resources: []string{"*", "destinationrules", "envoyfilters", "serviceentries", "sidecars", "gateways", "virtualservices"},
			Verbs: []string{
				"get",
				"watch",
				"list",
				"create",
				"update",
				"patch",
				"delete",
			},
		}

		// Check if userPermissions are available as part of payload, if not use default
		if lcData.Spec.Permissions == nil {
			lcData.Spec.Permissions = defaultUserPerm
		}

		// Create User Permissions for Standard Logical Cloud
		retCode, err := h.createUserPermissions(vars["projectName"], lcData.Name, lcData.Spec.Namespace,
			lcData.Spec.Permissions, lcDataRetPayload)
		if err != nil || retCode != http.StatusCreated {
			log.Errorf("Creating user permissions failed for logical cloud: %s: status code %d", lcData.Name, retCode)
			return retCode, err
		}

		// Default quota
		defaultUserQuota := &QuotaInfo{
			LimitsCPU:                   "400",
			LimitsMemory:                "1000Gi",
			RequestsCPU:                 "300",
			RequestsMemory:              "900Gi",
			RequestsStorage:             "500Gi",
			PersistentVolumeClaims:      "500",
			Pods:                        "500",
			ConfigMaps:                  "1000",
			ReplicationControllers:      "500",
			ResourceQuotas:              "500",
			Services:                    "500",
			ServicesLoadBalancers:       "500",
			ServicesNodePorts:           "500",
			Secrets:                     "500",
			CountReplicationControllers: "500",
			CountDeploymentsApps:        "500",
			CountReplicasetsApps:        "500",
			CountStatefulSets:           "500",
			CountJobsBatch:              "500",
			CountCronJobsBatch:          "500",
			CountDeploymentsExtensions:  "500",
		}

		if lcData.Spec.Quotas == nil {
			lcData.Spec.Quotas = defaultUserQuota
		}

		// Create User Quotas for Standard Logical Cloud
		retCode, err = h.createUserQuota(vars["projectName"], lcData.Name, lcData.Spec.Quotas, lcDataRetPayload)
		if err != nil || retCode != http.StatusCreated {
			log.Errorf("Updating user quota failed for logical cloud: %s", lcData.Name)
			return retCode, err
		}

	} else {
		retCode, err := h.addPrivilegedPermisions(lcData, vars["projectName"], lcDataRetPayload)
		if err != nil || retCode != http.StatusCreated {
			log.Errorf("Updating user quota failed for logical cloud: %s", lcData.Name)
			return retCode, err
		}
	}
	return http.StatusCreated, nil
}

func (h *OrchestrationHandler) CreateAmcopSystemLogicalCloud(w http.ResponseWriter, ClusterProviderName string, jsonData ClusterMetadata) bool {
	// Final Result ByDefault Considered True
	Result := true
	// Variable for Logical Cloud
	var provider ClusterProviders
	var cluster Clusters
	h.InitializeResponseMap()

	// for provider Metadata
	provider.Metadata.Name = ClusterProviderName
	provider.Metadata.Description = ""

	// for cluster Metadata
	cluster.Metadata.Name = jsonData.Metadata.Name
	cluster.Metadata.Description = ""

	// Array for ClusterProvider and Cluster List
	provider.Spec.ClustersList = append(provider.Spec.ClustersList, cluster)

	// Initialing the Logical Cloud Struct with Payload
	lcData := logicalCloudsPayload{
		Name:        "operator-logical-cloud-" + jsonData.Metadata.Name,
		Description: "operator-logical-cloud",
		Spec: LogicalCloudSpec{
			ClusterProvidersList: []ClusterProviders{provider},
		},
	}

	h.client = http.Client{}

	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h
	// Creating the Logical Cloud for Monitoring Service
	lcData.CloudType = "admin"
	lcDataRetPayload := LogicalClouds{}
	lcStatus := lcHandler.createLogicalCloud(lcData, &lcDataRetPayload)
	if lcStatus != http.StatusCreated {
		log.Errorf("%s(): Failed to create logical cloud %s", PrintFunctionName(), lcData.Name)
		lcRollback(lcHandler, &lcDataRetPayload)
		Result = false
	}
	return Result
}

func lcRollback(lcHandler *logicalCloudHandler, lcData *LogicalClouds) {
	vars := lcHandler.orchInstance.Vars
	projectName := vars["projectName"]
	lcName := lcData.Metadata.Name
	if lcData.Spec.Level != "0" {
		for _, p := range lcData.Spec.UserPerminssionMetadata {
			retval, _ := lcHandler.deleteUserPermissions(projectName, lcName, p.UserPermissionName)
			if retval != http.StatusNoContent {
				log.Debugf("%s(): Failed to delete user permissions for lc %s", PrintFunctionName(), lcName)
			}
		}
		retval, _ := lcHandler.deleteUserQuota(projectName, lcName, lcData.Spec.UserQuotaMetadata.QuotaName)
		if retval != http.StatusNoContent {
			log.Debugf("%s(): Failed to delete quota info for lc %s", PrintFunctionName(), lcName)
		}
	}
	for _, n := range lcData.Spec.ClusterReferences.Metadata.ClusterRefenceNames {
		retval, _ := lcHandler.deleteClusterReference(projectName, lcName, n)
		if retval != http.StatusNoContent {
			log.Debugf("%s(): Failed to delete lc reference for %s", PrintFunctionName(), lcName)
		}
	}
	retval, _ := lcHandler.deleteLogicalCloud(projectName, lcName)
	if retval != http.StatusNoContent {
		log.Debugf("%s(): Failed to delete lc %s", PrintFunctionName(), lcName)
	}
	log.Infof("%s(): Deleted Logical cloud %s", PrintFunctionName(), lcName)
}

// HandleLCCreateRequest CreateLogicalCloud, creates the logical clouds (level 0/level 1)
func (h *OrchestrationHandler) HandleLCCreateRequest(w http.ResponseWriter, r *http.Request) {
	var lcData logicalCloudsPayload
	h.Vars = mux.Vars(r)
	h.InitializeResponseMap()
	lcDataRetPayload := LogicalClouds{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lcData)
	if err != nil {
		log.Errorf("%s(): failed to parse json: %s", PrintFunctionName(), err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h
	lcStatus := lcHandler.createLogicalCloud(lcData, &lcDataRetPayload)
	if lcStatus != http.StatusCreated {
		log.Errorf("%s(): Failed to create logical cloud %s", PrintFunctionName(), lcData.Name)
		lcRollback(lcHandler, &lcDataRetPayload)
	}
	log.Infof("---------- %s", lcDataRetPayload)
	w.WriteHeader(lcStatus)
	retVal, _ := json.Marshal(lcDataRetPayload)
	if _, err := w.Write(retVal); err != nil {
		log.Error(err, PrintFunctionName())
	}
}

func (h *OrchestrationHandler) GetLogicalCloudsStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.InitializeResponseMap()
	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h

	// Get the logical cloud list
	lcStatus, err := lcHandler.getLogicalCloudsStatus(vars["projectName"], vars["logicalCloud"])
	if err != nil {
		log.Infof("Failed to get logical cloud %s status, error %s", vars["logicalCloud"], err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			log.Error(err, PrintFunctionName())
		}
		return
	}
	js, _ := json.Marshal(lcStatus)
	if _, err := w.Write(js); err != nil {
		log.Error(err, PrintFunctionName())
	}
}

func getCloudProperties(lcHandler *logicalCloudHandler, lcList []LogicalClouds) error {
	log.Infof("%s(): lcList: %+v", PrintFunctionName(), lcList)
	var wg sync.WaitGroup
	ERR := &globalErr{}
	for k := range lcList {
		wg.Add(1)
		k := k
		go func(k int) {
			defer wg.Done()
			nestedRef, err := lcHandler.getLogicalCloudReferences(lcList[k].Metadata.Name)
			if err != nil {
				log.Errorf("%s(): Failed to get lcReferences: for LC %s: %+v", PrintFunctionName(),
					lcList[k].Metadata.Name, nestedRef)
				ERR.Error(err)
				return
			}
			log.Infof("%s(): lcReferences: for LC %s: %+v", PrintFunctionName(), lcList[k].Metadata.Name, nestedRef)

			lcList[k].Spec.ClusterReferences = nestedRef
			if lcList[k].Spec.Level != "0" {
				// Fetch logical cloud permissions, if it is standard/privileged logical cloud
				usrPm, err := lcHandler.GetUserPermissions(lcList[k].Metadata.Name)
				if err != nil {
					log.Errorf("%s(): Unable to fetch user permissions for L1 logical cloud: %s",
						PrintFunctionName(), lcList[k].Metadata.Name)
					ERR.Error(err)
					return
				}
				if len(usrPm) != 0 {
					for _, p := range usrPm {
						lcList[k].Spec.UserPermissionSpec = append(lcList[k].Spec.UserPermissionSpec, p.Specification)
						lcList[k].Spec.UserPerminssionMetadata = append(lcList[k].Spec.UserPerminssionMetadata, p.MetaData)
					}
				}
				// Fetch logical cloud quota info, if it is standard/privileged logical cloud
				quota, err := lcHandler.GetClusterQuotas(lcList[k].Metadata.Name)
				if err != nil {
					log.Errorf("%s(): Unable to fetch Quota for L1 logical cloud: %s",
						PrintFunctionName(), lcList[k].Metadata.Name)
					ERR.Error(err)
					return
				}
				if len(quota) != 0 { // There only one quota FIXME
					for _, q := range quota {
						lcList[k].Spec.UserQuota = q.Specification
						lcList[k].Spec.UserQuotaMetadata = q.MetaData
					}
				}
			}
			// Get the logical cloud list
			lcStatus, err := lcHandler.getLogicalCloudsStatus(lcHandler.orchInstance.Vars["projectName"],
				lcList[k].Metadata.Name)
			if err != nil {
				log.Infof("Failed to get logical cloud %s  error: %s", lcList[k].Metadata.Name,
					err)
				ERR.Error(err)
				return
			}
			lcList[k].Spec.Status = lcStatus.DeployedStatus
		}(k)

	}
	wg.Wait()
	return ERR.Errors()
}

// GetLogicalClouds Get LC information
func (h *OrchestrationHandler) GetLogicalClouds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h.Vars = vars
	h.InitializeResponseMap()
	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h
	// Get the logical cloud list
	lcList, err := lcHandler.getLogicalClouds()
	if err != nil {
		log.Infof("Failed to get logical clouds : %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = getCloudProperties(lcHandler, lcList)
	if err != nil {
		log.Infof("%s(): Failed to get logical clouds properties : %s", PrintFunctionName(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf("%s(): LC list after filling the permissions and quotas : %+v", PrintFunctionName(), lcList)

	retval, err := json.Marshal(lcList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf("%s(): retval of GetLogicalCloud date: %s", PrintFunctionName(), retval)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(retval); err != nil {
		log.Error(err, PrintFunctionName())
	}
}

// GetUserPermissions Fetch User Permissions for L1 Logical Cloud
func (h *logicalCloudHandler) GetUserPermissions(lcName string) ([]UserPermission, error) {
	userPermList := []UserPermission{}
	var url string
	orch := h.orchInstance
	url = "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		orch.Vars["projectName"] + "/logical-clouds/" + lcName + "/user-permissions"
	reply, err := orch.apiGet(url, lcName+"-permissions")
	if err != nil {
		log.Errorf("%s(): Failed to get userpermission for LC %s", PrintFunctionName(), lcName)
		return userPermList, err
	}
	err = json.Unmarshal(reply.Data, &userPermList)
	log.Debugf("%s(): lc user permission: %+v", PrintFunctionName(), userPermList)
	return userPermList, err
}

// GetClusterQuotas Fetch Cluster Quotas for L1 logical cloud
func (h *logicalCloudHandler) GetClusterQuotas(lcName string) ([]Quota, error) {
	quotas := []Quota{}
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		orch.Vars["projectName"] + "/logical-clouds/" + lcName + "/cluster-quotas"
	reply, err := orch.apiGet(url, lcName+"-quotas")
	if err != nil {
		log.Errorf("%s(): Failed to get quota info for LC %s", PrintFunctionName(), lcName)
		return quotas, err
	}
	err = json.Unmarshal(reply.Data, &quotas)
	log.Debugf("%s(): LC quotas: %+v", PrintFunctionName(), quotas)
	return quotas, err
}

// DeleteLogicalCloud deletes the logical clouds (level 0/level 1)
func (h *OrchestrationHandler) DeleteLogicalCloud(w http.ResponseWriter, r *http.Request) {
	h.Vars = mux.Vars(r)
	h.InitializeResponseMap()
	// There will be just one element in the list
	lcList := []LogicalClouds{}
	lcName := h.Vars["logicalCloud"]
	projectName := h.Vars["projectName"]
	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h

	// Fetch logical cloud information
	lc, err := lcHandler.getLogicalCloud(lcName)
	if err != nil {
		log.Infof("Failed to get logical clouds : %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lcList = append(lcList, lc)
	err = getCloudProperties(lcHandler, lcList)
	if err != nil {
		log.Errorf("%s(): Failed to get logical clouds properties : %s", PrintFunctionName(), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugf("%s(): LC list after filling the permissions and quotas : %+v", PrintFunctionName(), lcList)

	if lcList[0].Spec.Level != "0" {
		for _, p := range lcList[0].Spec.UserPerminssionMetadata {
			retval, err := lcHandler.deleteUserPermissions(projectName, lcName, p.UserPermissionName)
			if retval != http.StatusNoContent {
				log.Errorf("%s(): Failed to delete user permissions for lc %s", PrintFunctionName(), lcName)
				w.WriteHeader(retval)
				if err != nil {
					if _, err := w.Write([]byte(err.Error())); err != nil {
						log.Error(err, PrintFunctionName())
					}
				}
				return
			}
		}
		if len(lcList[0].Spec.UserQuota) != 0 {
			retval, err := lcHandler.deleteUserQuota(projectName, lcName, lcList[0].Spec.UserQuotaMetadata.QuotaName)
			if retval != http.StatusNoContent {
				log.Errorf("%s(): Failed to delete quota info for lc %s", PrintFunctionName(), lcName)
				w.WriteHeader(retval)
				if err != nil {
					if _, err := w.Write([]byte(err.Error())); err != nil {
						log.Error(err, PrintFunctionName())
					}
				}
				return
			}
		}
	}
	retval, _ := lcHandler.terminateLogicalCloud(projectName, lcName)
	if retval != http.StatusAccepted {
		log.Errorf("%s(): Failed to tetminate lc %s", PrintFunctionName(), lcName)
		//w.WriteHeader(retval)
		//if err != nil {
		//	w.Write([]byte(err.Error()))
		//}
		//return do not return and go ahead with ref deletion
	}
	// Terminate logical cloud
	for _, n := range lcList[0].Spec.ClusterReferences.Metadata.ClusterRefenceNames {
		count := 0
		retval := http.StatusConflict
		for retval != http.StatusNoContent {
			retval, err = lcHandler.deleteClusterReference(projectName, lcName, n)
			log.Infof("Count %d", count)
			count += 1
			time.Sleep(time.Second)
			if count > 20 {
				log.Errorf("%s(): Failed to delete lc reference for %s", PrintFunctionName(), lcName)
				w.WriteHeader(retval)
				if err != nil {
					if _, err := w.Write([]byte(err.Error())); err != nil {
						log.Error(err, PrintFunctionName())
					}
				}
				return
			}
		}
	}
	retval, err = lcHandler.deleteLogicalCloud(projectName, lcName)
	if retval != http.StatusNoContent {
		log.Errorf("%s(): Failed to delete lc %s", PrintFunctionName(), lcName)
		w.WriteHeader(retval)
		if err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Error(err, PrintFunctionName())
			}
		}
		return
	}

	log.Infof("%s(): Deleted Logical cloud %s", PrintFunctionName(), lcName)
	w.WriteHeader(http.StatusNoContent)
}

// UpdateLogicalCloud updates the logical clouds (level 0/level 1)
func (h *OrchestrationHandler) UpdateLogicalCloud(w http.ResponseWriter, r *http.Request) {
	var lcData logicalCloudUpdatePayload
	lcDataRetPayload := LogicalClouds{}
	h.Vars = mux.Vars(r)
	h.InitializeResponseMap()
	lcHandler := &logicalCloudHandler{}
	lcHandler.orchInstance = h
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lcData)
	if err != nil {
		log.Errorf("failed to parse update json: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if lcData.CloudType == "standard" {
		// Delete user permissions for standard logical cloud
		retCode, err := lcHandler.deleteUserPermissions(h.Vars["projectName"], h.Vars["logicalCloud"], "")
		if retCode != http.StatusNoContent {
			log.Errorf("Deleting user permissions failed for logical cloud: %s", h.Vars["logicalCloud"])
			w.WriteHeader(retCode)
			if err != nil {
				if _, err := w.Write([]byte(err.Error())); err != nil {
					log.Error(err, PrintFunctionName())
				}
			}
			return
		}

		// Create user permissions for standard logical cloud based on updated user permissions
		retCode, err = lcHandler.createUserPermissions(h.Vars["projectName"], h.Vars["logicalCloud"], lcData.Namespace,
			lcData.Permissions, &lcDataRetPayload)
		if retCode != http.StatusCreated {
			log.Errorf("Updating user permissions failed for logical cloud: %s", h.Vars["logicalCloud"])
			w.WriteHeader(retCode)
			if err != nil {
				if _, err := w.Write([]byte(err.Error())); err != nil {
					log.Error(err, PrintFunctionName())
				}
			}
			return
		}

		// Delete user quotas for standard logical cloud
		retCode, err = lcHandler.deleteUserQuota(h.Vars["projectName"], h.Vars["logicalCloud"], "name")
		if retCode != http.StatusNoContent {
			log.Errorf("Deleting user quota failed for logical cloud: %s", h.Vars["logicalCloud"])
			w.WriteHeader(retCode)
			if err != nil {
				if _, err := w.Write([]byte(err.Error())); err != nil {
					log.Error(err, PrintFunctionName())
				}
			}
			return
		}

		// Create user quotas for standard logical cloud based on updated user quotas
		retCode, err = lcHandler.createUserQuota(h.Vars["projectName"], h.Vars["logicalCloud"], lcData.Quotas, &lcDataRetPayload)
		if retCode != http.StatusCreated {
			log.Errorf("Updating user quota failed for logical cloud: %s", h.Vars["logicalCloud"])
			w.WriteHeader(retCode)
			if err != nil {
				if _, err := w.Write([]byte(err.Error())); err != nil {
					log.Error(err, PrintFunctionName())
				}
			}
			return
		}
	}

	// Process cluster references for logical cloud update
	for _, clusterProviderInfo := range lcData.ClusterProvidersList {
		clusterProvider := clusterProviderInfo.Metadata.Name
		for _, clusterInfo := range clusterProviderInfo.Spec.ClustersList {
			if clusterInfo.Metadata.Operation == "add" {
				retCode := lcHandler.createClusterReference(h.Vars["projectName"], h.Vars["logicalCloud"],
					clusterProvider, clusterInfo.Metadata.Name)
				if retCode != http.StatusCreated {
					w.WriteHeader(retCode)
					return
				}
			}

			if clusterInfo.Metadata.Operation == "delete" {
				clusterReference := h.Vars["logicalCloud"] + "-" +
					clusterProvider + "-" + clusterInfo.Metadata.Name
				retCode, err := lcHandler.deleteClusterReference(h.Vars["projectName"], h.Vars["logicalCloud"], clusterReference)
				if retCode != http.StatusNoContent {
					w.WriteHeader(retCode)
					if err != nil {
						if _, err := w.Write([]byte(err.Error())); err != nil {
							log.Error(err, PrintFunctionName())
						}
					}
					return
				}
			}
		}
	}

	// Invoke logical cloud update to apply updated configuration
	var jsonLoad []byte
	url := "http://" + h.MiddleendConf.Dcm + "/v2/projects/" +
		h.Vars["projectName"] + "/logical-clouds/" + h.Vars["logicalCloud"] + "/update"
	resp, err := h.apiPost(jsonLoad, url, h.Vars["logicalCloud"]+"_update")
	if err != nil {
		log.Errorf("Encountered error while updating logical cloud: %s", h.Vars["logicalCloud"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(resp.(int))
}

func (h *logicalCloudHandler) createClusterReference(projectName string, lcName string, clusterProvider string,
	clusterName string,
) int {
	orch := h.orchInstance
	clusterReferencePayload := clusterReferenceFlat{}
	clusterReferencePayload.Metadata.Name = lcName + "-" +
		clusterProvider + "-" + clusterName
	clusterReferencePayload.Metadata.Description = "Cluster reference for cluster" +
		clusterProvider + ":" + clusterName
	clusterReferencePayload.Metadata.Userdata1 = "NA"
	clusterReferencePayload.Metadata.Userdata2 = "NA"
	clusterReferencePayload.Spec.ClusterProvider = clusterProvider
	clusterReferencePayload.Spec.ClusterName = clusterName
	clusterReferencePayload.Spec.LoadbalancerIP = "0.0.0.0"
	jsonLoad, _ := json.Marshal(clusterReferencePayload)
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projectName + "/logical-clouds/" + lcName + "/cluster-references"
	resp, err := orch.apiPost(jsonLoad, url, lcName+"-"+clusterName)
	if err != nil {
		log.Errorf("Encountered error while creating cluster reference: %s", err)
	}
	return resp.(int)
}

func (h *logicalCloudHandler) deleteClusterReference(projName string, lcName string, clusterReference string) (int, error) {
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projName + "/logical-clouds/" + lcName + "/cluster-references/" + clusterReference
	resp, err := orch.apiDel(url, lcName+"_lcrefdel")
	return resp.(int), err
}

func (h *logicalCloudHandler) createUserPermissions(projName string, lcName string, namespace string, perm *userPermissions,
	lcDataRetPayload *LogicalClouds,
) (int, error) {
	orch := h.orchInstance
	usrPermRetVal := UserPermission{}
	userPerm := UserPermission{
		MetaData: UPMetaData{
			UserPermissionName: projName + "_permissions",
			Description:        "User Permissions",
			UserData1:          "UserData1",
			UserData2:          "UserData2",
		},
		Specification: UPSpec{
			Namespace: namespace,
			APIGroups: perm.APIGroups,
			Resources: perm.Resources,
			Verbs:     perm.Verbs,
		},
	}
	jsonLoad, _ := json.Marshal(userPerm)
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + orch.Vars["projectName"] + "/logical-clouds/" +
		lcName + "/user-permissions"
	resp, err := orch.apiPost(jsonLoad, url, lcName+"_usrperm")
	_ = json.Unmarshal(h.orchInstance.response.payload[lcName+"_usrperm"], &usrPermRetVal)
	lcDataRetPayload.Spec.UserPermissionSpec = append(lcDataRetPayload.Spec.UserPermissionSpec,
		usrPermRetVal.Specification)
	lcDataRetPayload.Spec.UserPerminssionMetadata = append(lcDataRetPayload.Spec.UserPerminssionMetadata,
		usrPermRetVal.MetaData)
	return resp.(int), err
}

func (h *logicalCloudHandler) deleteUserPermissions(projName string, lcName string, permName string) (int, error) {
	orch := h.orchInstance
	if permName == "" {
		permName = projName
	}
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projName + "/logical-clouds/" + lcName + "/user-permissions/" + permName
	resp, err := orch.apiDel(url, lcName+"_permdel")
	return resp.(int), err
}

func (h *logicalCloudHandler) createUserQuota(projName string, lcName string, quotas *QuotaInfo, lcDataRetPayload *LogicalClouds) (int, error) {
	quotaRetVal := Quota{}
	orch := h.orchInstance
	quotaInfo := make(map[string]string)
	quotaInfo["limits.cpu"] = quotas.LimitsCPU
	quotaInfo["limits.memory"] = quotas.LimitsMemory
	quotaInfo["requests.cpu"] = quotas.RequestsCPU
	quotaInfo["requests.memory"] = quotas.RequestsMemory
	quotaInfo["requests.storage"] = quotas.RequestsStorage
	/*quotaInfo["limits.ephemeral-storage"] = lcData.Spec.Quotas.LimitsEphemeralStorage*/
	quotaInfo["persistentvolumeclaims"] = quotas.PersistentVolumeClaims
	quotaInfo["pods"] = quotas.Pods
	quotaInfo["configmaps"] = quotas.ConfigMaps

	quotaInfo["replicationcontrollers"] = quotas.ReplicationControllers
	quotaInfo["resourcequotas"] = quotas.ResourceQuotas
	quotaInfo["services"] = quotas.Services
	quotaInfo["services.loadbalancers"] = quotas.ServicesLoadBalancers
	quotaInfo["services.nodeports"] = quotas.ServicesNodePorts
	quotaInfo["secrets"] = quotas.Secrets
	quotaInfo["count/replicationcontrollers"] = quotas.CountReplicationControllers
	quotaInfo["count/deployments.apps"] = quotas.CountDeploymentsApps
	quotaInfo["count/replicasets.apps"] = quotas.CountReplicasetsApps
	quotaInfo["count/statefulsets.apps"] = quotas.CountStatefulSets
	quotaInfo["count/jobs.batch"] = quotas.CountJobsBatch
	quotaInfo["count/cronjobs.batch"] = quotas.CountCronJobsBatch
	quotaInfo["count/cronjobs.batch"] = quotas.CountDeploymentsExtensions
	quota := Quota{
		MetaData: QMetaData{
			QuotaName:   projName + "-quotas",
			Description: "User Quotas",
			UserData1:   "UserData1",
			UserData2:   "UserData2",
		},
		Specification: quotaInfo,
	}

	jsonLoad, _ := json.Marshal(quota)
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projName + "/logical-clouds/" + lcName + "/cluster-quotas"
	resp, err := orch.apiPost(jsonLoad, url, lcName+"_quota")
	_ = json.Unmarshal(h.orchInstance.response.payload[lcName+"_quota"], &quotaRetVal)
	lcDataRetPayload.Spec.UserQuota = quotaRetVal.Specification
	lcDataRetPayload.Spec.UserQuotaMetadata = quotaRetVal.MetaData
	return resp.(int), err
}

func (h *logicalCloudHandler) deleteUserQuota(projName string, lcName string, quotaName string) (int, error) {
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" +
		projName + "/logical-clouds/" + lcName + "/cluster-quotas/" + quotaName
	resp, err := orch.apiDel(url, lcName+"_quotadel")
	return resp.(int), err
}

func (h *logicalCloudHandler) deleteLogicalCloud(projName string, lcName string) (int, error) {
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + projName + "/logical-clouds/" + lcName
	resp, err := orch.apiDel(url, lcName+"_lcdel")
	log.Infof("Delete logical cloud %s : %d", lcName, resp)
	return resp.(int), err
}

func (h *logicalCloudHandler) terminateLogicalCloud(projName string, lcName string) (int, error) {
	orch := h.orchInstance
	url := "http://" + orch.MiddleendConf.Dcm + "/v2/projects/" + projName + "/logical-clouds/" + lcName + "/terminate"
	var jsonLoad []byte
	resp, err := orch.apiPost(jsonLoad, url, lcName+"-terminate")
	return resp.(int), err
}
