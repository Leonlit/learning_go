import { useState } from "react";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";
import { useNavigate } from "react-router-dom";

type NewProjectResponse = {
    projectID: string;
};

const createNewTeamMember = () => {
    const [teamMemberName, setTeamMemberName] = useState<string>("");
    const [teamMemberDepartment, setTeamMemberDepartment] =
        useState<string>("");
    const [teamMemberRole, setTeamMemberRole] = useState<string>("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!teamMemberName.trim()) {
            alert("Please provide a team member name.");
            return;
        }

        if (!teamMemberDepartment.trim()) {
            alert("Please provide a team member department.");
            return;
        }

        if (!teamMemberRole.trim()) {
            alert("Please provide a team member role.");
            return;
        }

        try {
            const response = await fetch(
                "http://localhost:8080/api/teamMembers/new",
                {
                    method: "POST",
                    credentials: "include",
                    body: JSON.stringify({
                        teamMemberName,
                        teamMemberDepartment,
                        teamMemberRole,
                    }),
                },
            );

            if (!response.ok) {
                throw new Error("Team Member creation failed.");
            }

            const res: NewProjectResponse = await response.json();
            console.log(res);

            alert("Project created successfully!");
            navigate("/users/team-member/info/" + res.projectID);
        } catch (err) {
            if (err instanceof Error) {
                console.error(err.message);
                alert(err.message);
            }
        }
    };

    return (
        <ProtectedLayout>
            <HeadMetadata title={"Create New Team Member"} />

            <h2>Create New Team Member</h2>

            <form onSubmit={handleSubmit}>
                <label>Name</label>

                <input
                    type="text"
                    name="teamMemberName"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setTeamMemberName(e.target.value)
                    }
                />

                <label>Department</label>

                <input
                    type="text"
                    name="teamMemberDepartment"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setTeamMemberDepartment(e.target.value)
                    }
                />

                <label>Role</label>

                <input
                    type="text"
                    name="teamMemberRole"
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setTeamMemberRole(e.target.value)
                    }
                />
                <button type="submit">Create</button>
            </form>
        </ProtectedLayout>
    );
};

export default createNewTeamMember;
