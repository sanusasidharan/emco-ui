// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation

package localstore

import (
	"example.com/middleend/db"
	pkgerrors "github.com/pkg/errors"
)

type InboundServerIntent struct {
	Metadata Metadata               `json:"metadata"`
	Spec     InbondServerIntentSpec `json:"spec"`
}

type InbondServerIntentSpec struct {
	AppName         string `json:"app"`
	AppLabel        string `json:"appLabel"`
	ServiceName     string `json:"serviceName"`
	ExternalName    string `json:"externalName", default:""`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	ExternalSupport bool   `json:"externalSupport", default:false`
	ServiceMesh     string `json:"serviceMesh", default:"none"`
}


type InboundServerIntentSpec struct {
	AppName         string `json:"app"`
        AppLabel        string `json:"appLabel"`
        ServiceName     string `json:"serviceName"`
        ExternalName    string `json:"externalName", default:""`
        Port            string    `json:"port"`
        Protocol        string `json:"protocol"`
        ExternalSupport bool   `json:"externalSupport", default:false`
        ServiceMesh     string `json:"serviceMesh", default:"none"`
}


type InboundServerIntentDbClient struct {
	db ClientDbInfo
}

// ServerInboundIntentKey is the key structure that is used in the database
type InboundServerIntentKey struct {
	Project                   string `json:"project"`
	CompositeApp              string `json:"compositeApp"`
	CompositeAppVersion       string `json:"compositeAppVersion"`
	DeploymentIntentGroupName string `json:"deploymentIntentGroup"`
	TrafficGroupIntentName    string `json:"trafficGroupIntent"`
	ServerInboundIntentName   string `json:"inboundServerIntent"`
}

func NewServerInboundIntentClient() *InboundServerIntentDbClient {
	return &InboundServerIntentDbClient{
		db: ClientDbInfo{
			storeName: "resources",
			tagMeta:   "serverintentmetadata",
		},
	}
}

func (v InboundServerIntentDbClient) CreateServerInboundIntent(isi InboundServerIntent, project, compositeapp, compositeappversion, deploymentintentgroupname, trafficintentgroupname string, exists bool) (InboundServerIntent, error) {

	//Construct key and tag to select the entry
	key := InboundServerIntentKey{
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: deploymentintentgroupname,
		TrafficGroupIntentName:    trafficintentgroupname,
		ServerInboundIntentName:   isi.Metadata.Name,
	}

	err := db.DBconn.Insert(v.db.storeName, key, nil, v.db.tagMeta, isi)
	if err != nil {
		return InboundServerIntent{}, pkgerrors.Wrap(err, "Creating DB Entry")
	}

	return isi, nil
}

// GetServerInboundIntent returns the ServerInboundIntent for corresponding name
func (v *InboundServerIntentDbClient) GetServerInboundIntent(name, project, compositeapp, compositeappversion, dig, trafficintentgroupname string) (InboundServerIntent, error) {

	//Construct key and tag to select the entry
	key := InboundServerIntentKey{
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: dig,
		TrafficGroupIntentName:    trafficintentgroupname,
		ServerInboundIntentName:   name,
	}

	value, err := db.DBconn.Find(v.db.storeName, key, v.db.tagMeta)
	if err != nil {
		return InboundServerIntent{}, err
	} else if len(value) == 0 {
		return InboundServerIntent{}, pkgerrors.New("Inbound server intent not found")
	}

	//value is a byte array
	if value != nil {
		wi := InboundServerIntent{}
		err = db.DBconn.Unmarshal(value[0], &wi)
		if err != nil {
			return InboundServerIntent{}, err
		}
		return wi, nil
	}

	return InboundServerIntent{}, pkgerrors.New("Unknown Error")
}

// GetServerInboundIntents returns all of the ServerInboundIntents
func (v *InboundServerIntentDbClient) GetServerInboundIntents(project, compositeapp, compositeappversion, deploymentintentgroupname, trafficintentgroupname string) ([]InboundServerIntent, error) {

	//Construct key and tag to select the entry
	key := InboundServerIntentKey{
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: deploymentintentgroupname,
		TrafficGroupIntentName:    trafficintentgroupname,
		ServerInboundIntentName:   "",
	}

	var resp []InboundServerIntent
	values, err := db.DBconn.Find(v.db.storeName, key, v.db.tagMeta)
	if err != nil {
		return []InboundServerIntent{}, err
	}

	for _, value := range values {
		is := InboundServerIntent{}
		err = db.DBconn.Unmarshal(value, &is)
		if err != nil {
			return []InboundServerIntent{}, err
		}
		resp = append(resp, is)
	}

	return resp, nil
}

// Delete the  ServerInboundIntents from database
func (v *InboundServerIntentDbClient) DeleteServerInboundIntent(name, project, compositeapp, compositeappversion, dig, trafficintentgroupname string) error {

	//Construct key and tag to select the entry
	key := InboundServerIntentKey{
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: dig,
		TrafficGroupIntentName:    trafficintentgroupname,
		ServerInboundIntentName:   name,
	}

	err := db.DBconn.Remove(v.db.storeName, key)
	return err
}
