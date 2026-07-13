import { useEffect } from 'react';
import { zServer } from './store/server'
import { useCurrentCity } from './hooks/useCurrentCity';
import { zForecast } from './store/weather-forecast';

import Main from './components/main';
import Sidebar from './components/sidebar';
import FaveCities from './components/faveCities';

function App() {
	const zSetServerUrl = zServer(state => state.setServerUrl);
    const zFetchWeatherForecast = zForecast(state => state.fetchWeatherForecast);

	// Get the current/nearest city, so by default it will display something in the dashboard
	const { city, status, error } = useCurrentCity();

	useEffect(() => {
		const serverUrl = 'http://localhost:8080';
		zSetServerUrl(serverUrl);

		// ------------------------------------------------------

		if(status === 'loading') console.log('Finding your location...');
		if(error === 'error') {
			console.log(`Couldn't get location: ${error}`);
			return;
		}
		if(!city) return;

		(async () => await zFetchWeatherForecast(serverUrl, city))();
	}, [city]);

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