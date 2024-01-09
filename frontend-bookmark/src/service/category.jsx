import api from './api';

export const getAllCategories = async () => {
  try {
    const response = await api.get('/category');
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Endpoint untuk mendapatkan kategori berdasarkan ID
export const getCategoryById = async (categoryId) => {
  try {
    const response = await api.get(`/category/${categoryId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Endpoint untuk membuat kategori baru
export const createCategory = async (formData) => {
  try {
    const response = await api.post('/category/create', formData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Endpoint untuk mengupdate kategori
export const updateCategory = async (categoryId, formData) => {
  try {
    const response = await api.put(`/category/update/${categoryId}`, formData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Endpoint untuk menghapus kategori
export const deleteCategory = async (categoryId) => {
  try {
    const response = await api.delete(`/category/delete/${categoryId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};
