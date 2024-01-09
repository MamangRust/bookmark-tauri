import React, { useEffect, useState } from 'react';
import { getAllCategories } from '@/service/category';
import { getPostById, updatePost } from '@/service/post';
import MDEditor from '@uiw/react-md-editor';
import { useNavigate, useParams } from 'react-router-dom';

function PostEditAdmin() {
  const [formData, setFormData] = useState({
    title: '',
    content: '',
    category_id: '', // Ganti dengan kategori yang sesuai
  });

  const { id } = useParams();

  const [categories, setCategories] = useState([]);

  const navigate = useNavigate();

  const { title, content, category_id } = formData;

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const postData = {
        title,
        content,
        category_id: parseInt(category_id),
      };

      console.log('Post data:', postData);

      const response = await updatePost(id, postData);
      console.log('Post created:', response);

      setFormData({ title: '', content: '', category_id: '', user_id: '' });

      navigate('/admin/post-list');
    } catch (error) {
      console.error('Error creating post:', error.message);
    }
  };

  useEffect(() => {
    getPostById(id)
      .then((post) => {
        setFormData({
          title: post.data.Title,
          content: post.data.Content,
          category_id: post.data.Category_id,
        });
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  useEffect(() => {
    getAllCategories()
      .then((response) => {
        console.log('Categories:', response.data);
        setCategories(response.data);
      })
      .catch((error) => {
        console.error('Error fetching categories:', error);
      });
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleEditorChange = (value) => {
    setFormData((prevData) => ({
      ...prevData,
      content: value,
    }));
  };

  return (
    <div className="p-4">
      <h2 className="text-2xl font-bold mb-4">Create Post</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block">Title:</label>
          <input
            type="text"
            name="title"
            value={title}
            onChange={handleChange}
            className="border border-gray-300 rounded px-3 py-2 w-full"
            required
          />
        </div>
        <div>
          <label className="block">Content:</label>
          <MDEditor value={content} onChange={handleEditorChange} />
        </div>
        <div>
          <label className="block">Category:</label>
          <select
            name="category_id"
            value={category_id}
            onChange={handleChange}
            className="border border-gray-300 rounded px-3 py-2 w-full"
            required
          >
            <option value="">Select a category</option>
            {categories.map((category) => (
              <option key={category.ID} value={category.ID}>
                {category.Name}
              </option>
            ))}
          </select>
        </div>
        <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Create
        </button>
      </form>
    </div>
  );
}

export default PostEditAdmin;
