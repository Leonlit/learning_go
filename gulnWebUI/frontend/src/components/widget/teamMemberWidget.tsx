import { useEffect, useState } from "react";
import ListingWidgetLayout from "../layouts/listingWidgetLayout";
import { useNavigate } from "react-router-dom";

type TeamMember = {
    team_member_uuid: string;
    team_member_name: string;
    team_member_department: string;
    team_member_project_role: string;
};

type TeamMemberCount = {
    count: number;
};

const TeamMemberWidget = () => {
    const navigate = useNavigate();
    const [teamMembers, setTeamMembers] = useState<TeamMember[]>([]);
    const [filteredTeamMembers, setFilteredTeamMembers] = useState<
        TeamMember[]
    >([]);
    const [loading, setLoading] = useState(true);
    const [authenticated, setAuthenticated] = useState(false);
    const [error, setError] = useState("");

    const [page, setPage] = useState(1);
    const [total, setTotal] = useState(0);

    useEffect(() => {
        const fetchTeamMembersList = async () => {
            try {
                const totalRes = await fetch(
                    "http://localhost:8080/api/team-member/count",
                    { credentials: "include" },
                );

                const countData: TeamMemberCount = await totalRes.json();
                setTotal(countData.count);

                const listRes = await fetch(
                    `http://localhost:8080/api/team-member/list/${page}`,
                    { credentials: "include" },
                );

                const data: TeamMember[] = await listRes.json();

                setTeamMembers(data);
                setFilteredTeamMembers(data); // reset filter on page change
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchTeamMembersList();
    }, [page]);

    if (loading) return <p>Loading...</p>;
    if (error) return <p className="error">{error}</p>;

    return (
        <ListingWidgetLayout
            data={teamMembers}
            onNewClick={() => navigate("/users/team/new")}
            newClickLabel="Add Team Member"
            onSearch={setFilteredTeamMembers}
            currentPage={page}
            total={total}
            onPageChange={setPage}
        >
            {(filteredTeamMembers.length === 0
                ? teamMembers
                : filteredTeamMembers
            ).length === 0 ? (
                <p>No Team Members.</p>
            ) : (
                (filteredTeamMembers.length === 0
                    ? teamMembers
                    : filteredTeamMembers
                ).map((teamMember) => (
                    <div key={teamMember.team_member_uuid}>
                        <div>{teamMember.team_member_name}</div>
                        <div>{teamMember.team_member_department}</div>
                        <div>{teamMember.team_member_project_role}</div>
                    </div>
                ))
            )}
        </ListingWidgetLayout>
    );
};

export default TeamMemberWidget;
