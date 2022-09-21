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
import Dialog from "@material-ui/core/Dialog";
import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import CloseIcon from "@material-ui/icons/Close";
import { Grid, Box, CssBaseline, Chip, Paper } from "@material-ui/core";
import Divider from "@material-ui/core/Divider";
import Card from "@material-ui/core/Card";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";

const useStyles = makeStyles((theme) => ({
  paper: {
    width: "100%",
    marginBottom: theme.spacing(2),
  },
  appBar: {
    position: "sticky",
    boxShadow: "none",
  },
  fieldName: {
    fontWeight: 500,
    flex: 1,
  },
  fieldValue: {
    flex: 1,
  },
  divider: {
    variant: "fullWidth",
    width: "90%",
    margin: 15,
  },
  card: {
    padding: "0 20px 20px 20px",
  },
}));

const LogicalCloudView = ({ open, handleClose, data, logicalCloud }) => {
  const classes = useStyles();
  return (
    <>
      {open && (
        <Dialog
          maxWidth={"md"}
          open={open}
          onClose={() => {
            handleClose(false);
          }}
        >
          <CssBaseline />
          <AppBar className={classes.appBar}>
            <Toolbar>
              <Typography variant="h6" className={classes.fieldName}>
                Logical Cloud Details:&nbsp;{logicalCloud.metadata.name}
              </Typography>
              <IconButton
                edge="end"
                color="inherit"
                onClick={() => {
                  handleClose(false);
                }}
                aria-label="close"
              >
                <CloseIcon />
              </IconButton>
            </Toolbar>
          </AppBar>
          <Box m={2}>
            <Card
              variant="outlined"
              style={{ marginBottom: 4 }}
              className={classes.card}
            >
              <Grid container spacing={2} style={{ marginTop: 2 }}>
                <Grid item xs={12}>
                  <Typography className={classes.fieldName} component="span">
                    Status:&nbsp;
                  </Typography>
                  <Typography className={classes.fieldValue} component="span">
                    {logicalCloud.spec.status}
                  </Typography>
                </Grid>

                <Grid item xs={6}>
                  <Typography component="span" className={classes.fieldName}>
                    Cloud Type:&nbsp;
                  </Typography>
                  <Typography component="span" className={classes.fieldValue}>
                    {logicalCloud.spec.level === "0" ? "Admin" : "User"}
                  </Typography>
                </Grid>
                <Grid item xs={6}>
                  <Typography component="span" className={classes.fieldName}>
                    Namespace:&nbsp;
                  </Typography>
                  <Typography component="span" className={classes.fieldValue}>
                    {logicalCloud.spec.namespace}
                  </Typography>
                </Grid>
              </Grid>
            </Card>

            <Card
              variant="outlined"
              style={{ marginBottom: 4 }}
              className={classes.card}
            >
              <Typography
                className={classes.fieldName}
                variant="h6"
                style={{ marginTop: "10px" }}
              >
                Cluster References
              </Typography>
              <List
                sx={{
                  width: "100%",
                  maxWidth: 360,
                  bgcolor: "background.paper",
                }}
              >
                {logicalCloud.spec.clusterReferences.spec.clusterProviders.map(
                  (clusterProvider, index, arr) => (
                    <React.Fragment key={index}>
                      <ListItem alignItems="flex-start">
                        <ListItemText
                          primary={`Cluster Provider : ${clusterProvider.metadata.name}`}
                          secondary={
                            <span>
                              Clusters:{" "}
                              {clusterProvider.spec.clusters
                                .map((x) => x.metadata.name)
                                .join(", ")}
                            </span>
                          }
                        />
                      </ListItem>
                      {index < arr.length - 1 ? (
                        <Divider component="li" />
                      ) : null}
                    </React.Fragment>
                  )
                )}
              </List>
            </Card>
            {logicalCloud.spec.level === "1" && (
              <>
                <Card
                  variant="outlined"
                  className={classes.card}
                  style={{ marginBottom: 4 }}
                >
                  <Grid container spacing={2} style={{ marginTop: 2 }}>
                    <Grid item xs={12}>
                      <Typography
                        className={classes.fieldName}
                        variant="h6"
                        style={{ marginTop: "10px" }}
                      >
                        Permissions
                      </Typography>
                    </Grid>

                    {logicalCloud.spec.userPermissions.map((userPermission) => (
                      <Grid item xs={12}>
                        <Paper
                          style={{ padding: "20px" }}
                          variant="outlined"
                          square
                        >
                          <Grid item xs={12} style={{ marginBottom: "10px" }}>
                            <Typography
                              gutterBottom
                              component="span"
                              className={classes.fieldName}
                            >
                              {userPermission.namespace ? "Namespace:" : "Cluster wide"}&nbsp;
                            </Typography>
                            <Typography
                              component="span"
                              className={classes.fieldValue}
                            >
                              {userPermission.namespace}
                            </Typography>
                          </Grid>
                          <Grid item xs={12}>
                            <Typography
                              display="block"
                              className={classes.fieldName}
                            >
                              API Groups
                            </Typography>
                            <Typography
                              component="span"
                              className={classes.fieldValue}
                            >
                              {userPermission.apiGroups.map(
                                (apiGroup, index) => (
                                  <Chip
                                    key={index}
                                    label={apiGroup}
                                    variant="outlined"
                                    color="primary"
                                    style={{
                                      marginRight: "4px",
                                      marginBottom: "2px",
                                    }}
                                  ></Chip>
                                )
                              )}
                            </Typography>
                          </Grid>
                          <Grid item xs={12}>
                            <Typography
                              display="block"
                              gutterBottom
                              className={classes.fieldName}
                            >
                              Resources
                            </Typography>
                            <Typography
                              component="span"
                              className={classes.fieldValue}
                            >
                              {userPermission.resources.map(
                                (resource, index) => (
                                  <Chip
                                    key={index}
                                    label={resource}
                                    variant="outlined"
                                    color="primary"
                                    style={{
                                      marginRight: "4px",
                                      marginBottom: "2px",
                                    }}
                                  ></Chip>
                                )
                              )}
                            </Typography>
                          </Grid>
                          <Grid item xs={12}>
                            <Typography
                              display="block"
                              gutterBottom
                              className={classes.fieldName}
                            >
                              Verbs
                            </Typography>
                            <Typography
                              component="span"
                              className={classes.fieldValue}
                            >
                              {userPermission.verbs.map((verb, index) => (
                                <Chip
                                  key={index}
                                  label={verb}
                                  variant="outlined"
                                  color="primary"
                                  style={{
                                    marginRight: "4px",
                                    marginBottom: "2px",
                                  }}
                                ></Chip>
                              ))}
                            </Typography>
                          </Grid>
                        </Paper>
                      </Grid>
                    ))}
                  </Grid>
                </Card>
                <Card variant="outlined" className={classes.card}>
                  <Grid container spacing={2} style={{ marginTop: 2 }}>
                    <Grid item xs={12}>
                      <Typography
                        className={classes.fieldName}
                        variant="h6"
                        style={{ marginTop: "10px" }}
                      >
                        Quotas
                      </Typography>
                    </Grid>
                    <Grid item container spacing={1}>
                      {Object.keys(logicalCloud.spec.userQuota || {}).map(
                        (key, index) => (
                          <Grid item xs={6} key={index}>
                            <Typography
                              className={classes.fieldName}
                              component="span"
                            >
                              {key}:
                            </Typography>
                            <Typography
                              className={classes.fieldValue}
                              component="span"
                            >
                              {logicalCloud.spec.userQuota[key]}
                            </Typography>
                          </Grid>
                        )
                      )}
                    </Grid>
                  </Grid>
                </Card>
              </>
            )}
          </Box>
        </Dialog>
      )}
    </>
  );
};

export default LogicalCloudView;