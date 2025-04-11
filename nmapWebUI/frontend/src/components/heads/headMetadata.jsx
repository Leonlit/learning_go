// components/HeadMetadata.jsx
import { Helmet } from 'react-helmet-async';

const HeadMetadata = ({ title }) => {
	const pageTitle = title ? `${title} | Nmap Management` : 'Nmap Management';
	const description = "Web UI to manage Nmap scan results.";
	const url = "http://localhost:5173"; // Can be updated to dynamic env var

	console.log(pageTitle);

	return (
		<Helmet>
			<title data-react-helmet="true">{pageTitle}</title>
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<meta charSet="UTF-8" />
			<meta name="keywords" content="Nmap, management, tools" />
			<meta name="author" content="Your Name or Brand" />
			<meta name="description" content={description} />

			{/* Open graph */}
			<meta property="og:title" content={pageTitle} data-react-helmet="true" />
			<meta property="og:description" content={description} />
			<meta property="og:type" content="website" />
			<meta property="og:url" content={url} />
			<meta property="og:image" content={`${url}/og-image.jpg`} />

			{/* Twitter */}
			<meta name="twitter:card" content="summary_large_image" />
			<meta name="twitter:title" content={pageTitle} data-react-helmet="true" />
			<meta name="twitter:description" content={description} />
			<meta name="twitter:image" content={`${url}/og-image.jpg`} />
		</Helmet>
	);
};

export default HeadMetadata;