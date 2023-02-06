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
import React, {useState} from "react";
import {
    Backdrop,
    Button,
    Chip,
    CircularProgress,
    DialogActions,
    DialogContent,
    DialogTitle,
    IconButton,
    Paper
} from "@material-ui/core";
import AddIcon from "@material-ui/icons/Add";
import {makeStyles} from "@material-ui/styles";
import KvPairForm from "./KvPairForm";
import apiService from "../../../services/apiService";
import DeleteIcon from "@material-ui/icons/DeleteOutline";
import CloseIcon from '@material-ui/icons/Close';
import EditIcon from "@material-ui/icons/EditOutlined";

const useStyles = makeStyles((theme) => ({
    cardRoot: {
        display: 'flex',
        '& > *': {
            marginBottom: theme.spacing(3),
            width: "100%",
        },
    },
    cardHeader: {
        padding: "9px",
        backgroundColor: "rgba(0, 0, 0, 0.12)",
        borderRadius: "7px 7px 0 0",
        marginBottom: "10px"
    },
    backdrop: {
        zIndex: theme.zIndex.drawer + 1,
        color: "#fff",
    },
    closeButton: {
        position: 'absolute',
        right: theme.spacing(1),
        top: theme.spacing(1),
        color: theme.palette.grey[500],
    }
}));


const ManageKvPair = ({onSubmit, handleKvPairFormClose, data, updateData, ...props}) => {
    const classes = useStyles();
    const [kvPairFormOpen, setKvPairFormOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [kvPairToEdit, setKvPairToEdit] = useState();
    const handleKvPairFormSubmit = (values, isEdit) => {
        let request = {providerName: props.providerName, clusterName: data.clusterName};
        request.payload = {
            metadata: {
                name: values.name,
                description: values.description,
            },
            spec: {
                kv: []
            }
        };
        values.kvPair.forEach(kv => request.payload.spec.kv.push({[kv.key]: kv.value}));
        if(isEdit){
            request.kvPairName = kvPairToEdit.metadata.name;
            apiService.updateKvPair(request).then(res => {
                updateData(initialData => {
                    let itemToEditIndex = initialData.findIndex(element => element.metadata.name === res.metadata.name);
                    let updatedData = initialData;
                    updatedData[itemToEditIndex] = res;
                    updatedData.clusterName = initialData.clusterName;
                    return updatedData;
                });
            }).finally(() => {
                setKvPairFormOpen(false);
            })
        } else{
            apiService.createKvPair(request).then(res => {
                updateData(initialData => {
                    initialData.push(res);
                    return initialData;
                });
            }).finally(() => {
                setKvPairFormOpen(false);
            })
        }
    };

    const handleDeleteKvPair = (kvPairName) => {
        setIsLoading(true);
        let request = {providerName: props.providerName, clusterName: data.clusterName, kvPairName: kvPairName};
        apiService.deleteKvPair(request).then(() => {
            updateData((initialData) => {
                let updatedData = initialData.filter(item => item.metadata.name !== kvPairName);
                updatedData.clusterName = initialData.clusterName;
                return updatedData;
            })
        }).finally(() => {
            setIsLoading(false);
        })
    }

    const handleEditKvPair = (kvPair) => {
        setKvPairToEdit(kvPair);
        setKvPairFormOpen(true);
    }

    return (
        <>
            <Backdrop className={classes.backdrop} open={isLoading}>
                <CircularProgress color="primary"/>
            </Backdrop>
            <KvPairForm
                kvPairToEdit={kvPairToEdit}
                existingKvPairs={kvPairToEdit ? null : data}
                open={kvPairFormOpen}
                handleSubmit={handleKvPairFormSubmit}
                handleClose={() => setKvPairFormOpen(false)}
            />
            <Paper elevation={0} style={{
                width: "600px",
                padding: "5px",
                backgroundColor: "white",
                display: "flex",
                flexDirection: "column",
                height: "100%"
            }}>
                <DialogTitle id="kvPairDialog" onClose={handleKvPairFormClose}>
                    Manage Key Value Pairs
                    <IconButton aria-label="close" className={classes.closeButton} onClick={handleKvPairFormClose}>
                        <CloseIcon/>
                    </IconButton>
                </DialogTitle>
                <DialogContent dividers>
                    <div>
                        <Button
                            style={{marginTop: "10px", marginBottom: "15px"}}
                            fullWidth
                            variant="outlined"
                            color="primary"
                            startIcon={<AddIcon/>}
                            onClick={() => {
                                setKvPairToEdit(null);
                                setKvPairFormOpen(true);
                            }}
                        >
                            Add KV Pair
                        </Button>
                    </div>
                    {data && data.map((kvPair) =>
                        <div className={classes.cardRoot} key={kvPair.metadata.name}>
                            <Paper variant="outlined">
                                <div className={classes.cardHeader}>
                                    {kvPair.metadata.name}
                                    <div style={{float: "right"}}>
                                        <IconButton
                                            style={{padding: "0"}}
                                            aria-label="delete" title="Delete"
                                            onClick={() => handleDeleteKvPair(kvPair.metadata.name)}
                                            color="secondary">
                                            <DeleteIcon fontSize="small"/>
                                        </IconButton>
                                        <IconButton
                                            style={{padding: "0", marginLeft:"10px"}}
                                            aria-label="delete" title="Delete"
                                            onClick={() => handleEditKvPair(kvPair)}
                                            color="primary">
                                            <EditIcon fontSize="small"/>
                                        </IconButton>
                                    </div>
                                </div>
                                <div style={{padding: "15px"}}>
                                    {kvPair.spec.kv.map(kvItem => {
                                        return (Object.keys(kvItem).map((kvEntryKey, index) =>
                                                <Chip key={index}
                                                      style={{margin: "5px"}}
                                                      label={`${kvEntryKey} : ${kvItem[kvEntryKey]}`}
                                                      color="primary"
                                                />
                                            )
                                        )
                                    })
                                    }
                                </div>
                            </Paper>
                        </div>
                    )}
                </DialogContent>
                <DialogActions>
                    <Button
                        autoFocus
                        onClick={handleKvPairFormClose}
                        color="primary"
                    >
                        Close
                    </Button>
                </DialogActions>
            </Paper>
        </>
    )
}

export default ManageKvPair;