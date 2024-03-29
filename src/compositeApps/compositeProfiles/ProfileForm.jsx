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
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  FormHelperText,
} from "@material-ui/core";
import * as Yup from "yup";
import { Formik } from "formik";
import FileUpload from "../../common/FileUpload";

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

const SUPPORTED_FORMATS = [
  ".tgz",
  ".tar.gz",
  ".tar",
  "application/x-tar",
  "application/x-tgz",
  "application/x-compressed",
  "application/x-gzip",
  "application/x-compressed-tar",
  "application/gzip",
];
const schema = Yup.object({
  name: Yup.string().required(),
  description: Yup.string(),
  appName: Yup.string().required(),
  file: Yup.mixed()
    .required("A file is required")
    .test(
      "fileFormat",
      "Unsupported file format",
      (value) => value && SUPPORTED_FORMATS.includes(value.type)
    ),
});

const ProfileForm = (props) => {
  const { onClose, item, open, onSubmit, appsData } = props;
  const buttonLabel = item ? "OK" : "Add";
  const title = item ? "Edit Profile" : "Add Profile";
  const handleClose = () => {
    onClose();
  };
  let initialValues = item
    ? {
        name: item.metadata.name,
        description: item.metadata.description,
        appName: item.spec["app-name"],
      }
    : {
        name: "",
        description: "",
        file: undefined,
        appName: appsData[0].metadata.name,
      };
  return (
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
        onSubmit={async (values) => {
          onSubmit(values);
        }}
        validationSchema={schema}
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
            <form
              encType="multipart/form-data"
              noValidate
              onSubmit={handleSubmit}
            >
              <DialogContent dividers>
                <TextField
                  style={{ float: "left", marginBottom: "10px" }}
                  id="name"
                  label="Name"
                  type="text"
                  value={values.name}
                  onChange={handleChange}
                  onBlur={handleBlur}
                  helperText={errors.name && touched.name && "Name is required"}
                  required
                  error={errors.name && touched.name}
                />
                <FormControl
                  style={{ width: "45%", float: "right", marginBottom: "10px" }}
                  error={errors.appName && touched.appName}
                >
                  <InputLabel id="appname-select-label">App</InputLabel>
                  <Select
                    labelId="appname-select-label"
                    id="appname-select-label"
                    name="appName"
                    value={values.appName}
                    onChange={handleChange}
                  >
                    {appsData.map((app) => (
                      <MenuItem
                        key={app.metadata.name}
                        value={app.metadata.name}
                      >
                        {app.metadata.name}
                      </MenuItem>
                    ))}
                  </Select>
                  {errors.appName && touched.appName && (
                    <FormHelperText>Required</FormHelperText>
                  )}
                </FormControl>

                <TextField
                  style={{ width: "100%", marginBottom: "25px" }}
                  name="description"
                  value={values.description}
                  onChange={handleChange}
                  onBlur={handleBlur}
                  id="description"
                  label="Description"
                  multiline
                  rowsMax={4}
                />
                <label
                  className="MuiFormLabel-root MuiInputLabel-root"
                  htmlFor="file"
                  id="file-label"
                >
                  Config override file
                  <span className="MuiFormLabel-asterisk MuiInputLabel-asterisk">
                     *
                  </span>
                </label>

                <FileUpload
                  setFieldValue={setFieldValue}
                  file={values.file}
                  name="file"
                  handleBlur={handleBlur}
                  accept={SUPPORTED_FORMATS}
                />
                {touched.file && (
                  <p style={{ color: "#f44336" }}>{errors.file}</p>
                )}
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

ProfileForm.propTypes = {
  onClose: PropTypes.func.isRequired,
  open: PropTypes.bool.isRequired,
  item: PropTypes.object,
  onSubmit: PropTypes.func.isRequired,
};

export default ProfileForm;
