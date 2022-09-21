// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation

package localstore

/*
 This files store/retrieves the application intents data in
 middleend db, in the same format as it is stored in EMCO.
*/

import (

	//"reflect"

	"example.com/middleend/db"
	pkgerrors "github.com/pkg/errors"
)

type TrafficGroupIntent struct {
	Metadata TrafficGroupMetadata `json:"metadata"`
}

type TrafficGroupMetadata struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"-"`
	UserData1   string `json:"userData1" yaml:"-"`
	UserData2   string `json:"userData2" yaml:"-"`
}


type TrafficGroupIntentDbClient struct {
	storeName   string
	tagMetaData string
}

// TrafficGroupIntentKey is the key structure that is used in the database
type TrafficGroupIntentKey struct {
	TrafficGroupIntentName    string `json:"trafficGroupIntent"`
	Project                   string `json:"project"`
	CompositeApp              string `json:"compositeApp"`
	CompositeAppVersion       string `json:"compositeAppVersion"`
	DeploymentIntentGroupName string `json:"deploymentIntentGroup"`
}

func NewTrafficGroupIntentClient() *TrafficGroupIntentDbClient {
	return &TrafficGroupIntentDbClient{

		storeName:   "resources",
		tagMetaData: "traffigroupintentmetadata",
	}
}

func (v TrafficGroupIntentDbClient) CreateTrafficGroupIntent(tci TrafficGroupIntent, project, compositeapp, compositeappversion, deploymentintentgroupname string, exists bool) (TrafficGroupIntent, error) {

	//Construct key and tag to select the entry
	key := TrafficGroupIntentKey{
		TrafficGroupIntentName:    tci.Metadata.Name,
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: deploymentintentgroupname,
	}
	//Check if this TrafficGroupIntent already exists
	//_, err := v.GetTrafficGroupIntent(tci.Metadata.Name, project, compositeapp, compositeappversion, deploymentintentgroupname)
	//if err == nil && !exists {
	//	return TrafficGroupIntent{}, pkgerrors.New("TrafficGroupIntent already exists")
	//}

	err := db.DBconn.Insert(v.storeName, key, nil, v.tagMetaData, tci)
	if err != nil {
		return TrafficGroupIntent{}, pkgerrors.Wrap(err, "Creating DB Entry")
	}

	return tci, nil
}

// GetTrafficGroupIntent returns the TrafficGroupIntent for corresponding name
func (v *TrafficGroupIntentDbClient) GetTrafficGroupIntent(name, project, compositeapp, compositeappversion, dig string) (TrafficGroupIntent, error) {

	//Construct key and tag to select the entry
	key := TrafficGroupIntentKey{
		TrafficGroupIntentName:    name,
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: dig,
	}

	value, err := db.DBconn.Find(v.storeName, key, v.tagMetaData)
	if err != nil {
		return TrafficGroupIntent{}, err
	} else if len(value) == 0 {
		return TrafficGroupIntent{}, pkgerrors.New("Traffic group intent not found")

	}

	//value is a byte array
	if value != nil {
		tgi := TrafficGroupIntent{}
		err = db.DBconn.Unmarshal(value[0], &tgi)
		if err != nil {
			return TrafficGroupIntent{}, err
		}
		return tgi, nil
	}

	return TrafficGroupIntent{}, pkgerrors.New("Unknown Error")

}

// GetTrafficGroupIntents returns all of the TrafficGroupIntents
func (v *TrafficGroupIntentDbClient) GetTrafficGroupIntents(project, compositeapp, compositeappversion, dig string) ([]TrafficGroupIntent, error) {

	//Construct key and tag to select the entry
	key := TrafficGroupIntentKey{
		TrafficGroupIntentName:    "",
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: dig,
	}

	var resp []TrafficGroupIntent
	values, err := db.DBconn.Find(v.storeName, key, v.tagMetaData)
	if err != nil {
		return []TrafficGroupIntent{}, err
	}

	for _, value := range values {
		tgi := TrafficGroupIntent{}
		err = db.DBconn.Unmarshal(value, &tgi)
		if err != nil {
			return []TrafficGroupIntent{}, err
		}
		resp = append(resp, tgi)
	}

	return resp, nil
}

// Delete the  TrafficGroupIntent from database
func (v *TrafficGroupIntentDbClient) DeleteTrafficGroupIntent(name, project, compositeapp, compositeappversion, dig string) error {

	//Construct key and tag to select the entry
	key := TrafficGroupIntentKey{
		TrafficGroupIntentName:    name,
		Project:                   project,
		CompositeApp:              compositeapp,
		CompositeAppVersion:       compositeappversion,
		DeploymentIntentGroupName: dig,
	}

	err := db.DBconn.Remove(v.storeName, key)
	return err
}
