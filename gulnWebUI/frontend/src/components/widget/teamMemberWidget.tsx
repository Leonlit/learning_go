import { useEffect, useState } from "react";
import ListingWidgetLayout from "../layouts/listingWidgetLayout";
import { useNavigate } from "react-router-dom";

type TeamMember = {
    team_member_uuid: string
    team_member_name: string;
    team_member_department: string;
    team_member_project_role: string;
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

    useEffect(() => {
        const fetchTeamMembersList = async () => {
            try {
                const res = await fetch(
                    "http://localhost:8080/api/users/team-member/list/1",
                    {
                        credentials: "include", // this sends the cookie
                    },
                );

                const data: TeamMember[] = await res.json();

                setTeamMembers(data);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchTeamMembersList();
    }, []);

    if (loading) return <p>Loading...</p>;
    if (error) return <p className="error">{error}</p>;

    return (
        <ListingWidgetLayout
            data={teamMembers}
            onNewClick={() => navigate("/users/team/new")}
            newClickLabel="Add Team Member"
            onSearch={setFilteredTeamMembers}
            paginationAPI="http://localhost:8080/api/users/team-member/list/"
        >
            {(filteredTeamMembers.length === 0 ? teamMembers : filteredTeamMembers)
                .length === 0 ? (
                <p>No Team Members.</p>
            ) : (
                (filteredTeamMembers.length === 0
                    ? teamMembers
                    : filteredTeamMembers
                ).map((teamMember) => (
                    <div key={teamMember.team_member_uuid}>
                        <div>
                            {teamMember.team_member_name}
                        </div>
                        <div>
                            {teamMember.team_member_department}
                        </div>
                        <div>
                            {teamMember.team_member_project_role}
                        </div>
                    </div>
                ))
            )}
        </ListingWidgetLayout>
    );
};

export default TeamMemberWidget;
