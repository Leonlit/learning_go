import { useEffect, useState } from "react";
import ListingWidgetLayout from "../layouts/listingWidgetLayout";
import Modal from "../../common/Modal";
import {
    TeamMember,
    TeamMemberCount
} from "../../types/teamMembers";
import { useNavigate } from "react-router-dom";

const TeamMemberWidget = () => {
    const navigate = useNavigate();
    const [teamMembers, setTeamMembers] = useState<TeamMember[]>([]);
    const [filteredTeamMembers, setFilteredTeamMembers] = useState<TeamMember[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

    const [page, setPage] = useState(1);
    const [total, setTotal] = useState(0);

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [selectedTeamMembers, setSelectedTeamMembers] = useState<TeamMember[]>([]);

    useEffect(() => {
        const fetchTeamMembersList = async () => {
            try {
                const totalRes = await fetch(
                    "http://localhost:8080/api/team-members/count",
                    { credentials: "include" },
                );
                const countData: TeamMemberCount = await totalRes.json();
                setTotal(countData.count);

                const listRes = await fetch(
                    `http://localhost:8080/api/team-members/list/${page}`,
                    { credentials: "include" },
                );
                const data: TeamMember[] = await listRes.json();

                setTeamMembers(data);
                setFilteredTeamMembers(data);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                }
            } finally {
                setLoading(false);
            }
        };

        fetchTeamMembersList();

        const handleTabFocus = () => {
            fetchTeamMembersList();
        };

        window.addEventListener("focus", handleTabFocus);
        return () => window.removeEventListener("focus", handleTabFocus);

    }, [page]);

    const displayList =
        filteredTeamMembers.length === 0 ? teamMembers : filteredTeamMembers;

    const isSelected = (teamMember: TeamMember) =>
        selectedTeamMembers.some(
            (tm) => tm.team_member_uuid === teamMember.team_member_uuid
        );

    const toggleSelect = (teamMember: TeamMember) => {
        setSelectedTeamMembers((prev) =>
            isSelected(teamMember)
                ? prev.filter((tm) => tm.team_member_uuid !== teamMember.team_member_uuid)
                : [...prev, teamMember]
        );
    };

    const removeSelected = (uuid: string) => {
        setSelectedTeamMembers((prev) =>
            prev.filter((tm) => tm.team_member_uuid !== uuid)
        );
    };

    return (
        <div className="team-member-selector">
            {selectedTeamMembers.map((tm) => (
                <input
                    key={tm.team_member_uuid}
                    type="hidden"
                    name="teamMemberIds"
                    value={tm.team_member_uuid}
                />
            ))}

            <button type="button" onClick={() => setIsModalOpen(true)}>
                {selectedTeamMembers.length === 0
                    ? "Select Team Members"
                    : `${selectedTeamMembers.length} member(s) selected`}
            </button>

            {selectedTeamMembers.length > 0 && (
                <ul className="selected-team-members">
                    {selectedTeamMembers.map((tm) => (
                        <li key={tm.team_member_uuid}>
                            {tm.team_member_name}
                            <button
                                type="button"
                                onClick={() => removeSelected(tm.team_member_uuid)}
                            >
                                &times;
                            </button>
                        </li>
                    ))}
                </ul>
            )}

            <Modal
                isOpen={isModalOpen}
                onClose={() => setIsModalOpen(false)}
                title="Select Team Members"
            >
                {loading ? (
                    <p>Loading...</p>
                ) : error ? (
                    <p className="error">{error}</p>
                ) : (
                    <>
                        <ListingWidgetLayout
                            data={teamMembers}
                            onNewClick={() => window.open("/users/team-members/new")}
                            headers={["selected","No.", "Name", "Department", "Project Role"]}
                            newClickLabel="Create New Team Member"
                            onSearch={setFilteredTeamMembers}
                            currentPage={page}
                            total={total}
                            onPageChange={setPage}
                        >
                            {displayList.length === 0 ? (
                                <p>No Team Members.</p>
                            ) : (
                                displayList.map((teamMember, idx) => (
                                    <tr
                                        key={teamMember.team_member_uuid}
                                        onClick={() => toggleSelect(teamMember)}
                                        style={{
                                            cursor: "pointer",
                                            background: isSelected(teamMember)
                                                ? "#e6f4ff"
                                                : undefined,
                                        }}
                                    >
                                        <td>
                                            <input
                                                type="checkbox"
                                                checked={isSelected(teamMember)}
                                                readOnly
                                            />
                                        </td>
                                        <td>{idx + 1}</td>
                                        <td>{teamMember.team_member_name}</td>
                                        <td>{teamMember.team_member_department}</td>
                                        <td>{teamMember.team_member_role}</td>
                                    </tr>
                                ))
                            )}
                        </ListingWidgetLayout>

                        <button type="button" onClick={() => setIsModalOpen(false)}>
                            Done ({selectedTeamMembers.length} selected)
                        </button>
                    </>
                )}
            </Modal>
        </div>
    );
};

export default TeamMemberWidget;