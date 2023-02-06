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

import React from "react";
import DashboardIcon from "@material-ui/icons/Dashboard";
import DeviceHubIcon from "@material-ui/icons/DeviceHub";
import DnsRoundedIcon from "@material-ui/icons/DnsRounded";
import PeopleIcon from "@material-ui/icons/People";
import SettingsIcon from "@material-ui/icons/SettingsRounded";
import LogicalCloudIcon from "@material-ui/icons/SettingsSystemDaydream";
import CloudSyncIcon from "@material-ui/icons/Input";
import BusinessIcon from "@material-ui/icons/Business";

const {ENABLE_RBAC} = window._env_ || {};
const SMO_ENABLED = window._env_ && window._env_.ENABLE_SMO === "true";

const adminMenu = [
    {
        id: "adminMenu",
        children: [
            {
                id: "Tenants",
                icon: <BusinessIcon/>,
                url: "/projects",
            },
            {
                id: "K8s Controllers",
                icon: <SettingsIcon/>,
                url: "/controllers",
            },
            {
                id: "Clusters",
                icon: <DnsRoundedIcon/>,
                url: "/clusters",
            },
        ],
    },
];

if (ENABLE_RBAC === 'true') {
    adminMenu[0].children.push({
        id: "Users",
        icon: <PeopleIcon/>,
        url: "/users",
    });
}

if (SMO_ENABLED) {
    adminMenu[0].children.push({
        id: "SMO",
        icon: <CloudSyncIcon/>,
        url: "/smo",
    });
}
const tenantMenu = [
    {
        id: "tenantMenu",
        children: [
            {
                id: "Dashboard",
                icon: <DashboardIcon/>,
                url: "/dashboard",
            },
            {
                id: "Services",
                icon: <DeviceHubIcon/>,
                url: "/services",
            },
            {
                id: "Logical Clouds",
                icon: <LogicalCloudIcon/>,
                url: "/logical-clouds",
            },
            {
                id: "Service Instances",
                icon: <DnsRoundedIcon/>,
                url: "/deployment-intent-groups",
            },
        ],
    },
];

const routes = [
    {path:'/app/admin/projects',name:'Tenants'},
    {path:'/app/admin/clusters',name:"Clusters"},
    {path:'/app/admin/controllers',name:"Controllers"},
    {path:'/app/admin/smo',name:"SMO"},
    {path:'/app/admin/users',name:"Users"},
    {path:'/app/projects/:projectName/Dashboard', name: "Dashboard"},
    {path:'/app/projects/:projectName/services', name: "Services"},
    {path:'/app/projects/:projectName/logical-clouds',name:'Logical Clouds'},
    {path:'/app/projects/:projectName/deployment-intent-groups',name:'Service Instances'},
    {path:'/app/projects/:projectName/deployment-intent-groups/:compositeAppName/:compositeAppVersion/:digName/status',name:'Service Instance Detail'},
    {path:'/app/projects/:projectName/deployment-intent-groups/:compositeAppName/:compositeAppVersion/:digName/checkout',name:'Service Instance Detail'},
    {path:'/app/projects/:projectName/services/:appname/:version',param:'appname'}
]

export {adminMenu};
export {tenantMenu};
export {SMO_ENABLED}
export {routes};
export default {adminMenu: adminMenu, tenantMenu: tenantMenu, SMO_ENABLED:SMO_ENABLED};