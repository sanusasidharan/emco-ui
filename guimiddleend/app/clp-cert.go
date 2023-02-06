package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (h *clpHandler) generateCaIntentName(cProvider string) string {
	return cProvider + "-certintent1" // FIXME:  There can be multiple intents under cluster provider
}

func (h *clpHandler) CreateCertClusters(caIntent, cProvider string, clusters []*ClusterWithLabel) error {
	for i, cl := range clusters {
		err := h.CreateCertCluster(caIntent, cProvider, cl)
		if err != nil {
			for j := i; j < len(clusters); j++ {
				err := h.DeleteCertCluster(caIntent, cProvider, cl.Metadata.Name)
				if err != nil {
					log.Println(err)
				}
			}
			return fmt.Errorf("Error occured during creating the clusters - %s. Rolling back", err)
		}
	}
	return nil
}

func (h *clpHandler) DeleteCertClusters(caIntent, cProvider string, clusters []string) error {
	for _, cl := range clusters {
		err := h.DeleteCertCluster(caIntent, cProvider, cl)
		if err != nil {
			return fmt.Errorf("Error occured during deleting the clusters - %s", err)
		}
	}
	return nil
}

func (h *clpHandler) DeleteCertCluster(caIntent, cProvider string, cluster string) error {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/clusters/" + cluster
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Deleting Cert cluster")
	statusCode, err := h.apiDel(url, "DeleteCertCluster")
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
	l.Debug("Cert cluster succesfully deleted")
	return nil
}

func (h *clpHandler) CreateCertCluster(caIntent, cProvider string, cluster *ClusterWithLabel) error {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/clusters"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Creating Cert Cluster")
	certIntentCluster := &CertIntentCluster{
		Metadata: apiMetaData{Name: cluster.Metadata.Name},
		Spec: CertIntentClusterSpec{
			Scope:           "name",
			Cluster:         cluster.Metadata.Name,
			ClusterProvider: cProvider,
		},
	}

	payload, err := json.Marshal(certIntentCluster)
	if err != nil {
		return err
	}

	l = l.WithField("request_payload", string(payload))

	statusCode, err := h.apiPost(payload, url, "payload")
	if err != nil {
		return err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	l.Debug("Cert Cluster successfully created")

	return nil
}

func (h *clpHandler) InstantiateEnrollment(caIntent, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/enrollment/instantiate"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Instantiating the enrollment")
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}

	l.Debug("Enrollment instantiated")

	return statusCode, nil
}

func (h *clpHandler) TerminateEnrollment(caIntent, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/enrollment/terminate"
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

func (h *clpHandler) InstantiateDistribution(caIntent, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/distribution/instantiate"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Instantiating the distribution")
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}
	l.Debug("Distribution instantiated")
	return statusCode, nil
}

func (h *clpHandler) TerminateDistribution(caIntent, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/distribution/terminate"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Terminating the distribution")
	sc, err := h.apiPost(nil, url, "payload")
	statusCode := sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s code - %d", h.response.payload["payload"], statusCode)
	}
	l.Debug("Distribution terminated")
	return statusCode, nil
}

func (h *clpHandler) GetCaCertEnrollmentStatus(caIntent, cProvider string) (*CertStatus, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/enrollment/status"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Fetching CA enrollment status")
	reply, err := h.apiGet(url, "GetCaCertEnrollmentStatus")
	if err != nil {
		return nil, err
	}

	status := &CertStatus{}
	if err := json.Unmarshal(reply.Data, status); err != nil {
		return nil, err
	}

	l.WithField("response_payload", string(reply.Data)).Debug("Ca enrollment status successfully fetched")

	return status, nil
}

func (h *clpHandler) GetCaCertDistributionStatus(caIntent, cProvider string) (*CertStatus, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/distribution/status"

	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Fetching CA distribution status")
	reply, err := h.apiGet(url, "GetCaCertDistributionStatus")
	if err != nil {
		return nil, err
	}

	status := &CertStatus{}
	if err := json.Unmarshal(reply.Data, status); err != nil {
		return nil, err
	}
	l.WithField("response_payload", string(reply.Data)).Debug("Ca distribution status successfully fetched")
	return status, nil
}

func (h *clpHandler) GetCaCert(caIntent, cProvider string) (*CaCert, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Fetching CA Cert")
	reply, err := h.apiGet(url, "caCertsGet")
	if err != nil {
		return nil, err
	}

	caCert := &CaCert{}
	if err := json.Unmarshal(reply.Data, caCert); err != nil {
		return nil, err
	}

	l.WithField("response_payload", string(reply.Data)).Debug("CA Cert succesfully fetched")

	return caCert, nil
}

func (h *clpHandler) PostCaCert(caCert *CaCert, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Posting CA Cert")
	statusCode := 500
	payload, err := json.Marshal(caCert)
	if err != nil {
		return statusCode, err
	}

	l = l.WithField("request_payload", string(payload))

	sc, err := h.apiPost(payload, url, "payload")
	statusCode = sc.(int)
	if err != nil {
		return statusCode, err
	}

	if !(statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted) {
		return statusCode, fmt.Errorf("%s", h.response.payload["payload"])
	}

	l.Debug("CA Cert successfully posted")

	return statusCode, nil
}

func (h *clpHandler) DeleteCaCert(caIntent, cProvider string) (int, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Deleting CA Cert")
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

	l.Debug("CA Cert succesfully deleted")

	return statusCode, nil
}

func (h *clpHandler) GetCAIntentClusters(caIntent, cProvider string) ([]*CertIntentCluster, error) {
	url := "http://" + h.MiddleendConf.Cert + "/v2/cluster-providers/" + cProvider + "/ca-certs/" + caIntent + "/clusters"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Fetching CA Intent Clusters")
	reply, err := h.apiGet(url, "GetCAIntentClusters")
	if err != nil {
		return nil, err
	}

	clusters := []*CertIntentCluster{}
	if err := json.Unmarshal(reply.Data, &clusters); err != nil {
		return nil, err
	}
	l.WithField("response_payload", string(reply.Data)).Debug("CA Intent Clusters succesfully fetched")
	return clusters, nil
}

func (h *clpHandler) GetClustersByProvider(cProvider string) ([]*ClusterWithLabel, error) {
	url := "http://" + h.MiddleendConf.Clm + "/v2/cluster-providers/" + cProvider + "/clusters?withLabels=true"
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName(), "emco_url": url})
	l.Debug("Fetching Clusters")
	reply, err := h.apiGet(url, "caCertsGet")
	if err != nil {
		return nil, err
	}

	clusters := []*ClusterWithLabel{}
	if err := json.Unmarshal(reply.Data, &clusters); err != nil {
		return nil, err
	}

	l.WithField("response_payload", string(reply.Data)).Debug("Clusters succesfully fetched")

	return clusters, nil
}

// For future use
// func (h *OrchestrationHandler) findClustersByNames(clusters []*ClusterWithLabel, names []string) ([]*ClusterWithLabel, error) {
// 	clMap := make(map[string]*ClusterWithLabel)
// 	filtered := []*ClusterWithLabel{}
// 	for _, cl := range clusters {
// 		clMap[cl.Metadata.Name] = cl
// 	}
// 	for _, name := range names {
// 		if cl, ok := clMap[name]; ok {
// 			filtered = append(filtered, cl)
// 		} else {
// 			return nil, fmt.Errorf("Couldn't find cluster '%s'", name)
// 		}
// 	}
// 	return filtered, nil
// }

func (h *clpHandler) caCertInstantiate(caIntent, clusterProvider string) (int, error) {
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName()})
	l.Debug("Instantiating CA Cert")
	statusCode, err := h.InstantiateEnrollment(caIntent, clusterProvider)
	if err != nil {
		return statusCode, err
	}

	for i := 0; i < 4; i++ {
		tmp_logger := l.WithFields(logrus.Fields{"enr_attempt": i + 1})
		tmp_logger.Debug("Checking enrollment status")
		enrStatus, err := h.GetCaCertEnrollmentStatus(caIntent, clusterProvider)
		if err != nil {
			return http.StatusBadRequest, err
		}
		tmp_logger.Debug("Enrollment status", enrStatus.ReadyStatus)
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

	l.Debug("CA Cert successfully instantiated")
	return 200, nil
}

func (h *clpHandler) caCertTerminate(caIntent, clusterProvider string) (int, error) {
	l := h.Logger.WithFields(logrus.Fields{"function": PrintFunctionName()})
	l.Debug("Instantiating CA Cert")
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

func (h *clpHandler) caCertReInstantiate(caIntent, clusterProvider string) (int, error) {
	statusCode, err := h.caCertTerminate(caIntent, clusterProvider)
	if err != nil {
		return statusCode, err
	}
	return h.caCertInstantiate(caIntent, clusterProvider)
}
