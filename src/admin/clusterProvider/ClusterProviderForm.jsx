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
import PropTypes from "prop-types";
import { withStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";

import Dialog from "@material-ui/core/Dialog";
import MuiDialogTitle from "@material-ui/core/DialogTitle";
import MuiDialogContent from "@material-ui/core/DialogContent";
import MuiDialogActions from "@material-ui/core/DialogActions";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Typography from "@material-ui/core/Typography";
import {
  FormControl,
  FormControlLabel,
  TextField,
  InputAdornment,
  FormHelperText,
  InputLabel,
  Input,
} from "@material-ui/core";
import * as Yup from "yup";
import { FieldArray, Formik, getIn } from "formik";
import Checkbox from "@material-ui/core/Checkbox";
import Visibility from "@material-ui/icons/Visibility";
import VisibilityOff from "@material-ui/icons/VisibilityOff";

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

const getSchema = (existingProviders) => {
  let schema;
  schema = Yup.object({
    metadata: Yup.object({
      name: Yup.string()
        .required("Name is required")
        .matches(
          /^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$/,
          "Name can only contain letters, numbers, '-', '_' and no spaces. Name must start and end with an alphanumeric character"
        )
        .test(
          "duplicate-test",
          "A cloud provider with same name exists, please use a different name",
          (name) => {
            return existingProviders
              ? existingProviders.findIndex((x) => x.metadata.name === name) ===
                  -1
              : true;
          }
        )
        .max(128, "Name cannot exceed 128 characters"),
      description: Yup.string(),
    }),
    spec: Yup.object({
      gitEnabled: Yup.boolean().required("This field is required"),
      kv: Yup.array()
        .of(
          Yup.object({
            gitType: Yup.string().required("Git type is required"),
            userName: Yup.string().required("Username is required"),
            gitToken: Yup.string().required("Git token is required"),
            repoName: Yup.string().required("Repo name is required"),
            branch: Yup.string().required("Branch is required"),
          })
        )
        .nullable(),
    }),
  });
  return schema;
};

const ClusterProviderForm = (props) => {
  const { onClose, item, open, onSubmit, existingProviders } = props;
  const buttonLabel = item ? "OK" : "Create";
  const title = item ? "Edit Cluster Provider" : "Register Cluster Provider";
  const handleClose = () => {
    onClose();
  };
  const [tokenShown, setTokenShown] = React.useState(false);

  let initialValues = item
    ? {
        metadata: {
          name: item.metadata.name,
          description: item.metadata.description,
        },
      }
    : {
        metadata: { name: "", description: "" },
        spec: {
          gitEnabled: false,
          kv: [],
        },
      };
  const toggleShowToken = () => {
    setTokenShown(!tokenShown);
  };
  return (
    <Dialog
      maxWidth={"xs"}
      fullWidth
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
        validationSchema={getSchema(existingProviders)}
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
          } = props;
          return (
            <form noValidate onSubmit={handleSubmit}>
              <DialogContent dividers>
                <TextField
                  style={{ width: "100%", marginBottom: "10px" }}
                  id="metadata.name"
                  label="Provider name"
                  type="text"
                  value={values.metadata.name}
                  onChange={handleChange}
                  onBlur={handleBlur}
                  required
                  helperText={
                    getIn(touched, "metadata.name") &&
                    getIn(errors, "metadata.name")
                  }
                  error={Boolean(
                    getIn(touched, "metadata.name") &&
                      getIn(errors, "metadata.name")
                  )}
                />
                <TextField
                  style={{ width: "100%", marginBottom: "25px" }}
                  name="metadata.description"
                  value={values.metadata.description}
                  onChange={handleChange}
                  onBlur={handleBlur}
                  id="description"
                  label="Description"
                  multiline
                  rowsMax={4}
                />

                <FieldArray
                  name="spec.kv"
                  render={(arrayHelpers) => (
                    <>
                      <FormControlLabel
                        control={
                          <Checkbox
                            checked={values.spec.gitEnabled}
                            onChange={(el) => {
                              if (el.target.checked) {
                                arrayHelpers.push({
                                  gitType: "",
                                  userName: "",
                                  gitToken: "",
                                  repoName: "",
                                  branch: "",
                                });
                              } else {
                                //TODO for now we are assuming there will be only one kv pair per cluster provider
                                arrayHelpers.remove(0);
                              }
                              handleChange(el);
                            }}
                            name="spec.gitEnabled"
                          />
                        }
                        label="Add Git Ops Support"
                      />
                      {values.spec.gitEnabled &&
                        values.spec.kv.map((kvPair, index) => (
                          <fieldset
                            key={index + "kv"}
                            style={{
                              marginBottom: "20px",
                              marginTop: "20px",
                              border: "1px solid rgba(0, 0, 0, 0.42)",
                              borderRadius: "5px",
                            }}
                          >
                            <legend>Git Ops Details</legend>
                            <>
                              <TextField
                                key={index + "gitType"}
                                style={{ width: "100%", marginBottom: "10px" }}
                                id={`spec.kv[${index}].gitType`}
                                label="Git Type"
                                type="text"
                                value={values.spec.kv[index].gitType}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                required
                                helperText={
                                  getIn(touched, `spec.kv[${index}].gitType`) &&
                                  getIn(errors, `spec.kv[${index}].gitType`)
                                }
                                error={Boolean(
                                  getIn(touched, `spec.kv[${index}].gitType`) &&
                                    getIn(errors, `spec.kv[${index}].gitType`)
                                )}
                              />
                              <TextField
                                key={index + "userName"}
                                style={{ width: "100%", marginBottom: "10px" }}
                                id={`spec.kv[${index}].userName`}
                                label="User Name"
                                type="text"
                                value={values.spec.kv[index].userName}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                required
                                helperText={
                                  getIn(
                                    touched,
                                    `spec.kv[${index}].userName`
                                  ) &&
                                  getIn(errors, `spec.kv[${index}].userName`)
                                }
                                error={Boolean(
                                  getIn(
                                    touched,
                                    `spec.kv[${index}].userName`
                                  ) &&
                                    getIn(errors, `spec.kv[${index}].userName`)
                                )}
                              />

                              <FormControl
                                required
                                fullWidth
                                error={Boolean(
                                  getIn(
                                    touched,
                                    `spec.kv[${index}].gitToken`
                                  ) &&
                                    getIn(errors, `spec.kv[${index}].gitToken`)
                                )}
                              >
                                <InputLabel htmlFor="standard-adornment-password">
                                  Git Token
                                </InputLabel>
                                <Input
                                  key={index + "gitToken"}
                                  id={`spec.kv[${index}].gitToken`}
                                  type={tokenShown ? "text" : "password"}
                                  value={values.spec.kv[index].gitToken}
                                  onChange={handleChange}
                                  onBlur={handleBlur}
                                  endAdornment={
                                    <InputAdornment position="end">
                                      <IconButton
                                        onClick={toggleShowToken}
                                        onMouseDown={(event) => {
                                          event.preventDefault();
                                        }}
                                      >
                                        {tokenShown ? (
                                          <Visibility />
                                        ) : (
                                          <VisibilityOff />
                                        )}
                                      </IconButton>
                                    </InputAdornment>
                                  }
                                />
                                <FormHelperText error>
                                  {getIn(
                                    touched,
                                    `spec.kv[${index}].gitToken`
                                  ) &&
                                    getIn(errors, `spec.kv[${index}].gitToken`)}
                                </FormHelperText>
                              </FormControl>

                              <TextField
                                key={index + "repoName"}
                                style={{ width: "100%", marginBottom: "10px" }}
                                id={`spec.kv[${index}].repoName`}
                                label="Repo Name"
                                type="text"
                                value={values.spec.kv[index].repoName}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                required
                                helperText={
                                  getIn(
                                    touched,
                                    `spec.kv[${index}].repoName`
                                  ) &&
                                  getIn(errors, `spec.kv[${index}].repoName`)
                                }
                                error={Boolean(
                                  getIn(
                                    touched,
                                    `spec.kv[${index}].repoName`
                                  ) &&
                                    getIn(errors, `spec.kv[${index}].repoName`)
                                )}
                              />
                              <TextField
                                key={index + "branch"}
                                style={{ width: "100%", marginBottom: "10px" }}
                                id={`spec.kv[${index}].branch`}
                                label="Branch"
                                type="text"
                                value={values.spec.kv[index].branch}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                required
                                helperText={
                                  getIn(touched, `spec.kv[${index}].branch`) &&
                                  getIn(errors, `spec.kv[${index}].branch`)
                                }
                                error={Boolean(
                                  getIn(touched, `spec.kv[${index}].branch`) &&
                                    getIn(errors, `spec.kv[${index}].branch`)
                                )}
                              />
                            </>
                          </fieldset>
                        ))}
                    </>
                  )}
                />
              </DialogContent>
              <DialogActions>
                <Button autoFocus onClick={handleClose} color="secondary">
                  Cancel
                </Button>
                <Button
                  autoFocus
                  type="submit"
                  color="primary"
                  disabled={isSubmitting}
                >
                  {buttonLabel}
                </Button>
              </DialogActions>
            </form>
          );
        }}
      </Formik>
    </Dialog>
  );
};

ClusterProviderForm.propTypes = {
  onClose: PropTypes.func.isRequired,
  open: PropTypes.bool.isRequired,
  item: PropTypes.object,
};

export default ClusterProviderForm;
