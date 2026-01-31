import { Helmet } from "react-helmet-async";

interface HeadMetadataProps {
  title?: string;
}

const HeadMetadata = ({ title }: HeadMetadataProps): JSX.Element => {
  const pageTitle = title
    ? `${title} | Guln Vulnerability Management`
    : "Guln Vulnerability Management";

  const description = "Web UI to manage vulnerabilities.";
  const url = import.meta.env.VITE_APP_URL ?? "http://localhost:5173";

  return (	
    <Helmet>
      <title>{pageTitle}</title>

      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <meta charSet="UTF-8" />
      <meta name="keywords" content="VAPT, Vulnerability Management, tools" />
      <meta name="author" content="Guln Vulnerability Management" />
      <meta name="description" content={description} />

      {/* Open Graph */}
      <meta property="og:title" content={pageTitle} />
      <meta property="og:description" content={description} />
      <meta property="og:type" content="website" />
      <meta property="og:url" content={url} />
      <meta property="og:image" content={`${url}/og-image.jpg`} />

      {/* Twitter */}
      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:title" content={pageTitle} />
      <meta name="twitter:description" content={description} />
      <meta name="twitter:image" content={`${url}/og-image.jpg`} />
    </Helmet>
  );
};

export default HeadMetadata;
