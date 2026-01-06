import { useEffect } from 'react';

const Modal = ({title='', message='', close=()=>{}}) => {
    useEffect(() => {
        const interId = setInterval(() => {
            close();
        }, 1000);

        return () => clearInterval(interId); // clean up
    }, [])

    return (
        <div className="w-full h-screen max-h-screen bg-slate-600/75 absolute z-50 top-0 leading-0 right-0 bottom-0 flex justify-center items-center">
            <div className="w-full max-w-[500px] p-8 rounded-sm bg-white flex justify-center items-center flex-col gap-4">
                <h2 className="text-3xl font-medium">{title}</h2>
                <p className="text-lg">{message}</p>
            </div>
        </div>
    );
}

export default Modal;