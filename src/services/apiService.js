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
import axios from "axios";

//axios.defaults.baseURL = "http://147.75.53.57"  || "";
//orchestrator
//projects

const instance = axios.create({
  baseURL: "http://147.75.53.57",
  withCredentials: false,
  headers: {
    'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods':'*',
  }
});
const createProject = (request) => {
  return instance.post("/v2/projects", { ...request })
      .then((res) => res.data);
};
const updateProject = (request) => {
  return instance
    .put(`/v2/projects/${request.metadata.name}`, {
      ...request })
    .then((res) => res.data);
};
const deleteProject = (projectName) => {
  return instance.delete(`/v2/projects/${projectName}`);
};
const getProjectDetails = (projectName) => {
  return instance.get(`/v2/projects/${projectName}`).then((res) => res.data);
};
const getAllProjects = () => {
  return instance.get("/v2/projects").then((response) => {
    return response.data;
  });
};

//composite apps
const getCompositeApps = (request) => {
  return instance
    .get(`/v2/projects/${request.projectName}/composite-apps`)
    .then((res) => {
      return res.data;
    });
};

const createCompositeApp = ({ projectName, ...request }) => {
  return instance
    .post(`/v2/projects/${projectName}/composite-apps`, request.payload)
    .then((res) => {
      return res.data;
    });
};
const updateCompositeApp = (request) => {
  return instance
    .put(
      `/v2/projects/${request.projectName}/composite-apps/${request.payload.name}/${request.compositeAppVersion}`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

//apps
const getApps = (request) => {
  return instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/apps`
    )
    .then((res) => {
      return res.data;
    });
};
const addApp = (request) => {
  return instance
    .post(
      `/v2/projects/${request.get("projectName")}/composite-apps/${request.get(
        "compositeAppName"
      )}/${request.get("compositeAppVersion")}/apps`,
      request
    )
    .then((res) => {
      return res.data;
    });
};
const updateApp = (request) => {
  return instance
    .put(
      `/v2/projects/${request.get("projectName")}/composite-apps/${request.get(
        "compositeAppName"
      )}/${request.get("compositeAppVersion")}/apps/${request.get("appName")}`,
      request
    )
    .then((res) => {
      return res.data;
    });
};
const deleteApp = (request) => {
  return instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/apps/${request.appName}`
    )
    .then((res) => {
      return res.data;
    });
};

//profiles
const createCompositeProfile = (request) => {
  return instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/composite-profiles`,
      request.payload
    )
    .then((res) => res.data);
};
const getCompositeProfiles = (request) => {
  return instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/composite-profiles`
    )
    .then((res) => res.data);
};
const deleteCompositeProfile = (request) => {
  return instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/composite-profiles/${request.compositeProfileName}`
    )
    .then((res) => res.data);
};
const addProfile = (request) => {
  return instance
    .post(
      `/v2/projects/${request.get("projectName")}/composite-apps/${request.get(
        "compositeAppName"
      )}/${request.get("compositeAppVersion")}/composite-profiles/${request.get(
        "compositeProfileName"
      )}/profiles`,
      request
    )
    .then((res) => {
      return res.data;
    });
};
const getProfiles = (request) => {
  return instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/composite-profiles/${request.compositeProfileName}/profiles`
    )
    .then((res) => res.data);
};
const deleteProfile = (request) => {
  return instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/composite-profiles/${request.compositeProfileName}/profiles/${request.profileName}`
    )
    .then((res) => res.data);
};

//placement intents
const getGenericPlacementIntents = (request) => {
  return instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents`
    )
    .then((res) => res.data);
};
const createGenericPlacementIntent = (request) => {
  return instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents`,
      request.payload
    )
    .then((res) => res.data);
};
const deleteGenericPlacementIntent = (request) => {
  return instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents/${request.genericPlacementIntentName}`
    )
    .then((res) => res.data);
};
const getAppPlacementIntents = (request) => {
  return instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents/${request.genericPlacementIntentName}/app-intents`
    )
    .then((res) => res.data);
};
const addAppPlacementIntent = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents/${request.genericPlacementIntentName}/app-intents`,
      request.payload
    )
    .then((res) => res.data);
};
const deleteAppPlacementIntent = (request) => {
  return  instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/generic-placement-intents/${request.genericPlacementIntentName}/app-intents/${request.appPlacementIntentName}`
    )
    .then((res) => res.data);
};

//network intents
const getNetworkControllerIntents = (request) => {
  return  instance
    .get(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent`
    )
    .then((res) => {
      return res.data;
    });
};
const addNetworkControllerIntent = (request) => {
  return  instance
    .post(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const deleteNetworkControllerIntent = (request) => {
  return  instance
    .delete(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}`
    )
    .then((res) => {
      return res.data;
    });
};
const getWorkloadIntents = (request) => {
  return  instance
    .get(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents`
    )
    .then((res) => {
      return res.data;
    });
};
const addWorkloadIntent = (request) => {
  return  instance
    .post(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const deleteWorkloadIntent = (request) => {
  return  instance
    .delete(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents/${request.workloadIntentName}`
    )
    .then((res) => {
      return res.data;
    });
};
const getInterfaces = (request) => {
  return  instance
    .get(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents/${request.workloadIntentName}/interfaces`
    )
    .then((res) => {
      return res.data;
    });
};
const addInterface = (request) => {
  return  instance
    .post(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents/${request.workloadIntentName}/interfaces`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const deleteInterface = (request) => {
  return  instance
    .delete(
      `/v2/ovnaction/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/network-controller-intent/${request.networkControllerIntentName}/workload-intents/${request.workloadIntentName}/interfaces/${request.interfaceName}`
    )
    .then((res) => {
      return res.data;
    });
};

//deployment intent group
const addIntentsToDeploymentIntentGroup = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/intents`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const editDeploymentIntentGroup = (request) => {
  return  instance
    .put(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const getDeploymentIntentGroupIntents = (request) => {
  return  instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/intents`
    )
    .then((res) => {
      return res.data;
    });
};
const approveDeploymentIntentGroup = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/approve`
    )
    .then((res) => {
      return res.data;
    });
};
const instantiate = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/instantiate`
    )
    .then((res) => {
      return res.data;
    });
};

const terminateDeploymentIntentGroup = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/terminate`
    )
    .then((res) => res.data);
};
//cluster-clm
const getClusterProviders = () => {
  return  instance.get("/v2/cluster-providers").then((res) => {
    return res.data;
  });
};
const getClusterProvider = (providerName) => {
  return  instance.get(`/v2/cluster-providers/${providerName}`).then((res) => {
    return res.data;
  });
};
const getClusters = (providerName) => {
  return  instance
    .get(`/v2/cluster-providers/${providerName}/clusters`)
    .then((res) => {
      return res.data;
    });
};

const getClustersByLabel = (providerName, label) => {
  return  instance
    .get(`/v2/cluster-providers/${providerName}/clusters`, {
      params: { label: `${label}` },
    })
    .then((res) => {
      return res.data;
    });
};
//updated cluster provider api with middle end api which takes care of gitops
// const registerClusterProvider = (request) => {
//     return  instance.post(`/v2/cluster-providers`, {...request}).then((res) => {
//         return res.data;
//     });
// };

//updated cluster provider api with middle end api which takes care of gitops
// const deleteClusterProvider = (providerName) => {
//   return  instance.delete(`/v2/cluster-providers/${providerName}`).then((res) => {
//     return res.data;
//   });
// };
const updateClusterProvider = (request) => {
  return  instance
    .put(`/v2/cluster-providers/${request.providerName}`, request.payload)
    .then((res) => {
      return res.data;
    });
};
const updateCluster = (request) => {
  return  instance
    .put(
      `/v2/cluster-providers/${request.get("providerName")}/clusters/${
        JSON.parse(request.get("metadata")).metadata.name
      }`,
      request
    )
    .then((res) => {
      return res.data;
    });
};
const addClusterLabel = (request) => {
  return  instance
    .post(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/labels`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const getClusterLabels = (request) => {
  return  instance
    .get(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/labels`
    )
    .then((res) => {
      return res.data;
    });
};

const deleteClusterLabel = (request) => {
  return  instance
    .delete(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/labels/${request.clusterLabel}`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const deleteCluster = (request) => {
  return  instance
    .delete(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}`
    )
    .then((res) => {
      return res.data;
    });
};

const getAllKvPairs = (request) => {
  return  instance
    .get(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/kv-pairs`
    )
    .then((res) => res.data);
};

const createKvPair = (request) => {
  return  instance
    .post(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/kv-pairs`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

const updateKvPair = (request) => {
  return  instance
    .put(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/kv-pairs/${request.kvPairName}`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

const deleteKvPair = (request) => {
  return  instance
    .delete(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/kv-pairs/${request.kvPairName}`
    )
    .then((res) => {
      return res.data;
    });
};

//network-ncm
const getClusterProviderNetworks = (request) => {
  return  instance
    .get(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/provider-networks`
    )
    .then((res) => {
      return res.data;
    });
};

const getClusterNetworks = (request) => {
  return  instance
    .get(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/networks`
    )
    .then((res) => {
      return res.data;
    });
};

const addNetwork = (request) => {
  return  instance
    .post(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/${request.networkType}`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

const deleteClusterNetwork = (request) => {
  return  instance
    .delete(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/${request.networkType}/${request.networkName}`
    )
    .then((res) => {
      return res.data;
    });
};
const applyNetworkConfig = (request) => {
  return  instance
    .post(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/apply`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};
const terminateNetworkConfig = (request) => {
  return  instance
    .post(
      `/v2/ncm/${request.providerName}/clusters/${request.clusterName}/terminate`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

//controller
const getControllers = () => {
  return  instance.get(`/v2/controllers`).then((res) => {
    return res.data;
  });
};

const addController = (request) => {
  return  instance.post(`/v2/controllers`, request).then((res) => {
    return res.data;
  });
};

const addITSController = (request) => {
  return  instance.post(`/v2/dtc-controllers`, request).then((res) => {
    return res.data;
  });
};
const updateController = (request) => {
  return  instance.put(`/v2/controllers`, request).then((res) => {
    return res.data;
  });
};

const removeITSController = () => {
  return  instance.delete(`/v2/dtc-controllers/its`).then((res) => {
    return res.data;
  });
};

const removeController = (controllerName) => {
  return  instance.delete(`/v2/controllers/${controllerName}`).then((res) => {
    return res.data;
  });
};

//config servce
const getBlueprintConfig = () => {
  return  instance.get(`/configsvc/getWorkflows`).then((res) => {
    return res.data;
  });
};

const getAppBlueprintConfig = (request) => {
  let req = `/configsvc/${request.compositeAppName}/${request.compositeAppVersion}/${request.appName}/bp`;
  if (request.type) req = req + `?type=${request.type}`;
  return  instance.get(req).then((res) => {
    return res.data;
  });
};

const executeWorkflow = ({ type, ...request }) => {
  let req = {
    method: "get",
    url: `/configsvc/${request.blueprintName}/${request.blueprintVersion}/${request.actionName}`,
  };
  if (request.payload) {
    req.data = request.payload;
    req.method = "post";
  }
  if (type && type === "DELETE") req.method = "delete";
  return  instance(req).then((res) => {
    return res.data;
  });
};

//DCM
const getLogicalCloudClusterReferences = (request) => {
  return  instance
    .get(
      `/v2/dcm/projects/${request.projectName}/logical-clouds/${request.logicalCloudName}/cluster-references`
    )
    .then((res) => res.data);
};

const terminateLogicalCloud = (request) => {
  return  instance
    .post(
      `/v2/dcm/projects/${request.projectName}/logical-clouds/${request.logicalCloudName}/terminate`
    )
    .then((res) => res.data);
};

const deleteLogicalCloudClusterReference = (request) => {
  return  instance
    .delete(
      `/v2/dcm/projects/${request.projectName}/logical-clouds/${request.logicalCloudName}/cluster-references/${request.clusterReferenceName}`
    )
    .then((res) => res.data);
};

const deleteLogicalCloud = (request) => {
  let deleteUrl = `/v2/dcm/projects/${request.projectName}/logical-clouds/${request.logicalCloudName}`;
  return  instance.delete(deleteUrl);
};

//middleend
const addCluster = (request) => {
  return  instance
    .post(
      `/v2/cluster-providers/${request.get("providerName")}/clusters`,
      request
    )
    .then((res) => {
      return res.data;
    });
};

const addService = ({ projectName, ...request }) => {
  return  instance
    .post(`/v2/projects/${projectName}/composite-apps`, request.payload)
    .then((res) => {
      return res.data;
    });
};

const deleteCompositeApp = (request) => {
  return  instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}`
    )
    .then((res) => {
      return res.data;
    });
};

const checkoutService = (request) => {
  var req = `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/checkout`;
  return  instance.post(req).then((res) => {
    return res.data;
  });
};

const checkInService = (request) => {
  var req = `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/update`;
  return  instance.post(req).then((res) => {
    return res.data;
  });
};

const getAllCompositeApps = (request) => {
  return  instance
    .get(`/v2/projects/${request.projectName}/composite-apps`, {
      params: { filter: "depthAll" },
    })
    .then((res) => {
      return res.data;
    });
};

const getCreatedCompositeApps = (request) => {
  return  instance
    .get(`/v2/projects/${request.projectName}/composite-apps`, {
      params: { filter: "depthAll", status: "created" },
    })
    .then((res) => {
      return res.data;
    });
};

const getCompositeAppDetails = (request) => {
  return  instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}`,
      { params: { filter: "depthAll" } }
    )
    .then((res) => {
      return res.data;
    });
};

const getCompositeAppVersions = (request) => {
  return  instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/versions`,
      { params: request.state ? { state: "created" } : {} }
    )
    .then((res) => res.data);
};

const removeAppFromService = (request) => {
  return  instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/apps/${request.appName}`
    )
    .then((res) => {
      return res.data;
    });
};

const updateService = (request) => {
  let params;
  if (request.get("operation") && request.get("operation") === "updateApp") {
    params = { operation: "updateApp" };
  }
  return  instance
    .post(
      `/v2/projects/${request.get(
        "projectName"
      )}/composite-apps/${request.get("compositeAppName")}/${request.get(
        "compositeAppVersion"
      )}/app`,
      request,
      { params: params }
    )
    .then((res) => {
      return res.data;
    });
};

const createDeploymentIntentGroup = ({
  projectName,
  compositeAppName,
  compositeAppVersion,
  payload,
}) => {
  return  instance
    .post(
      `/v2/projects/${projectName}/composite-apps/${compositeAppName}/${compositeAppVersion}/deployment-intent-groups`,
      payload
    )
    .then((res) => {
      return res.data;
    });
};
const getDeploymentIntentGroups = (request) => {
  return  instance
    .get(`/v2/projects/${request.projectName}/deployment-intent-groups`)
    .then((res) => {
      return res.data;
    });
};
const deleteDeploymentIntentGroup = (request) => {
  return  instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}`,
      { params: { operation: "deleteAll" } }
    )
    .then((res) => {
      return res.data;
    });
};

const getDeploymentIntentGroupStatus = (request) => {
  return  instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/status`
    )
    .then((res) => {
      return res.data;
    });
};

const getDashboardData = (projectName) => {
  return  instance
    .get(`/v2/projects/${projectName}/dashboard`)
    .then((res) => res.data);
};

const createLogicalCloud = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/logical-clouds`,
      request.payload
    )
    .then((res) => {
      return res.data;
    });
};

const getAllClusters = () => {
  return  instance.get(`/v2/all-clusters`).then((res) => {
    return res.data;
  });
};

const getAllClusterNetworks = (request) => {
  return  instance
    .get(
      `/v2/cluster-providers/${request.providerName}/clusters/${request.clusterName}/networks`
    )
    .then((res) => res.data);
};

const getCheckoutServiceInstance = (request) => {
  return  instance
    .get(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/checkout
    `
    )
    .then((res) => res.data);
};

const checkoutServiceInstance = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/checkout`,
      null,
      {
        params: { operation: "update" },
      }
    )
    .then((res) => res.data);
};

const saveCheckoutServiceInstance = (request) => {
  return  instance
    .put(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/checkout`,
      request.payload,
      {
        params: { operation: "save" },
      }
    )
    .then((res) => res.data);
};

const deleteCheckoutServiceInstance = (request) => {
  return  instance
    .delete(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}`,
      {
        params: { operation: "local" },
      }
    )
    .then((res) => res);
};

const submitCheckoutServiceInstance = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/checkout/submit`
    )
    .then((res) => res.data);
};

const migrateServiceInstance = (request) => {
  return  instance
    .post(
      `/v2/projects/${request.projectName}/composite-apps/${request.compositeAppName}/${request.compositeAppVersion}/deployment-intent-groups/${request.deploymentIntentGroupName}/checkout`,
      null,
      {
        params: { operation: "migrate", targetVersion: request.targetVersion },
      }
    )
    .then((res) => res.data);
};
const terminateAndDeleteLogicalCloud = (request) => {
  let deleteUrl = `/v2/projects/${request.projectName}/logical-clouds/${request.logicalCloudName}`;
  return  instance.delete(deleteUrl);
};

const registerClusterProvider = (request) => {
  return  instance
    .post(`/v2/cluster-providers`, { ...request })
    .then((res) => {
      return res.data;
    });
};

const deleteClusterProvider = (providerName) => {
  return  instance.delete(`/v2/cluster-providers/${providerName}`).then((res) => {
    return res.data;
  });
}
const getLogicalClouds = (projectName) => {
    return  instance
        .get(`/v2/projects/${projectName}/logical-clouds`)
        .then((res) => res.data);
};

const getUserDetails = () => {
  return  instance.get("/api/user/me").then((res) => res.data);
};

const getAllUsers = () => {
  return  instance.get("/api/users").then((res) => res.data);
};

const addUser = (request) => {
  return  instance.post("/api/user/add", request).then((res) => res.data);
};

const deleteUser = (userId) => {
  return  instance.delete(`/api/user/${userId}`).then((res) => res.data);
};

const updateUserPassword = ({ userId, ...request }) => {
  return  instance
    .put(`/api/user/${userId}/account/password`, request)
    .then((res) => res.data);
};

const updateUserDetails = ({ userId, ...request }) => {
  return  instance.put(`/api/user/${userId}`, request).then((res) => res.data);
};

const vimService = {
  createProject,
  updateProject,
  deleteProject,
  getProjectDetails,
  getAllProjects,
  getClusterProviders,
  getClusterProvider,
  getClusters,
  getClustersByLabel,
  registerClusterProvider,
  getClusterNetworks,
  getClusterProviderNetworks,
  addCluster,
  updateCluster,
  getClusterLabels,
  addNetwork,
  deleteClusterNetwork,
  applyNetworkConfig,
  getCompositeApps,
  getProfiles,
  createCompositeApp,
  addService,
  updateCompositeApp,
  deleteCompositeApp,
  getApps,
  addApp,
  updateApp,
  deleteApp,
  createCompositeProfile,
  getCompositeProfiles,
  deleteCompositeProfile,
  addProfile,
  deleteProfile,
  getGenericPlacementIntents,
  createGenericPlacementIntent,
  deleteGenericPlacementIntent,
  getAppPlacementIntents,
  addAppPlacementIntent,
  deleteAppPlacementIntent,
  getNetworkControllerIntents,
  addNetworkControllerIntent,
  deleteNetworkControllerIntent,
  getWorkloadIntents,
  addWorkloadIntent,
  deleteWorkloadIntent,
  getInterfaces,
  addInterface,
  deleteInterface,
  createDeploymentIntentGroup,
  addIntentsToDeploymentIntentGroup,
  getDeploymentIntentGroups,
  editDeploymentIntentGroup,
  deleteDeploymentIntentGroup,
  getDeploymentIntentGroupIntents,
  deleteClusterProvider,
  updateClusterProvider,
  deleteCluster,
  deleteClusterLabel,
  addClusterLabel,
  getAllKvPairs,
  createKvPair,
  updateKvPair,
  deleteKvPair,
  approveDeploymentIntentGroup,
  instantiate,
  terminateDeploymentIntentGroup,
  getDeploymentIntentGroupStatus,
  getControllers,
  addController,
  addITSController,
  updateController,
  removeController,
  removeITSController,
  getBlueprintConfig,
  getAppBlueprintConfig,
  executeWorkflow,
  checkoutService,
  getAllCompositeApps,
  getCreatedCompositeApps,
  getCompositeAppDetails,
  getCompositeAppVersions,
  removeAppFromService,
  updateService,
  checkInService,
  getDashboardData,
  getLogicalClouds,
  createLogicalCloud,
  getLogicalCloudClusterReferences,
  terminateLogicalCloud,
  deleteLogicalCloudClusterReference,
  deleteLogicalCloud,
  getAllClusters,
  getAllClusterNetworks,
  terminateNetworkConfig,
  checkoutServiceInstance,
  getCheckoutServiceInstance,
  saveCheckoutServiceInstance,
  deleteCheckoutServiceInstance,
  submitCheckoutServiceInstance,
  migrateServiceInstance,
  terminateAndDeleteLogicalCloud,
  getUserDetails,
  getAllUsers,
  addUser,
  deleteUser,
  updateUserPassword,
  updateUserDetails,
};
export default vimService;
