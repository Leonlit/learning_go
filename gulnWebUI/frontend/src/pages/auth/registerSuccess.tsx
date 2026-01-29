import React from "react";
import { useNavigate } from "react-router-dom";
import "../../css/App.css";

function RegisterPage () {
    const navigate = useNavigate();
    
    return  (
		<NormalPageLayout title="Register Successful">
			<div className="login-container">
				<h1>Guln Vulnerability Management</h1><br />
				<h2>Registration Successful</h2>
				<form className="login-form">
					<button 
						type="button"
						className="submit-btn"
						onClick={() => navigate("/")}
					>
					Back to Login
					</button>
				</form>
			</div>
		</NormalPageLayout>
	);
}
export default RegisterPage;