export const sendJson = async (url, payload, method='POST') => {
    const res = await fetch(url, {
        credentials: 'include', // for cookies/auth
        method: method,
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });

    const resPayload = await res.json();
    if(!res.ok) {
        const message = resPayload?.message || 'Failed to fetch';
        throw new Error(message, {cause: { status: res.status, payload: resPayload }})
    }

    return resPayload
}