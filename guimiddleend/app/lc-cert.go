package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (h *lcHandler) generateCaIntentName(cProject string) string {
	return cProject + "-certintent1" // FIXME:  There can be multiple intents under cluster provider
}

func (h *lcHandler) CreateCertLogicalClouds(caIntent, cProject string, clouds []*CaCertLogicalCloud) error {
	for i, cl := range clouds {
		err := h.CreateCertLogicalCloud(caIntent, cProject, cl)
		if err != nil {
			for j := i; j < len(clouds); j++ {
				err := h.DeleteCertLogicalCloud(caIntent, cProject, cl.Metadata.Name)
				if err != nil {
					log.Println(err)
				}
			}
			return fmt.Errorf("Error occured during creating the logical clouds - %s. Rolling back", err)
		}
	}
	return nil
}

func (h *lcHandler) DeleteCertLogicalClouds(caIntent, cProject string, clusters []string) error {
	for _, cl := range clusters {
		err := h.DeleteCertLogicalCloud(caIntent, cProject, cl)
		if err != nil {
			return fmt.Errorf("Error occured during deleting the logical clouds - %s", err)
		}
	}
	return nil
}

func (h *lcHandler) DeleteCertLogicalCloud(caIntent, cProject string, cluster string) error {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/logical-clouds/" + cluster

	statusCode, err := h.apiDel(url, "DeleteCertLogicalCloud")
	if err != nil {
		return err
	}

	if !(statusCode == http.StatusOK ||
		statusCode == http.StatusCreated ||
		statusCode == http.StatusAccepted ||
		statusCode == http.StatusNoContent ||
		statusCode == http.StatusNotFound) {
		return fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}
	return nil
}

func (h *lcHandler) CreateCertLogicalCloud(caIntent, cProject string, cloud *CaCertLogicalCloud) error {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/logical-clouds"
	cloud.Spec.LogicalCloud = cloud.Metadata.Name
	payload, err := json.Marshal(cloud)
	if err != nil {
		return err
	}

	fmt.Println("payload", string(payload))

	statusCode, err := h.apiPost(payload, url, "payload")
	if err != nil {
		return err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return nil
}

func (h *lcHandler) InstantiateEnrollment(caIntent, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/enrollment/instantiate"
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return statusCode, nil
}

func (h *lcHandler) TerminateEnrollment(caIntent, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/enrollment/terminate"
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return statusCode, nil
}

func (h *lcHandler) InstantiateDistribution(caIntent, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/distribution/instantiate"
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return statusCode, nil
}

func (h *lcHandler) TerminateDistribution(caIntent, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/distribution/terminate"
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return statusCode, nil
}

func (h *lcHandler) GetCaCertEnrollmentStatus(caIntent, cProject string) (*CertStatus, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/enrollment/status"

	reply, err := h.apiGet(url, "GetCaCertEnrollmentStatus")
	if err != nil {
		return nil, err
	}

	status := &CertStatus{}
	if err := json.Unmarshal(reply.Data, status); err != nil {
		return nil, err
	}

	return status, nil
}

func (h *lcHandler) GetCaCertDistributionStatus(caIntent, cProject string) (*CertStatus, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/distribution/status"

	reply, err := h.apiGet(url, "GetCaCertDistributionStatus")
	if err != nil {
		return nil, err
	}

	status := &CertStatus{}
	if err := json.Unmarshal(reply.Data, status); err != nil {
		return nil, err
	}

	return status, nil
}

func (h *lcHandler) GetCaCert(caIntent, cProject string) (*CaCert, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent

	reply, err := h.apiGet(url, "caCertsGet")
	if err != nil {
		return nil, err
	}

	caCert := &CaCert{}
	if err := json.Unmarshal(reply.Data, caCert); err != nil {
		return nil, err
	}

	return caCert, nil
}

func (h *lcHandler) PostCaCert(caCert *CaCert, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs"
	statusCode := 500
	payload, err := json.Marshal(caCert)
	if err != nil {
		return statusCode, err
	}

	sc, err := h.apiPost(payload, url, "payload")
	statusCode = sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s", h.response.payload["payload"])
	}

	return statusCode, nil
}

func (h *lcHandler) DeleteCaCert(caIntent, cProject string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent

	sc, err := h.apiDel(url, "caCertDel")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK ||
		statusCode == http.StatusCreated ||
		statusCode == http.StatusAccepted ||
		statusCode == http.StatusNoContent ||
		statusCode == http.StatusNotFound) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	return statusCode, nil
}

func (h *lcHandler) GetCAIntentLogicalClouds(caIntent, cProject string) ([]*CaCertLogicalCloud, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/logical-clouds"

	reply, err := h.apiGet(url, "payload")
	if err != nil {
		return nil, err
	}

	clouds := []*CaCertLogicalCloud{}
	if err := json.Unmarshal(reply.Data, &clouds); err != nil {
		return nil, err
	}

	return clouds, nil
}

func (h *lcHandler) GetCAIntentClusters(caIntent, cProject string) ([]*CertIntentCluster, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/projects/" + cProject + "/ca-certs/" + caIntent + "/clusters"

	reply, err := h.apiGet(url, "GetCAIntentClusters")
	if err != nil {
		return nil, err
	}

	clusters := []*CertIntentCluster{}
	if err := json.Unmarshal(reply.Data, &clusters); err != nil {
		return nil, err
	}

	return clusters, nil
}

func (h *lcHandler) GetLogicalCloudsByProject(cProject string) ([]*CaCertLogicalCloud, error) {
	url := "http://" + h.MiddleendConf.Dcm + "/v2/projects/" + cProject + "/logical-clouds"

	reply, err := h.apiGet(url, "payload")
	if err != nil {
		return nil, err
	}

	clouds := []*CaCertLogicalCloud{}
	if err := json.Unmarshal(reply.Data, &clouds); err != nil {
		return nil, err
	}

	return clouds, nil
}

func findClustersByNames(clusters []*ClusterWithLabel, names []string) ([]*ClusterWithLabel, error) {
	clMap := make(map[string]*ClusterWithLabel)
	filtered := []*ClusterWithLabel{}
	for _, cl := range clusters {
		clMap[cl.Metadata.Name] = cl
	}
	for _, name := range names {
		if cl, ok := clMap[name]; ok {
			filtered = append(filtered, cl)
		} else {
			return nil, fmt.Errorf("Couldn't find cluster '%s'", name)
		}
	}
	return filtered, nil
}

func findCloudsByNames(clouds []*CaCertLogicalCloud, names []string) ([]*CaCertLogicalCloud, error) {
	clMap := make(map[string]*CaCertLogicalCloud)
	filtered := []*CaCertLogicalCloud{}
	for _, cl := range clouds {
		clMap[cl.Metadata.Name] = cl
	}
	for _, name := range names {
		if cl, ok := clMap[name]; ok {
			filtered = append(filtered, cl)
		} else {
			return nil, fmt.Errorf("Couldn't find logical cloud '%s'", name)
		}
	}
	return filtered, nil
}

func (h *lcHandler) caCertInstantiate(caIntent, clusterProvider string) (int, error) {
	statusCode, err := h.InstantiateEnrollment(caIntent, clusterProvider)
	if err != nil {
		return statusCode, err
	}

	for i := 0; i < 4; i++ {
		enrStatus, err := h.GetCaCertEnrollmentStatus(caIntent, clusterProvider)
		if err != nil {
			return http.StatusBadRequest, err
		}
		if enrStatus.ReadyStatus == "Ready" {
			break
		}
		time.Sleep(time.Duration(2+i) * time.Second)
	}

	statusCode, err = h.InstantiateDistribution(caIntent, clusterProvider)
	if err != nil {
		err = fmt.Errorf("Coudn't instantiate the distribution. Enrollment state is not ready")
		return statusCode, err
	}
	return 200, nil
}

func (h *lcHandler) caCertTerminate(caIntent, clusterProvider string) (int, error) {
	dstStatus, err := h.GetCaCertDistributionStatus(caIntent, clusterProvider)
	if err != nil {
		return http.StatusBadRequest, err
	}

	dstLastStatus := ""
	if len(dstStatus.States.Actions) > 0 {
		lstState := dstStatus.States.Actions[len(dstStatus.States.Actions)-1]
		if lstState.Revision == len(dstStatus.States.Actions)-1 {
			dstLastStatus = lstState.State
		}
	}

	if dstLastStatus == "Instantiated" {
		statusCode, err := h.TerminateDistribution(caIntent, clusterProvider)
		if err != nil {
			return statusCode, err
		}
	}

	enrStatus, err := h.GetCaCertEnrollmentStatus(caIntent, clusterProvider)
	if err != nil {
		return http.StatusBadRequest, err
	}

	enrLastStatus := ""
	if len(enrStatus.States.Actions) > 0 {
		lstState := enrStatus.States.Actions[len(enrStatus.States.Actions)-1]
		if lstState.Revision == len(enrStatus.States.Actions)-1 {
			enrLastStatus = lstState.State
		}
	}

	if enrLastStatus == "Instantiated" {
		statusCode, err := h.TerminateEnrollment(caIntent, clusterProvider)
		if err != nil {
			return statusCode, err
		}
	}

	return 200, nil
}

func (h *lcHandler) caCertReInstantiate(caIntent, clusterProvider string) (int, error) {
	statusCode, err := h.caCertTerminate(caIntent, clusterProvider)
	if err != nil {
		return statusCode, err
	}
	return h.caCertInstantiate(caIntent, clusterProvider)
}
