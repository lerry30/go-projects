import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

export const zForecast = create(
    devtools(
        persist((set) => ({
            cityName: '',
            country: '',
            sunrise: null,
            sunset: null,
            list: [],

            setWeatherForecast: (cityName, country, sunrise, sunset, list) => {
                if(!cityName || !country) return;
                if(!list || list?.length === 0) return;
                set((state) => ({ cityName, country, sunrise, sunset, list }));
            },
        }), {
            name: 'store-weather-forecast'
        })
    )
);