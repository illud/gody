import {  useState } from 'react';
import { TextField, Button, Container, Typography, Box } from "@mui/material";
import { LoginApi } from './api'
import {  useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';
import { useTokenStore } from '../../services/zustand/zustand';

interface FormData {
    username: string;
    password: string;
}

function Login() {
    const {  setToken } = useTokenStore();

    const [formData, setFormData] = useState<FormData>({
        username: "",
        password: "",
    });

    const [error, setError] = useState<string>("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        if (!formData.username || !formData.password) {
            setError("Both fields are required!");
            return;
        }

        var result = await LoginApi(formData.username, formData.password);
        if (result.data) {
            if (result.data === "") {
                toast("Login failed!", { type: "error" });
                return;
            }else{
                setToken(result.data);
                toast("Login successful!", { type: "success" });
            }
            navigate("/home"); // Redirect to home page
        } else {
           toast("Login failed!", { type: "error" });
        }
    };

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value,
        });
    };

    return (
        <>
            <Container component="main" maxWidth="xs">
                <Box
                    sx={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        marginTop: 8,
                    }}
                >
                    <Typography variant="h5">Login</Typography>

                    <form onSubmit={handleSubmit} style={{ width: "100%" }}>
                        <TextField
                            variant="standard"
                            margin="normal"
                            required
                            fullWidth
                            label="Username"
                            name="username"
                            value={formData.username}
                            onChange={handleInputChange}
                            autoComplete="username"
                            autoFocus
                            slotProps={{
                                input: {
                                    style: { color: 'white', width: '400px' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                        />

                        <TextField
                            variant="standard"
                            margin="normal"
                            required
                            fullWidth
                            label="Password"
                            name="password"
                            type="password"
                            value={formData.password}
                            onChange={handleInputChange}
                            autoComplete="current-password"
                            slotProps={{
                                input: {
                                    style: { color: 'white', width: '400px' },
                                },
                                inputLabel: {
                                    style: { color: 'grey' },
                                },
                            }}
                        />

                        {/* Error message */}
                        {error && (
                            <Typography color="error" variant="body2" align="center" gutterBottom>
                                {error}
                            </Typography>
                        )}

                        {/* Submit Button */}
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            color="primary"
                            sx={{ marginTop: 3 }}
                        >
                            Sign In
                        </Button>
                    </form>
                </Box>
            </Container>
            <ToastContainer />
        </>
    );
}

export default Login;
