import { getCategoryById } from '@/service/category';
import React, { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';

function Category() {
  const { id } = useParams();
  const [category, setCategory] = useState({});

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

  return (
    <div className="flex justify-center items-center h-screen">
      <div className="container mx-auto">
        <div className="bg-white border rounded-md p-6">
          <img
            src={'http://localhost:5000/' + category.Image}
            alt={category.Name}
            className="w-50 h-50 object-contain mb-4 rounded-md"
          />
          <h2 className="text-2xl font-semibold mb-2">{category.Name}</h2>
          <p className="text-gray-600 mb-4">{category.Description}</p>
        </div>

        <hr className="my-8 border-gray-300" />
        <h3 className="text-xl font-semibold mb-2">Posts</h3>
        <div className="grid gap-4">
          {category.Posts &&
            category.Posts.map((post) => (
              <Link
                to={`/post-id/${post.ID}`}
                key={post.ID}
                className="text-black"
              >
                <div className="bg-white border rounded-md p-4 shadow-md">
                  <h4 className="font-semibold">{post.Title}</h4>
                </div>
              </Link>
            ))}
        </div>
      </div>
    </div>
  );
}

export default Category;
