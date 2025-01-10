

import '../../App.css'
import { DeleteActionsApi } from './api'
import { useEffect, useState } from 'react';
import { GetActionsApi } from './api';

import { Link, useNavigate } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import Menu from '../../components/menu/menu'
import RunIcon from '@mui/icons-material/PlayCircle';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { Card, CardContent, CardActions, Button, Typography, Grid } from '@mui/material';

function Actions() {
    const navigation = useNavigate();

    const [actions, setActions] = useState<any[]>([]);
    // const [setLoading] = useState(false);

    const getActions = async () => {
        let result: any = await GetActionsApi();
        setActions(result.data);
    }

    const deleteAction = async (actionId: number) => {
        if (confirm('¿Are you sure you want to delete this action?')) {
            let result: any = await DeleteActionsApi(actionId);
            if (result.data === 'actions deleted') {
                getActions()
            }
        }
    }

    // const RunAction = async (actionId: number) => {
    //     setLoading(true);
    //     let result: any = await Run(actionId);
    //     if (result.error) {
    //         toast(result.error, { type: 'error' });
    //     } else if (result.data === "actions executed successfully") {
    //         toast('Actions executed successfully', { type: 'success' });
    //         getActions()
    //     }
    //     setLoading(false);
    // }

    const goEditAction = (actionId: number, actionName: String, actionPath: String, steps: String) => {
        navigation('/edit-action', { state: { actionId, actionName, actionPath, steps } })
    }

    useEffect(() => {
        getActions()
    }, [])

    return (
        <>
            <Menu />
            <br></br>
            <div className="card">
                <Link to="/create-action">
                    <button style={{ marginLeft: '10px', float: 'right' }} className="primary-btn" >
                        New action
                    </button>
                </Link>
            </div>

            <Grid container spacing={2} style={{ backgroundColor: '#242424' }}>
                {actions.map((row) => (
                    <Grid item xs={12} sm={6} md={4} key={row.id}>
                        <Card style={{ backgroundColor: '#333', color: 'white' }}>
                            <CardContent>
                                <Typography variant="h6" component="div" style={{ textAlign: 'center' }}>
                                    {row.action_name}
                                </Typography>
                                <Typography variant="body2" style={{ textAlign: 'center' }}>
                                    <strong>Created at:</strong> {row.created_at}
                                </Typography>
                                <Typography variant="body2" style={{ textAlign: 'center' }}>
                                    <strong>Updated at:</strong> {row.updated_at}
                                </Typography>
                            </CardContent>
                            <CardActions style={{ justifyContent: 'center' }}>
                                <Link to={`/actions-details/${row.id}/${row.action_name}`} style={{ textDecoration: 'none' }}>
                                    <Button variant="contained" color="primary" startIcon={<RunIcon />}>
                                        Run
                                    </Button>
                                </Link>
                                <Button
                                    variant="contained"
                                    color="secondary"
                                    startIcon={<EditIcon />}
                                    style={{ marginLeft: '10px' }}
                                    onClick={() => goEditAction(row.id, row.action_name, row.project_path, row.steps)}
                                >
                                    Edit
                                </Button>
                                <Button
                                    variant="contained"
                                    color="error"
                                    startIcon={<DeleteIcon />}
                                    style={{ marginLeft: '10px' }}
                                    onClick={() => deleteAction(row.id)}
                                >
                                    Delete
                                </Button>
                            </CardActions>
                        </Card>
                    </Grid>
                ))}
            </Grid>
            <ToastContainer />
        </>
    )
}

export default Actions
