import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../../css/App.css";
import LoginPageLayout from "../../components/layouts/loginPageLayout";

const RegisterPage = (): JSX.Element => {
	const navigate = useNavigate();

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [repeatPassword, setRepeatPassword] = useState("");
	const [error] = useState("");

	const handleSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
		e.preventDefault();

		fetch('http://localhost:8080/register', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ username: username, password: password, repeatPassword: repeatPassword }),
		})
			.then((response) => {
				if (!response.ok) {
					throw new Error(`HTTP error! Status: ${response.status}`);
				}
				return response.json();
			})
			.then((data) => {
				console.log(data);
				if (data.status == 201) {
					navigate("/registerSuccess")
				}
			})
			.catch((error) => {
				console.error('Error fetching data:', error);
			});
	};

	return (
		<LoginPageLayout title="Register">
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
				<div className="input-group">
					<label htmlFor="repeatPassword">Repeat Password:</label>
					<input
						type="password"
						id="repeatPassword"
						value={repeatPassword}
						onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
							setRepeatPassword(e.target.value)
						}
						required
					/>
				</div>
				{error && <p className="error">{error}</p>}
				<button type="submit" className="submit-btn">Register</button>
				<button
					type="button"
					className="submit-btn"
					onClick={() => navigate("/login")}
				>
					Back to Login
				</button>
			</form>
		</LoginPageLayout>
	);
}
export default RegisterPage;