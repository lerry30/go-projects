// {credentials: 'include'} for auth or cookies

export const getData = async (urlPath) => {
    const res = await fetch(urlPath, {credentials: 'include'});

    const payload = await res.json();
    console.log(payload, '---------');
    if(!res.ok) {
        const message = payload?.message || 'No response';
        throw new Error(message, {cause: {response: res, payload}});
    }

    return payload;
}
