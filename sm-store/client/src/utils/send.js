export const sendJSON = async (url, objData, method="POST") => {
    const res = await fetch(url, {
        method: method,
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(objData)
    });

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.Error || 'Failed to fetch';
        throw new Error(message, { cause: { status: res.status, payload: payload } });
    }

    const token = res.headers.get('Authorization');
    if(token) localStorage.setItem('token', token);

    return payload;
}

export const sendForm = async (url, form, method="POST") => {
    const res = await fetch(url, {
        method: method,
        body: form
    });

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.Error || 'Failed to fetch';
        throw new Error(message, { cause: { response: res, payload } });
    }
    
    return payload;
}

/*
    apiRequest('/api/data', {
    method: 'POST',
    body: JSON.stringify({ name: 'John' }),
    auth: true,
    })
  */
export const apiRequest = async (url, options={}) => {
    const headers = {
        'Content-Type': 'application/json',
        ...(options?.headers || {})
    }

    if(options?.auth) {
        const token = localStorage.getItem('token');
        if(token) {
            headers['Authorization'] = token;
        }
    }

    const res = await fetch(url, {
        ...options,
        headers,
    });

    const payload = await res.json();
    if(!res?.ok) {
        const message = payload?.Error || 'No response';
        throw new Error(message, { cause: { response: res, payload } });
    }

    return payload;
}