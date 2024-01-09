import api from './api';

export const getAllPosts = async () => {
  try {
    const response = await api.get('/posts');
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Mengambil post berdasarkan ID
export const getPostById = async (postId) => {
  try {
    const response = await api.get(`/posts/${postId}`);
    console.log(response.data);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Membuat post baru
export const createPost = async (postData) => {
  try {
    const response = await api.post('/posts/create', postData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Mengupdate post
export const updatePost = async (postId, postData) => {
  try {
    const response = await api.put(`/posts/update/${postId}`, postData);
    return response.data;
  } catch (error) {
    throw error;
  }
};

// Menghapus post
export const deletePost = async (postId) => {
  try {
    const response = await api.delete(`/posts/delete/${postId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};
