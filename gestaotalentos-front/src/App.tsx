import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Candidatos from './pages/Candidatos';
import RegistrarAdm from './pages/RegistrarAdm';
import PrivateRoute from './components/privateRoute/privateRoute';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<Navigate to="/login" />} />
        <Route path="/registrar-adm" element={<RegistrarAdm />} />
        
        <Route 
          path="/candidatos" 
          element={
            <PrivateRoute>
              <Candidatos />
            </PrivateRoute>
          } 
        />
      </Routes>
    </Router>
  );
}

export default App;
