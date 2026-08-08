import { useState } from "react";
import { useParams, useLocation } from "react-router-dom";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";

const ProjectUpload = () => {
	const [file, setFile] = useState<File | null>(null);
    const [scanName, setScanName] = useState("");
    const { projectUUID } = useParams();
	const {state} = useLocation()

	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const selectedFile = e.target.files?.[0]; // safely handles null FileList
		if (selectedFile) {
			setFile(selectedFile);
		}
	};

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();

		if (!file) {
			alert("Please select a file to upload.");
			return;
		}

        const formData = new FormData(e.currentTarget);

		try {
			const response = await fetch("http://localhost:8080/projects/upload/" + projectUUID, {
				method: "POST",
				credentials: "include",
				body: formData,
			});

			if (response.ok) {
				alert("File uploaded successfully!");
			} else {
				alert("Upload failed.");
			}
		} catch (err) {
			console.error(err);
			alert("Error uploading file.");
		}
	};

	return (
		<ProtectedLayout>
			<HeadMetadata title={state.project_name + " - Project Scan Upload"}/>
			<form onSubmit={handleSubmit}>
				<h2>{state.project_name + " - Scan Upload"}</h2>
                <label>Scan Name</label>
                <input
                    type="text"
                    name="scanName"
                    value={scanName}
                    onChange={(e) => setScanName(e.target.value)}
                />
                <label> File Upload</label>
				<input type="file" name="file" onChange={handleFileChange} accept=".xml" />
				<button type="submit">Upload</button>
			</form>
		</ProtectedLayout>
	);
};

export default ProjectUpload;