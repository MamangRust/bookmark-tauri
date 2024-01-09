import { getAllPosts, deletePost } from '@/service/post';
import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

function PostListAdmin() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    getAllPosts()
      .then((post) => {
        setPosts(post.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const handleDeletePost = async (id) => {
    deletePost(id)
      .then(() => {
        getAllPosts()
          .then((post) => {
            setPosts(post.data);
          })
          .catch((error) => {
            console.log(error);
          });
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div>
      <h2 className="text-2xl font-semibold mb-4">Post List</h2>
      <Link
        to="/admin/post-add"
        className="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 mb-4 inline-block"
      >
        Add
      </Link>
      <table className="w-full table-auto">
        <thead>
          <tr className="bg-gray-200">
            <th className="px-4 py-2">ID</th>
            <th className="px-4 py-2">Title</th>
            <th className="px-4 py-2">Content</th>
            <th className="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {posts.map((post) => (
            <tr className="text-center" key={post.ID}>
              <td className="border px-4 py-2">{post.ID}</td>
              <td className="border px-4 py-2">{post.Title}</td>
              <td className="border px-4 py-2">{post.Content}</td>

              <td className="border px-4 py-2">
                <Link
                  to={'/admin/post-edit/' + post.ID}
                  className="mr-2 bg-blue-500 text-white px-4 py-1 rounded hover:bg-blue-600"
                >
                  Edit
                </Link>
                <button
                  onClick={() => handleDeletePost(post.ID)}
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

export default PostListAdmin;
