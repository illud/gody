

import '../../App.css'
import { Run, DeleteActionsApi } from './api'
import { useEffect, useState } from 'react';
import { GetActionsApi } from './api';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Link, useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';
import Menu from '../../components/menu/menu'
import RunIcon from '@mui/icons-material/PlayCircle';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

function Actions() {
    const navigation = useNavigate();

    const [actions, setActions] = useState<any[]>([]);
    const [loading, setLoading] = useState(false);

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

    const RunAction = async (actionId: number) => {
        setLoading(true);
        let result: any = await Run(actionId);
        if (result.error) {
            toast(result.error, { type: 'error' });
        } else if (result.data === "actions executed successfully") {
            toast('Actions executed successfully', { type: 'success' });
            getActions()
        }
        setLoading(false);
    }

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

            <TableContainer component={Paper} style={{ backgroundColor: '#242424' }}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table" >
                    <TableHead>
                        <TableRow>
                            <TableCell style={{ color: 'white', textAlign: 'center' }}>Action name</TableCell>
                            <TableCell align="right" style={{ color: 'white', textAlign: 'center' }}>Created at</TableCell>
                            <TableCell align="right" style={{ color: 'white', textAlign: 'center' }}>updated at</TableCell>
                            <TableCell align="right" style={{ color: 'white', textAlign: 'center' }}>Actions</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {actions.map((row) => (
                            <TableRow
                                key={row.id}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row" style={{ color: 'white' }}>
                                    {row.action_name}
                                </TableCell>
                                <TableCell align="right" style={{ color: 'white' }}>{row.created_at}</TableCell>
                                <TableCell align="right" style={{ color: 'white' }}>{row.updated_at}</TableCell>
                                <TableCell align="right" style={{ color: 'white' }}>
                                    {
                                        loading ? <button style={{ marginLeft: '10px' }} className="primary-btn">Running...</button> : <button style={{ marginLeft: '10px' }} className="primary-btn" onClick={() => RunAction(row.id)}>
                                            <RunIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
                                        </button>
                                    }
                                    <button style={{ marginLeft: '10px' }} className="primary-btn" onClick={() => goEditAction(row.id, row.action_name, row.project_path, row.steps)}>
                                        <EditIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
                                    </button>
                                    <button style={{ marginLeft: '10px' }} className="primary-btn" onClick={() => deleteAction(row.id)}>
                                        <DeleteIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
                                    </button>
                                </TableCell>

                                {/* <TableCell align="right">{row.calories}</TableCell>
                                <TableCell align="right">{row.fat}</TableCell>
                                <TableCell align="right">{row.carbs}</TableCell>
                                <TableCell align="right">{row.protein}</TableCell> */}
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer >
            <ToastContainer />
        </>
    )
}

export default Actions
