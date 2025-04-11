import { HelmetProvider } from 'react-helmet-async'; 
import { BrowserRouter as Router} from "react-router-dom";
import { createRoot } from 'react-dom/client'
import App from './App';

createRoot(document.getElementById('root')).render(
		<HelmetProvider>
			<Router>
				<App/>
			</Router>
		</HelmetProvider>
)
