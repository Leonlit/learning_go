import { useEffect, useState } from "react";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";

const ProjectDashboard = () => {
	const [projects, setprojects] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");

	useEffect(() => {
		const fetchprojects = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/list/1", {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch projects");
				}

				const data = await res.json();
				setprojects(data);
			} catch (err) {
				setError(err.message);
			} finally {
				setLoading(false);
			}
		};

		fetchprojects();
	}, []);

	if (loading) return <p>Loading...</p>;
	if (error) return <p className="error">{error}</p>;

	return (
		<ProtectedLayout>
			<HeadMetadata title={"Project Dashboard"}/>
			<button><a href="/users/projects/new">Create New Project</a></button>
			{<div className="dashboard">
				<h2>Project Dashboard</h2>
				{!projects || projects.length === 0 ? (
					<p>No Projects.</p>
				) : (
					<table className="project-table">
						<thead>
							<tr>
								<th>Project Name</th>
								<th>Created On</th>
							</tr>
						</thead>
						<tbody>
							{projects.map((project) => (
								<tr key={project.project_uuid}>
									<td><a href={"/users/projects/info/" + project.project_uuid}>{project.project_name}</a></td>
									<td>{new Date(project.project_created).toLocaleString()}</td>
								</tr>
							))}
						</tbody>
					</table>
				)}
			</div>}
		</ProtectedLayout>
	);
};

export default ProjectDashboard;
