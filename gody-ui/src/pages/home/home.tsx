

import { Link } from 'react-router-dom'
import Menu from '../../components/menu/menu'
import { Typography } from "@mui/material";

function Home() {
    return (
        <>
            <Menu />
            <div className="bodyContainer" style={{ marginTop: '-100px' }}>
                <div className="container">
                    <Typography
                        variant="h6"
                        noWrap
                        component="a"
                        // href="#app-bar-with-responsive-menu"
                        sx={{
                            mr: 2,
                            scale: 1.5,
                            fontFamily: 'monospace',
                            fontWeight: 700,
                            letterSpacing: '.3rem',
                            color: 'inherit',
                            textDecoration: 'none',
                        }}
                        style={{ color: 'white', fontSize: '2rem' }}
                    >
                        GODY
                    </Typography>
                    <p className="subtitle">
                        A minimal GitHub and FTP action runner for deploying your projects.
                    </p>
                    <div className="action-container">
                        <Link to="/actions">
                            <button className="primary-btn">
                                Go to Actions
                            </button>
                        </Link>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Home