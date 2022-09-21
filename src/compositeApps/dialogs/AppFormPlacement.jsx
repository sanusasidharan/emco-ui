import React from "react";
import PropTypes from "prop-types";
import { Grid, Paper, Typography } from "@material-ui/core";
import ClustersTable from "./DIGPlacementTable";

function AppFormPlacement({
  formikProps,
  index,
  clusterProviders,
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
          scrollbarWidth: "thin",
        }}
      >
        {clusterProviders &&
          clusterProviders.length > 0 &&
          clusterProviders.map((clusterProvider) => (
            <Grid key={clusterProvider.name} item xs={12}>
              <Paper>
                <ClustersTable
                  key={clusterProvider.name}
                  tableName={clusterProvider.name}
                  clusters={clusterProvider.clusters}
                  formikValues={formikProps.values.apps[index].clusters}
                  onRowSelect={handleRowSelect}
                />
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
