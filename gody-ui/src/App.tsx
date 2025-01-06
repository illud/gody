
// import { useState } from 'react'
import { Routes, HashRouter, Route } from 'react-router-dom'
import RunAction from './pages/actions/runAction'
import CreateAction from './pages/actions/createAction'
import Home from './pages/home/home'
import Actions from './pages/actions/actions'
import EditAction from './pages/actions/editAction'
import Login from './pages/login/login'

function App() {
  return (
    <>
      <HashRouter>
        <Routes>Login
          <Route path="/" element={<Login />} />
          <Route path="/home" element={<Home />} />
          <Route path="/run-action" element={<RunAction />} />
          <Route path="/create-action" element={<CreateAction />} />
          <Route path="/edit-action" element={<EditAction />} />
          <Route path="/actions" element={<Actions />} />
        </Routes>
      </HashRouter>
    </>
  )
}

export default App
