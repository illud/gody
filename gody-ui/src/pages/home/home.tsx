

import '../../App.css'
import { Link } from 'react-router-dom'
import Menu from '../../components/menu/menu'

function RunAction() {
    return (
        <>
            <Menu />
<br></br>
            {/* <div>
                <a href="https://vite.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo" />
                </a>
            </div> */}
            <h1>Gody</h1>
            <p className="read-the-docs">
                A minimal GitHub, Gitlab and Gitbucket action runner for deploying your projects.
            </p>
            <div className="card">
                <Link to="/actions">
                    <button   >
                        View actions
                    </button>
                </Link>
            </div>
            <p className="read-the-docs">
                Created by illud.
            </p>
        </>
    )
}

export default RunAction
