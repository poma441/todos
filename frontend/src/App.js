import cl from './styles/App.css'
import React, {useState,userRef} from 'react';
import Button from './components/UI/button/Button'
import MyCalendar from './components/MyCalendar';
import TaskItem from './components/TaskItem';
import TaskForm from './components/TaskForm';
import { createStore } from 'redux';
function App() {

  const [tasks, setTasks] = useState ([
    {id: 1, task:'make front',active: true,  isred: false},
    {id:2, task: 'make back',active: true,  isred: false},
    {id:3, task: 'make back2',active: false,  isred: false},
    {id:4, task: 'make back3',active: false,  isred: false}
  ])
  
  //const reducer = (state,action) => {
  //  switch (action.type){

   //   default:
   //     return state
  //  }

  //}

  //const store = createStore()

  const removeTask = (task)=>{
    setTasks(tasks.filter(tsk=>tsk.id!==task.id))
  }

  const createTask= (newTask)=> {
    setTasks([...tasks, newTask])
  }

  const editTask = (task, param,value) => {
    setTasks (tasks.map(tsk => {
      if (tsk.id === task.id) return {...tsk, [param]:value }
      else return tsk;
    }))
    console.log(task)
  }

  const getDead = (task) => {
    if (new Date(task.deadline)>new Date()) return 'task good';
      else {if(new Date(task.deadline)<new Date()) return 'task bad';
            else return 'task near' ;
            } 
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
        getDead = {getDead} 
        number={index+1}/>
        
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
        getDead = {getDead} 
        number={index+1}/>
        
    )}
    </div>
  );
}

export default App;
