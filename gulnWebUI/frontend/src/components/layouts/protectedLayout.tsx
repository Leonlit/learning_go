import { ReactNode } from "react";

interface ProtectedLayoutProps {
	children: ReactNode;
}

const handleLogout = async (): Promise<void> => {
	await fetch("http://localhost:8080/api/logout", {
		method: "POST",
		credentials: "include"
	});

	window.location.href = "/";
};

const ProtectedLayout = ({ children }: ProtectedLayoutProps): JSX.Element => {
	return (
		<div className="app-container">
			<header className="header">
				<h1>Guln Vulnerability Management</h1>

				<nav className="header-nav">
					<button>
						<a href="/users/settings">Settings</a>
					</button>
					<button onClick={handleLogout}>Logout</button>
				</nav>
			</header>

			<div className="main-content">
				<aside className="sidebar">
					<p>
						<a href="/users/dashboard">Dashboard</a>
					</p>
					<p>
						<a href="/users/projects">Projects</a>
					</p>
				</aside>

				<main className="content-area">{children}</main>
			</div>

			<footer className="footer">
				<p>© 2026 Guln Vulnerability Management Web UI</p>
			</footer>
		</div>
	);
};

export default ProtectedLayout;
