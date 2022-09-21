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
import React, { useEffect, useState, useContext } from "react";
import PropTypes from "prop-types";
import MuiDialogActions from "@material-ui/core/DialogActions";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import {
  Backdrop,
  Button,
  Checkbox,
  Chip,
  CircularProgress,
  Dialog,
  FormControl,
  FormHelperText,
  IconButton,
  Input,
  InputLabel,
  ListItemText,
  makeStyles,
  MenuItem,
  Select,
  TextField,
  Typography,
  withStyles,
} from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Close";
import apiService from "../../services/apiService";
import * as Yup from "yup";
import { Formik, getIn } from "formik";
import LoadingButton from "../../common/LoadingButton";
import FormLabel from "@material-ui/core/FormLabel";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Radio from "@material-ui/core/Radio";
import { UserContext } from "../../UserContext";

const useStyles = makeStyles((theme) => ({
  chips: {
    display: "flex",
    flexWrap: "wrap",
  },
  chip: {
    margin: 2,
  },
  backdrop: {
    zIndex: theme.zIndex.drawer + 9999,
    color: "#fff",
  },
  formControl: {
    marginTop: "20px",
  },
}));

const styles = (theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(2),
  },
  closeButton: {
    position: "absolute",
    right: theme.spacing(1),
    top: theme.spacing(1),
    color: theme.palette.grey[500],
  },
});

const DialogTitle = withStyles(styles)((props) => {
  const { children, classes, onClose, ...other } = props;
  return (
    <MuiDialogTitle disableTypography className={classes.root} {...other}>
      <Typography variant="h6">{children}</Typography>
      {onClose ? (
        <IconButton className={classes.closeButton} onClick={onClose}>
          <CloseIcon />
        </IconButton>
      ) : null}
    </MuiDialogTitle>
  );
});

const DialogActions = withStyles((theme) => ({
  root: {
    margin: 0,
    padding: theme.spacing(1),
  },
}))(MuiDialogActions);

const DialogContent = withStyles((theme) => ({
  root: {
    padding: theme.spacing(2),
  },
}))(MuiDialogContent);

const getSchema = (existingLogicalClouds) => {
  let schema;
  schema = Yup.object({
    name: Yup.string()
      .required("Name is required")
      .max(50, "Name cannot exceed more than 50 characters")
      .matches(
        /^[a-zA-Z0-9_-]+$/,
        "Name can only contain letters, numbers, '-' and '_' and no spaces."
      )
      .matches(/^[a-zA-Z0-9]/, "Name must start with an alphanumeric character")
      .matches(/[a-zA-Z0-9]$/, "Name must end with an alphanumeric character")
      .test(
        "duplicate-test",
        "Logical Cloud with same name exists, please use a different name",
        (name) => {
          return existingLogicalClouds
            ? existingLogicalClouds.findIndex(
              (x) => x.metadata.name === name
            ) === -1
            : true;
        }
      ),
    description: Yup.string(),
    cloudType: Yup.string().required("This field is required"),
    enableServiceDiscovery: Yup.boolean().required("This is required"),
    spec: Yup.object()
      .when("cloudType", {
        is: (value) => value === "user" || value === 'privileged',
        then: Yup.object({
          namespace: Yup.string()
            .max(63, "Namespace cannot exceed more than 63 characters")
            .matches(
              /^[a-z0-9-]+$/,
              "Namespace can only contain lowercase alphanumeric characters or '-' and no spaces."
            )
            .matches(
              /^[a-zA-Z0-9]/,
              "Namespace must start with an alphanumeric character"
            )
            .matches(
              /[a-zA-Z0-9]$/,
              "Namespace must end with an alphanumeric character"
            )
            .required("Namespace is required"),
          clusterproviders: Yup.array()
            .of(Yup.object({}))
            .required("At least one cluster is required"),
          permissions: Yup.object({
            apiGroups: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Api Groups values, expected array of string"),
            resources: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Resources values, expected array of string"),
            verbs: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Verbs values, expected array of string"),
          }),
          quotas: Yup.object().typeError("Invalid Quotas values, expected JSON")
        }),
        otherwise: Yup.object({
          clusterproviders: Yup.array()
            .of(Yup.object({}))
            .required("At least one cluster is required"),
          permissions: Yup.object({
            apiGroups: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Api Groups values, expected array of string"),
            resources: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Resources values, expected array of string"),
            verbs: Yup.array()
              .of(Yup.string())
              .typeError("Invalid Verbs values, expected array of string"),
          }),
          quotas: Yup.object().typeError("Invalid Quotas values, expected JSON")
        })
      })
  });
  return schema;
};

const LogicalCloudForm = (props) => {
  const classes = useStyles();
  const { onClose, item, open, onSubmit } = props;
  const [selectedClusters, setSelectedClusters] = React.useState([]);
  const [isLoading, setIsloading] = useState(true);
  const [clusterProviders, setClusterProviders] = useState([]);
  const buttonLabel = item ? "OK" : "Create";
  const title = item ? "Edit Logical Cloud" : "Create Logical Cloud";
  const { user } = useContext(UserContext);
  const handleClose = () => {
    onClose();
  };
  const isRbacEnabled = window._env_ && window._env_.ENABLE_RBAC === "true";
  useEffect(() => {
    apiService
      .getAllClusters()
      .then((res) => {
        //filter out the providers with no clusters
        let clusterProviders = res.filter(
          (cp) => cp.spec.clusters && cp.spec.clusters.length > 0
        );
        setClusterProviders(clusterProviders);
      })
      .catch((err) => {
        console.log("error getting all clusters : " + err);
      })
      .finally(() => {
        setIsloading(false);
      });
  }, []);

  let initialValues = item
    ? {
      name: item.metadata.name,
      description: item.metadata.description,
      spec: {
        clusters: [],
      },
    }
    : {
      name: "",
      description: "",
      cloudType: "admin",
      spec: {
        namespace: "",
        user: { userName: user ? user.email : "" },
        clusterproviders: [],
        permissions: {
          apiGroups: undefined,
          resources: undefined,
          verbs: undefined,
        },
        quotas: undefined,
      },
      enableServiceDiscovery: false,
    };

  const selectCluster = (provider, cluster, setFieldValue) => {
    const existingList = selectedClusters.filter(
      (item) => item.metadata.name !== provider
    );
    const p = selectedClusters.filter(
      (item) => item.metadata.name === provider
    );

    //check if current selected cluster is from a provider in selectedClusters list, if not go to else
    if (p.length > 0) {
      const c = p[0].spec.clusters.filter(
        (item) => item.metadata.name === cluster.metadata.name
      );
      //if the cluster is already selected then remove it from selected, otherwise add to selected
      if (c.length > 0) {
        p[0].spec.clusters = p[0].spec.clusters.filter(
          (item) => item.metadata.name !== cluster.metadata.name
        );
      } else {
        p[0].spec.clusters.push(cluster);
      }
      //update the selected clusters with new entry and existing entries
      if (p[0].spec.clusters.length < 1) {
        setSelectedClusters([...existingList]);
        setFieldValue("spec.clusterproviders", [...existingList]);
      } else {
        setSelectedClusters([...existingList, p[0]]);
        setFieldValue("spec.clusterproviders", [...existingList, p[0]]);
      }
    } else {
      let newEntry = {
        metadata: {
          name: provider,
        },
        spec: {
          clusters: [cluster],
        },
      };
      if (selectedClusters.length > 0) {
        setSelectedClusters((selectedClusters) => [
          ...selectedClusters,
          newEntry,
        ]);
        setFieldValue("spec.clusterproviders", [...selectedClusters, newEntry]);
      } else {
        setSelectedClusters([newEntry]);
        setFieldValue("spec.clusterproviders", [newEntry]);
      }
    }
  };

  //function to handle bulk select/ unselect
  const selectClusters = (provider, setFieldValue) => {
    let newProvider = {
      metadata: provider.metadata,
      spec: {
        clusters: [...provider.spec.clusters],
      },
    };
    let providerIndex = selectedClusters.findIndex(
      (x) => x.metadata.name === provider.metadata.name
    );
    //if selected provider is not in selectedClusters list, then add that provider and select all it's clusters, else remove that entry from selectedClusters list
    if (providerIndex === -1) {
      setSelectedClusters([...selectedClusters, newProvider]);
      setFieldValue("spec.clusterproviders", [
        ...selectedClusters,
        newProvider,
      ]);
    } else {
      setSelectedClusters((selectedClusters) => {
        return selectedClusters.filter(
          (entry) => entry.metadata.name !== provider.metadata.name
        );
      });
      setFieldValue(
        "spec.clusterproviders",
        selectedClusters.filter(
          (entry) => entry.metadata.name !== provider.metadata.name
        )
      );
    }
  };

  const getIsChecked = (provider, cluster) => {
    let providerIndex = selectedClusters.findIndex(
      (x) => x.metadata.name === provider
    );
    if (providerIndex !== -1) {
      let clusterIndex = selectedClusters[
        providerIndex
      ].spec.clusters.findIndex(
        (y) => y.metadata.name === cluster.metadata.name
      );
      return clusterIndex !== -1;
    }
    return false;
  };

  const getIsIndeterminate = (provider, totalClusters) => {
    let providerIndex = selectedClusters.findIndex(
      (x) => x.metadata.name === provider
    );
    if (providerIndex !== -1) {
      return (
        selectedClusters[providerIndex].spec.clusters.length !== totalClusters
      );
    }
    return false;
  };

  const getIsAllChecked = (provider, totalClusters) => {
    let providerIndex = selectedClusters.findIndex(
      (x) => x.metadata.name === provider
    );
    if (providerIndex !== -1) {
      return (
        selectedClusters[providerIndex].spec.clusters.length === totalClusters
      );
    }
    return false;
  };

  return (
    <>
      <Backdrop className={classes.backdrop} open={isLoading}>
        <CircularProgress color="primary" />
      </Backdrop>
      <Dialog
        maxWidth={"xs"}
        onClose={handleClose}
        aria-labelledby="customized-dialog-title"
        open={open}
        disableBackdropClick
      >
        <DialogTitle id="simple-dialog-title">{title}</DialogTitle>
        <Formik
          initialValues={initialValues}
          onSubmit={(values) => {
            onSubmit(values);
          }}
          validationSchema={getSchema(props.existingLogicalClouds)}
        >
          {(props) => {
            const {
              values,
              touched,
              errors,
              isSubmitting,
              handleChange,
              handleBlur,
              handleSubmit,
              setFieldValue,
            } = props;
            return (
              <form noValidate onSubmit={handleSubmit}>
                <DialogContent dividers>
                  <TextField
                    fullWidth
                    id="name"
                    label="Logical Cloud name"
                    type="text"
                    value={values.name}
                    onChange={handleChange}
                    onBlur={handleBlur}
                    helperText={touched.name && errors.name}
                    required
                    disabled={item}
                    error={errors.name && touched.name}
                  />
                  <TextField
                    fullWidth
                    name="description"
                    value={values.description}
                    onChange={handleChange}
                    onBlur={handleBlur}
                    id="description"
                    label="Description"
                    multiline
                    rowsMax={4}
                  />
                  {isRbacEnabled && (
                    <FormControl
                      component="fieldset"
                      style={{ marginTop: "30px" }}
                    >
                      <FormLabel component="legend">Cloud Type</FormLabel>
                      <RadioGroup
                        row
                        aria-label="cloudType"
                        name="cloudType"
                        value={values.cloudType}
                        onChange={handleChange}
                        onBlur={handleBlur}
                      >
                        <FormControlLabel
                          value="admin"
                          control={<Radio />}
                          label="Admin"
                        />
                        <FormControlLabel
                          value="user"
                          control={<Radio />}
                          label="User"
                        />
                        <FormControlLabel
                          value="privileged"
                          control={<Radio />}
                          label="Privileged"
                        />
                      </RadioGroup>
                    </FormControl>
                  )}

                  {(values.cloudType === "user" || values.cloudType === "privileged") && (
                    <TextField
                      fullWidth
                      name="spec.namespace"
                      value={values.spec.namespace}
                      onChange={handleChange}
                      onBlur={handleBlur}
                      id="namespace"
                      label="Namespace"
                      helperText={
                        getIn(touched, "spec.namespace") &&
                        getIn(errors, "spec.namespace")
                      }
                      error={Boolean(
                        getIn(touched, "spec.namespace") &&
                        getIn(errors, "spec.namespace")
                      )}
                      required
                    />
                  )}
                  <FormControl
                    fullWidth
                    className={classes.formControl}
                    required
                    error={Boolean(
                      getIn(touched, "spec.clusterproviders") &&
                      getIn(errors, "spec.clusterproviders")
                    )}
                  >
                    <InputLabel id="demo-mutiple-chip-label">
                      Select Clusters
                    </InputLabel>
                    <Select
                      labelId="demo-mutiple-chip-label"
                      id="demo-mutiple-chip"
                      multiple
                      value={selectedClusters}
                      input={<Input id="select-multiple-chip" />}
                      renderValue={(selected) => (
                        <div className={classes.chips}>
                          {selected.map((provider) =>
                            provider.spec.clusters.map((cluster) => (
                              <Chip
                                color="primary"
                                key={cluster.metadata.name}
                                label={cluster.metadata.name}
                                className={classes.chip}
                              />
                            ))
                          )}
                        </div>
                      )}
                    >
                      <div
                        style={{
                          padding: "0 20px",
                          maxHeight: "200px",
                          overflow: "auto",
                        }}
                      >
                        {clusterProviders.map((provider) => {
                          return (
                            <React.Fragment key={provider.metadata.name}>
                              <Typography
                                variant="body1"
                                style={{ display: "inline-flex" }}
                                key={provider.metadata.name}
                              >
                                {provider.metadata.name}
                              </Typography>
                              <Checkbox
                                checked={getIsAllChecked(
                                  provider.metadata.name,
                                  provider.spec.clusters.length
                                )}
                                indeterminate={getIsIndeterminate(
                                  provider.metadata.name,
                                  provider.spec.clusters.length
                                )}
                                onClick={(e) => {
                                  selectClusters(
                                    { ...provider },
                                    setFieldValue
                                  );
                                  e.stopPropagation();
                                }}
                              />
                              {provider.spec.clusters.map((cluster) => (
                                <MenuItem
                                  key={cluster.metadata.name}
                                  value={cluster.metadata.name}
                                  onClick={(e) => {
                                    selectCluster(
                                      provider.metadata.name,
                                      cluster,
                                      setFieldValue,
                                      values
                                    );
                                    e.stopPropagation();
                                  }}
                                >
                                  <Checkbox
                                    checked={getIsChecked(
                                      provider.metadata.name,
                                      cluster
                                    )}
                                  />
                                  <ListItemText
                                    primary={cluster.metadata.name}
                                  />
                                </MenuItem>
                              ))}
                            </React.Fragment>
                          );
                        })}
                      </div>
                    </Select>
                    {values.cloudType === "user" && (
                      <>
                        <fieldset
                          style={{
                            marginBottom: "20px",
                            marginTop: "20px",
                            border: "1px solid rgba(0, 0, 0, 0.42)",
                            borderRadius: "5px",
                          }}
                        >
                          <legend>Permissions</legend>
                          <TextField
                            fullWidth
                            style={{ marginBottom: "20px", marginTop: "20px" }}
                            id="apiGroups"
                            label="Api Groups"
                            type="text"
                            name="spec.permissions.apiGroups"
                            value={values.spec.permissions.apiGroups}
                            onChange={handleChange}
                            onBlur={handleBlur}
                            multiline
                            rows={1}
                            variant="outlined"
                            helperText={
                              getIn(touched, "spec.permissions.apiGroups") &&
                              getIn(errors, "spec.permissions.apiGroups")
                            }
                            error={Boolean(
                              getIn(touched, "spec.permissions.apiGroups") &&
                              getIn(errors, "spec.permissions.apiGroups")
                            )}
                          />

                          <TextField
                            fullWidth
                            style={{ marginBottom: "20px" }}
                            id="resources"
                            label="Resources"
                            type="text"
                            name="spec.permissions.resources"
                            value={values.spec.permissions.resources}
                            onChange={handleChange}
                            onBlur={handleBlur}
                            multiline
                            rows={1}
                            variant="outlined"
                            helperText={
                              getIn(touched, "spec.permissions.resources") &&
                              getIn(errors, "spec.permissions.resources")
                            }
                            error={Boolean(
                              getIn(touched, "spec.permissions.resources") &&
                              getIn(errors, "spec.permissions.resources")
                            )}
                          />
                          <TextField
                            fullWidth
                            style={{ marginBottom: "20px" }}
                            id="verbs"
                            label="Verbs"
                            type="text"
                            name="spec.permissions.verbs"
                            value={values.spec.permissions.verbs}
                            onChange={handleChange}
                            onBlur={handleBlur}
                            multiline
                            rows={1}
                            variant="outlined"
                            helperText={
                              getIn(touched, "spec.permissions.verbs") &&
                              getIn(errors, "spec.permissions.verbs")
                            }
                            error={Boolean(
                              getIn(touched, "spec.permissions.verbs") &&
                              getIn(errors, "spec.permissions.verbs")
                            )}
                          />
                        </fieldset>
                      </>
                    )}
                    {values.cloudType === "user" && (
                      <TextField
                        fullWidth
                        style={{ marginBottom: "20px" }}
                        id="quotas"
                        label="Quotas"
                        type="text"
                        name="spec.quotas"
                        value={values.spec.quotas}
                        onChange={handleChange}
                        onBlur={handleBlur}
                        multiline
                        rows={4}
                        variant="outlined"
                        helperText={
                          getIn(touched, "spec.quotas") &&
                          getIn(errors, "spec.quotas")
                        }
                        error={Boolean(
                          getIn(touched, "spec.quotas") &&
                          getIn(errors, "spec.quotas")
                        )}
                      />
                    )}

                    {errors.spec && touched.spec && (
                      <FormHelperText>{errors.spec.clusters}</FormHelperText>
                    )}
                  </FormControl>

                  {isRbacEnabled && (values.cloudType === "user" || values.cloudType === "privileged") && (
                    <div style={{ marginTop: "15px" }}>
                      <FormControlLabel
                        control={
                          <Checkbox
                            checked={values.enableServiceDiscovery}
                            onChange={handleChange}
                            name="enableServiceDiscovery"
                          />
                        }
                        label="Enable Service Discovery"
                      />
                    </div>
                  )}
                </DialogContent>
                <DialogActions>
                  <Button
                    autoFocus
                    onClick={(e) => {
                      setSelectedClusters([]);
                      handleClose(e);
                    }}
                    disabled={isSubmitting}
                    color="secondary"
                  >
                    Cancel
                  </Button>
                  <LoadingButton
                    type="submit"
                    buttonLabel={buttonLabel}
                    loading={isSubmitting}
                  />
                </DialogActions>
              </form>
            );
          }}
        </Formik>
      </Dialog>
    </>
  );
};

LogicalCloudForm.propTypes = {
  onClose: PropTypes.func.isRequired,
  open: PropTypes.bool.isRequired,
};

export default LogicalCloudForm;
