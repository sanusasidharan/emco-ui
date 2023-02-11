
import React from "react";
import { withStyles } from "@material-ui/core/styles";


const styles = (theme) => ({
  paper: {
    maxWidth: 936,
    margin: "auto",
    overflow: "hidden",
  },
  searchBar: {
    borderBottom: "1px solid rgba(0, 0, 0, 0.12)",
  },
  searchInput: {
    fontSize: theme.typography.fontSize,
  },
  block: {
    display: "block",
  },
  addUser: {
    marginRight: theme.spacing(1),
  },
  contentWrapper: {
    margin: "0",
    background: "#000",
    color: "#fff",
    zIndex: "10",
    height: "30px",
    paddingLeft: "20px",
    paddingTop: "5px",

  },
  content: {
    background: "#000",
    color: "#fff"
  },
});

function Footer(props) {
  const { classes } = props;

  return (
    <React.Fragment>
   
      <div className={classes.contentWrapper}>
     
       
      <label
                  className="MuiFormLabel-root MuiInputLabel-root"
                  htmlFor="file"
                  id="file-label"
                >
                  Copyright © 2023 Infosys Limited     </label>
 
                 
      </div>
      </React.Fragment>
  );
}



export default withStyles(styles)(Footer);
