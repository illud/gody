
import config from "../../../public/config.json"

interface IActionStep {
    step_type: Number,
    step: String
}

export const CreateActionApi = async (actionName: String, github: {
    githubExecute: Boolean,
    githubToken: String,
    repositoryOwner: String,
    repositoryName: String,
    branchName: String,
    githubProjectPath: String
},
    ftp: {
        ftpExecute: Boolean
        ftpServer: String
        username: String
        password: String
        projectPath: String
        ftpDirectory: String
    },
    stepsPath: String,
    steps: Array<String>) => {

    var stepsArray: Array<IActionStep> = steps.map(step => {
        return {
            "step_type": 1,
            "step": step
        }
    })


    var stepsAction: String = JSON.stringify({
        "github": {
            "github_execute": github.githubExecute,
            "ftp_execute": true,
            "github_token": github.githubToken,
            "repository_owner": github.repositoryOwner,
            "repository_name": github.repositoryName,
            "branch_name": github.branchName,
            "github_project_path": github.githubProjectPath
        },
        "ftp": {
            "ftp_execute": ftp.ftpExecute,
            "ftp_server": ftp.ftpServer,
            "username": ftp.username,
            "password": ftp.password,
            "project_path": ftp.projectPath,
            "ftp_directory": ftp.ftpDirectory
        },
        "steps_path": stepsPath,
        "steps": stepsArray
    })

    var body = {
        action_name: actionName,
        action_type: 1, // github
        steps: stepsAction
    }

    var apilUrl = config.url + ":" + config.port
    try {

        let rawResult = await fetch(`${apilUrl}/actions`, {
            method: 'POST',
            credentials: 'omit',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        }).then((response) => response)
        let result = await rawResult.json()
        if (result.data) return result
        else return { data: null }
    } catch (error) {
        console.log(error)
    }
}

export const EditActionApi = async (actionId: number, actionName: String, github: {
    githubExecute: Boolean,
    githubToken: String,
    repositoryOwner: String,
    repositoryName: String,
    branchName: String,
    githubProjectPath: String
},
    ftp: {
        ftpExecute: Boolean
        ftpServer: String
        username: String
        password: String
        projectPath: String
        ftpDirectory: String
    },
    stepsPath: String,
    steps: Array<String>) => {

    var stepsArray: Array<IActionStep> = steps.map(step => {
        return {
            "step_type": 1,
            "step": step
        }
    })


    var stepsAction: String = JSON.stringify({
        "github": {
            "github_execute": github.githubExecute,
            "ftp_execute": true,
            "github_token": github.githubToken,
            "repository_owner": github.repositoryOwner,
            "repository_name": github.repositoryName,
            "branch_name": github.branchName,
            "github_project_path": github.githubProjectPath
        },
        "ftp": {
            "ftp_execute": ftp.ftpExecute,
            "ftp_server": ftp.ftpServer,
            "username": ftp.username,
            "password": ftp.password,
            "project_path": ftp.projectPath,
            "ftp_directory": ftp.ftpDirectory
        },
        "steps_path": stepsPath,
        "steps": stepsArray
    })

    var body = {
        action_name: actionName,
        action_type: 1, // github
        steps: stepsAction
    }

    var apilUrl = config.url + ":" + config.port
    try {

        let rawResult = await fetch(`${apilUrl}/actions/${actionId}`, {
            method: 'PUT',
            credentials: 'omit',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        }).then((response) => response)
        let result = await rawResult.json()
        if (result.data) return result
        else return { data: null }
    } catch (error) {
        console.log(error)
    }
}

export const Run = async (actionId: number) => {
    var apilUrl = config.url + ":" + config.port
    try {
        var body = {
            action_id: actionId,
        }
        let rawResult = await fetch(`${apilUrl}/actions/run`, {
            method: 'POST',
            credentials: 'omit',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        }).then((response) => response)
        let result = await rawResult.json()
        if (result.data) return result
        else return result
    } catch (error) {
        console.log(error)
    }
}

export const GetActionsApi = async () => {
    var apilUrl = config.url + ":" + config.port
    try {
        let rawResult = await fetch(`${apilUrl}/actions`, {
            method: 'GET',
            credentials: 'omit',
            headers: {
                'Content-Type': 'application/json',
            },
        }).then((response) => response)
        let result = await rawResult.json()
        if (result.data) return result
        else return { data: null }
    } catch (error) {
        console.log(error)
    }
}

export const DeleteActionsApi = async (actionsId: number) => {
    var apilUrl = config.url + ":" + config.port
    try {
        let rawResult = await fetch(`${apilUrl}/actions/${actionsId}`, {
            method: 'DELETE',
            credentials: 'omit',
            headers: {
                'Content-Type': 'application/json',
            },
        }).then((response) => response)
        let result = await rawResult.json()
        if (result.data) return result
        else return { data: null }
    } catch (error) {
        console.log(error)
    }
}