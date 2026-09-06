import { useEffect, useState } from "react";
import HeadMetadata from "../../components/heads/headMetadata";
import ProtectedLayout from "../../components/layouts/protectedLayout";

type DashboardData = {
    projectCount: number;
    assessmentCount: number;
};

type CountResponse = {
    count: number
}

const UserDashboard = () => {
    const [userDashboard, setUserDashboard] = useState<DashboardData | null>(
        null,
    );
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

    useEffect(() => {
        const fetchProjectsAndAssessmentCount = async () => {
            try {
                const projectRes = await fetch(
                    "http://localhost:8080/api/projects/count",
                    {
                        credentials: "include", // Send JWT cookie
                    },
                );

                const assessmentRes = await fetch(
                    "http://localhost:8080/api/stats/assessments/count",
                    {
                        credentials: "include", // Send JWT cookie
                    },
                );

                if (!projectRes.ok || !assessmentRes.ok) {
                    throw new Error("Failed to get core info");
                }

                const projectCount: CountResponse = await projectRes.json();
                const assessmentCount: CountResponse = await assessmentRes.json();
                const data = {
                    projectCount: projectCount.count,
                    assessmentCount: assessmentCount.count,
                };
                setUserDashboard(data);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchProjectsAndAssessmentCount();
    }, []);

    if (loading) return <p>Loading...</p>;
    if (error) return <p className="error">{error}</p>;

    return (
        <ProtectedLayout>
            <HeadMetadata title={"User dashboard"} />
            {
                <div className="dashboard">
                    <h2>User Dashboard</h2>
                    {!userDashboard ? (
                        <p>No data in database.</p>
                    ) : (
                        <div>
                            <table className="styled-table">
                                <thead>
                                    <tr>
                                        <th>Projects</th>
                                        <th>Hosts</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <td>{userDashboard.projectCount}</td>
                                        <td>{userDashboard.assessmentCount}</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    )}
                </div>
            }
        </ProtectedLayout>
    );
};

export default UserDashboard;
