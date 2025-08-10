import { useState } from "react";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";
import ProtectedLayout from "../../../components/layouts/protectedLayout";

const ScanUpload = () => {
	const [file, setFile] = useState(null);

	const handleFileChange = (e) => {
		setFile(e.target.files[0]);  // Only taking the first file
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		if (!file) {
			alert("Please select a file to upload.");
			return;
		}

		const formData = new FormData();
		formData.append("file", file); // key name 'file' should match what your backend expects

		try {
			const response = await fetch("http://localhost:8080/scans/upload", {
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
			<HeadMetadata title={"Scan Upload"}/>
			<form onSubmit={handleSubmit}>
				<input type="file" onChange={handleFileChange} accept=".xml" />
				<button type="submit">Upload</button>
			</form>
		</ProtectedLayout>
	);
};

export default ScanUpload;