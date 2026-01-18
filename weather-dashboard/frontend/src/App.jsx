import { useEffect } from 'react';
import { getData } from './utils/send';

function App() {
  const urlPath = 'http://localhost:8080';

  const getMyName = async () => {
    try {
      const data = await getData(`${urlPath}/name`);
      console.log(data);
    } catch(error) {
      console.log(error)
    }
  }

  useEffect(() => {
    getMyName();
  }, [])

  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center">
      <h1 className="text-4xl font-bold text-white">Hello Tailwind + React!</h1>
    </div>
  )
}

export default App;