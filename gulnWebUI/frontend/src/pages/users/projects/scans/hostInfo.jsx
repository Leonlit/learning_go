import { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import HeadMetadata from "../../../../components/heads/headMetadata";
import ProtectedLayout from "../../../../components/layouts/protectedLayout";

const ProjectScanHostInfo = () => {
	const navigate = useNavigate();
	const [ports, setPorts] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
	const { projectUUID, scanUUID, hostUUID } = useParams();
	const {state: hostState} = useLocation()

	const navigateToPortInfo = (port) => {
		navigate("/users/projects/info/"+ projectUUID +"/scan/info/" + scanUUID + 
			"/host/" + hostUUID + "/port/" + port.port_uuid , {
			state: {ipAddr: hostState.ipAddr, }
		})
	}
	
	useEffect(() => {
		const fetchprojects = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/scans/host/info/" + projectUUID + "/" + scanUUID + "/" + hostUUID , {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch host info");
				}

				const data = await res.json();
				setPorts(data);
				console.log(data);
				
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
			<HeadMetadata title={hostState.ipAddr + " - Host details"} />
			{<div className="dashboard">
				<h2>{hostState.ipAddr} - Host Details</h2>
				{!ports ? (
					<p>No data in database.</p>
				) : (
					<table className="styled-table">
						<thead>
							<tr>
								<th>Port Number</th>
								<th>Port Protocol</th>
								<th>Port State</th>
								<th>Port Reason</th>
								<th>Port Service Name</th>
								<th>Port Service Product</th>
								<th>Port Service Version</th>
							</tr>
						</thead>
						<tbody>
							{ports.map((port) => (
                            <tr key={port.port_uuid}>
								<td><a onClick={() => navigateToPortInfo(port)}>{port.port_number}</a></td>
								<td>{port.protocol}</td>
								<td>{port.state}</td>
								<td>{port.reason}</td>
								<td>{port.service_name}</td>
								<td>{port.service_product}</td>
								<td>{port.service_version}</td>
                            </tr>
							))}
						</tbody>
					</table>
				)}
			</div>}
		</ProtectedLayout>
	);
};

export default ProjectScanHostInfo;
