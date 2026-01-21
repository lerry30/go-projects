import { useState, useEffect } from 'react';
import { zForecast } from '../store/weather-forecast';

const Main = () => {
    const zForecastCityName = zForecast(state => state.cityName);
    const zForecastCountry = zForecast(state => state.country);
    const zForecastSunrise = zForecast(state => state.sunrise);
    const zForecastSunset = zForecast(state => state.sunset);
    const zForecastList = zForecast(state => state.list);

    const [ forecastList, setForecastList ] = useState([]);

    useEffect(() => {
        
    }, []);

    return (
        <main className="w-[70%] h-full bg-red-300 overflow-hidden rounded-tr-lg rounded-br-lg">

        </main>
    )
}

export default Main;