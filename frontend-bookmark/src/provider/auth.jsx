import { createContext, useEffect } from 'react';
import { useAuthStore } from '../store/auth';
import { useNavigate } from 'react-router-dom';

function AuthProvider({ children }) {
  const AuthContext = createContext();
  const authStore = useAuthStore();
  const navigate = useNavigate();

  const handleTokenExpiration = () => {
    const token = authStore.token;
    const now = Math.floor(Date.now() / 1000);

    const expirationTime = now + 60 * 60;

    if (expirationTime && expirationTime < Date.now() / 1000) {
      authStore.setToken('');
      navigate('/');
    }
  };

  useEffect(() => {
    handleTokenExpiration();
  }, [authStore.token, navigate]);

  return (
    <AuthContext.Provider value={authStore}>{children}</AuthContext.Provider>
  );
}

export default AuthProvider;
