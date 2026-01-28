import { useEffect } from 'react';
import { zServer } from './store/server'

import Main from './components/main';
import Sidebar from './components/sidebar';

function App() {
	const zSetServerUrl = zServer(state => state.setServerUrl);

	useEffect(() => {
		zSetServerUrl('http://localhost:8080')
	}, []);

	return (
		<div className="h-screen max-h-screen bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center p-4 sm:flex-col lg:flex-row">
			<Sidebar />
			<Main />
		</div>
	)
}

export default App;