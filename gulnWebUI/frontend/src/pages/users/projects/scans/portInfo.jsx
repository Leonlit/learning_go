import { useEffect, useState } from "react";
import { useParams, useLocation } from "react-router-dom";
import HeadMetadata from "../../../../components/heads/headMetadata";
import ProtectedLayout from "../../../../components/layouts/protectedLayout";

const ProjectScanHostInfo = () => {
	const [portDetails, setPortInfo] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
	const { projectUUID, scanUUID, hostUUID, portUUID } = useParams();
	const {state: hostState} = useLocation()
	
	useEffect(() => {
		const fetchprojects = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/scans/host/port/info/" + projectUUID + "/" + scanUUID + "/" + hostUUID + "/" + portUUID, {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch port info");
				}

				const data = await res.json();
				setPortInfo(data);
				
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
			<HeadMetadata title={hostState.ipAddr + ", port " +  + " - Host details"} />
			{<div className="dashboard">
				<h2>{hostState.ipAddr} - Host Details</h2>
				{!portDetails ? (
					<p>No data in database.</p>
				) : (
                    <div>
						<table className="styled-table">
							<thead>
								<tr>
									<th>Item</th>
									<th>Value</th>
								
								</tr>
							</thead>
							<tbody>
								<tr>
									<td>Port</td>
									<td>{ portDetails.port_number }</td>
								</tr>
								
								<tr>
									<td>Protocol</td>
									<td>{ portDetails.protocol }</td>
								</tr>
								
								<tr>
									<td>State</td>
									<td>{ portDetails.state } ({ portDetails.reason })</td>
								</tr>

								{portDetails.service_product && (
									<tr>
										<td>Product</td>
										<td>{portDetails.service_product}</td>
									</tr>
								)}
								
								{portDetails.service_name && (
									<tr>
										<td>Service Name</td>
										<td>{portDetails.service_name}</td>
									</tr>
								)}

								{portDetails.service_version && (
									<tr>
										<td>Version</td>
										<td>{portDetails.service_version}</td>
									</tr>
								)}
								
								{portDetails.service_fp && (
									<tr>
										<td>Service Fingerprint</td>
										<td>{portDetails.service_fp}</td>
									</tr>
								)}

								{portDetails.service_cpe && (
									<tr>
										<td>Service CPE</td>
										<td>{portDetails.service_cpe}</td>
									</tr>
								)}
								
							</tbody>
						</table>
                        <table className="styled-table">
							<thead>
								<tr>
									<th>Script ID</th>
									<th>Script Output</th>
								
								</tr>
							</thead>
							<tbody>
								{portDetails.scripts.map((script) => (
									<tr>
										<td>{script.script_id}</td>
										<td>{script.script_output}</td>
									</tr>
								))}
							</tbody>
                        </table>
                    </div>
				)}
			</div>}
		</ProtectedLayout>
	);
};

export default ProjectScanHostInfo;
