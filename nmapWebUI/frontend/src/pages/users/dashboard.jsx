import { useEffect, useState } from "react";
import HeadMetadata from "../../components/heads/headMetadata";

const Dashboard = () => {
	const [scans, setScans] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");

	useEffect(() => {
		const fetchScans = async () => {
			try {
				const res = await fetch("/api/scans/", {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch scans");
				}

				const data = await res.json();
				setScans(data);
			} catch (err) {
				setError(err.message);
			} finally {
				setLoading(false);
			}
		};

		fetchScans();
	}, []);

	if (loading) return <p>Loading...</p>;
	if (error) return <p className="error">{error}</p>;

	return (
		<>
			<HeadMetadata />
			<div className="dashboard">
				<h2>Scan Dashboard</h2>
				<table className="scan-table">
					<thead>
						<tr>
							<th>ID</th>
							<th>Scan Name</th>
							<th>Status</th>
							<th>Date</th>
						</tr>
					</thead>
					<tbody>
						{scans.map((scan) => (
							<tr key={scan.id}>
								<td>{scan.id}</td>
								<td>{scan.name}</td>
								<td>{scan.status}</td>
								<td>{new Date(scan.createdAt).toLocaleString()}</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</>
	);
};

export default Dashboard;
