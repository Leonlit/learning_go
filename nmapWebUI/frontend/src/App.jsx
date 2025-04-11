import React from 'react';
import { Route, Routes } from 'react-router-dom';
import LoginPage from './pages/auth/Login.jsx'
import RegisterPage from './pages/auth/Register.jsx'
import RegisterSuccessPage from "./pages/auth/RegisterSuccess.jsx"

const App = () => {
    return (
        <Routes>
            <Route path="/" element={<LoginPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/registerSuccess" element={<RegisterSuccessPage />} />
        </Routes>
    );
};

export default App;