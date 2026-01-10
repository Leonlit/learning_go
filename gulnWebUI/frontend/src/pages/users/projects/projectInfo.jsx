import { useEffect, useState } from "react";
import { useParams, useLocation, useNavigate } from "react-router-dom";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";

const ProjectInfo = () => {
	const [scans, setScans] = useState([]);
	const [infoHeader, setInfoHeader] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
    const { projectUUID } = useParams();
	const navigate = useNavigate();
	const {state} = useLocation()

	const navigateToProjectUpload = (project) => {
		navigate("/users/projects/upload/" + project.project_uuid, {
			state: {projectName: project.projectName}
		})
	}

	const navigateToScanInfo = (projectUUID, scan) => {
		navigate("/users/projects/info/"+ projectUUID +"/scan/info/" + scan.scan_uuid, {
			state: {scanName: scan.scan_name}
		})
	}

	useEffect(() => {
		let cancelled = false;

		const fetchAll = async () => {
			try {
				const [scansRes, headerRes] = await Promise.all([
					fetch(`http://localhost:8080/projects/info/scans/${projectUUID}/1`, {
						credentials: "include",
					}),
					fetch(`http://localhost:8080/projects/info/header/${projectUUID}`, {
						credentials: "include",
					}),
				]);

				if (!scansRes.ok || !headerRes.ok) {
					throw new Error("Failed to fetch project info");
				}

				const [scansData, headerData] = await Promise.all([
					scansRes.json(),
					headerRes.json(),
				]);

				if (!cancelled) {
					setScans(scansData);
					setInfoHeader(headerData);
				}
			} catch (err) {
				if (!cancelled) setError(err.message);
			} finally {
				if (!cancelled) setLoading(false);
			}
		};

		fetchAll();

		return () => {
			cancelled = true;
		};
	}, [projectUUID]);

	if (loading) return <p>Loading...</p>;
	if (error) return <p className="error">{error}</p>;

	return (
		<ProtectedLayout>
			<HeadMetadata title={state.projectName + " - Project Scans"}/>
			<button><a onClick={() => navigateToProjectUpload(state)}>Add Scan</a></button>
			<div className="dashboard">
				<h2>{state.projectName} - Project Scans</h2>
				{ !infoHeader ? (
					<p>No data in database.</p>
				) : (
					<div>
						<table className="styled-table">
							<thead>
								<tr>
									<th>Scans</th>
									<th>Hosts</th>
								</tr>
							</thead>
							<tbody>
								<tr>
									<td>{ infoHeader.scan_count }</td>
									<td>{ infoHeader.hosts_count }</td>
								</tr>
							</tbody>
						</table>
					</div>
				)}

				{!scans || scans.length === 0 ? (
					<p>No data in database.</p>
				) : (
					<table className="styled-table">
						<thead>
							<tr>
								<th>Scan Name</th>
								<th>Total Hosts</th>
								<th>Hosts Up</th>
								<th>Hosts Down</th>
								<th>Start Date</th>
								<th>Finish Date</th>
							</tr>
						</thead>
						<tbody>
							{scans.map((scan) => (
								<tr key={scan.scan_uuid}>
									<td><a onClick={() => navigateToScanInfo(state.projectUUID, scan)}>{scan.scan_name}</a></td>
									<td>{scan.total_hosts}</td>
									<td>{scan.hosts_up}</td>
									<td>{scan.hosts_down}</td>
									<td>{new Date(scan.scan_start_time).toLocaleString()}</td>
									<td>{new Date(scan.scan_finish_time).toLocaleString()}</td>
								</tr>
							))}
						</tbody>
					</table>
				)}
			</div>
		</ProtectedLayout>
	);
};

export default ProjectInfo;
