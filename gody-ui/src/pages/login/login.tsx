import { useState } from 'react';
import { TextField, Container, Typography, Box, IconButton, InputAdornment } from "@mui/material";
import { LoginApi } from './api'
import { useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';
import { useTokenStore } from '../../services/zustand/zustand';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

interface FormData {
    username: string;
    password: string;
}

function Login() {
    const { setToken } = useTokenStore();

    const [formData, setFormData] = useState<FormData>({
        username: "",
        password: "",
    });

    const [error, setError] = useState<string>("");

    const navigate = useNavigate();

    const [showPassword, setShowPassword] = useState(false);
    // Toggle visibility of password
    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

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
            } else {
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
                <div className="bodyContainer" style={{ marginTop: "-100px" }}>
                    <div className="container">
                        <Box
                            sx={{
                                display: "flex",
                                flexDirection: "column",
                                alignItems: "center",
                            }}
                        >
                            <Typography
                                variant="h6"
                                noWrap
                                component="a"
                                // href="#app-bar-with-responsive-menu"
                                sx={{
                                    mr: 2,
                                    scale: 1.5,
                                    display: { xs: 'none', md: 'flex' },
                                    fontFamily: 'monospace',
                                    fontWeight: 700,
                                    letterSpacing: '.3rem',
                                    color: 'inherit',
                                    textDecoration: 'none',
                                }}
                                style={{ color: 'white' }}
                            >
                                GODY
                            </Typography>

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
                                            style: { color: 'white', width: '100%' },
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
                                    type={showPassword ? "text" : "password"}
                                    value={formData.password}
                                    onChange={handleInputChange}
                                    autoComplete="current-password"
                                    slotProps={{
                                        input: {
                                            style: { color: 'white', width: '100%' },
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
                                />

                                {/* Error message */}
                                {error && (
                                    <Typography color="error" variant="body2" align="center" gutterBottom>
                                        {error}
                                    </Typography>
                                )}

                                {/* Submit Button */}
                                <button
                                    type="submit"
                                    className="primary-btn"
                                    style={{ width: '100%' }}
                                >
                                    Sign In
                                </button>
                            </form>
                        </Box>
                    </div>
                </div>
            </Container>
            <ToastContainer />
        </>
    );
}

export default Login;
