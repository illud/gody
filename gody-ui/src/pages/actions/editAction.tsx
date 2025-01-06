import '../../App.css';
import TextField from '@mui/material/TextField';
import { useEffect, useState } from 'react';
import { EditActionApi } from './api';
import { ToastContainer, toast } from 'react-toastify';
import {  useNavigate, useLocation } from 'react-router-dom';
import Menu from '../../components/menu/menu'
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';

function EditAction() {
    const navigate = useNavigate();
    const location = useLocation();

    //actionId, actionName, actionPath, steps 
    var steps = JSON.parse(location.state.steps)

    var stepsStringArray = []
    var stepsArrayLength = []
    for (var i = 0; i < steps.steps.length; i++) {
        stepsStringArray.push(steps.steps[i].step)
        stepsArrayLength.push(i)
    }

    const [actionId] = useState(location.state.actionId || 0);
    const [githubExecute, setGithubExecute] = useState(steps.github.github_execute || false);
    const [actionName, setActionName] = useState(location.state.actionName || '');
    const [githubToken, setGithubToken] = useState(steps.github.github_token || '');
    const [repositoryOwner, setRepositoryOwner] = useState(steps.github.repository_owner || '');
    const [repositoryName, setRepositoryName] = useState(steps.github.repository_name || '');
    const [branchName, setBranchName] = useState(steps.github.branch_name || '');
    const [githubProjectPath, setGithubProjectPath] = useState(steps.github.github_project_path || '');

    const [ftpExecute, setftpExecute] = useState(steps.ftp.ftp_execute || false);
    const [ftpServer, setFtpServer] = useState(steps.ftp.ftp_server || '');
    const [ftpUsername, setFtpUsername] = useState(steps.ftp.username || '');
    const [ftpPassword, setFtpPassword] = useState(steps.ftp.password || '');
    const [ftpProjectPath, setFtpProjectPath] = useState(steps.ftp.project_path || '');
    const [ftpDirectory, setFtpDirectory] = useState(steps.ftp.ftp_directory || '');

    const [stepsPath, setStepsPath] = useState(steps.steps_path || '');
    const [stepsCount, setStepsCount] = useState<number[]>(stepsArrayLength); // Track steps here
    const [stepsText, setStepsText] = useState<string[]>(stepsStringArray); // Track each step's value

    const [loading, setLoading] = useState(false);

    const EditAction = async () => {
        // if (!actionName || !githubToken || !repositoryOwner || !repositoryName || !branchName || !projectPath || stepsText.length === 0) {
        //     toast('Please fill all the fields', { type: 'warning' });
        //     return;
        // }
        var githubBody = {
            githubExecute: githubExecute,
            githubToken: githubToken,
            repositoryOwner: repositoryOwner,
            repositoryName: repositoryName,
            branchName: branchName,
            githubProjectPath: githubProjectPath
        }

        var ftpBody = {
            ftpExecute: ftpExecute,
            ftpServer: ftpServer,
            username: ftpUsername,
            password: ftpPassword,
            projectPath: ftpProjectPath,
            ftpDirectory: ftpDirectory
        }

        setLoading(true);
        let result: any = await EditActionApi(actionId, actionName, githubBody, ftpBody, stepsPath, stepsText);
        if (result.data === 'actions updated') {
            toast('Actions updated', { type: 'success' });
            // Delay navigation by a short time to allow toast to show
            setTimeout(() => {
                navigate('/actions');
            }, 1000);
        } else {
            toast('Error editing actions', { type: 'error' });
        }
        setLoading(true);
    };

    // Add a new step when the user clicks "Add Step"
    const addStep = () => {
        setStepsCount((prevSteps) => [...prevSteps, prevSteps.length]); // Only add a step when clicking the button
        setStepsText((prevSteps) => [...prevSteps, '']); // Add an empty string to stepsText as a placeholder
    };

    // Remove a step when the user clicks "Delete Step"
    const deleteStep = (index: number) => {
        const newStepsCount = stepsCount.filter((_, i) => i !== index); // Remove the step from stepsCount
        const newStepsText = stepsText.filter((_, i) => i !== index); // Remove the step from stepsText

        // Re-index the steps
        setStepsCount(newStepsCount.map((_, i) => i)); // Re-index stepsCount to reset indexes
        setStepsText(newStepsText); // Set updated stepsText
    };

    // Update the value of a specific step when the user types something
    const handleStepChange = (index: number, value: string) => {
        const newSteps = [...stepsText];
        newSteps[index] = value; // Update the step at the given index
        setStepsText(newSteps);
    };

    const handleGithubCheckedChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setGithubExecute(event.target.checked);
    };

    const handleFtpCheckedChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setftpExecute(event.target.checked);
    };

    const StepsTexts = (index: number) => (
        <div key={index}>
            <p>{`Step ${index + 1}`}</p>
            <TextField
                id="filled-multiline-static"
                label={`Step ${index + 1}`}
                multiline
                rows={4}
                variant="filled"
                slotProps={{
                    input: {
                        style: { color: 'white', width: '400px' },
                    },
                    inputLabel: {
                        style: { color: 'grey' },
                    },
                }}
                value={stepsText[index] || ''}
                onChange={(e) => handleStepChange(index, e.target.value)}
            />
            <button onClick={() => deleteStep(index)} style={{ marginLeft: '10px' }}>
                X
            </button>
        </div>
    );
    useEffect(() => {

    }, [])
    return (
        <>
            <Menu />
            <br></br>
            <br></br>

            <h3>Adding GitHub step</h3>
            <p className="read-the-docs">
                By default the action will always check if theres a new change commited to the repository and will run the steps you define.
            </p>
            <p className="read-the-docs">
                Project must be already pulled to a location in the server.
            </p>
            <FormControlLabel
                control={
                    <Checkbox checked={githubExecute} onChange={handleGithubCheckedChange} name="gilad" />
                }
                label="Github"
            />
            <FormControlLabel
                control={
                    <Checkbox checked={ftpExecute} onChange={handleFtpCheckedChange} name="gilad" />
                }
                label="Ftp"
            />
            <div className="input-row">
                <TextField
                    id="standard-basic"
                    label="Action Name"
                    variant="standard"
                    slotProps={{
                        input: {
                            style: { color: 'white' },
                        },
                        inputLabel: {
                            style: { color: 'grey' },
                        },
                    }}
                    value={actionName}
                    onChange={(e) => setActionName(e.target.value)}
                />
            </div>
            {/* github */}
            {
                githubExecute ? <div>
                    <h4>GitHub step</h4>
                    <div className="input-row">
                        {/* Wrap the inputs in a flex container */}

                        <div className="input-row">
                            <TextField
                                id="standard-basic"
                                label="Project path (Example: C:/var/www/html)"
                                variant="standard"
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
                                value={githubProjectPath}
                                onChange={(e) => setGithubProjectPath(e.target.value)}
                            />

                            <TextField
                                id="standard-basic"
                                label="Github token"
                                variant="standard"
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
                                value={githubToken}
                                onChange={(e) => setGithubToken(e.target.value)}
                            />
                            <TextField
                                id="standard-basic"
                                label="Repository owner"
                                variant="standard"
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
                                value={repositoryOwner}
                                onChange={(e) => setRepositoryOwner(e.target.value)}
                            />
                            <TextField
                                id="standard-basic"
                                label="Repository name"
                                variant="standard"
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
                                value={repositoryName}
                                onChange={(e) => setRepositoryName(e.target.value)}
                            />
                        </div>

                        <div className="input-row">
                            <TextField
                                id="standard-basic"
                                label="Branch name"
                                variant="standard"
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
                                value={branchName}
                                onChange={(e) => setBranchName(e.target.value)}
                            />
                        </div>
                    </div>
                </div> : <></>
            }


            <br />
            <br />
            {/* ftp */}
            {
                ftpExecute ? <div>
                    <h4>Ftp step</h4>
                    <div className="input-row">
                        <TextField
                            id="standard-basic"
                            label="FTP Project path (Example: C:/username/project)"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={ftpProjectPath}
                            onChange={(e) => setFtpProjectPath(e.target.value)}
                        />
                        <TextField
                            id="standard-basic"
                            label="FTP Server"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={ftpServer}
                            onChange={(e) => setFtpServer(e.target.value)}
                        />

                        <TextField
                            id="standard-basic"
                            label="FTP Username"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={ftpUsername}
                            onChange={(e) => setFtpUsername(e.target.value)}
                        />

                        <TextField
                            id="standard-basic"
                            label="FTP Password"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={ftpPassword}
                            onChange={(e) => setFtpPassword(e.target.value)}
                        />

                        <TextField
                            id="standard-basic"
                            label="FTP Directory (Example: /godytest) in FTP Server"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={ftpDirectory}
                            onChange={(e) => setFtpDirectory(e.target.value)}
                        />

                    </div>

                </div> : <></>}
            <br />
            {
                stepsCount.length > 0 ? <div>
                    <h4>Scripts Step</h4>
                    <div className="input-row">
                        <TextField
                            id="standard-basic"
                            label="STEPS path (Example: C:/username/project)"
                            variant="standard"
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                            value={stepsPath}
                            onChange={(e) => setStepsPath(e.target.value)}
                        />
                    </div>
                </div> : <></>
            }
            <br />
            {/* Only render steps if there are steps in the array */}
            {
                stepsCount.length > 0 && stepsCount.map((index) => (
                    <div key={index}>{StepsTexts(index)}</div>
                ))
            }

            <br />
            <br />
            <button onClick={addStep}>Add Step</button>

            <br />
            <br />
            {loading ? <div>Loading...</div> : <button onClick={EditAction}>Edit Action</button>}
            <ToastContainer />
        </>
    );
}

export default EditAction;
