import { useEffect, useState } from "react";
import HeadMetadata from "../../components/heads/headMetadata";

const Dashboard = () => {
	const [scans, setScans] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");

	useEffect(() => {
		const fetchScans = async () => {
			try {
				const res = await fetch("http://localhost:8080/scans/list/1", {
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
							<tr key={scan.scan_uuid}>
								<td>{scan.scan_uuid}</td>
								<td>{scan.total_hosts}</td>
								<td>{scan.hosts_up}</td>
								<td>{scan.hosts_down}</td>
								<td>{new Date(scan.scan_time).toLocaleString()}</td>
							</tr>
						))}
					</tbody>
				</table>
			</div>
		</>
	);
};

export default Dashboard;
