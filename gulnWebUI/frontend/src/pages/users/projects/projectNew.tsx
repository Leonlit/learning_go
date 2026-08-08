import { useState } from "react";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";
import TeamMemberWidget from "../../../components/widget/teamMemberWidget"
import { useNavigate } from "react-router-dom";
import { 
    NewProjectResponse 
} from "../../../types/projects"; 

const createNewProject = () => {
    const [projectName, setProjectName] = useState<string>("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!projectName.trim()) {
            alert("Please provide a project name.");
            return;
        }

        const formData = new FormData(e.currentTarget);

        try {
            const response = await fetch("http://localhost:8080/api/projects/new", {
                method: "POST",
                credentials: "include",
                body: formData,
            });

            if (!response.ok) {
                throw new Error("Project creation failed.");
            }

            const res: NewProjectResponse = await response.json();
            console.log(res);

            alert("Project created successfully!");
            navigate("/users/projects/edit/" + res.projectID);
        } catch (err) {
            if (err instanceof Error) {
                console.error(err.message);
                alert(err.message);
            }
        }
    };

    return (
        <ProtectedLayout>
            <HeadMetadata title={"Create New Project"} />

            <h2>Create New Project</h2>

            <form onSubmit={handleSubmit}>
                <label>Project Name</label>

                <input
                    type="text"
                    name="projectName"
                    value={projectName}
                    onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                        setProjectName(e.target.value)
                    }
                />

                <label>Team Member:</label>

                <TeamMemberWidget/>

                <button type="submit">Create New Project</button>
            </form>
        </ProtectedLayout>
    );
};

export default createNewProject;