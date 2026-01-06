import { useState } from 'react'
import { sendJson } from './utils/send'

import Modal from './components/modals/modal';

const App = () => {
  const [task, setTask] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  const [resModal, setResModal] = useState(false)
  const [modalMessage, setModalMessage] = useState({title: '', message: ''})

  const submitNewTask = async () => {
    try {
      if(!task) throw new Error('Todo should not be empty.');

      const response = await sendJson('http://localhost:8080/newtask', {task});
      if(response) {
        setResModal(true);
        setModalMessage({title: 'Awesome', message: response?.message});
      }
    } catch(error) {
      setErrorMessage(error?.message)
    } finally {
      setTask('')
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center flex-col gap-2">
      <h1 className="text-4xl font-bold text-white">Hello Todoers</h1>
      <div className="w-full max-w-[500px] p-8 bg-white flex justify-center items-center flex-col gap-4 rounded-sm">
        <input 
          id="task" 
          value={task} 
          onChange={elem => setTask(elem.target?.value)}
          className="w-full max-w-[400px] outline-none border-2 border-neutral-400 rounded-lg py-2 px-4"
          placeholder="Add new todo"
        />
        <button 
          id="submit-task"
          onClick={submitNewTask}
          className="bg-black text-white w-full max-w-[400px] outline-none rounded-lg py-2 px-4"
        >
          Submit
        </button>
        <p
          className="text-red-700"
        >
          {errorMessage}
        </p>
      </div>
      {resModal && (
        <Modal 
          title={modalMessage.title} 
          message={modalMessage.message} 
          close={() => setResModal(false)}
        />
      )}

      
    </div>
  )
}

export default App;