import config from "../../../public/config.json"

export const LoginApi = async (username: String, password: String) => {
    var apilUrl = config.url + ":" + config.port 
    try {
        var body = {
            username: username,
            password: password,
        }
        let rawResult = await fetch(`${apilUrl}/users/login`, {
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