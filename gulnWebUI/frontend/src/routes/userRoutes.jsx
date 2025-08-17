import { Routes, Route, Outlet } from "react-router-dom";
import Dashboard from "../pages/users/dashboard";
import ProjectDashboard from "../pages/users/projects/projectsDashboard";
import ProjectNew from "../pages/users/projects/projectNew"
import ProjectInfo from "../pages/users/projects/projectInfo"
import ProjectScanInfo from "../pages/users/projects/scans/scanInfo"
import ProjectUpload from "../pages/users/projects/projectUpload"
import ProtectedRoute from "../components/auth/protectedRoutes";

const UserRoutes = () => {
	return (
		<Routes>
			<Route path="dashboard" element={
				<ProtectedRoute>
					<Dashboard />
				</ProtectedRoute>
			}/>

			<Route path="projects" element={
				<ProtectedRoute>
					<ProjectDashboard />
				</ProtectedRoute>
			}/>

			<Route path="projects/new" element={
				<ProtectedRoute>	
					<ProjectNew />
				</ProtectedRoute>
			} />

			<Route path="projects/info/:projectUUID" element={
				<ProtectedRoute>	
					<ProjectInfo />
				</ProtectedRoute>
			} />

			<Route path="projects/upload/:projectUUID" element={
				<ProtectedRoute>	
					<ProjectUpload />
				</ProtectedRoute>
			} />

			<Route path="projects/info/:projectUUID/scan/info/:scanUUID" element={
				<ProtectedRoute>	
					<ProjectScanInfo />
				</ProtectedRoute>
			} />
		</Routes>
	);
};

export default UserRoutes;