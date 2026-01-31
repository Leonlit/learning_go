import { useState } from "react";
import "../../css/App.css";
import { useNavigate } from "react-router-dom";
import LoginPageLayout from "../../components/layouts/loginPageLayout";

const LoginPage = (): JSX.Element => {
	const navigate = useNavigate();

	const [username, setUsername] = useState<string>("");
	const [password, setPassword] = useState<string>("");
	const [error, setError] = useState<string>("");

	const handleSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
		e.preventDefault();

		fetch("http://localhost:8080/login", {
			method: "POST",
			credentials: "include",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({ username, password })
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`HTTP error! Status: ${response.status}`);
				}
				return response.json();
			})
			.then(() => {
				navigate("/users/dashboard");
			})
			.catch((err: unknown) => {
				console.error("Error fetching data:", err);
				setError("Invalid username or password.");
			});
	};

	return (
		<div>
			<LoginPageLayout title="Login">
				<form onSubmit={handleSubmit} className="login-form">
					<div className="input-group">
						<label htmlFor="username">Username:</label>
						<input
							type="text"
							id="username"
							value={username}
							onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
								setUsername(e.target.value)
							}
							required
						/>
					</div>

					<div className="input-group">
						<label htmlFor="password">Password:</label>
						<input
							type="password"
							id="password"
							value={password}
							onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
								setPassword(e.target.value)
							}
							required
						/>
					</div>

					{error && <p className="error">{error}</p>}

					<button type="submit" className="submit-btn">
						Login
					</button>

					<button
						type="button"
						className="submit-btn"
						onClick={() => navigate("/register")}
					>
						Register
					</button>
				</form>
			</LoginPageLayout>
		</div>
	);
};

export default LoginPage;
