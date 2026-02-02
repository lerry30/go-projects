import { useState, useEffect } from 'react';
import { Heart } from 'lucide-react';
import { zFaveCities } from '../store/fave-cities';

const SaveCity = ({city}) => {
    const add = zFaveCities(fave => fave.add);
    const remove = zFaveCities(fave => fave.remove);
    const isSaved = zFaveCities(fave => fave.isSaved);
    
    const [ saved, setSaved ] = useState(false);

    const save = () => {
        if(!String(city).trim()) return;
        setSaved(s => {
            if(!s) {
                add(city);
                return true;
            }

            remove(city);
            return false;
        })
    }

    useEffect(() => {
        if(!String(city).trim()) return;
        setSaved(isSaved(city));
    }, []);

    return (
        <Heart 
            fill={`${saved ? '#b11' : 'none'}`}
            stroke={`${saved ? 'none' : 'white'}`}
            onClick={save} 
        />
    );
}

export default SaveCity;