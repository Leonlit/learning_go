import { Routes, Route, Outlet } from "react-router-dom";
import Dashboard from "../pages/users/dashboard";
import ScanUpload from "../pages/users/scan/scanUpload";
import ProtectedRoute from "../components/auth/protectedRoutes";

const UserRoutes = () => {
	return (
		<Routes>
			<Route path="dashboard" element={
				<ProtectedRoute>
					<Dashboard />
				</ProtectedRoute>
			}

			/>
			<Route path="scanUpload" element={
				<ProtectedRoute>
					<ScanUpload />
				</ProtectedRoute>} />
		</Routes>
	);
};

export default UserRoutes;