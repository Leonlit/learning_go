import { useNavigate } from "react-router-dom";
import "../../css/App.css";
import LoginPageLayout from "../../components/layouts/loginPageLayout";

const RegisterPage = (): JSX.Element => {
	const navigate = useNavigate();

	return (
		<LoginPageLayout title="Register Successful">
			<form className="login-form">
				<button
					type="button"
					className="submit-btn"
					onClick={() => navigate("/")}
				>
					Back to Login
				</button>
			</form>
		</LoginPageLayout>
	);
};

export default RegisterPage;
