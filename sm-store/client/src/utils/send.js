export const sendJSON = async (urlPath, objData, method="POST") => {
    const res = await fetch(urlPath, {
        credentials: 'include',
        method: method,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(objData)
    });

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.message || 'Failed to fetch';
        throw new Error(message, { cause: { status: res.status, payload: payload } });
    }

    const token = res.headers.get('Authorization');
    if(token) localStorage.setItem('token', token);

    return payload;
}

export const sendForm = async (urlPath, form, method="POST") => {
    const res = await fetch(urlPath, {
        credentials: 'include',
        method: method,
        body: form
    });

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.message || 'Failed to fetch';
        throw new Error(message, { cause: { response: res, payload } });
    }
    
    return payload;
}

export const getData = async (urlPath) => {
    const res = await fetch(urlPath, {credentials: 'include'});

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.message || 'No response';
        throw new Error(message, { cause: { response: res, payload } });
    }

    return payload;
}