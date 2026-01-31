import { ReactNode } from "react";
import HeadMetadata from "../heads/headMetadata";

interface NormalPageLayoutProps {
	title: string;
	children: ReactNode;
}

const LoginPageLayout = ({
	title,
	children
}: NormalPageLayoutProps): JSX.Element => {
	return (
		<>
			<HeadMetadata title={title} />
			<main className="public-container">
				<div className="login-container">
				<h1>Guln Vulnerability Management</h1>
				<br />
				<h2>{title}</h2>
				{children}
				</div>
			</main>
		</>
	);
};

export default LoginPageLayout;
