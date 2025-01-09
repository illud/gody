import '../../App.css';
import { TextField, IconButton, InputAdornment } from '@mui/material';
import { useState } from 'react';
import { CreateActionApi } from './api';
import { ToastContainer, toast } from 'react-toastify';
import { useNavigate } from 'react-router-dom';
import Menu from '../../components/menu/menu'
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import BackIcon from '@mui/icons-material/ArrowBack';
import { Link } from 'react-router-dom'

function CreateGithubAction() {
    const navigate = useNavigate();

    const [githubExecute, setGithubExecute] = useState(false);
    const [actionName, setActionName] = useState('');
    const [githubToken, setGithubToken] = useState('');
    const [repositoryOwner, setRepositoryOwner] = useState('');
    const [repositoryName, setRepositoryName] = useState('');
    const [branchName, setBranchName] = useState('');
    const [githubProjectPath, setGithubProjectPath] = useState('');

    const [ftpExecute, setftpExecute] = useState(false);
    const [ftpServer, setFtpServer] = useState('');
    const [ftpUsername, setFtpUsername] = useState('');
    const [ftpPassword, setFtpPassword] = useState('');
    const [ftpProjectPath, setFtpProjectPath] = useState('');
    const [ftpDirectory, setFtpDirectory] = useState('');

    const [stepsPath, setStepsPath] = useState('');
    const [stepsCount, setStepsCount] = useState<number[]>([]); // Track steps here
    const [stepsText, setStepsText] = useState<string[]>([]); // Track each step's value

    const [loading, setLoading] = useState(false);

    const handleGithubCheckedChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setGithubExecute(event.target.checked);
    };

    const handleFtpCheckedChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setftpExecute(event.target.checked);
    };

    const CreateAction = async () => {
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
        let result: any = await CreateActionApi(actionName, githubBody, ftpBody, stepsPath, stepsText);
        if (result.data === 'actions created') {
            toast('Actions created', { type: 'success' });
            // Delay navigation by a short time to allow toast to show
            setTimeout(() => {
                navigate('/actions');
            }, 1000);
        } else {
            toast('Error creating actions', { type: 'error' });
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

    const [showPassword, setShowPassword] = useState(false);
    // Toggle visibility of password
    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
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
            <button className="primary-btn" onClick={() => deleteStep(index)} style={{ marginLeft: '10px' }}>
                X
            </button>
        </div>
    );

    return (
        <>
            <Menu />
            <br></br>
            <br></br>
            <br></br>
            <Link to="/actions">
                <BackIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
            </Link>
            <h3>Creating new action</h3>
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
                                onChange={(e) => setGithubProjectPath(e.target.value)}
                            />
                            <TextField
                                id="standard-basic"
                                label="Github token"
                                variant="standard"
                                type={showPassword ? 'text' : 'password'}
                                slotProps={{
                                    input: {
                                        style: { color: 'white' },
                                        endAdornment: (
                                            <InputAdornment position="end">
                                                <IconButton
                                                    aria-label="toggle password visibility"
                                                    onClick={handleClickShowPassword}
                                                    edge="end"
                                                    style={{ color: 'white' }}
                                                >
                                                    {showPassword ? <VisibilityOff /> : <Visibility />}
                                                </IconButton>
                                            </InputAdornment>
                                        ),
                                    },
                                    inputLabel: {
                                        style: { color: 'grey' },
                                    },
                                }}
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
                            onChange={(e) => setFtpUsername(e.target.value)}
                        />

                        <TextField
                            id="standard-basic"
                            label="FTP Password"
                            variant="standard"
                            type={showPassword ? 'text' : 'password'}
                            slotProps={{
                                input: {
                                    style: { color: 'white' },
                                    endAdornment: (
                                        <InputAdornment position="end">
                                            <IconButton
                                                aria-label="toggle password visibility"
                                                onClick={handleClickShowPassword}
                                                edge="end"
                                                style={{ color: 'white' }}
                                            >
                                                {showPassword ? <VisibilityOff /> : <Visibility />}
                                            </IconButton>
                                        </InputAdornment>
                                    ),
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
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
                            onChange={(e) => setFtpDirectory(e.target.value)}
                        />
                    </div>

                </div> : <></>}
            <br />
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
            <button className="primary-btn" onClick={addStep}>Add Step</button>

            <br />
            <br />
            {loading ? <div>Loading...</div> : <button className="primary-btn" onClick={CreateAction}>Create Action</button>}
            <ToastContainer />
        </>
    );
}

export default CreateGithubAction;
