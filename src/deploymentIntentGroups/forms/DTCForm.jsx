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
import {getIn, useFormikContext} from 'formik';
import {FormControlLabel, Grid, TextField} from "@material-ui/core";
import Checkbox from "@material-ui/core/Checkbox";

const DTCForm = (props) => {
    const {values, handleBlur, handleChange, errors, setFieldValue} = useFormikContext();
    return (<>
        <div>
            <FormControlLabel
                control={<Checkbox checked={values.apps[props.index].dtcEnabled}
                                   onChange={(el) => {
                                       if (el.target.checked){
                                           setFieldValue(`apps[${props.index}].inboundServerIntent`, {"serviceName":"","port":"","protocol":""});
                                           handleChange(el);
                                       }
                                       else{
                                           handleChange(el);
                                          // setFieldValue(`apps[${props.index}].inboundServerIntent`, {});
                                       }
                                   }}
                                   name={`apps[${props.index}].dtcEnabled`}
                />}
                label="Expose Service"
            />
        </div>

        {values.apps[props.index].dtcEnabled && <Grid container spacing={2}>
            <Grid item xs={4}>
                <TextField
                    fullWidth
                    name={`apps[${props.index}].inboundServerIntent.serviceName`}
                    onBlur={handleBlur}
                    label="Service Name"
                    value={values.apps[props.index].inboundServerIntent.serviceName}
                    onChange={handleChange}
                    helperText={getIn(errors, `apps[${props.index}].inboundServerIntent.serviceName`)}
                    error={Boolean(getIn(errors, `apps[${props.index}].inboundServerIntent.serviceName`))}
                />
            </Grid>
            <Grid item xs={4}>
                <TextField
                    fullWidth
                    name={`apps[${props.index}].inboundServerIntent.port`}
                    onBlur={handleBlur}
                    label="Port"
                    value={values.apps[props.index].inboundServerIntent.port}
                    onChange={handleChange}
                    helperText={getIn(errors, `apps[${props.index}].inboundServerIntent.port`)}
                    error={Boolean(getIn(errors, `apps[${props.index}].inboundServerIntent.port`))}
                />
            </Grid>
            <Grid item xs={4}>
                <TextField
                    fullWidth
                    name={`apps[${props.index}].inboundServerIntent.protocol`}
                    onBlur={handleBlur}
                    label="Protocol"
                    value={values.apps[props.index].inboundServerIntent.protocol}
                    onChange={handleChange}
                    helperText={getIn(errors, `apps[${props.index}].inboundServerIntent.protocol`)}
                    error={Boolean(getIn(errors, `apps[${props.index}].inboundServerIntent.protocol`))}
                />
            </Grid>
        </Grid>}
    </>)
}

export default DTCForm;