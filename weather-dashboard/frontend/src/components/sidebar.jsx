import { Search } from 'lucide-react';
import { useState, useRef } from 'react';

import { getData } from '../utils/send';
import { zServer } from '../store/server';
import { zForecast } from '../store/weather-forecast';

const dummy = {
    city_name: "City of Gapan",
    country: "PH",
    sunrise: "2026-01-21T06:26:11+08:00",
    sunset: "2026-01-21T17:47:54+08:00",
    list: {
        Friday: [
            {
                temp: "23.32°C",
                feels_like: "23.51°C",
                humidity: "69%",
                pressure: "1014% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "7%",
                wind_speed: "6.63 m/s",
                wind_direction: "93°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-23 09:00:00",
                local_date_time: "2026-01-23T17:00:00+08:00"
            },
            {
                temp: "21.16°C",
                feels_like: "21.24°C",
                humidity: "73%",
                pressure: "1017% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "4%",
                wind_speed: "4.86 m/s",
                wind_direction: "78°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-23 12:00:00",
                local_date_time: "2026-01-23T20:00:00+08:00"
            },
            {
                temp: "19.18°C",
                feels_like: "19.32°C",
                humidity: "83%",
                pressure: "1017% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "5%",
                wind_speed: "1.81 m/s",
                wind_direction: "13°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-23 15:00:00",
                local_date_time: "2026-01-23T23:00:00+08:00"
            },
            {
                temp: "18.10°C",
                feels_like: "18.26°C",
                humidity: "88%",
                pressure: "1016% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "4%",
                wind_speed: "1.62 m/s",
                wind_direction: "348°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-23 18:00:00",
                local_date_time: "2026-01-24T02:00:00+08:00"
            },
            {
                temp: "17.21°C",
                feels_like: "17.36°C",
                humidity: "91%",
                pressure: "1016% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "0%",
                wind_speed: "0.92 m/s",
                wind_direction: "9°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-23 21:00:00",
                local_date_time: "2026-01-24T05:00:00+08:00"
            },
            {
                temp: "20.34°C",
                feels_like: "20.47°C",
                humidity: "78%",
                pressure: "1018% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "0%",
                wind_speed: "1.10 m/s",
                wind_direction: "353°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-24 00:00:00",
                local_date_time: "2026-01-24T08:00:00+08:00"
            },
            {
                temp: "25.66°C",
                feels_like: "25.80°C",
                humidity: "58%",
                pressure: "1017% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03d",
                clouds_cover: "30%",
                wind_speed: "1.98 m/s",
                wind_direction: "14°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-24 03:00:00",
                local_date_time: "2026-01-24T11:00:00+08:00"
            },
            {
                temp: "26.42°C",
                feels_like: "26.42°C",
                humidity: "55%",
                pressure: "1014% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03d",
                clouds_cover: "50%",
                wind_speed: "3.71 m/s",
                wind_direction: "93°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-24 06:00:00",
                local_date_time: "2026-01-24T14:00:00+08:00"
            }
        ],
        Saturday: [
            {
                temp: "23.32°C",
                feels_like: "23.46°C",
                humidity: "67%",
                pressure: "1014% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "56%",
                wind_speed: "6.60 m/s",
                wind_direction: "92°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-24 09:00:00",
                local_date_time: "2026-01-24T17:00:00+08:00"
            },
            {
                temp: "21.18°C",
                feels_like: "21.21°C",
                humidity: "71%",
                pressure: "1016% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "27%",
                wind_speed: "5.25 m/s",
                wind_direction: "76°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-24 12:00:00",
                local_date_time: "2026-01-24T20:00:00+08:00"
            },
            {
                temp: "19.09°C",
                feels_like: "19.17°C",
                humidity: "81%",
                pressure: "1017% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "0%",
                wind_speed: "2.18 m/s",
                wind_direction: "9°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-24 15:00:00",
                local_date_time: "2026-01-24T23:00:00+08:00"
            },
            {
                temp: "17.38°C",
                feels_like: "17.50°C",
                humidity: "89%",
                pressure: "1016% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "34%",
                wind_speed: "1.22 m/s",
                wind_direction: "335°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-24 18:00:00",
                local_date_time: "2026-01-25T02:00:00+08:00"
            },
            {
                temp: "17.09°C",
                feels_like: "17.15°C",
                humidity: "88%",
                pressure: "1015% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "98%",
                wind_speed: "1.52 m/s",
                wind_direction: "352°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-24 21:00:00",
                local_date_time: "2026-01-25T05:00:00+08:00"
            },
            {
                temp: "19.90°C",
                feels_like: "19.98°C",
                humidity: "78%",
                pressure: "1016% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "97%",
                wind_speed: "1.07 m/s",
                wind_direction: "322°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-25 00:00:00",
                local_date_time: "2026-01-25T08:00:00+08:00"
            },
            {
                temp: "26.31°C",
                feels_like: "26.31°C",
                humidity: "58%",
                pressure: "1016% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "74%",
                wind_speed: "2.32 m/s",
                wind_direction: "358°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-25 03:00:00",
                local_date_time: "2026-01-25T11:00:00+08:00"
            },
            {
                temp: "28.69°C",
                feels_like: "29.16°C",
                humidity: "49%",
                pressure: "1013% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "74%",
                wind_speed: "1.89 m/s",
                wind_direction: "74°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-25 06:00:00",
                local_date_time: "2026-01-25T14:00:00+08:00"
            }
        ],
        Sunday: [
            {
                temp: "25.32°C",
                feels_like: "25.60°C",
                humidity: "65%",
                pressure: "1013% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "62%",
                wind_speed: "3.22 m/s",
                wind_direction: "121°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-25 09:00:00",
                local_date_time: "2026-01-25T17:00:00+08:00"
            },
            {
                temp: "23.15°C",
                feels_like: "23.45°C",
                humidity: "74%",
                pressure: "1015% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "81%",
                wind_speed: "3.22 m/s",
                wind_direction: "69°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-25 12:00:00",
                local_date_time: "2026-01-25T20:00:00+08:00"
            },
            {
                temp: "22.28°C",
                feels_like: "22.65°C",
                humidity: "80%",
                pressure: "1015% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "100%",
                wind_speed: "2.51 m/s",
                wind_direction: "356°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-25 15:00:00",
                local_date_time: "2026-01-25T23:00:00+08:00"
            },
            {
                temp: "20.27°C",
                feels_like: "20.62°C",
                humidity: "87%",
                pressure: "1014% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "85%",
                wind_speed: "2.03 m/s",
                wind_direction: "330°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-25 18:00:00",
                local_date_time: "2026-01-26T02:00:00+08:00"
            },
            {
                temp: "19.89°C",
                feels_like: "20.26°C",
                humidity: "89%",
                pressure: "1014% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "84%",
                wind_speed: "1.60 m/s",
                wind_direction: "351°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-25 21:00:00",
                local_date_time: "2026-01-26T05:00:00+08:00"
            },
            {
                temp: "22.12°C",
                feels_like: "22.50°C",
                humidity: "81%",
                pressure: "1015% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "92%",
                wind_speed: "2.02 m/s",
                wind_direction: "343°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-26 00:00:00",
                local_date_time: "2026-01-26T08:00:00+08:00"
            },
            {
                temp: "27.52°C",
                feels_like: "29.05°C",
                humidity: "63%",
                pressure: "1015% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "100%",
                wind_speed: "1.97 m/s",
                wind_direction: "335°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-26 03:00:00",
                local_date_time: "2026-01-26T11:00:00+08:00"
            },
            {
                temp: "28.75°C",
                feels_like: "30.62°C",
                humidity: "60%",
                pressure: "1012% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "100%",
                wind_speed: "1.80 m/s",
                wind_direction: "112°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-26 06:00:00",
                local_date_time: "2026-01-26T14:00:00+08:00"
            }
        ],
        Thursday: [
            {
                temp: "23.75°C",
                feels_like: "23.90°C",
                humidity: "66%",
                pressure: "1013% hPa",
                description: "few clouds",
                icon: "http://openweathermap.org/img/w/02d",
                clouds_cover: "12%",
                wind_speed: "4.94 m/s",
                wind_direction: "98°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-22 09:00:00",
                local_date_time: "2026-01-22T17:00:00+08:00"
            },
            {
                temp: "21.81°C",
                feels_like: "21.95°C",
                humidity: "73%",
                pressure: "1015% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "33%",
                wind_speed: "6.10 m/s",
                wind_direction: "86°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-22 12:00:00",
                local_date_time: "2026-01-22T20:00:00+08:00"
            },
            {
                temp: "19.99°C",
                feels_like: "20.16°C",
                humidity: "81%",
                pressure: "1016% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "41%",
                wind_speed: "2.20 m/s",
                wind_direction: "30°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-22 15:00:00",
                local_date_time: "2026-01-22T23:00:00+08:00"
            },
            {
                temp: "18.36°C",
                feels_like: "18.58°C",
                humidity: "89%",
                pressure: "1016% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "44%",
                wind_speed: "1.74 m/s",
                wind_direction: "359°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-22 18:00:00",
                local_date_time: "2026-01-23T02:00:00+08:00"
            },
            {
                temp: "17.76°C",
                feels_like: "17.89°C",
                humidity: "88%",
                pressure: "1016% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01n",
                clouds_cover: "2%",
                wind_speed: "1.72 m/s",
                wind_direction: "18°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-22 21:00:00",
                local_date_time: "2026-01-23T05:00:00+08:00"
            },
            {
                temp: "20.51°C",
                feels_like: "20.63°C",
                humidity: "77%",
                pressure: "1018% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "3%",
                wind_speed: "1.63 m/s",
                wind_direction: "10°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-23 00:00:00",
                local_date_time: "2026-01-23T08:00:00+08:00"
            },
            {
                temp: "25.78°C",
                feels_like: "25.82°C",
                humidity: "54%",
                pressure: "1018% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "5%",
                wind_speed: "2.41 m/s",
                wind_direction: "40°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-23 03:00:00",
                local_date_time: "2026-01-23T11:00:00+08:00"
            },
            {
                temp: "27.17°C",
                feels_like: "27.58°C",
                humidity: "50%",
                pressure: "1014% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "10%",
                wind_speed: "3.78 m/s",
                wind_direction: "87°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-23 06:00:00",
                local_date_time: "2026-01-23T14:00:00+08:00"
            }
        ],
        Wednesday: [
            {
                temp: "27.33°C",
                feels_like: "29.33°C",
                humidity: "69%",
                pressure: "1010% hPa",
                description: "clear sky",
                icon: "http://openweathermap.org/img/w/01d",
                clouds_cover: "4%",
                wind_speed: "6.31 m/s",
                wind_direction: "98°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-21 09:00:00",
                local_date_time: "2026-01-21T17:00:00+08:00"
            },
            {
                temp: "26.55°C",
                feels_like: "26.55°C",
                humidity: "71%",
                pressure: "1011% hPa",
                description: "few clouds",
                icon: "http://openweathermap.org/img/w/02n",
                clouds_cover: "15%",
                wind_speed: "4.23 m/s",
                wind_direction: "80°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-21 12:00:00",
                local_date_time: "2026-01-21T20:00:00+08:00"
            },
            {
                temp: "24.98°C",
                feels_like: "25.60°C",
                humidity: "79%",
                pressure: "1013% hPa",
                description: "scattered clouds",
                icon: "http://openweathermap.org/img/w/03n",
                clouds_cover: "40%",
                wind_speed: "2.25 m/s",
                wind_direction: "16°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-21 15:00:00",
                local_date_time: "2026-01-21T23:00:00+08:00"
            },
            {
                temp: "22.27°C",
                feels_like: "22.67°C",
                humidity: "81%",
                pressure: "1013% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04n",
                clouds_cover: "52%",
                wind_speed: "2.39 m/s",
                wind_direction: "29°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-21 18:00:00",
                local_date_time: "2026-01-22T02:00:00+08:00"
            },
            {
                temp: "21.20°C",
                feels_like: "21.52°C",
                humidity: "82%",
                pressure: "1014% hPa",
                description: "few clouds",
                icon: "http://openweathermap.org/img/w/02n",
                clouds_cover: "14%",
                wind_speed: "2.28 m/s",
                wind_direction: "41°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "night",
                date_time: "2026-01-21 21:00:00",
                local_date_time: "2026-01-22T05:00:00+08:00"
            },
            {
                temp: "22.66°C",
                feels_like: "22.91°C",
                humidity: "74%",
                pressure: "1016% hPa",
                description: "few clouds",
                icon: "http://openweathermap.org/img/w/02d",
                clouds_cover: "19%",
                wind_speed: "2.18 m/s",
                wind_direction: "28°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-22 00:00:00",
                local_date_time: "2026-01-22T08:00:00+08:00"
            },
            {
                temp: "26.09°C",
                feels_like: "26.09°C",
                humidity: "54%",
                pressure: "1015% hPa",
                description: "overcast clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "98%",
                wind_speed: "3.63 m/s",
                wind_direction: "55°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-22 03:00:00",
                local_date_time: "2026-01-22T11:00:00+08:00"
            },
            {
                temp: "27.25°C",
                feels_like: "27.73°C",
                humidity: "51%",
                pressure: "1013% hPa",
                description: "broken clouds",
                icon: "http://openweathermap.org/img/w/04d",
                clouds_cover: "52%",
                wind_speed: "3.88 m/s",
                wind_direction: "85°",
                visibility: "10.00 km",
                chance_of_rain_or_snow: "0.00%",
                part_of_day: "day",
                date_time: "2026-01-22 06:00:00",
                local_date_time: "2026-01-22T14:00:00+08:00"
            }
        ]
    }
};

const Sidebar = () => {
    const zServerUrl = zServer(state => state.serverUrl);
    const zSetWeatherForecastData = zForecast(state => state.setWeatherForecast);

    // Using useRef to avoid unnecessary re-renders
    const cityInputField = useRef(null);

    const [ isSuggestionOpen, setIsSuggestionOpen ] = useState(false);
    const [ suggestionData, setSuggestionData ] = useState([]);

    const openSuggestion = async (city) => {
        const cityInput = city.trim();
        if(cityInput) {
            try {
                const response = await getData(`${zServerUrl}/filter/${cityInput}`);
                if(response && response?.data?.length > 0) {
                    setSuggestionData(response.data);
                    setIsSuggestionOpen(true);
                }
            } catch(error) {
                console.log(error)
            }
        } else {
            setIsSuggestionOpen(false);
        }
    }

    const getWeatherForecast = async (city) => {
        try {
            if(suggestionData?.length == 0) {
                throw new Error('City not found.')
            }
            
            city = city.trim().toLowerCase();
            if(!suggestionData.includes(city)) {
                const firstCity = suggestionData[0];
                if(city.length > firstCity.length) {
                    throw new Error('City not found.');
                }
                // this is the only purpose for this 
                // useRef, to autofill the input field
                cityInputField.current.value = firstCity;
                city = firstCity;
            } else {
                cityInputField.current.value = city;
            }
            
            setIsSuggestionOpen(false);

            zSetWeatherForecastData(
                dummy.city_name,
                dummy.country,
                dummy.sunrise,
                dummy.sunset,
                dummy.list,
            );

            //const response = await getData(`${zServerUrl}/forecast/${city}`);
            //if(response) {
            //    zSetWeatherForecastData(
            //        response?.city_name,
            //        response?.country,
            //        new Date(response?.sunrise),
            //        new Date(response?.sunset),
            //        response?.list
            //    );
            //}
        } catch(error) {
            console.log(error);
        }
    }

    const BoxSuggestion = () => {
        if(!isSuggestionOpen) {
            return null;
        }

        return (
            <div className="w-full max-h-[50vh] overflow-y-auto overflow-x-clip bg-slate-300 absolute top-full left-0 pb-4 px-4 rounded-b-xl">
                {suggestionData.map((s, i) => (
                    <div key={i} onClick={() => {
                        getWeatherForecast(s);
                    }}>
                        {s}
                    </div>
                ))}
            </div>
        )
    }

    return (
        <aside className="w-[30%] h-full bg-white overflow-hidden p-4 rounded-tl-lg rounded-bl-lg">
            <div className={`w-full bg-slate-300 flex py-3 px-4 relative ${isSuggestionOpen ? "rounded-t-xl" : "rounded-xl"}`}>
                <input
                    ref={cityInputField}
                    className="w-full outline-0 font-sans"
                    placeholder="Search city"
                    onChange={elem => openSuggestion(elem.target.value)}
                    onKeyDown={e => e.key === "Enter" && getWeatherForecast(e.target.value)}
                />
                <Search />
                <BoxSuggestion />
            </div>
        </aside>
    )
}

export default Sidebar;