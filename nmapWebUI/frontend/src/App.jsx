import React, { useState } from "react";
import "./App.css";

const App = () => {
	// States to manage form inputs
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [error] = useState("");

	// Handle form submission
	const HandleSubmit = (e) => {
		e.preventDefault();

		fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ username: username, password: password }),
		})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.json();
		})
		.then((data) => {
			console.log(data);
			// Handle successful login here
		})
		.catch((error) => {
			console.error('Error fetching data:', error);
		});
	};

	return (
		<div className="login-container">
			<h1>Nmap Management</h1><br />
			<h2>Login</h2>
			<form onSubmit={HandleSubmit} className="login-form">
				<div className="input-group">
					<label htmlFor="username">Username:</label>
					<input
						type="text"
						id="username"
						value={username}
						onChange={(e) => setUsername(e.target.value)}
						required
					/>
				</div>
				<div className="input-group">
					<label htmlFor="password">Password:</label>
					<input
						type="password"
						id="password"
						value={password}
						onChange={(e) => setPassword(e.target.value)}
						required
					/>
				</div>
				{error && <p className="error">{error}</p>}
				<button type="submit" className="submit-btn">Login</button>
			</form>
		</div>
	);
};

export default App;
