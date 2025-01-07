

const getConfigFile = async () => {
    try {
        const response = await fetch('/config');  // The path to the config.json file inside the public folder
        const data = await response.json();  // Parse the JSON
        return data.data;  // Return the configuration data
    } catch (err) {
        throw err;  // Throw an error if the fetch fails
    }
}

export const LoginApi = async (username: String, password: String) => {
    const config = await getConfigFile(); // Wait for the config to load
    const apilUrl = config.url + ":" + config.port;
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