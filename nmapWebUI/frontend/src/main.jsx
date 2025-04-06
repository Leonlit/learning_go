import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import LoginPage from './pages/auth/Login.jsx'
import RegisterPage from './pages/auth/Register.jsx'
import RegisterSuccessPage from "./pages/auth/RegisterSuccess.jsx"

createRoot(document.getElementById('root')).render(
	<StrictMode>
		<Router>
			<Routes>
				<Route path="/" element={<LoginPage />} />
				<Route path="/register" element={<RegisterPage />} />
				<Route path="/registerSuccess" element={<RegisterSuccessPage />} />
			</Routes>
		</Router>
	</StrictMode>,
)
