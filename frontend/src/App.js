import cl from './styles/App.css'
import React, {useState,userRef} from 'react';
import Button from './components/UI/button/Button'
import TaskItem from './components/TaskItem';
import TaskForm from './components/TaskForm';
import { createStore } from 'redux';
import { useDispatch, useSelector } from 'react-redux';

function App() {
  const dispatch = useDispatch()
  const tasks = useSelector(state => state.tasks)
  

  const createTask = (newTask)=>{
    dispatch({type:"ADD_TASK",payload: newTask})
  }

  const removeTask = (task) => {
    dispatch({type: "REMOVE_TASK",payload: task})
  }



  const editTask = (task, param,value) => {
    dispatch({type: "EDIT_TASK",payload: {task:task, param:param, value:value}})
    console.log(task)
  }

  return (
    <div>
      <TaskForm create={createTask}/>



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

export default App;
