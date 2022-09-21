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
	"net/http"
	"strconv"
	"strings"

	"example.com/middleend/localstore"
	log "github.com/sirupsen/logrus"
)

// plamcentIntentHandler implements the orchworkflow interface
type dtcIntentHandler struct {
	orchURL      string
	orchInstance *OrchestrationHandler
}

func (h *localStoreIntentHandler) DeleteClientsInboundIntent(clientIntentName string, p string, ca string, v string,
	digName string, trafficIntentName string, inboundIntentName string,
) (interface{}, interface{}) {
	// Get the local store handler.
	c := localstore.NewClientsInboundIntentClient()
	deleteErr := c.DeleteClientsInboundIntent(clientIntentName, p, ca, v, digName, trafficIntentName, inboundIntentName)
	if deleteErr != nil {
		log.Error(deleteErr.Error(), log.Fields{})
		if strings.Contains(deleteErr.Error(), "not found") {
			return http.StatusNotFound, deleteErr
		} else if strings.Contains(deleteErr.Error(), "conflict") {
			return http.StatusConflict, deleteErr
		} else {
			return http.StatusInternalServerError, deleteErr
		}
	}
	return http.StatusNoContent, deleteErr
}

func (h *remoteStoreIntentHandler) DeleteClientsInboundIntent(clientIntentName string, p string, ca string, v string,
	digName string, trafficIntentName string, inboundIntentName string,
) (interface{}, interface{}) {
	orch := h.orchInstance
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName +
		"/inbound-intents/" + inboundIntentName + "/clients/" + clientIntentName
	status, err := orch.apiDel(url, clientIntentName)
	return status, err
}

func (h *remoteStoreIntentHandler) GetClientsInboundIntents(p string, ca string,
	v string, digName string, trafficIntentName string, inboundIntentName string,
) ([]byte, error) {
	orch := h.orchInstance
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName + "/inbound-intents/" + inboundIntentName + "/clients"
	reply, err := orch.apiGet(url, ca+"_inboundclient")
	return reply.Data, err
}

func (h *localStoreIntentHandler) GetClientsInboundIntents(p string, ca string,
	v string, digName string, trafficIntentName string, inboundIntentName string,
) ([]byte, error) {
	var retval []byte
	c := localstore.NewClientsInboundIntentClient()
	inboundclientIntent, err := c.GetClientsInboundIntents(p, ca, v, digName, trafficIntentName, inboundIntentName)
	if err != nil {
		log.Error(err.Error(), log.Fields{})
		return retval, err
	}
	retval, _ = json.Marshal(inboundclientIntent)
	return retval, err
}

func (h *localStoreIntentHandler) CreateClientsInboundIntent(g localstore.InboundClientsIntent, p string, ca string,
	v string, digName string, trafficIntentName string, serverName string, exist bool,
) (interface{}, interface{}) {
	// Get the local store handler
	c := localstore.NewClientsInboundIntentClient()
	_, createErr := c.CreateClientsInboundIntent(g, p, ca, v, digName, trafficIntentName, serverName, exist)
	if createErr != nil {
		log.Error(createErr.Error(), log.Fields{})
		if strings.Contains(createErr.Error(), "Unable to find the project") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the composite-app") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the deploymentIntentGroupName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the TrafficGroupIntentName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the InboundIntentName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Client Intent already exists") {
			return http.StatusConflict, createErr
		} else {
			return http.StatusInternalServerError, createErr
		}
	}

	return http.StatusCreated, createErr
}

func (h *remoteStoreIntentHandler) CreateClientsInboundIntent(g localstore.InboundClientsIntent, p string, ca string,
	v string, digName string, trafficIntentName string, serverName string, exist bool,
) (interface{}, interface{}) {
	orch := h.orchInstance
	jsonLoad, _ := json.Marshal(g)
	dtcintName := ca + "_inboundclient"
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName + "/inbound-intents/" + serverName + "/clients"
	resp, err := orch.apiPost(jsonLoad, url, dtcintName)

	return resp, err
}

func (h *localStoreIntentHandler) DeleteServerInboundIntent(serverIntentName string, p string, ca string, v string,
	digName string, trafficIntentName string,
) (interface{}, interface{}) {
	// Get the local store handler.
	c := localstore.NewServerInboundIntentClient()
	deleteErr := c.DeleteServerInboundIntent(serverIntentName, p, ca, v, digName, trafficIntentName)
	if deleteErr != nil {
		log.Error(deleteErr.Error(), log.Fields{})
		if strings.Contains(deleteErr.Error(), "not found") {
			return http.StatusNotFound, deleteErr
		} else if strings.Contains(deleteErr.Error(), "conflict") {
			return http.StatusConflict, deleteErr
		} else {
			return http.StatusInternalServerError, deleteErr
		}
	}
	return http.StatusNoContent, deleteErr
}

func (h *remoteStoreIntentHandler) DeleteServerInboundIntent(serverIntentName string, p string, ca string, v string,
	digName string, trafficIntentName string,
) (interface{}, interface{}) {
	orch := h.orchInstance
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName + "/inbound-intents/" + serverIntentName
	status, err := orch.apiDel(url, serverIntentName)
	return status, err
}

func (h *remoteStoreIntentHandler) GetServerInboundIntents(p string, ca string,
	v string, digName string, trafficIntentName string,
) ([]byte, error) {
	orch := h.orchInstance
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName + "/inbound-intents"

	reply, err := orch.apiGet(url, ca+"_inboundserver")
	return reply.Data, err
}

func (h *localStoreIntentHandler) GetServerInboundIntents(p string, ca string,
	v string, digName string, trafficIntentName string,
) ([]byte, error) {
	var retval []byte
	c := localstore.NewServerInboundIntentClient()
	inboundserverIntent, err := c.GetServerInboundIntents(p, ca, v, digName, trafficIntentName)
	if err != nil {
		log.Error(err.Error(), log.Fields{})
		if strings.Contains(err.Error(), "db Find error") {
			return retval, err
		} else {
			return retval, err
		}
	}
	retval, _ = json.Marshal(inboundserverIntent)
	return retval, nil
}

func (h *localStoreIntentHandler) CreateServerInboundIntent(g localstore.InboundServerIntent, p string, ca string,
	v string, digName string, trafficIntentName string, exist bool,
) (interface{}, interface{}) {
	// Get the local store handler
	c := localstore.NewServerInboundIntentClient()
	_, createErr := c.CreateServerInboundIntent(g, p, ca, v, digName, trafficIntentName, exist)
	if createErr != nil {
		log.Error(createErr.Error(), log.Fields{})
		if strings.Contains(createErr.Error(), "Unable to find the project") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the composite-app") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the deploymentIntentGroupName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the TrafficGroupIntentName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Server Intent already exists") {
			return http.StatusConflict, createErr
		} else {
			return http.StatusInternalServerError, createErr
		}
	}

	return http.StatusCreated, createErr
}

func (h *remoteStoreIntentHandler) CreateServerInboundIntent(g localstore.InboundServerIntent, p string, ca string,
	v string, digName string, trafficIntentName string, exist bool,
) (interface{}, interface{}) {
	orch := h.orchInstance
	jsonLoad, _ := json.Marshal(g)
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents/" + trafficIntentName + "/inbound-intents"
	resp, err := orch.apiPost(jsonLoad, url, g.Metadata.Name)

	return resp, err
}

func (h *localStoreIntentHandler) DeleteTrafficGroupIntent(dtintName string, p string, ca string,
	v string, digName string,
) (interface{}, interface{}) {
	c := localstore.NewTrafficGroupIntentClient()

	err := c.DeleteTrafficGroupIntent(dtintName, p, ca, v, digName)
	if err != nil {
		log.Error(err.Error(), log.Fields{})
		if strings.Contains(err.Error(), "not found") {
			return http.StatusNotFound, err
		} else if strings.Contains(err.Error(), "conflict") {
			return http.StatusConflict, err
		} else {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusNoContent, nil
}

func (h *remoteStoreIntentHandler) DeleteTrafficGroupIntent(dtintName string, p string, ca string,
	v string, digName string,
) (interface{}, interface{}) {
	orch := h.orchInstance
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName + "/traffic-group-intents/" + dtintName
	resp, err := orch.apiDel(orchURL, dtintName)
	return resp, err
}

func (h *localStoreIntentHandler) GetTrafficGroupIntents(project string, compositeAppName string, version string,
	digName string,
) ([]byte, error) {
	var retval []byte
	c := localstore.NewTrafficGroupIntentClient()
	dTIntent, err := c.GetTrafficGroupIntents(project, compositeAppName, version, digName)
	if err != nil {
		log.Error(err.Error(), log.Fields{})
		if strings.Contains(err.Error(), "Unable to find") {
			return retval, err
		} else if strings.Contains(err.Error(), "db Find error") {
			return retval, err
		} else {
			return retval, err
		}
	}
	log.Infof("Get All dtcint localstore Composite app %s dig %s value %+v", compositeAppName,
		digName, dTIntent)
	retval, _ = json.Marshal(dTIntent)
	return retval, err
}

func (h *remoteStoreIntentHandler) GetTrafficGroupIntents(project string, compositeAppName string, version string,
	digName string,
) ([]byte, error) {
	orch := h.orchInstance

	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" +
		project + "/composite-apps/" + compositeAppName +
		"/" + version +
		"/deployment-intent-groups/" + digName + "/traffic-group-intents"
	// retcode, retval, err := orch.apiGet(orchURL, compositeAppName+"_dtcint")
	reply, err := orch.apiGet(orchURL, "testdtc")
	log.Infof("Get Dtint in Composite app %s dig %s status: %d", compositeAppName,
		digName, reply.StatusCode)
	return reply.Data, err
}

func (h *localStoreIntentHandler) CreateTrafficGroupIntent(g localstore.TrafficGroupIntent, p string, ca string,
	v string, digName string, exist bool,
) (interface{}, interface{}) {
	c := localstore.NewTrafficGroupIntentClient()

	_, createErr := c.CreateTrafficGroupIntent(g, p, ca, v, digName, true)
	if createErr != nil {
		log.Error(createErr.Error(), log.Fields{})
		if strings.Contains(createErr.Error(), "Unable to find the project") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the composite-app") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Unable to find the deploymentIntentGroupName") {
			return http.StatusNotFound, createErr
		} else if strings.Contains(createErr.Error(), "Intent already exists") {
			return http.StatusConflict, createErr
		} else {
			return http.StatusInternalServerError, createErr
		}
	}
	return http.StatusCreated, nil
}

func (h *remoteStoreIntentHandler) CreateTrafficGroupIntent(g localstore.TrafficGroupIntent, p string, ca string,
	v string, digName string, exist bool,
) (interface{}, interface{}) {
	orch := h.orchInstance
	// dtcintName := ca + "_dtcint"
	dtcintName := "testdtc"

	jsonLoad, _ := json.Marshal(g)
	orchURL := "http://" + orch.MiddleendConf.Dtc + "/v2/projects/" + p +
		"/composite-apps/" + ca + "/" + v +
		"/deployment-intent-groups/" + digName
	url := orchURL + "/traffic-group-intents"
	log.Infof("url traffic %s", url)
	resp, err := orch.apiPost(jsonLoad, url, dtcintName)

	return resp, err
}

func (h *dtcIntentHandler) createAnchor() interface{} {
	orch := h.orchInstance
	intentData := h.orchInstance.DigData
	// gPintName := intentData.CompositeAppName + "_dtcint"
	gPintName := "testdtc"

	vars := orch.Vars
	projectName := vars["projectName"]
	version := vars["version"]
	digName := intentData.Name

	dpi := localstore.TrafficGroupIntent{
		Metadata: localstore.TrafficGroupMetadata{
			Name:        gPintName,
			Description: "Traffic Group intent created from middleend",
			UserData1:   "data 1",
			UserData2:   "data2",
		},
	}

	// POST the Dtc placement intent
	log.Infof("compositeAppName %s", intentData.CompositeAppName)
	log.Infof("dpi %s", dpi)
	exist := false
	resp, err := orch.bstore.CreateTrafficGroupIntent(dpi, projectName, intentData.CompositeAppName, version, digName, exist)
	if err != nil {
		return err
	}
	if resp != http.StatusCreated {
		return resp
	}
	log.Infof("Dtc Placement intent response: %d", resp)
	jsonLoad, _ := json.Marshal(dpi)
	orch.response.payload["testdtc"] = jsonLoad
	orch.response.status["testdtc"] = resp.(int)
	return nil
}

func (h *dtcIntentHandler) createObject() interface{} {
	orch := h.orchInstance
	vars := orch.Vars
	projectName := vars["projectName"]
	version := vars["version"]
	// compositeAppName := vars["compositeAppName"]
	intentData := orch.DigData
	digName := intentData.Name
	// dTintName := intentData.CompositeAppName + "_dtcint"
	dTintName := "testdtc"
	exist := false
	var serverName string
	// var serviceName string

	for _, appData := range orch.DigData.Spec.Apps {
		if (appData.InboundServerIntent.ServiceName != "" && appData.InboundServerIntent.Protocol != "") && appData.InboundServerIntent.Port != "0" {
			serverName = appData.Metadata.Name + "-inboundserver"
			Port, _ := strconv.Atoi(appData.InboundServerIntent.Port)
			// serviceName = appData.Inboundserver.Spec.ServiceName
			dpi := localstore.InboundServerIntent{
				Metadata: localstore.Metadata{
					Name:        serverName,
					Description: "NA",
					UserData1:   "data1",
					UserData2:   "data2",
				},

				Spec: localstore.InbondServerIntentSpec{
					AppName:         appData.Metadata.Name,
					AppLabel:        "app=" + appData.Metadata.Name,
					ServiceName:     appData.InboundServerIntent.ServiceName,
					ExternalName:    "",
					Port:            Port,
					Protocol:        appData.InboundServerIntent.Protocol,
					ExternalSupport: false,
					ServiceMesh:     "istio",
				},
			}

			retcode, err := orch.bstore.CreateServerInboundIntent(dpi, projectName, intentData.CompositeAppName, version, digName, dTintName, exist)
			log.Infof("Creation of inbound server intent response: %s", retcode)
			if err != nil {
				return err
			}
			if retcode != nil && retcode.(int) != http.StatusCreated {
				return retcode.(int)
			}

			for _, appData := range orch.DigData.Spec.Apps {
				if serverName == appData.Metadata.Name+"-inboundserver" {
					continue
				}
				clientName := appData.Metadata.Name + "-inboundclient"
				dpi := localstore.InboundClientsIntent{
					Metadata: localstore.Metadata{
						Name:        clientName,
						Description: "NA",
						UserData1:   "data1",
						UserData2:   "data2",
					},
					Spec: localstore.InboundClientsIntentSpec{
						AppName:     appData.Metadata.Name,
						AppLabel:    "app=" + appData.Metadata.Name,
						ServiceName: appData.Metadata.Name + "-client-svc",
						Namespaces:  []string{},
						IpRange:     []string{},
					},
				}
				retcode, err := orch.bstore.CreateClientsInboundIntent(dpi, projectName, intentData.CompositeAppName, version, digName, dTintName, serverName, exist)
				log.Infof("Creation of inbound client intent response: %s", retcode)
				if err != nil {
					return err
				}
				if retcode != nil && retcode.(int) != http.StatusCreated {
					return retcode.(int)
				}
			}
		}
	}
	return nil
}

func (h *dtcIntentHandler) getObject() error {
	orch := h.orchInstance
	vars := orch.Vars
	retcode := 200
	dataRead := h.orchInstance.dataRead
	dtcData := make(map[string][]string)
	project := vars["projectName"]
	for _, compositeAppValue := range dataRead.compositeAppMap {
		if compositeAppValue.Status == "checkout" {
			continue
		}
		compositeAppMetadata := compositeAppValue.Metadata.Metadata
		compositeAppSpec := compositeAppValue.Metadata.Spec

		Dig := compositeAppValue.DigMap
		// AppData := compositeAppValue.AppsDataArray

		for digName, digValue := range Dig {
			for dtintName, dtintValue := range digValue.DtintMap {
				serverPint := []localstore.InboundServerIntent{}
				retval_server, err_server := orch.bstore.GetServerInboundIntents(project, compositeAppMetadata.Name,
					compositeAppSpec.Version, digName, dtintName)
				log.Infof("Get Dtint Dtc intent in Composite app %s dig %s Dtint %s status: %d",
					compositeAppMetadata.Name, digName, dtintName, retcode)
				if err_server != nil {
					log.Error("Failed to read dtc dint\n")
					return err_server
				}
				err := json.Unmarshal(retval_server, &serverPint)
				if err != nil {
					log.Errorf("Failed to unmarshal json %s\n", err)
					return err
				}
				dtintValue.ServerIntentArray = serverPint
				log.Infof("dtint pratik %v", dtintValue.ServerIntentArray)

				for _, servername := range serverPint {
					clientPint := []localstore.InboundClientsIntent{}
					retval, err := orch.bstore.GetClientsInboundIntents(project, compositeAppMetadata.Name,
						compositeAppSpec.Version, digName, dtintName, servername.Metadata.Name)
					log.Infof("Get Dtint Dtc intent in Composite app %s dig %s Dtint %s status: %d", compositeAppMetadata.Name, digName, dtintName, retcode)
					if err != nil {
						log.Error("Failed to read dtc dint\n")
						return err
					}
					err = json.Unmarshal(retval, &clientPint)
					if err != nil {
						log.Errorf("Failed to unmarshal json %s\n", err)
						return err
					}
					if clientPint != nil {
						for _, clientname := range clientPint {
							dtcData[servername.Metadata.Name] = append(dtcData[servername.Metadata.Name], clientname.Metadata.Name)
						}
					} else {
						dtcData[servername.Metadata.Name] = append(dtcData[servername.Metadata.Name], "no")
					}
				}
			}
		}
	}
	orch.dtck8sInfo = dtcData
	log.Infof("dtc data %v", orch.dtck8sInfo)
	return nil
}

func (h *dtcIntentHandler) getAnchor() error {
	orch := h.orchInstance
	vars := orch.Vars
	dataRead := h.orchInstance.dataRead
	project := vars["projectName"]

	for _, compositeAppValue := range dataRead.compositeAppMap {
		if compositeAppValue.Status == "checkout" {
			continue
		}
		compositeAppMetadata := compositeAppValue.Metadata.Metadata
		compositeAppSpec := compositeAppValue.Metadata.Spec
		Dig := compositeAppValue.DigMap
		for digName, digValue := range Dig {
			var dtcintList []localstore.TrafficGroupIntent
			retval, err := orch.bstore.GetTrafficGroupIntents(project, compositeAppMetadata.Name,
				compositeAppSpec.Version, digName)
			// log.Infof("Get Dtcint in Composite app %s dig %s status: %d", vars["compositeAppName"], digName, retcode)
			if err != nil {
				log.Error("Failed to read dtcint\n")
				return err
			}
			if err := json.Unmarshal(retval, &dtcintList); err != nil {
				log.Error(err, PrintFunctionName())
			}
			log.Info("anchor payload", &dtcintList)
			log.Info("anchor payload", dtcintList)
			digValue.DtintMap = make(map[string]*DtintData, len(dtcintList))
			for _, value := range dtcintList {
				var DtintDataInstance DtintData
				DtintDataInstance.Dpint = value
				digValue.DtintMap[value.Metadata.Name] = &DtintDataInstance
			}
		}
	}
	return nil
}

func (h *dtcIntentHandler) deleteObject() interface{} {
	orch := h.orchInstance
	vars := orch.Vars
	dataRead := h.orchInstance.dataRead
	dtcData := h.orchInstance.dtck8sInfo
	project := vars["projectName"]
	log.Info(project)
	for _, compositeAppValue := range dataRead.compositeAppMap {
		if compositeAppValue.Status == "checkout" {
			continue
		}
		compositeAppMetadata := compositeAppValue.Metadata.Metadata
		compositeAppSpec := compositeAppValue.Metadata.Spec
		Dig := compositeAppValue.DigMap
		log.Info(compositeAppMetadata)
		log.Info(compositeAppSpec)
		// AppData := compositeAppValue.AppsDataArray

		// loop through all app intens in the dtint
		for digName, digValue := range Dig {
			for dtintName := range digValue.DtintMap {
				for server, clientlist := range dtcData {
					log.Infof("client %v, server %v", clientlist, server)
					log.Infof("client pratik %v", clientlist)
					log.Infof("server pratik %v", server)
					log.Infof("client len %v", len(clientlist))
					for _, client := range clientlist {
						if client == "no" {
							continue
						}
						resp, err := orch.bstore.DeleteClientsInboundIntent(client, project,
							compositeAppMetadata.Name, compositeAppSpec.Version,
							digName, dtintName, server)
						if err != nil {
							return err
						}
						if resp != http.StatusNoContent {
							return resp
						}
						log.Infof("Delete client dpint intents response: %d", resp)

					}

					// query based on app name.
					resp_server, err_server := orch.bstore.DeleteServerInboundIntent(server, project,
						compositeAppMetadata.Name, compositeAppSpec.Version,
						digName, dtintName)
					if err_server != nil {
						return err_server
					}
					if resp_server != http.StatusNoContent {
						return resp_server
					}
					log.Infof("Delete server dpint intents response: %d", resp_server)
				}
			}
		}
	}
	return nil
}

func (h dtcIntentHandler) deleteAnchor() interface{} {
	orch := h.orchInstance
	vars := orch.Vars
	dataRead := h.orchInstance.dataRead
	for _, compositeAppValue := range dataRead.compositeAppMap {
		if compositeAppValue.Status == "checkout" {
			continue
		}
		compositeAppMetadata := compositeAppValue.Metadata.Metadata
		compositeAppSpec := compositeAppValue.Metadata.Spec
		Dig := compositeAppValue.DigMap

		// loop through all app intens in the dtint
		for digName, digValue := range Dig {
			for dtintName := range digValue.DtintMap {
				log.Infof("Delete dtint  %s", h.orchURL)
				resp, err := orch.bstore.DeleteTrafficGroupIntent(dtintName, vars["projectName"],
					compositeAppMetadata.Name, compositeAppSpec.Version, digName)
				if err != nil {
					return err
				}
				if resp != http.StatusNoContent {
					return resp
				}
				log.Infof("Delete dtint response: %d", resp)
			}
		}
	}
	return nil
}

func addDtcIntent(I orchWorkflow) interface{} {
	// 1. Create the Anchor point
	err := I.createAnchor()
	if err != nil {
		return err
	}
	// 2. Create the Objects
	err = I.createObject()
	if err != nil {
		return err
	}
	return nil
}
