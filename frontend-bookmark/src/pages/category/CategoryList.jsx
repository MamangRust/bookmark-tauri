import { getAllCategories } from '@/service/category';
import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';

function CategoryList() {
  const [categories, setCategories] = React.useState([]);

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

  return (
    <div className="flex justify-center items-center h-screen">
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mb-4">Category List</h1>
        <div className="grid grid-cols-3 gap-4">
          {categories.map((category) => (
            <Link to={`/category-id/${category.ID}`} key={category.ID}>
              <div className="bg-white border rounded-md p-4 shadow-md">
                <img
                  src={'http://localhost:5000/' + category.Image}
                  alt={category.Name}
                  className="w-full h-64 object-cover mb-4 rounded-t-md"
                />
                <h2 className="text-lg font-semibold mb-2">{category.Name}</h2>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
}

export default CategoryList;
