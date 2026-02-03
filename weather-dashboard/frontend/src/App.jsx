import { useEffect } from 'react';
import { zServer } from './store/server'

import Main from './components/main';
import Sidebar from './components/sidebar';
import FaveCities from './components/faveCities';

function App() {
	const zSetServerUrl = zServer(state => state.setServerUrl);

	useEffect(() => {
		zSetServerUrl('http://localhost:8080')
	}, []);

	return (
		<div className="h-screen max-h-screen bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center p-4 flex-col lg:flex-row">
			<Sidebar />
			<Main />
			<div 
				className="w-full p-4 bg-gray-700 graflex lg:hidden 
					rounded-bl-lg rounded-br-lg"
			>
				<FaveCities />
			</div>
		</div>
	)
}

export default App;