import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";
import ListingWidgetLayout from "../../../components/layouts/listingWidgetLayout";

type Project = {
	project_uuid: string
	project_name: string
	project_created: string
}

const ProjectDashboard = () => {
    const navigate = useNavigate();
    const [projects, setProjects] = useState<Project[]>([]);
	const [filteredProjects, setFilteredProjects] = useState<Project[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

    const navigateToProjectInfo = (project: Project) => {
        navigate("/users/projects/info/" + project.project_uuid, {
            state: {
                projectUUID: project.project_uuid,
                projectName: project.project_name,
            },
        });
    };

    useEffect(() => {
        const fetchProjectList = async () => {
            try {
                const res = await fetch(
                    "http://localhost:8080/api/projects/list/1",
                    {
                        credentials: "include", // Send JWT cookie
                    },
                );

                if (!res.ok) {
                    throw new Error("Failed to fetch projects");
                }

                const data: Project[] = await res.json();
                setProjects(data);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchProjectList();
    }, []);

    if (loading) return <p>Loading...</p>;
    if (error) return <p className="error">{error}</p>;

    return (
        <ProtectedLayout>
            <HeadMetadata title={"Project Dashboard"} />
            <h2>Project Dashboard</h2>
            <ListingWidgetLayout
                data={projects}
                onNewClick={() => navigate("/users/projects/new")}
                onSearch={setFilteredProjects}
                paginationAPI="http://localhost:8080/api/projects/list/"
            >
                {(filteredProjects.length === 0 ? projects : filteredProjects)
                    .length === 0 ? (
                    <p>No Projects.</p>
                ) : (
                    <table className="styled-table">
                        <thead>
                            <tr>
                                <th>Project Name</th>
                                <th>Created On</th>
                            </tr>
                        </thead>
                        <tbody>
                            {(filteredProjects.length === 0
                                ? projects
                                : filteredProjects
                            ).map((project) => (
                                <tr key={project.project_uuid}>
                                    <td>
                                        <a
                                            onClick={() =>
                                                navigateToProjectInfo(project)
                                            }
                                        >
                                            {project.project_name}
                                        </a>
                                    </td>
                                    <td>
                                        {new Date(
                                            project.project_created,
                                        ).toLocaleString()}
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </ListingWidgetLayout>
        </ProtectedLayout>
    );
};

export default ProjectDashboard;
