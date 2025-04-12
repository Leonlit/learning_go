import { useEffect, useState } from "react";
import HeadMetadata from "../../components/heads/headMetadata";
import ProtectedLayout from "../../components/layouts/protectedLayout";

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
		<ProtectedLayout>
			<HeadMetadata title={"Dashboard"}/>
			<div className="dashboard">
				<h2>Scan Dashboard</h2>
				{!scans || scans.length === 0 ? (
					<p>No data in database.</p>
				) : (
					<table className="scan-table">
						<thead>
							<tr>
								<th>ID</th>
								<th>Scan Name</th>
								<th>Status</th>
								<th>Hosts Down</th>
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
				)}
			</div>
		</ProtectedLayout>
	);
};

export default Dashboard;
