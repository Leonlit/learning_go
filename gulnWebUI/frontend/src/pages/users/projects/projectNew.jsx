import { useState } from "react";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";
// import { useNavigate } from "react-router-dom";

const ScanUpload = () => {
    const [projectName, setProjectName] = useState("");

    // const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!projectName.trim()) {
            alert("Please provide a project name.");
            return;
        }

        const formData = new FormData(e.target);

        try {
            const response = await fetch("http://localhost:8080/projects/new", {
                method: "POST",
                credentials: "include",
                body: formData,
            });

            if (response.ok) {
                const res = await response.json();
                console.log(res);
                alert("Project created successfully!");
                // navigate("/users/projects");
            } else {
                alert("Project creation failed.");
            }
        } catch (err) {
            console.error(err);
            alert("Error creating project.");
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
                    onChange={(e) => setProjectName(e.target.value)}
                />
                <button type="submit">Create New Project</button>
            </form>
        </ProtectedLayout>
    );
};

export default ScanUpload;
