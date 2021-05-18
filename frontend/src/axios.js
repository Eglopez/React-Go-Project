import axios from 'axios';

const instance = axios.create({
    baseURL: "../../backend/services",
});

export default instance;