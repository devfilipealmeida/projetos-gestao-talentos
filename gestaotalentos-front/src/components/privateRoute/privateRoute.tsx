import React from 'react';
import { Navigate } from 'react-router-dom';
import { decodeToken } from '../../utils/decodeToken';

interface PrivateRouteProps {
  children: React.ReactNode;
}

const isTokenValid = (token: string | null): boolean => {
  if (!token) return false;

  const decodedToken = decodeToken(token);
  if (!decodedToken) return false;

  const currentTime = Date.now() / 1000;
  return decodedToken.exp > currentTime;
};

const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
  const token = localStorage.getItem('token');
  
  return isTokenValid(token) ? <>{children}</> : <Navigate to="/login" />;
};

export default PrivateRoute;