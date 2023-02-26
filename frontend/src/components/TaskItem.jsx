import React from "react";
import Button from "./UI/button/Button";
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import IconButton from '@mui/material/IconButton';
import ClearIcon from '@mui/icons-material/Clear';
import AddTaskIcon from '@mui/icons-material/AddTask';
import { useState } from "react";

const TaskItem = (props) =>{
    const [state, setState] = useState ('')

    if(props.task.active === props.setshow){



    if (props.task.isred === false) return (
       
        <div className= 'task good'>
             <input
             type='checkbox'
             checked={!props.task.active}
             onChange = {event => 
                props.editTask(props.task,'active',!props.task.active)
            }
             />
        {props.task.task}
        <div className= 'task_buttons>'>         
        <IconButton aria-label="edit" onClick = {
          () => {props.editTask(props.task,'isred',!props.task.isred)
          setState(props.task.task)
        }
      }>
        <EditIcon />
        </IconButton>
        <IconButton aria-label="delete" onClick = {() => props.removeTask(props.task)}>
        <DeleteIcon />
        </IconButton>

        </div>
      </div>
    )

    else return (
       <div className= 'task good'>
             <input
             type='checkbox'
             checked={!props.task.active}
             onChange = {event => 
                props.editTask(props.task,'active',!props.task.active)
            }
             />
      <input
      type='text'
      value = {props.task.task}
      disabled = {!props.task.isred}
      onChange = {event => 
          props.editTask(props.task,'task',event.target.value)
      }
      />
 <div className= 'task_buttons>'>  

        <IconButton aria-label="edit" onClick = {() => props.editTask(props.task,'isred',!props.task.isred)}>
        <AddTaskIcon />
        </IconButton>
        <IconButton aria-label="edit" onClick = {
          () => {
            props.editTask(props.task,'task',state)
            props.editTask(props.task,'isred',!props.task.isred)
          }}>
        <ClearIcon />
        </IconButton>
        <IconButton aria-label="delete" onClick = {() => props.removeTask(props.task)}>
        <DeleteIcon />
        </IconButton>

        </div>

      </div>

      
    )
    }
}
export default TaskItem;