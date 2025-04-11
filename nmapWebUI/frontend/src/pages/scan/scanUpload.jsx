import React, { useState } from "react";
import "../../css/App.css";
import { useNavigate } from "react-router-dom";
import HeadMetadata from "../../components/heads/headMetadata";

const LoginPage = () => {
	// States to manage form inputs
	const navigate = useNavigate();

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [error, setError] = useState("");

	const HandleSubmit = (e) => {
		e.preventDefault();

		fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ username, password }),
		})
		.then((response) => {
			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}
			return response.json();
		})
		.then((data) => {
			console.log(data);
			// You can add a success state or redirect after successful login
		})
		.catch((error) => {
			console.error('Error fetching data:', error);
			setError("Invalid username or password.");
		});
	};

	return (
		<div>
			<HeadMetadata title={"Login"}/>
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
					<button
						type="button"
						className="submit-btn"
						onClick={() => navigate("/register")}
					>
						Register
					</button>
				</form>
			</div>
		</div>
	);
};

export default LoginPage;
