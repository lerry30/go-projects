import { create } from 'zustand';
import { persist } from 'zustand/middleware';

export const zForecast = create(
    persist((set, get) => ({
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
        name: 'store-weather-forecast',
    })
);