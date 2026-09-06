import { useEffect, useState } from "react";
import { useParams, useLocation, useNavigate } from "react-router-dom";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";
import ListingWidgetLayout from "../../../components/layouts/listingWidgetLayout";

import {
    Project,
    ProjectCount
} from "../../../types/projects";

import {
    AssessmentMinimal,
} from "../../../types/assessments";

const ProjectInfo = () => {
	const [assessmentList, setAssessmentList] = useState<AssessmentMinimal[]>([]);
	const [filteredAssessments, setFilteredAssessments] = useState<AssessmentMinimal[]>([]);
	const [projectInfo, setProjectInfo] = useState<Project | null>(null);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState("");
    const { projectUUID } = useParams();
	const navigate = useNavigate();
	const { state } = useLocation() as { state: Project | null };

	const [page, setPage] = useState(1);
    const [total, setTotal] = useState(0);

	if (!state) {
		return <p>No project data found.</p>;
	}

	const navigateToAddNewAssessment = (project: Project) => {
		//TODO: Change to new URL
		navigate("/users/projects/" + project.uuid + "/assessment/new", {
			state: project
		})
	}

	const navigateToAssessmentInfo = (project: Project, assessment: AssessmentMinimal) => {
		navigate("/users/projects/" + project.uuid + "/assessment/new", {
			state: project
		})
	}

	useEffect(() => {
		let cancelled = false;

		const fetchAll = async () => {
			try {
				const [projectInfoRes] = await Promise.all([
					fetch(`http://localhost:8080/api/projects/${projectUUID}/info`, {
						credentials: "include",
					}),
				]);

				const [assessmentListRes] = await Promise.all([
					fetch(`http://localhost:8080/api/projects/${projectUUID}/assessments/1`, {
						credentials: "include",
					}),
				]);

				if (!projectInfoRes.ok) {
					throw new Error("Failed to fetch project info");
				}

				if (!assessmentListRes) {
					throw new Error("Failed to fetch assessments info");
				}

				const [projectInfoData, assessmentListData] : [Project, AssessmentMinimal[]] = await Promise.all([
					projectInfoRes.json(),
					assessmentListRes.json()
				]);

				if (!cancelled) {
					setProjectInfo(projectInfoData);
					setAssessmentList(assessmentListData);
				}
			} catch (err) {
				if (!cancelled)
					if (err instanceof Error) {
						setError(err.message);
					}
			} finally {
				if (!cancelled) setLoading(false);
			}
		};

		fetchAll();

		return () => {
			cancelled = true;
		};
	}, [page]);

	if (loading) return <p>Loading...</p>;
	if (error) return <p className="error">{error}</p>;
	if (!projectInfo) {
		return <p>Loading...</p>; // or "Project not found"
	}

	return (
        <ProtectedLayout>
            <HeadMetadata title={projectInfo?.name + " - Project"} />
			

            <button>
                <a onClick={() => navigateToAddNewAssessment(state)}>
                    Add Assessment
                </a>
            </button>
            <div className="dashboard">
                <h2>Project - {state.name}</h2>

                <ListingWidgetLayout
                    data={assessmentList}
                    onNewClick={() => navigate("/users/projects/" + state.uuid + "/assessment/new", {state: state})}
                    headers={["Name", "Type", "Date Started", "Date Updated", "Date Ended", "Actions"]}
                    newClickLabel="Add New Project"
                    onSearch={setFilteredAssessments}
                    currentPage={page}
                    total={total}
                    onPageChange={setPage}
                >
                    {(filteredAssessments.length === 0
                        ? assessmentList
                        : filteredAssessments
                    ).length === 0 ? (
                        <p>No Projects.</p>
                    ) : (
                        (filteredAssessments.length === 0
                            ? assessmentList
                            : filteredAssessments
                        ).map((assessment: AssessmentMinimal) => (
                            <tr key={assessment.uuid}>
                                <td>
                                    <a
                                        onClick={() =>
                                            navigateToAssessmentInfo(projectInfo, assessment)
                                        }
                                    >
                                        {assessment.name}
                                    </a>
                                </td>
                                <td>
                                    {new Date(
                                        assessment.start_datetime,
                                    ).toLocaleString()}
                                </td>
                            </tr>
                        ))
                    )}
                </ListingWidgetLayout>
            </div>
        </ProtectedLayout>
    );
};

export default ProjectInfo;
