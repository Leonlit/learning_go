import { useState } from "react";
import "../../../css/App.css";
import HeadMetadata from "../../../components/heads/headMetadata";

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
			const response = await fetch("/api/upload", {
				method: "POST",
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
		<>
			<HeadMetadata title={"Scan Upload"}/>
			<form onSubmit={handleSubmit}>
				<input type="file" onChange={handleFileChange} />
				<button type="submit">Upload</button>
			</form>
		</>
	);
};

export default ScanUpload;