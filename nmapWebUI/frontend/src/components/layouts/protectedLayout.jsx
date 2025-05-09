const ProtectedLayout = ({ children }) => {
	return (
		<div className="app-container">
			<header className="header"> {/* Placeholder for Nav */}
				<h1>Nmap Management</h1>
                <nav className="header-nav">
					<a href="/users/settings">Settings</a>
					<button onClick={handleLogout}>Logout</button>
				</nav>
			</header>

			<div className="main-content">
				<aside className="sidebar"> {/* Placeholder for Sidebar */}
					<p>Sidebar</p>
				</aside>

				<main className="content-area">
					{children}
				</main>
			</div>

			<footer className="footer"> {/* Placeholder for Footer */}
				<p>© 2025 Nmap Management Web UI</p>
			</footer>
		</div>
	);
};

// Function to handle logout (clear cookies, local storage, etc.)
const handleLogout = async () => {
    await fetch('http://localhost:8080/logout', {
      method: 'POST',
      credentials: 'include', // Ensure cookies are sent
    });
  };
export default ProtectedLayout;