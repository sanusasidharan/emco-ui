import React from "react";
import ExpandableCard from "../../common/ExpandableCard";
import { Grid, Paper, TextField, Typography } from "@material-ui/core";
import FileUpload from "../../common/FileUpload";

function AppDetailsForm({ formikProps, ...props }) {
  return (
    <Grid container spacing={3}>
      <Grid item xs={6}>
        <TextField
          fullWidth
          value={formikProps.values.apps[props.index].appName}
          name={`apps[${props.index}].appName`}
          id="app-name"
          label="App name"
          size="small"
          onChange={formikProps.handleChange}
          onBlur={formikProps.handleBlur}
          required
          helperText={
            formikProps.errors.apps &&
            formikProps.errors.apps[props.index] &&
            formikProps.errors.apps[props.index].appName
          }
          error={
            formikProps.errors.apps &&
            formikProps.errors.apps[props.index] &&
            formikProps.errors.apps[props.index].appName &&
            true
          }
        />
      </Grid>
      <Grid item xs={6}>
        <TextField
          fullWidth
          value={formikProps.values.apps[props.index].description}
          name={`apps[${props.index}].description`}
          id="app-description"
          label="Description"
          multiline
          onChange={formikProps.handleChange}
          onBlur={formikProps.handleBlur}
          rowsMax={4}
        />
      </Grid>
      <Grid item xs={6}>
        <label
          style={{ marginTop: "20px" }}
          className="MuiFormLabel-root MuiInputLabel-root"
          htmlFor="file"
          id="file-label"
        >
          App tgz file
          <span className="MuiFormLabel-asterisk MuiInputLabel-asterisk">
             *
          </span>
        </label>
        <FileUpload
          setFieldValue={formikProps.setFieldValue}
          file={formikProps.values.apps[props.index].file}
          onBlur={formikProps.handleBlur}
          name={`apps[${props.index}].file`}
          accept={".tgz"}
        />
        {formikProps.errors.apps &&
          formikProps.errors.apps[props.index] &&
          formikProps.errors.apps[props.index].file && (
            <p style={{ color: "#f44336" }}>
              {formikProps.errors.apps[props.index].file}
            </p>
          )}
      </Grid>
      <Grid item xs={6}>
        <label
          style={{ marginTop: "20px" }}
          className="MuiFormLabel-root MuiInputLabel-root"
          htmlFor="file"
          id="file-label"
        >
          Profile tar file
          <span className="MuiFormLabel-asterisk MuiInputLabel-asterisk">
             *
          </span>
        </label>
        <FileUpload
          setFieldValue={formikProps.setFieldValue}
          file={formikProps.values.apps[props.index].profilePackageFile}
          onBlur={formikProps.handleBlur}
          name={`apps[${props.index}].profilePackageFile`}
          accept={".tar.gz, .tar"}
        />
        {formikProps.errors.apps &&
          formikProps.errors.apps[props.index] &&
          formikProps.errors.apps[props.index].profilePackageFile && (
            <p style={{ color: "#f44336" }}>
              {formikProps.errors.apps[props.index].profilePackageFile}
            </p>
          )}
      </Grid>
    </Grid>
  );
}

const AppForm = (props) => {
  return (
    <ExpandableCard
      error={
        props.formikProps.errors.apps &&
        props.formikProps.errors.apps[props.index]
      }
      title={<Typography variant="h6">{props.name}</Typography>}
      description={props.description}
      content={
        <AppDetailsForm
          formikProps={props.formikProps}
          name={props.name}
          index={props.index}
        />
      }
    />
  );
};
export default AppForm;
