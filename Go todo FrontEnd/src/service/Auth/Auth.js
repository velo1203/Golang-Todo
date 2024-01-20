import apiClient from "../api";

const AuthLogin = async (username, password) => {
    const data = {
        username: username,
        password: password
    }
    console.log(data)
    try {
        const response = await apiClient.post('/auth/login', data);
        return response.data;
    } catch (error) {
        console.error('POST Request Error:', error);
        throw error;
    }
};

const AuthRegister = async (username, password) => {
    const data = {
        username: username,
        password: password
    }
    console.log(data)
    try {
        const response = await apiClient.post('/auth/register', data);
        return response.data;
    } catch (error) {
        console.error('POST Request Error:', error);
        throw error;
    }
};

export {
    AuthLogin,
    AuthRegister
}