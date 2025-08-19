import { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import HeadMetadata from "../../../../components/heads/headMetadata";
import ProtectedLayout from "../../../../components/layouts/protectedLayout";

const ProjectScanInfo = () => {
	const navigate = useNavigate();
	const [hosts, setHosts] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
	const { projectUUID, scanUUID } = useParams();
	const {state: scanState} = useLocation()

	const navigateToHostInfo = (host) => {
		navigate("/users/projects/info/"+ projectUUID +"/scan/info/" + scanUUID + "/host/" + host.host_uuid, {
			state: {ipAddr: host.ip_address}
		})
	}

	useEffect(() => {
		const fetchprojects = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/scans/info/" + projectUUID + "/" + scanUUID + "/1" , {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch scan hosts");
				}

				const data = await res.json();
				setHosts(data);
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
			<HeadMetadata title={scanState.scanName + " - Scan Dashboard"} />
			{<div className="dashboard">
				<h2>{scanState.scanName} - Scan Dashboard</h2>
				{!hosts || hosts.length === 0 ? (
					<p>No data in database.</p>
				) : (
					<table className="project-table">
						<thead>
							<tr>
								<th>IP Address</th>
								<th>Hostname</th>
								<th>Status</th>
							</tr>
						</thead>
						<tbody>
							{hosts.map((host) => (
								<tr key={host.host_uuid}>
									<td><a onClick={() => navigateToHostInfo(host)}>{host.ip_address}({host.addr_type})</a></td>
									<td>{host.hostname}</td>
									<td>{host.status}</td>
								</tr>
							))}
						</tbody>
					</table>
				)}
			</div>}
		</ProtectedLayout>
	);
};

export default ProjectScanInfo;
