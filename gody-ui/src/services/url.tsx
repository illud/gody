export const getConfigFile = async () => {
    try {
        const response = await fetch('/config');  // use http://localhost:5000/config if you want to run the server locally
        const data = await response.json();  // Parse the JSON
        return data.data;  // Return the configuration data
    } catch (err) {
        throw err;  // Throw an error if the fetch fails
    }
}