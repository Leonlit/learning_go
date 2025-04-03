import React, { useState } from "react";
import "./App.css";

const App = () => {
	// States to manage form inputs
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [error] = useState("");

	// Handle form submission
	const handleSubmit = (e) => {
		e.preventDefault();

	};

	return (
		<div className="login-container">
			<h1>Nmap Management</h1><br/>
			<h2>Login</h2>
			<form onSubmit={handleSubmit} className="login-form">
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
