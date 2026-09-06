import { Routes, Route, Outlet } from "react-router-dom";
import UserDashboard from "../pages/users/userDashboard";
import ProjectDashboard from "../pages/users/projects/projectsDashboard";
import ProjectNew from "../pages/users/projects/projectNew"
import ProjectInfo from "../pages/users/projects/projectInfo"
import CreateNewTeamMember from "../pages/users/teamMembers/newTeamMembers"

import ProtectedRoute from "./protectedRoutes";


const UserRoutes = () => {
	return (
		<Routes>
			<Route path="dashboard" element={
				<ProtectedRoute>
					<UserDashboard />
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

			<Route path="projects/:projectUUID/info" element={
				<ProtectedRoute>	
					<ProjectInfo />
				</ProtectedRoute>
			} />

			<Route path="team-members/new" element={
				<ProtectedRoute>	
					<CreateNewTeamMember />
				</ProtectedRoute>
			} />
		</Routes>
	);
};

export default UserRoutes;