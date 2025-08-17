import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";

const ProjectInfo = () => {
	const [project, setprojects] = useState([]);
	const [scans, setScans] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
    const { projectUUID } = useParams();

	useEffect(() => {
		const fetchProjectInfo = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/info/" + projectUUID, {
					credentials: "include", // Send JWT cookie
				});

				if (!res.ok) {
					throw new Error("Failed to fetch scans");
				}

				const data = await res.json();
				setprojects(data);
			} catch (err) {
				setError(err.message);
			}
		};

		fetchProjectInfo();

		const fetchScans = async () => {
			try {
				const res = await fetch("http://localhost:8080/projects/info/scans/" + projectUUID + "/1", {
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
			<HeadMetadata title={project.project_name + " - Project Info"}/>
			<button><a href={"/users/projects/upload/" + projectUUID}>Add Scan</a></button>
			<div className="dashboard">
				<h2>{project.project_name} - Project Info</h2>
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
								<th>Start Date</th>
								<th>Finish Date</th>
							</tr>
						</thead>
						<tbody>
							{scans.map((scan) => (
								<tr key={scan.scan_uuid}>
									<td>{scan.scan_uuid}</td>
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
