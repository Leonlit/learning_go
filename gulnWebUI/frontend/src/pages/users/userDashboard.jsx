import { useEffect, useState } from "react";
import HeadMetadata from "../../components/heads/headMetadata";
import ProtectedLayout from "../../components/layouts/protectedLayout";

const UserDashboard = () => {
	const [userDashboard, setuserDashboard] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");

	useEffect(() => {
		const fetchprojects = async () => {
			try {
				const res = await fetch("http://localhost:8080/core/projects-hosts-counts", {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to get core info");
				}

				const data = await res.json();
				setuserDashboard(data);
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
			<HeadMetadata title={"User dashboard"} />
			{<div className="dashboard">
				<h2>User Dashboard</h2>
				{!userDashboard ? (
					<p>No data in database.</p>
				) : (
                    <div>
						<table className="styled-table">
							<thead>
								<tr>
									<th>Projects</th>
									<th>Hosts</th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td>{ userDashboard.project_count }</td>
									<td>{ userDashboard.host_count }</td>
								</tr>
							</tbody>
						</table>
						</div>
				)}
			</div>}
		</ProtectedLayout>
	);
};

export default UserDashboard;
