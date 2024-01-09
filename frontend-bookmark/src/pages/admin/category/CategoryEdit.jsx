import { getCategoryById, updateCategory } from '@/service/category';
import { useNavigate, useParams } from 'react-router-dom';
import { useState, useEffect } from 'react';

function CategoryEditAdmin() {
  const { id } = useParams();
  const [category, setCategory] = useState({
    Name: '',
    Description: '',
    Image: '',
  });

  const { Name, Image, Description } = category;
  const navigate = useNavigate();

  useEffect(() => {
    getCategoryById(id)
      .then((response) => {
        console.log(response.data);
        setCategory(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCategory((prevCategory) => ({
      ...prevCategory,
      [name]: value,
    }));
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    setCategory((prevCategory) => ({
      ...prevCategory,
      Image: file,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const formDataCategory = new FormData();

      formDataCategory.append('name', Name);
      formDataCategory.append('file', Image);
      formDataCategory.append('description', Description);

      const response = await updateCategory(id, formDataCategory);

      setCategory({ Name: '', Image: '', Description: '' });

      navigate('/admin');
    } catch (error) {
      console.error('Error update category:', error);
    }
  };

  return (
    <div>
      <h2 className="text-2xl font-semibold mb-4">Category Edit</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block">Name:</label>
          <input
            type="text"
            name="Name"
            value={Name || ''}
            onChange={handleChange}
            className="border border-gray-300 rounded px-3 py-2 w-full"
            required
          />
        </div>
        <div>
          <label className="block">Description:</label>
          <textarea
            name="Description"
            value={Description || ''}
            onChange={handleChange}
            className="border border-gray-300 rounded px-3 py-2 w-full"
            required
          ></textarea>
        </div>
        <div>
          <label className="block">Current Image:</label>
          {Image && (
            <img
              src={`http://localhost:5000/${Image}`}
              alt={Name}
              className="border border-gray-300 rounded px-3 py-2 w-full"
              style={{ maxWidth: '50%', height: 'auto' }}
            />
          )}

          <input
            type="file"
            accept="image/*"
            name="Image"
            onChange={handleImageChange}
            className="border border-gray-300 rounded px-3 py-2 w-full"
          />
        </div>

        <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Submit
        </button>
      </form>
    </div>
  );
}

export default CategoryEditAdmin;
