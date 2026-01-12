import { useState, useEffect } from 'react'
import { Trash2 } from 'lucide-react';
import { sendJson, getData } from './utils/send'

import Modal from './components/modals/modal';

const App = () => {
  const urlPath = 'http://localhost:8080';

  const [task, setTask] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  const [todos, setTodos] = useState([])

  const [resModal, setResModal] = useState(false)
  const [modalMessage, setModalMessage] = useState({title: '', message: ''})

  const submitNewTask = async () => {
    try {
      if(!task) throw new Error('Todo should not be empty.');

      const response = await sendJson(`${urlPath}/newtask`, {task});
      if(response) {
        setResModal(true);
        setModalMessage({title: 'Awesome', message: response?.message});
      }
    } catch(error) {
      setErrorMessage(error?.message);
    } finally {
      setTask('');
    }
  }

  const getTodos = async () => {
    try {
      const data = await getData(`${urlPath}/todos`);
      if(Array.isArray(data)) {
        setTodos(data.reverse());
      }
    } catch(error) {
      console.log('No todos.');
    }
  }

  const updateTodo = async (id, status) => {
    try {
      if(!id) throw new Error('Undefined todo.');

      const response = await sendJson(`${urlPath}/status`, {id, completed: !status}, 'PUT');
      if(response) {
        setTodos(t => (
          t.map((item) => (
            item.id === id ? {...item, completed: !item?.completed} : item
          )
        )))
        setResModal(true);
        setModalMessage({title: 'Awesome', message: response?.message});
      }
    } catch(error) {
      setErrorMessage(error?.message)
    }
  }

  const removeTodo = async (id) => {
    try {
      if(!id) throw new Error('Undefined todo.');

      const response = await sendJson(`${urlPath}/remove`, {id}, 'DELETE');
      if(response) {
        setResModal(true);
        setModalMessage({title: 'Awesome', message: response?.message});
      }
    } catch(error) {
      setErrorMessage(error?.message)
    }
  }

  useEffect(() => {
    getTodos()
  }, [resModal]);

  return (
    <div className="min-h-screen bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center">
      <div className="w-[500px] flex justify-center items-center flex-col gap-2">
        <h1 className="text-4xl font-bold text-white">Hello Todoers</h1>
        <div className="w-full p-8 bg-white flex justify-center items-center flex-col gap-4 rounded-sm">
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
            className="bg-black text-white w-full max-w-[400px] outline-none rounded-lg py-2 px-4 cursor-pointer"
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

        <div className="w-full">
          {
            todos.map((todo, index) => (
              <div 
                key={index} 
                className="flex justify-between items-center my-2 bg-slate-500 p-2 rounded-sm"
              >
                <div className="flex items-center gap-4">
                  <input 
                    type="checkbox" 
                    checked={todo.completed}
                    onChange={() => updateTodo(todo?.id, todo?.completed)}
                    className="w-6 h-6 cursor-pointer"
                  />
                  <p className="text-white text-xl">{todo.task}</p>
                </div>
                <button 
                  className="cursor-pointer"
                  onClick={() => removeTodo(todo?.id)}
                >
                  <Trash2 className="text-red-400" />
                </button>
              </div>
            ))
          }
        </div>
      </div>
    </div>
  )
}

export default App;