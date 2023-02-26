import cl from '../styles/App.css'
import React, {useState,userRef, useEffect} from 'react';
import TaskItem from '../components/TaskItem';
import TaskForm from '../components/TaskForm';
import { createStore } from 'redux';
import { useDispatch, useSelector } from 'react-redux';
import MyModal from '../components/UI/MyModal/MyModal';
import axios from 'axios';


const style = {
  textAlign: 'center',
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

function Tasks() {




  const dispatch = useDispatch()
  const tasks = useSelector(state => state.tasks)
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const createTask = (newTask)=>{
    dispatch({type:"ADD_TASK",payload: newTask})
    handleClose()
  }
  const removeTask = (task) => {
    dispatch({type: "REMOVE_TASK",payload: task})
  }
  const editTask = (task, param,value) => {
    dispatch({type: "EDIT_TASK",payload: {task:task, param:param, value:value}})
    console.log(task)
  }

  async function fetchTasks() {   
    const response = await axios.get('http://192.168.189.158:5100/todos/1')
    console.log(response.data)
  }

  useEffect( () => {
    fetchTasks()
  }, [])


  return (
    <div>
      <MyModal visible={open} titlebox={'Новое задание'} style={style} titlebutton={'Напишите задание'} handleOpen={handleOpen} handleClose={handleClose}>
        <TaskForm create ={createTask}/>
      </MyModal>
      <h1 style={{textAlign: 'center'}}>
         Текущие задачи 
         </h1>
      {
        tasks.map((task,index)=>
        
        <TaskItem
        setshow = {true}
        removeTask = {removeTask} 
        task = {task}
        editTask ={editTask}
        />
        
    )}
    <h1 style={{textAlign: 'center'}}>
         Выполненные задачи 
         </h1>
         {
        tasks.map((task,index)=>
        
        <TaskItem
        setshow = {false}
        removeTask = {removeTask} 
        task = {task}
        editTask ={editTask}
        />
        
    )}
    </div>
  );
}

export default Tasks;