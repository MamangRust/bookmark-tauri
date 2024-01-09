import React, { useState } from 'react';
import { createCategory } from '../../service/category';
import { useNavigate } from 'react-router-dom';

function CategoryCreate() {
  const [formData, setFormData] = useState({
    name: '',
    image: '',
    description: '',
  });

  const { name, image, description } = formData;

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const formDataCategory = new FormData();

      formDataCategory.append('name', name);
      formDataCategory.append('file', image);
      formDataCategory.append('description', description);

      const response = await createCategory(formDataCategory);

      setFormData({ name: '', image: '', description: '' });

      navigate('/category-list');
    } catch (error) {
      console.error('Error creating category:', error);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    setFormData((prevData) => ({
      ...prevData,
      image: file,
    }));
  };

  return (
    <div className="flex justify-center items-center h-screen">
      <div className="w-1/2">
        <div className="bg-white border rounded-md p-6">
          <h2 className="text-2xl font-bold mb-4">Create Category</h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block">Name:</label>
              <input
                type="text"
                name="name"
                value={name}
                onChange={handleChange}
                className="border border-gray-300 rounded bg-gray-50 px-3 py-2 w-full"
                required
              />
            </div>
            <div>
              <label className="block">Image:</label>
              <input
                type="file"
                accept="image/*"
                name="image"
                onChange={handleImageChange}
                className="border border-gray-300 rounded bg-gray-50 px-3 py-2 w-full"
                required
              />
            </div>
            <div>
              <label className="block">Description:</label>
              <textarea
                name="description"
                value={description}
                onChange={handleChange}
                className="border border-gray-300 rounded bg-gray-50 px-3 py-2 w-full"
                required
              />
            </div>
            <button
              type="submit"
              className="bg-blue-700 hover:bg-blue-800 text-white px-4 py-2 rounded"
            >
              Create
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default CategoryCreate;
