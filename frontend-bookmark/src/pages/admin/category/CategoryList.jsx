import { deleteCategory, getAllCategories } from '@/service/category';
import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

function CategoryListAdmin() {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    getAllCategories()
      .then((category) => {
        console.log('Categories:', category.data);
        setCategories(category.data);
      })
      .catch((error) => {
        console.error('Error fetching categories:', error);
      });
  }, []);

  const handleDeleteCategory = async (id) => {
    deleteCategory(id)
      .then(() => {
        getAllCategories()
          .then((category) => {
            console.log('Categories:', category.data);
            setCategories(category.data);
          })
          .catch((error) => {
            console.error('Error fetching categories:', error);
          });
      })
      .catch((error) => {
        console.error('Error deleting category:', error);
      });
  };

  return (
    <div>
      <h2 className="text-2xl font-semibold mb-4">Category List</h2>
      <Link
        to="/admin/category-add" // Ganti dengan rute yang sesuai untuk halaman tambah kategori
        className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 mb-4 inline-block"
      >
        Add
      </Link>
      <table className="w-full table-auto">
        <thead>
          <tr className="bg-gray-200">
            <th className="px-4 py-2">ID</th>
            <th className="px-4 py-2">Name</th>
            <th className="px-4 py-2">Description</th>
            <th className="px-4 py-2">Image</th>
            <th className="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {categories.map((category) => (
            <tr className="text-center" key={category.ID}>
              <td className="border px-4 py-2">{category.ID}</td>
              <td className="border px-4 py-2">{category.Name}</td>
              <td className="border px-4 py-2">{category.Description}</td>
              <td className="border px-4 py-2">
                <img
                  src={'http://localhost:5000/' + category.Image}
                  alt={category.Name}
                  className="h-12"
                />
              </td>
              <td className="border px-4 py-2">
                <Link
                  to={'/admin/category-edit/' + category.ID}
                  className="mr-2 bg-blue-500 text-white px-4 py-1 rounded hover:bg-blue-600"
                >
                  Edit
                </Link>
                <button
                  onClick={() => handleDeleteCategory(category.ID)}
                  className="bg-red-500 text-white px-4 py-1 rounded hover:bg-red-600"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default CategoryListAdmin;
