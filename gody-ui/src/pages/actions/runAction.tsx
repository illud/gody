
// import { useState } from 'react'
import reactLogo from '../../../src/assets/react.svg'
import viteLogo from '../../../public/vite.svg'
import '../../App.css'

function RunAction() {
  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Gody</h1>
      <div className="card">
        {/* <button onClick={() => Run()}>
          Run Action
        </button> */}
      </div>
      <p className="read-the-docs">
        Click on the run button to run the action
      </p>
    </>
  )
}

export default RunAction
