import myApi from './api';

export const registerUser = async (userData) => {
  try {
    const response = await myApi.post(`/auth/register`, userData);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

// Fungsi untuk login pengguna
export const loginUser = async (userData) => {
  try {
    const response = await myApi.post(`/auth/login`, userData);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};
