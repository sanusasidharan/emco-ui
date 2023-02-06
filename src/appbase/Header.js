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
import React, {useContext} from "react";
import PropTypes from "prop-types";
import AppBar from "@material-ui/core/AppBar";
import Hidden from "@material-ui/core/Hidden";
import MenuIcon from "@material-ui/icons/Menu";
import Toolbar from "@material-ui/core/Toolbar";
import {withStyles} from "@material-ui/core/styles";
import {withRouter,Link,matchPath} from "react-router-dom";
import {
    Typography,
    Grid,
    IconButton,
    Popover,
    Divider,
    Button,
    Avatar,
} from "@material-ui/core";
import Breadcrumbs from "@material-ui/core/Breadcrumbs";
import {UserContext} from "../UserContext";
import {ExitToApp as LogoutIcon, NavigateNext, LockOpen} from "@material-ui/icons";
import {routes} from '../config/uiConfig';

const lightColor = "rgba(255, 255, 255, 0.7)";
const {ENABLE_RBAC} = window._env_ || {};
const styles = (theme) => ({
    root: {
        boxShadow:
        "0px 2px 4px -1px rgb(0 0 0 / 20%), 0px 4px 5px 0px rgb(0 0 0 / 14%), 0px 1px 10px 0px rgb(0 0 0 / 12%)",
    },
    secondaryBar: {
        zIndex: 0,
    },
    menuButton: {
        marginLeft: -theme.spacing(1),
    },
    iconButtonAvatar: {
        padding: 4,
    },
    link: {
        textDecoration: "none",
        color: lightColor,
        "&:hover": {
            color: theme.palette.common.white,
        },
    },
    button: {
        borderColor: lightColor,
    },
    breadcrumbLink: {
        color: "#FFF",
        textDecoration: "none",
        '&:hover': {
            textDecoration: "underline",
        }
      }
});

const getBreadcrumbs = ({ routes, pathname }) => {
    const matches = [];
    pathname
      .replace(/\/$/, '')
      .split('/')
      .reduce((previous, current) => {
        const pathSection = `${previous}/${current}`;
        let breadcrumbMatch;
        routes.some(({ name, path, param }) => {
          const match = matchPath(pathSection, { exact: true, path });
          if (match) {
            breadcrumbMatch = {
              breadcrumb: name,
              path,
              match,
              param
            };
            return true;
          }
          return false;
        });
        if (breadcrumbMatch) {
          matches.push(breadcrumbMatch);
        }
        return pathSection;
    });
    return matches;
};

function Header(props) {
    const {classes, onDrawerToggle, location} = props;
    const breadcrumbs = getBreadcrumbs({
        pathname: location.pathname,
        routes,
      });

    //set website title to current page
    breadcrumbs.forEach((breadcrumb, index) => {
        if (index === 0) {
            document.title = breadcrumb.breadcrumb;
        } else {
            document.title = document.title + " - " + (breadcrumb.param ? breadcrumb.match.params[breadcrumb.param] :breadcrumb.breadcrumb);
        }
    });

    const {user} = useContext(UserContext);
    const [anchorEl, setAnchorEl] = React.useState(null);
    const open = Boolean(anchorEl);
    const id = open ? "user-popover" : undefined;
    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <React.Fragment>
            {ENABLE_RBAC === 'true' &&
                <Popover
                    id={id}
                    open={open}
                    anchorEl={anchorEl}
                    onClose={handleClose}
                    anchorOrigin={{
                        vertical: "bottom",
                        horizontal: "center",
                    }}
                    transformOrigin={{
                        vertical: "top",
                        horizontal: "center",
                    }}
                >
                    <Typography
                        variant={"h6"}
                        style={{textAlign: "center", padding: "7px"}}
                    >
                        {user.displayName}
                    </Typography>
                    <Typography
                        variant={"subtitle2"}
                        color="textSecondary"
                        style={{textAlign: "center", padding: "0 0 7px 0"}}
                    >
                        {user.role}
                    </Typography>
                    <Divider variant="middle"/>
                    <Button style={{display: "block", padding: "10px 10px"}} onClick={() => {
                        handleClose();
                        props.onChangePasswordClick()
                    }} color="primary">
                        <LockOpen style={{verticalAlign: "bottom"}}/>
                        &nbsp;&nbsp;Change Password
                    </Button>
                    <Button style={{display: "block", padding: "10px 10px"}}
                            href="/logout" color="primary">
                        <LogoutIcon style={{verticalAlign: "bottom"}}/>
                        &nbsp;&nbsp;Logout
                    </Button>
                </Popover>}
            <AppBar
                className={classes.root}
                color="primary"
                position="sticky"
                elevation={0}
            >
                <Toolbar>
                    <Grid
                        container
                        spacing={1}
                        alignItems="center"
                        justify="space-between"
                    >
                        <Hidden smUp implementation="js">
                            <Grid item>
                                <IconButton
                                    color="inherit"
                                    onClick={onDrawerToggle}
                                    className={classes.menuButton}
                                >
                                    <MenuIcon/>
                                </IconButton>
                            </Grid>
                        </Hidden>
                        <Grid item>
                            <Breadcrumbs
                                color="white"
                                separator={<NavigateNext fontSize="small"/>}
                                aria-label="breadcrumb"
                            >
                            {breadcrumbs.map((breadcrumb, index) =>{
                                    if(index===breadcrumbs.length-1) 
                                return (
                                    <Typography
                                        underline="hover"
                                        key={breadcrumb.name + index}
                                        color="inherit"
                                    >
                                        {breadcrumb.param ? breadcrumb.match.params[breadcrumb.param] :breadcrumb.breadcrumb}
                                    </Typography>
                                ) 
                                else return (
                                    <Link
                                        className={classes.breadcrumbLink}
                                        key={breadcrumb.name + index}
                                        to={breadcrumb.match.url}
                                    >
                                        {breadcrumb.param ? breadcrumb.match.params[breadcrumb.param] :breadcrumb.breadcrumb}
                                    </Link>
                                ) } )}
                            </Breadcrumbs>
                        </Grid>
                        {ENABLE_RBAC === 'true' && <Grid item>
                            <IconButton color="inherit" onClick={handleClick}>
                                <Avatar src={user.image || null} alt="profile_image"/>
                            </IconButton>
                        </Grid>}
                    </Grid>
                </Toolbar>
            </AppBar>
        </React.Fragment>
    );
}

Header.propTypes = {
    classes: PropTypes.object.isRequired,
    onDrawerToggle: PropTypes.func.isRequired,
};

export default withStyles(styles)(withRouter(Header));
