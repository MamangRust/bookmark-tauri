import {
  VideoCameraIcon,
  ArrowUpTrayIcon,
  UserCircleIcon,
} from '@heroicons/react/24/outline';
import { Link } from 'react-router-dom';

const AdminLayout = ({ children }) => {
  return (
    <div className="flex h-screen bg-gray-100">
      <div className="bg-gray-900 text-white w-64 flex-none">
        <div className="p-4">
          <div className="mb-8">Bookmark</div>
          <ul>
            <Link to={'/admin'}>
              <li className="py-2 hover:bg-gray-800 flex items-center">
                <VideoCameraIcon className="w-6 h-6 mr-2" />
                <span>Category</span>
              </li>
            </Link>
            <Link to={'/admin/post-list'}>
              <li className="py-2 hover:bg-gray-800 flex items-center">
                <ArrowUpTrayIcon className="w-6 h-6 mr-2" />
                <span>Posts</span>
              </li>
            </Link>
          </ul>
        </div>
      </div>

      <div className="flex-1 bg-gray-100">
        <header className="bg-gray-700 text-white p-4 flex items-center justify-between">
          <div className="flex items-center">
            <UserCircleIcon className="w-8 h-8 mr-2" />
            <h1 className="text-lg font-semibold">Admin Panel</h1>
          </div>
        </header>

        <main className="p-8 mt-10 bg-white rounded-lg shadow-lg m-8 flex-1">
          {children}
        </main>
      </div>
    </div>
  );
};

export default AdminLayout;
