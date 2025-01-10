

import '../../App.css'
import { Run } from './api'
import { useEffect, useState } from 'react';
import { GetActionExecutionHistoryApi } from './api';
import { ToastContainer, toast } from 'react-toastify';
import Menu from '../../components/menu/menu'
import { useParams } from "react-router-dom";
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';

function ActionDetails() {
    const { id } = useParams();
    const [actionId] = useState<number>(id ? parseInt(id) : 0);

    // const navigation = useNavigate();

    const [actions, setActions] = useState<any[]>([]);
    const [loading, setLoading] = useState(false);

    const getActionExecutionHistory = async (actionId: number) => {
        let result: any = await GetActionExecutionHistoryApi(actionId);
        setActions(result.data);
    }

    // const deleteAction = async (actionId: number) => {
    //     if (confirm('¿Are you sure you want to delete this action?')) {
    //         let result: any = await DeleteActionsApi(actionId);
    //         if (result.data === 'actions deleted') {
    //             getActionExecutionHistory(actionId)
    //         }
    //     }
    // }

    const RunAction = async (actionId: number) => {
        setLoading(true);
        let result: any = await Run(actionId);
        if (result.error) {
            toast(result.error, { type: 'error' });
        } else if (result.data === "actions executed successfully") {
            toast('Actions executed successfully', { type: 'success' });
            getActionExecutionHistory(actionId)
        }
        setLoading(false);
    }

    // const goEditAction = (actionId: number, actionName: String, actionPath: String, steps: String) => {
    //     navigation('/edit-action', { state: { actionId, actionName, actionPath, steps } })
    // }

    const getTotalTime = (steps: any) => {
        // Parse the steps from JSON
        steps = JSON.parse(steps);

        let totalTimeInMs = 0; // Total time in milliseconds

        // Calculate the total execution time in milliseconds
        for (let i = 0; i < steps.length; i++) {
            let executionTime = steps[i].execution_time;
            totalTimeInMs += parseInt(executionTime);
        }

        // Convert milliseconds to seconds
        let totalSeconds = totalTimeInMs;

        // Calculate hours, minutes, and seconds
        let hours = Math.floor(totalSeconds / 3600);
        let minutes = Math.floor((totalSeconds % 3600) / 60);
        let seconds = Math.round(totalSeconds % 60);

        // Build the formatted string


        return `${hours}h ${minutes}m ${seconds}s`;
    };


    const getTotalTimeForEachStep = (executionTime: any) => {

        // Convert milliseconds to seconds
        let totalSeconds = executionTime;

        // Calculate hours, minutes, and seconds
        let hours = Math.floor(totalSeconds / 3600);
        let minutes = Math.floor((totalSeconds % 3600) / 60);
        let seconds = Math.round(totalSeconds % 60);

        // Build the formatted string


        return `${hours}h ${minutes}m ${seconds}s`;
    };

    useEffect(() => {
        getActionExecutionHistory(actionId)
    }, [])

    return (
        <>
            <Menu />
            <br></br>
            <br></br>
            <br></br>

            {loading ? <button style={{ marginLeft: '10px' }} className="primary-btn">Running...</button> : <button style={{}} className="primary-btn" onClick={() => RunAction(actionId)}>
                {/* <RunIcon sx={{ display: { xs: 'none', md: 'flow' }, mr: 1 }} />  */}
                Run Action
            </button>}
            <br></br>
            <br></br>
            {actions.map((row) => (
                <div>
                    {<p style={{ color: 'white', }}>Total Execution Time: {getTotalTime(row.step)} </p>}
                    <Stepper alternativeLabel>
                        {JSON.parse(row.step).map((step: any, index: number) => {
                            return (
                                <Step key={index} active>
                                    <StepLabel>
                                        <div>
                                            <p
                                                style={{
                                                    color: 'white',
                                                    whiteSpace: 'nowrap',
                                                    overflow: 'hidden',
                                                    textOverflow: 'ellipsis'
                                                }}
                                                title={step.execution_name}  // Show full text on hover
                                            >
                                                {step.execution_name.length > 15 ? `${step.execution_name.substring(0, 15)}...` : step.execution_name}
                                            </p>
                                            <p style={{ color: 'white', }}>{getTotalTimeForEachStep(step.execution_time)} </p>
                                            {/* <p style={{ color: 'white', }}>{row.created_at}</p> */}
                                        </div>
                                    </StepLabel>
                                    {/* <button style={{ marginLeft: '10px' }} className="primary-btn" >
                                    Logs
                                </button> */}
                                </Step>
                            );
                        })}

                    </Stepper>
                    <br></br>
                    <br></br>
                </div>
            ))}
            <ToastContainer />
        </>
    )
}

export default ActionDetails
