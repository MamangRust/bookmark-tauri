import axios from 'axios';

export const myApi = axios.create({
  baseURL: 'http://localhost:5000/',
});

// Menambahkan interceptor untuk setiap request
myApi.interceptors.request.use(
  (config) => {
    // Mendapatkan token dari local storage
    const token = localStorage.getItem('token');

    // Jika token ada, atur header Authorization
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Menambahkan interceptor untuk menangani response error
myApi.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default myApi;
