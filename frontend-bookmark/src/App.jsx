import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Register from './pages/auth/Register';
import Login from './pages/auth/Login';
import CategoryList from './pages/category/CategoryList';
import CategoryById from './pages/category/CategoryById';
import PostsId from './pages/posts/PostsId';
import AuthProvider from './provider/auth';
import CategoryCreate from './pages/category/CategoryCreate';
import AdminLayout from './component/AdminLayout';
import CategoryListAdmin from './pages/admin/category/CategoryList';
import CategoryEditAdmin from './pages/admin/category/CategoryEdit';
import CategoryCreateAdmin from './pages/admin/category/CategoryCreate';
import PostCreateAdmin from './pages/admin/post/PostCreate';
import PostListAdmin from './pages/admin/post/PostList';
import PostEditAdmin from './pages/admin/post/PostEdit';
import HomeLayout from './component/HomeLayout';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/admin"
          element={
            <AdminLayout>
              <CategoryListAdmin />
            </AdminLayout>
          }
        />
        <Route
          path="/admin/category-add"
          element={
            <AdminLayout>
              <CategoryCreateAdmin />
            </AdminLayout>
          }
        />

        <Route
          path="/admin/category-edit/:id"
          element={
            <AdminLayout>
              <CategoryEditAdmin />
            </AdminLayout>
          }
        />

        <Route
          path="/admin/post-list"
          element={
            <AdminLayout>
              <PostListAdmin />
            </AdminLayout>
          }
        />

        <Route
          path="/admin/post-add"
          element={
            <AdminLayout>
              <PostCreateAdmin />
            </AdminLayout>
          }
        />

        <Route
          path="/admin/post-edit/:id"
          element={
            <AdminLayout>
              <PostEditAdmin />
            </AdminLayout>
          }
        />
        <Route
          path="/post-id/:id"
          element={
            <HomeLayout>
              <PostsId />
            </HomeLayout>
          }
        />
        <Route
          path="/category-id/:id"
          element={
            <HomeLayout>
              <CategoryById />
            </HomeLayout>
          }
        />
        <Route
          path="/category-list"
          element={
            <HomeLayout>
              <CategoryList />
            </HomeLayout>
          }
        />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
