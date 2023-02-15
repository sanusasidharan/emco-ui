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
import React, { useState } from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Navigator from "../common/Navigator";
import Header from "../appbase/Header";
import Footer from "../appbase/Footer";
import theme from "../theme/Theme";
import Projects from "./projects/Projects";
import Users from "./users/Users";
import ClusterProviders from "./clusterProvider/ClusterProviders";
import Controllers from "./controllers/Controllers";
import {adminMenu, SMO_ENABLED} from "../config/uiConfig";
import { Switch, Route, Redirect } from "react-router-dom";
import Spinner from "../common/Spinner";
import PageNotFound from "../common/PageNotFound";
import {makeStyles} from "@material-ui/styles";
const useAppStyles = makeStyles({
  root: {
    display: "flex",
    minHeight: "100vh",
  },
  app: {
    flex: 1,
    display: "flex",
    flexDirection: "column",
  },
  main: {
    flex: 1,
    padding: theme.spacing(3, 4, 6, 4),
    background: "#eaeff1",
    width: "80%",
    float: "right",
    display: "inline-block",
    marginLeft: "20%"
  },
});

const Admin = (props) => {
  const classes = useAppStyles();
  const [mobileOpen, setMobileOpen] = useState(false);
  const handleDrawerToggle = () => {
    setMobileOpen(() => !mobileOpen);
  };
  return (
      <div className={classes.root}>
        <CssBaseline />
        <div className={classes.app}>
          <Header onDrawerToggle={handleDrawerToggle} onChangePasswordClick={props.handlePasswordFormOpen}/>
          <Navigator menu={adminMenu} handleDrawerToggle={handleDrawerToggle} mobileOpen={mobileOpen}/>
          <main
              className={classes.main}
              style={
                props.location.pathname === `${props.match.url}/smo`
                    ? { padding: 0 }
                    : null
              }
          >
            <Switch>
              <Route
                  path={`${props.match.url}/projects`}
                  component={Projects}
              />
              <Route
                  path={`${props.match.url}/clusters`}
                  component={ClusterProviders}
              />
              <Route
                  path={`${props.match.url}/controllers`}
                  component={Controllers}
              />
              {SMO_ENABLED && (
                  <Route
                      path={`${props.match.url}/smo`}
                      component={() => {
                        return <SmoIframe  />;
                      }}
                  />
              )}
              <Route
                  path={`${props.match.url}/users`}
                  component={Users}
              />
              <Route
                  path={`${props.match.url}/404`}
                  component={() => <PageNotFound />}
              />
              <Route
                  path="/"
                  component={() => (
                      <Redirect
                          exact
                          from={`${props.match.path}`}
                          to={`${props.match.path}/404`}
                      />
                  )}
              />
            </Switch>
          </main>
          <Footer/>
        </div>
      </div>
  );
}

const SmoIframe = () => {
  const [isLoading, setIsLoading] = useState(true);
  const smo_url = `http://${document.location.hostname}:30181/odlux/index.html#`;
  return (
      <>
        {isLoading && <Spinner />}
        <iframe
            src={smo_url}
            onLoad={() => {
              setIsLoading(false);
            }}
            style={{ height: "99%", width: "100%", borderWidth: 0 }}
        />
      </>
  );
};

export default Admin;
