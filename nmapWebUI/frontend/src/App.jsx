import React from 'react';
import { Route, Routes } from 'react-router-dom';
import LoginPage from './pages/auth/login.jsx'
import RegisterPage from './pages/auth/register.jsx'
import RegisterSuccessPage from "./pages/auth/registerSuccess.jsx"
import UserRoutes from './routes/userRoutes.jsx';

const App = () => {
    return (
        <Routes>
            <Route path="/" element={<LoginPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/registerSuccess" element={<RegisterSuccessPage />} />

            /* users routes */
            <Route path="/users/*" element={<UserRoutes />} />
        </Routes>
    );
};

export default App;