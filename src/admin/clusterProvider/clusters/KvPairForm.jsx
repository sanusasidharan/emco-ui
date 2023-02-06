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
import * as Yup from "yup";
import {FieldArray, Formik, getIn, useFormikContext} from "formik";
import {
    DialogActions,
    DialogContent,
    DialogTitle,
    FormControl,
    Grid,
    IconButton,
    MenuItem,
    Select,
    TextField
} from "@material-ui/core";
import Button from "@material-ui/core/Button";
import Dialog from "@material-ui/core/Dialog";
import InputLabel from "@material-ui/core/InputLabel";
import FormHelperText from "@material-ui/core/FormHelperText";
import DeleteIcon from "@material-ui/icons/DeleteTwoTone";
import AddIcon from "@material-ui/icons/Add";
import {Alert} from "@material-ui/lab";
import { nanoid } from 'nanoid'

const getSchema = (existingKvPairs) => {
    const commonSchema = (fieldName) => {
        const name = fieldName || "";
        return Yup.string()
            .max(50, `${name} cannot exceed more than 20 characters`)
            .matches(
                /^[a-zA-Z0-9_-]+$/,
                `${name} can only contain letters, numbers, '-' and '_' and no spaces.`
            )
            .matches(
                /^[a-zA-Z0-9]/,
                `${name} must start with an alphanumeric character`
            )
            .matches(/[a-zA-Z0-9]$/, `${name} must end with an alphanumeric character`);
    };
    Yup.addMethod(Yup.array, 'unique', function (message, mapper = a => {
        return a.key;
    }) {
        return this.test('unique', message, function (list) {
            return list.length === new Set(list.map(mapper)).size;
        });
    });

    return Yup.object(
        {
            type: Yup.string().required("Type is required"),
            name: Yup.string().required("Name is required").test(
                "duplicate-test",
                "KV pair with same name exists, please use a different name",
                (name) => {
                    return existingKvPairs
                        ? existingKvPairs.findIndex((x) => x.metadata.name === name) === -1
                        : true;
                }
            ).concat(commonSchema("Name")),
            kvPair: Yup.array().of(Yup.object({
                    key: Yup.string().concat(commonSchema("Key")).required("Key is required"),
                    value: Yup.string().required("Value is required")
                })
            ).unique('Key should be unique').required("At least one KV pair is required"),
        })
}
const KvPairFields = ({isEdit, ...props}) => {
    const {
        values: {type},
        touched,
        errors,
        values,
        handleBlur,
        handleChange,
        setFieldValue,
    } = useFormikContext();
    React.useEffect(() => {
        if (
            type.trim() === 'istioingresskvpairs' && !isEdit
        ) {
            setFieldValue("kvPair",
                [
                    {key: "istioingressgatewayaddress", value: ""},
                    {key: "istioingressgatewayport", value: ""},
                    {key: "istioingressgatewayinternalport", value: ""}
                ]);
            setFieldValue("name", "istioingresskvpairs")
        }
    }, [type, setFieldValue, props.name, isEdit]);

    return (
        <>
            <FieldArray
                name="kvPair"
                render={arrayHelpers => (
                    <>
                        {values.kvPair.map((kvPair, index) => (
                            <Grid container spacing={2} key={kvPair.id || index}>
                                <Grid item xs={5}>
                                    <TextField
                                        fullWidth
                                        autoFocus
                                        disabled={values.type !== "custom"}
                                        margin="dense"
                                        name={`kvPair[${index}].key`}
                                        label="Key"
                                        type="text"
                                        value={values.kvPair[index].key}
                                        onChange={handleChange}
                                        onBlur={handleBlur}
                                        variant="outlined"
                                        helperText={
                                            getIn(touched, `kvPair[${index}].key`) &&
                                            getIn(errors, `kvPair[${index}].key`)}
                                        error={Boolean(getIn(touched, `kvPair[${index}].key`) &&
                                            getIn(errors, `kvPair[${index}].key`))}
                                    />
                                </Grid>
                                <Grid item xs={5}>
                                    <TextField
                                        fullWidth
                                        margin="dense"
                                        name={`kvPair[${index}].value`}
                                        label="Value"
                                        type="text"
                                        value={values.kvPair[index].value}
                                        onChange={handleChange}
                                        onBlur={handleBlur}
                                        variant="outlined"
                                        helperText={
                                            getIn(touched, `kvPair[${index}].value`) &&
                                            getIn(errors, `kvPair[${index}].value`)}
                                        error={Boolean(getIn(touched, `kvPair[${index}].value`) &&
                                            getIn(errors, `kvPair[${index}].value`))}
                                    />
                                </Grid>
                                <Grid item xs={2}>
                                    <IconButton
                                        style={{marginTop: "10px"}}
                                        aria-label="delete" title="Delete"
                                        onClick={() => arrayHelpers.remove(index)}
                                        color="secondary"
                                        disabled={values.type !== "custom"}>
                                        <DeleteIcon fontSize="small"/>
                                    </IconButton>
                                </Grid>
                            </Grid>
                        ))}
                        <Grid item xs={10} style={{marginTop: "10px"}}>
                            {errors && errors.kvPair && typeof (errors.kvPair) === 'string' &&
                                <Alert severity="error">{errors.kvPair}</Alert>
                            }
                            <Button
                                disabled={(Boolean(errors && errors.kvPair) || values.type !== "custom") && values.kvPair.length > 0}
                                style={{marginTop: "20px"}}
                                variant="outlined"
                                color="primary"
                                startIcon={<AddIcon/>}
                                onClick={() => {
                                    arrayHelpers.push({key: '', value: '', id: nanoid()})
                                }}
                            >
                                Add KV Pair
                            </Button>
                        </Grid>
                    </>
                )}
            />
        </>
    );
};

const getInitValues = (kvPairToEdit) => {
    let initValues;
    if(kvPairToEdit){
        let kvPairObjectArray = [];
        let type = kvPairToEdit.metadata.name === "istioingresskvpairs" ? "istioingresskvpairs" : "custom";
        kvPairToEdit.spec.kv.map(kvItem => {
            return (Object.keys(kvItem).forEach((kvEntryKey) => {
                        kvPairObjectArray.push({key: kvEntryKey, value: kvItem[kvEntryKey]});
                    }
                )
            )
        })
        initValues = {type: type, name: kvPairToEdit.metadata.name, kvPair: kvPairObjectArray};
    } else {
        initValues = {type: "custom", name: "", kvPair: [{key: "", value: "", id: nanoid()}]};
    }
    return initValues;
}
const KvPairForm = ({open, handleSubmit, existingKvPairs, handleClose, ...props}) => {
    const KV_PAIR_TYPES = {custom: "Custom", istioingresskvpairs: "Istio Ingress"};
    const isEdit = !!props.kvPairToEdit;
    const title = isEdit ? "Edit KV Pair" : "Add New KV Pair";
    return (
        <Dialog
            fullWidth
            maxWidth="md"
            onClose={handleClose}
            aria-labelledby="customized-dialog-title"
            open={open}
            disableBackdropClick
        >
            <DialogTitle id="simple-dialog-title">{title}</DialogTitle>
            <Formik
                initialValues={getInitValues(props.kvPairToEdit)}
                onSubmit={(values) => {
                    handleSubmit(values, isEdit);
                }}
                validationSchema={getSchema(existingKvPairs)}
            >
                {(props) => {
                    const {
                        touched,
                        errors,
                        isSubmitting,
                        handleChange,
                        handleBlur,
                        handleSubmit,
                        values
                    } = props;
                    return (
                        <form noValidate onSubmit={handleSubmit}>
                            <DialogContent
                                dividers>
                                <Grid container spacing={2}>
                                    <Grid item xs={5}>
                                        <FormControl fullWidth variant="outlined" margin="dense">
                                            <InputLabel id="kvPair-select-label">Type</InputLabel>
                                            <Select
                                                disabled={isEdit}
                                                labelId="kvPair-select-label"
                                                label="Type"
                                                name="type"
                                                value={values.type}
                                                onChange={handleChange}
                                                onBlur={handleBlur}
                                            >
                                                {Object.keys(KV_PAIR_TYPES).map(kvPairType =>
                                                    <MenuItem
                                                        key={kvPairType} value={kvPairType}>{KV_PAIR_TYPES[kvPairType]}
                                                    </MenuItem>)}
                                            </Select>
                                            <FormHelperText style={{marginTop: "12px"}}>
                                                {errors.type && touched.type}
                                            </FormHelperText>
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={5}>
                                        <TextField
                                            fullWidth
                                            margin="dense"
                                            name="name"
                                            label="Name"
                                            type="text"
                                            value={values.name}
                                            onChange={handleChange}
                                            onBlur={handleBlur}
                                            variant="outlined"
                                            disabled={values.type !== "custom" || isEdit}
                                            helperText={
                                                getIn(touched, "name") &&
                                                getIn(errors, "name")}
                                            error={Boolean(getIn(touched, "name") &&
                                                getIn(errors, "name"))}
                                        />
                                    </Grid>
                                    <Grid item xs={12}>
                                        <KvPairFields name="kvPair" isEdit={isEdit}/>
                                    </Grid>
                                </Grid>
                            </DialogContent>
                            <DialogActions>
                                <Button
                                    disabled={isSubmitting}
                                    onClick={handleClose} color="secondary">
                                    Cancel
                                </Button>
                                <Button
                                    autoFocus
                                    type="submit"
                                    color="primary"
                                    disabled={isSubmitting}
                                >
                                    Submit
                                </Button>
                            </DialogActions>
                        </form>
                    );
                }}
            </Formik>
        </Dialog>
    )
}

export default KvPairForm;