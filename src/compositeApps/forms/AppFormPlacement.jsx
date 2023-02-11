import React from "react";
import PropTypes from "prop-types";
import { Grid, Paper, Typography } from "@material-ui/core";
import ClustersTable from "./DIGPlacementTable";
import LabelsTable from "./DIGPlacementTableLabels";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import FormControl from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormHelperText from "@material-ui/core/FormHelperText";
import Select from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";

function RadioButtonsGroup({ formikProps, index, ...props }) {
  return (
    <FormControl component="fieldset">
      <FormLabel component="legend">Criterion</FormLabel>
      <RadioGroup
        row
        aria-label="criterion"
        name={`apps[${index}].placementCriterion`}
        value={formikProps.values.apps[index].placementCriterion}
        onChange={formikProps.handleChange}
        onBlur={formikProps.handleBlur}
      >
        <FormControlLabel value="allOf" control={<Radio />} label="All Of" />
        <FormControlLabel value="anyOf" control={<Radio />} label="Any Of" />
      </RadioGroup>
      <FormHelperText>Criterion for the app placement.</FormHelperText>
    </FormControl>
  );
}

function CustomSelect({ formikProps, index, ...props }) {
  return (
    <div>
      <FormControl>
        <InputLabel id="app-placement-type">Type</InputLabel>
        <Select
          labelId="app-placement-type-helper-label"
          name={`apps[${index}].placementType`}
          value={formikProps.values.apps[index].placementType}
          onChange={formikProps.handleChange}
          onBlur={formikProps.handleBlur}
        >
          <MenuItem value="labels">Labels</MenuItem>
          <MenuItem value="clusters">Specific Clusters</MenuItem>
        </Select>
        <FormHelperText style={{ marginTop: "12px" }}>
          {`Select targets based on the ${formikProps.values.apps[index].placementType}.`}
        </FormHelperText>
      </FormControl>
    </div>
  );
}

function AppFormPlacement({
  formikProps,
  index,
  logicalCloud,
  handleRowSelect,
  ...props
}) {
  return (
    <>
      <Typography variant="subtitle1" style={{ float: "left" }}>
        Select Clusters
        <span className="MuiFormLabel-asterisk MuiInputLabel-asterisk"> *</span>
      </Typography>
      {formikProps.errors.apps &&
        formikProps.errors.apps[index] &&
        formikProps.errors.apps[index].clusters && (
          <span
            style={{
              color: "#f44336",
              marginRight: "35px",
              float: "right",
            }}
          >
            {typeof formikProps.errors.apps[index].clusters === "string" &&
              formikProps.errors.apps[index].clusters}
          </span>
        )}
      <Grid
        container
        spacing={3}
        style={{
          height: "400px",
          overflowY: "auto",
          width: "100%",
        }}
      >
        <Grid key={"placementTypeGrid"} item xs={6}>
          <Paper style={{ padding: "14px 10px" }}>
            <CustomSelect formikProps={formikProps} index={index}/>
          </Paper>
        </Grid>
        <Grid key={"placementCriterionGrid"} item xs={6}>
          <Paper style={{ padding: "14px 10px" }}>
            <RadioButtonsGroup formikProps={formikProps} index={index} />
          </Paper>
        </Grid>
        {logicalCloud &&
          logicalCloud.spec.clusterReferences.spec.clusterProviders.length > 0 &&
          logicalCloud.spec.clusterReferences.spec.clusterProviders.map((clusterProvider) => (
            <Grid key={clusterProvider.metadata.name} item xs={12}>
              <Paper>
                {formikProps.values.apps[index].placementType === "clusters" ? (
                  <ClustersTable
                    key={clusterProvider.metadata.name}
                    tableName={clusterProvider.metadata.name}
                    clusters={clusterProvider.spec.clusters}
                    formikValues={formikProps.values.apps[index].clusters}
                    onRowSelect={handleRowSelect}
                  />
                ) : (
                  <LabelsTable
                    key={clusterProvider.metadata.name}
                    tableName={clusterProvider.metadata.name}
                    labels={clusterProvider.spec.labels}
                    formikValues={formikProps.values.apps[index].clusters}
                    onRowSelect={handleRowSelect}
                  />
                )}
              </Paper>
            </Grid>
          ))}
      </Grid>
    </>
  );
}

AppFormPlacement.propTypes = {
  formikProps: PropTypes.object,
  index: PropTypes.number,
  handleRowSelect: PropTypes.func,
};

export default AppFormPlacement;
