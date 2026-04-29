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

    const token = res.headers.get('Authorization');
    if(token) localStorage.setItem('token', token);

    return payload;
}