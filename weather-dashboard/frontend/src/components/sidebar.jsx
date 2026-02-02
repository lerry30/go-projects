import { Search } from 'lucide-react';
import { useState, useRef } from 'react';

import { getData } from '../utils/send';
import { zServer } from '../store/server';
import { zForecast } from '../store/weather-forecast';

import FaveCities from './faveCities';

const Sidebar = () => {
    const zServerUrl = zServer(state => state.serverUrl);
    const zFetchWeatherForecast = zForecast(state => state.fetchWeatherForecast);

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
                // if item is selected from the list
                cityInputField.current.value = city;
            }
            
            setIsSuggestionOpen(false);
            await zFetchWeatherForecast(zServerUrl, city);
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
        <aside 
            className="sm:w-full lg:w-[20%] 
                h-full bg-gray-700 overflow-hidden p-4 border-r border-gray-300/20 flex flex-col gap-4
                sm:rounded-tl-lg sm:rounded-tr-lg lg:rounded-tr-none lg:rounded-bl-lg
            "
        >
            <div 
                className={
                    `w-full h-12 bg-slate-300 flex items-center py-3 px-4 relative 
                    ${isSuggestionOpen ? "rounded-t-xl" : "rounded-xl"}`
                }
            >
                <input
                    ref={cityInputField}
                    className="w-full h-6 outline-0 font-sans"
                    placeholder="Search city"
                    onChange={elem => openSuggestion(elem.target.value)}
                    onKeyDown={e => e.key === "Enter" && getWeatherForecast(e.target.value)}
                />
                <Search />
                <BoxSuggestion />
            </div>
            <div className="w-full h-full flex flex-col justify-end">
                <FaveCities />
            </div>
        </aside>
    )
}

export default Sidebar;