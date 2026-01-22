import { useState, useEffect } from 'react';
import { zForecast } from '../store/weather-forecast';

const Main = () => {
    const zForecastCityName = zForecast(state => state.cityName);
    const zForecastCountry = zForecast(state => state.country);
    const zForecastSunrise = zForecast(state => state.sunrise);
    const zForecastSunset = zForecast(state => state.sunset);
    const zForecastList = zForecast(state => state.list);

    useEffect(() => {

    }, []);

    return (
        <main className="w-[70%] h-full bg-red-300 overflow-hidden rounded-tr-lg rounded-br-lg p-4">
            <div className="bg-red-400 p-2">
                <span className="text-xl">{zForecastCityName} City, {zForecastCountry}</span>
            </div>
            <div className="bg-blue-400 px-2 py-6">
                <span className="text-sm">Weather Forecast</span>
            </div>
            <div className="">
                {
                    zForecastList.map((fc, i) => {
                        const updates = fc?.hour_weather_updates[0];

                        return (
                            <div key={i}>
                                <span>{fc?.weekday}</span>
                                <img src={updates.icon} alt="Weather Icon" />
                            </div>
                        )
                    })
                }
            </div>
        </main>
    )
}

export default Main;